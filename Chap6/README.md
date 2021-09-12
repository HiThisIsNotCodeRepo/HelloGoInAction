# Chap6

*process and thread*

Process is a container of computational resources that an application needs to run.

Thread is a space to run code, the operating system will schedule thread on processor, it doesn't need to be the
processor that run application.

*concurrency and parallelism*

Parallelism is to run different snippet on different processor, parallelism needs multicore.

Concurrency is to do a lot of things at the same time, concurrency can happen on single core.

*Configure logic core*

```
runtime.GOMAXPROCS(runtime.NumCPU())
```

*Race condition*

If more than one goroutine read/write shared resource that will cause race condition. It should be avoided.

*Race condition solution*

There are couple ways to prevent race condition.

1. Lock resource, `atomic.AddInt64`,`atomic.LoadInt64`,`atomic.StoreInt64`.
2. Mutex, `mutex.Lock()`, `mutex.Unlock()`.
