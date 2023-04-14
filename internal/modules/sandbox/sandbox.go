package sandbox

import (
	"fmt"
	"go.uber.org/atomic"
	"lucky/internal/modules/locker"
	timeutils "lucky/pkg/time"
	"strconv"
	"sync"
	"time"
)

type SandBoxImpl struct {
	locker locker.Locker
}

func New(locker locker.Locker) *SandBoxImpl {
	return &SandBoxImpl{
		locker: locker,
	}
}

func (m *SandBoxImpl) Run() {
	go timeutils.PrintTicker(1 * time.Second)
	var wg sync.WaitGroup

	// In the scenario, we have 100 goroutines and each of them will
	// try to get the lock the same key by 10 inner goroutines and then Unlock them
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

					err := m.locker.Lock(key)
					if err != nil {
						//fmt.Printf("[Fail] - %s: Can't lock. err: %s\n", key, err)
						return
					}
					successLocks.Inc()
				}()
			}
			lockWG.Wait()

			if successLocks.Load() == int32(1) {
				fmt.Printf("[SUCCESS] - %s: Different goroutines can't lock the same key\n", key)
			}

			time.Sleep(200 * time.Millisecond)

			err := m.locker.Unlock(key)
			if err != nil {
				fmt.Printf("[Fail] - %s: Can't Unlock lock. err: %s\n", key, err)
			}

			err = m.locker.Lock(key)
			if err != nil {
				fmt.Printf("[Fail] - %s: Can't lock. err: %s\n", key, err)
			}

			fmt.Printf("[SUCCESS] - %s: Locked the key again after Unlocking it\n", key)
		}(key)
	}
	wg.Wait()
}
