package tests

import (
    "testing"
    ch "github.com/example/cpuhelper"
)

func TestNumCPU14(t *testing.T) {
    _ = ch.NumCPU()
}
