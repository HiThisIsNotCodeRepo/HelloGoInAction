# Chap 2

*Structure of program*

1. Main goroutine
2. Search goroutine
3. Trace goroutine

*Steps*

1. Main goroutine receive data.
2. Main goroutine create search go routine to search.
3. Search go routine will send data back to main goroutine to display.
4. When search go routine complete job it will send message to trace goroutine to notify.
5. Trace goroutine will keep an internal counter when all search go routine complete it will send end signal to main go
   routine.