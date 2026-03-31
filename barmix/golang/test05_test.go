package tests

import (
    "testing"
    ch "github.com/example/cpuhelper"
)

func TestNumCPU05(t *testing.T) {
    if ch.NumCPU() <= 0 {
        t.Error("NumCPU must be positive")
    }
}
