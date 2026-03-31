package tests

import (
    "testing"
    ch "github.com/example/cpuhelper"
)

func TestNumCPU04(t *testing.T) {
    for i := 0; i < 5; i++ {
        _ = ch.NumCPU()
    }
}
