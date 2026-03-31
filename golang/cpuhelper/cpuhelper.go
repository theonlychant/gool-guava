package cpuhelper

import "runtime"

// NumCPU returns the number of logical CPUs usable by the current process.
func NumCPU() int {
    return runtime.NumCPU()
}
