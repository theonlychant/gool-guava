package tests

import (
    "testing"
    ch "github.com/example/cpuhelper"
)

func TestNumCPU12(t *testing.T) {
    if ch.NumCPU() <= 0 {
        t.Fatal("unexpected non-positive CPU count")
    }
}
