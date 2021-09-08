# 1. Goroutine
## 1. go
## 2. wg
1. var wg sync.WorkGroup
    - wg.Add(1)
    - wg.Done()
    - wg.Wait()
# 2. Channel
1. describe
    - var c1:=make(chan int,10)
    - c1 <- 8 # send
    - <- c1 # receive
    - close(c1)
    - func f1(ch1 <-chan int)
2. select
# 3. concurrent and lock
1. sync.Mutex
    - var lock sync.Mutex
    - lock.Lock()
    - lock.Unlock()
2. sync.RWLock
    - var rwLock sync.RWLock
    - rwLock.RLock()
    - rwLock.RUnLock()
    - rwLock.Lock()
    - rwLock.Unlock()
3. sync.Map
4. sync.Once
===============
# 1. unit test

