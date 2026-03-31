package profiling

import (
    "os"
    "runtime/pprof"
    "sync"
    "time"
)

// WorkerPoolProcess runs `jobs` CPU-bound tasks using a fixed number of workers.
// It uses a sync.Pool to reuse buffers and returns the combined result.
func WorkerPoolProcess(jobs int, workers int) int64 {
    if workers <= 0 {
        workers = 1
    }
    jobsCh := make(chan int, jobs)
    results := make(chan int64, jobs)

    var pool = sync.Pool{
        New: func() interface{} { b := make([]byte, 1024); return &b },
    }

    var wg sync.WaitGroup
    worker := func() {
        defer wg.Done()
        for j := range jobsCh {
            // get buffer from pool (reduces allocations)
            bufp := pool.Get().(*[]byte)
            buf := *bufp
            _ = buf // pretend to use buffer for I/O or formatting

            // small CPU-heavy work: sum of squares for a small range
            var acc int64
            for i := 0; i < 1000; i++ {
                v := int64((j+i)&0xffff)
                acc += v * v
            }

            // return buffer to pool
            pool.Put(bufp)
            results <- acc
        }
    }

    // start workers
    for w := 0; w < workers; w++ {
        wg.Add(1)
        go worker()
    }

    // enqueue jobs
    go func() {
        for j := 0; j < jobs; j++ {
            jobsCh <- j
        }
        close(jobsCh)
    }()

    // wait for workers in a background goroutine to close results when done
    go func() {
        wg.Wait()
        close(results)
    }()

    var total int64
    for r := range results {
        total += r
    }
    return total
}

// StartCPUProfile captures a CPU profile for the given duration and writes it to path.
// Returns the path written or an error.
func StartCPUProfile(path string, duration time.Duration) error {
    f, err := os.Create(path)
    if err != nil {
        return err
    }
    defer f.Close()
    if err := pprof.StartCPUProfile(f); err != nil {
        return err
    }
    time.Sleep(duration)
    pprof.StopCPUProfile()
    return nil
}

// SumSquaresNaive computes sum of squares 1..n using a loop.
func SumSquaresNaive(n int) int64 {
    var s int64
    for i := 1; i <= n; i++ {
        s += int64(i) * int64(i)
    }
    return s
}

// SumSquaresFormula computes sum of squares using formula n(n+1)(2n+1)/6.
func SumSquaresFormula(n int) int64 {
    nn := int64(n)
    return nn * (nn + 1) * (2*nn + 1) / 6
}
