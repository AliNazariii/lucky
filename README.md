# Lucky - A Distributed Lock Implementation in Golang

## Introduction

Lucky is a Golang implementation of a distributed lock, which is used in distributed systems to ensure that only one
process can access a critical section at a time. The critical section is the code that needs to be executed by only one
process at a time.

## Goal

Distributed locking is a common problem in distributed systems, and there are many solutions and algorithms to handle
it. Most of them use third-party tools such as ZooKeeper, Redis, etc. However, the goal of Lucky is to provide a
lightweight and simple library that solves this problem without any external dependencies.

## Resources

Resources that were referred to during the development of Lucky include:

- [Distributed Locks with Redis](https://redis.io/docs/manual/patterns/distributed-locks/)
- [How to do distributed locking](https://martin.kleppmann.com/2016/02/08/how-to-do-distributed-locking.html)
- [Is Redlock safe?](http://antirez.com/news/101)
- [Everything I know about distributed locks](https://davidecerbo.medium.com/everything-i-know-about-distributed-locks-2bf54de2df71)
- [Distributed Locks are Dead; Long Live Distributed Locks!](https://hazelcast.com/blog/long-live-distributed-locks/)
- [The Technical Practice of Distributed Locks in a Storage System](https://www.alibabacloud.com/blog/the-technical-practice-of-distributed-locks-in-a-storage-system_597141)
- [Control concurrency for shared resources in distributed systems with DLM (Distributed Lock Manager)](https://m-qafouri.medium.com/serialize-access-to-a-shared-resource-in-distributed-systems-with-dlm-distributed-lock-manager-5abf5e393e15)

## Roadmap

The current roadmap for Lucky includes:

- [X] Implementing a module using the RedLock algorithm and Redis as a backend
- [ ] Implementing a module using ZooKeeper as a backend
- [X] Implementing a module using etcd as a backend
- [ ] Implementing a module using MySQL as a backend
- [ ] Implementing a module using PostgreSQL as a backend
- [ ] Implementing a module using MongoDB as a backend
- [ ] Implementing a module using a simple file as a backend
- [ ] Implementing a module using simple in-memory storage as a backend

## Contributing

Contributions to Lucky are always welcome! Please feel free to open issues or submit pull requests to help improve the
project.

## License

Lucky is licensed under the [MIT License](LICENCE).
