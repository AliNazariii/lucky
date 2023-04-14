package etcdlock

import (
	"context"
	"errors"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"lucky/internal/config"
	"lucky/internal/modules/locker"
	"sync"
)

type EtcdLockImpl struct {
	client *concurrency.Session
	data   sync.Map
}

func New(config config.EtcdConfig) locker.Locker {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: config.Endpoints,
	})
	if err != nil {
		panic(err)
	}

	// Test client
	status, err := client.Status(context.Background(), config.Endpoints[0])
	fmt.Println("Etcd client status:", status)
	if err != nil {
		panic(err)
	}

	// create a sessions to acquire a lock
	s, err := concurrency.NewSession(client)
	if err != nil {
		panic(err)
	}

	return &EtcdLockImpl{
		client: s,
	}
}

func (m *EtcdLockImpl) Lock(key string) error {
	// Obtain a new mutex by using the same name for all instances wanting the same lock.
	mutex := concurrency.NewMutex(m.client, key)

	ctx := context.Background()
	if err := mutex.Lock(ctx); err != nil {
		return err
	}

	m.data.Store(key, mutex)
	return nil
}

func (m *EtcdLockImpl) Unlock(key string) error {
	// Obtain a new mutex by using the same name for all instances wanting the same lock.
	mutex, exists := m.data.Load(key)
	if !exists {
		return errors.New("key not found")
	}

	ctx := context.Background()
	if err := mutex.(*concurrency.Mutex).Unlock(ctx); err != nil {
		return err
	}

	m.data.Delete(key)
	return nil
}
