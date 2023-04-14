package redlock

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	redislib "github.com/redis/go-redis/v9"
	"lucky/internal/config"
	"lucky/internal/modules/locker"
	"sync"
)

type RedLockImpl struct {
	locker *redsync.Redsync
	data   sync.Map
}

func New(config config.RedisConfig) locker.Locker {
	client := redislib.NewClient(&redislib.Options{
		Addr: config.Host,
	})
	pong := client.Ping(context.Background())
	fmt.Println("Redis client Ping response:", pong)
	pool := goredis.NewPool(client)

	// Create an instance of redisync to be used to obtain a mutual exclusion lock.
	rs := redsync.New(pool)

	return &RedLockImpl{
		locker: rs,
	}
}

func (m *RedLockImpl) Lock(key string) error {
	// Obtain a new mutex by using the same name for all instances wanting the same lock.
	mutex := m.locker.NewMutex(key)

	// Obtain a lock for our given mutex. After this is successful, no one else
	// can obtain the same lock (the same mutex name) until we unlock it.
	if err := mutex.Lock(); err != nil {
		return err
	}

	m.data.Store(key, mutex)
	return nil
}

func (m *RedLockImpl) Unlock(key string) error {
	// Obtain a new mutex by using the same name for all instances wanting the same lock.
	mutex, exists := m.data.Load(key)
	if !exists {
		return errors.New("key not found")
	}

	// Release the lock so other processes or threads can obtain a lock.
	if ok, err := mutex.(*redsync.Mutex).Unlock(); !ok || err != nil {
		return err
	}

	m.data.Delete(key)
	return nil
}
