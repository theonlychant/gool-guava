package tests

import (
    "testing"
    ch "github.com/example/cpuhelper"
)

func TestNumCPU15(t *testing.T) {
    if n := ch.NumCPU(); n == 0 {
        t.Fatal("NumCPU returned 0")
    }
}
