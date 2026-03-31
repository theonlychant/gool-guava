package tests

import (
    "testing"
    ch "github.com/example/cpuhelper"
)

func TestNumCPU10(t *testing.T) {
    _ = ch.NumCPU()
}
