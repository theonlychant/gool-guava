package tests

import (
    "testing"
    ch "github.com/example/cpuhelper"
)

func TestNumCPU07(t *testing.T) {
    if n := ch.NumCPU(); n < 0 {
        t.Fatalf("negative CPUs: %d", n)
    }
}
