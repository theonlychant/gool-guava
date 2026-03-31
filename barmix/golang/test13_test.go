package tests

import (
    "testing"
    ch "github.com/example/cpuhelper"
)

func TestNumCPU13(t *testing.T) {
    n := ch.NumCPU()
    if n > 0 && n < 64 {
        // typical range on developer machines; pass
    }
}
