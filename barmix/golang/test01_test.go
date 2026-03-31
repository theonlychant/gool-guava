package tests

import (
    "testing"
    ch "github.com/example/cpuhelper"
)

func TestNumCPU01(t *testing.T) {
    if n := ch.NumCPU(); n < 1 {
        t.Fatalf("expected >=1, got %d", n)
    }
}
