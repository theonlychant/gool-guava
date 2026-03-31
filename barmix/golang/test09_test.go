package tests

import (
    "testing"
    ch "github.com/example/cpuhelper"
)

func TestNumCPU09(t *testing.T) {
    if ch.NumCPU() > 65536 {
        t.Fatalf("CPU count implausible: %d", ch.NumCPU())
    }
}
