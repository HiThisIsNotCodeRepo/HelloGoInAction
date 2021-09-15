# Chap7

*What runner can do?*

1. All tasks finish, program exit.
2. Program running overtime and then exit.
3. Program receive system interrupt signal then exit.

*What value can `start()` return?*

1. `nil` means program exit normally.
2. `ErrTimeout` means program running timeout.
3. `ErrInterrupt` means program being interrupt.

*`select` and `default`*

In `default` branch we can put non-blocking code to exit from `select`.

*Why `make(chan os.Signal, 1)`?*

The buffer channel is to ensure the program cache the first interrupt signal.

*Package `pool`*

Every time when request resource from pool, it will check if it has available cached resource. If not then create new
resource.

When release resource, it won't release immediately, the pool will check if its cache got vacancy if so the cache that
resource otherwise release it.

*Package `work`*

In `work` a set of goroutine will run tasks concurrently, this is achieved by unbuffered channel.
