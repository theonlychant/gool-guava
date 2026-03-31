package tests

import (
    "testing"
    ch "github.com/example/cpuhelper"
)

func TestNumCPU03(t *testing.T) {
    if n := ch.NumCPU(); n > 1024 {
        t.Fatalf("unrealistic CPU count: %d", n)
    }
}
