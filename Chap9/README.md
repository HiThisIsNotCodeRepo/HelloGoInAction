# Chap9

*Unit test*

1. Basic test
2. Table test
3. Mock test

*When will unit test be considered as fail*

When test doesn't call `t.Fatal` or `t.Error`

*Purpose of mock test*

It's not good for your test depends on external resource, because external resource is not guaranteed and without it
your can't run your test as your test will fail definitely. 
