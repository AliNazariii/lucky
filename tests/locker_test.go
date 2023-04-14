package tests

import (
	"go.uber.org/atomic"
	"lucky/internal/config"
	"lucky/internal/modules/locker"
	"lucky/internal/modules/locker/redlock"
	"strconv"
	"sync"
	"testing"
)

func TestScenario(t *testing.T) {
	configs := config.GetConfig("lucky", "../config.yaml")

	redisLocker := redlock.New(configs.Redis)
	//etcdLocker := etcdlock.New(configs.Etcd)
	modules := map[string]locker.Locker{
		"redis": redisLocker,
		//"etcd": etcdLocker,
	}

	for name, lockerModule := range modules {
		t.Run(name, func(t *testing.T) {
			var wg sync.WaitGroup

			wg.Add(100)
			for i := 0; i < 100; i++ {
				key := "key-" + strconv.Itoa(i)
				go func(key string) {
					defer wg.Done()

					var successLocks atomic.Int32
					var lockWG sync.WaitGroup

					lockWG.Add(10)
					for j := 0; j < 10; j++ {
						go func() {
							defer lockWG.Done()
							err := lockerModule.Lock(key)
							if err != nil {
								//t.Errorf("Can't lock key: %s. err: %s", key, err)
								return
							}
							successLocks.Inc()
						}()
					}
					lockWG.Wait()

					if successLocks.Load() != int32(1) {
						t.Errorf("Different goroutines shouldn't lock the same key: %s. successLocks: %v", key, successLocks.Load())
					}

					err := lockerModule.Unlock(key)
					if err != nil {
						t.Errorf("Can't unlock key: %s. err: %s", key, err)
					}

					err = lockerModule.Lock(key)
					if err != nil {
						t.Errorf("Can't lock key: %s. err: %s", key, err)
					}
				}(key)
			}
			wg.Wait()
		})
	}
}
