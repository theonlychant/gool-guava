package tests

import (
    "io/ioutil"
    "os"
    "path/filepath"
    "testing"
    pr "github.com/example/cpuhelper_tests/profiling"
)

func TestWorkerPoolProcess(t *testing.T) {
    // run a modest number of jobs with worker pool
    res := pr.WorkerPoolProcess(50, 4)
    if res <= 0 {
        t.Fatalf("unexpected result: %d", res)
    }
}

func TestSumSquaresOptimized(t *testing.T) {
    n := 1000
    naive := pr.SumSquaresNaive(n)
    formula := pr.SumSquaresFormula(n)
    if naive != formula {
        t.Fatalf("naive %d != formula %d", naive, formula)
    }
}

func TestCPUProfileCapture(t *testing.T) {
    dir, err := ioutil.TempDir("", "pproftest")
    if err != nil {
        t.Skipf("cannot create temp dir: %v", err)
    }
    defer os.RemoveAll(dir)
    p := filepath.Join(dir, "cpu.pprof")
    // capture a short profile (0.5s)
    if err := pr.StartCPUProfile(p, 500*1e6); err != nil {
        t.Skipf("pprof not available: %v", err)
    }
    // ensure file exists
    if _, err := os.Stat(p); err != nil {
        t.Fatalf("profile not created: %v", err)
    }
}
