package tests

import (
    "testing"
    ch "github.com/example/cpuhelper"
)

func TestNumCPU02(t *testing.T) {
    _ = ch.NumCPU()
}
