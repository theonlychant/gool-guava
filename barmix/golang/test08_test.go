package tests

import (
    "testing"
    ch "github.com/example/cpuhelper"
)

func TestNumCPU08(t *testing.T) {
    _ = ch.NumCPU() // simple smoke
}
