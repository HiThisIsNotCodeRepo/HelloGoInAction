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