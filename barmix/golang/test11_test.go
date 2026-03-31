package tests

import (
    "testing"
    ch "github.com/example/cpuhelper"
)

func TestNumCPU11(t *testing.T) {
    _ = ch.NumCPU()
}
