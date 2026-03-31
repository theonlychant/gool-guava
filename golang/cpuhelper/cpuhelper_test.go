package cpuhelper

import "testing"

func TestNumCPU(t *testing.T) {
    n := NumCPU()
    if n < 1 {
        t.Fatalf("NumCPU returned %d; expected >= 1", n)
    }
}
