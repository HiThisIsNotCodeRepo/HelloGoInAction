# Chap 1

*Application Performance vs Rapid Development*

C/C++ can provide good application performance.

Ruby/Python is suitable for rapid development.

Go is good at both.

*Go Core Feature*

1. Rapid compiling.
2. Good concurrency support.
    1. goroutine similar to other language's thread but easier to create.
    2. `channel` allows different goroutine exchange message it makes goroutine sync easy.
    3. `channel` won't ensure the sync when different goroutine write to same data.
3. Go doesn't have inheritance. Go use composition design pattern. There's no extra code to declare interface
   implementation just implement method will do.
4. Go will handler garbage collection and this makes development faster.