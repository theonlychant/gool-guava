package tests

import (
    "testing"
    ch "github.com/example/cpuhelper"
)

func TestNumCPU06(t *testing.T) {
    n := ch.NumCPU()
    if n != int(n) {
        t.Fatalf("NumCPU not integer: %v", n)
    }
}
