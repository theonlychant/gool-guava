package cpuhelper

import (
    "bufio"
    "fmt"
    "os"
    "runtime"
    "strings"
)

// NumCPU returns the number of logical CPUs usable by the current process.
func NumCPU() int {
    return runtime.NumCPU()
}

// NumGoroutine returns the number of goroutines currently existing.
func NumGoroutine() int {
    return runtime.NumGoroutine()
}

// CPUInfo is a compact snapshot of local CPU/runtime info.
type CPUInfo struct {
    NumCPU       int
    NumGoroutine int
    GOOS         string
    GOARCH       string
    ModelName    string
    IsDocker     bool
}

// GetCPUInfo collects CPU and runtime information. It's best-effort and
// safe to call on all supported platforms.
func GetCPUInfo() CPUInfo {
    info := CPUInfo{
        NumCPU:       NumCPU(),
        NumGoroutine: NumGoroutine(),
        GOOS:         runtime.GOOS,
        GOARCH:       runtime.GOARCH,
    }

    // Try to read model name from /proc/cpuinfo when available
    proc := ReadProcCpuinfo()
    if proc != "" {
        scanner := bufio.NewScanner(strings.NewReader(proc))
        for scanner.Scan() {
            line := scanner.Text()
            lower := strings.ToLower(line)
            if strings.HasPrefix(lower, "model name") || strings.HasPrefix(lower, "processor") || strings.HasPrefix(lower, "cpu part") {
                parts := strings.SplitN(line, ":", 2)
                if len(parts) == 2 {
                    info.ModelName = strings.TrimSpace(parts[1])
                    break
                }
            }
        }
    }

    info.IsDocker = isLikelyDocker()
    return info
}

func isLikelyDocker() bool {
    if _, err := os.Stat("/.dockerenv"); err == nil {
        return true
    }
    // check cgroup for docker/kubepods
    if data, err := os.ReadFile("/proc/1/cgroup"); err == nil {
        s := string(data)
        if strings.Contains(s, "docker") || strings.Contains(s, "kubepods") || strings.Contains(s, "containerd") {
            return true
        }
    }
    return false
}

// PrettySummary returns a one-line summary of CPU info.
func (c CPUInfo) PrettySummary() string {
    model := c.ModelName
    if model == "" {
        model = "unknown"
    }
    return fmt.Sprintf("os=%s arch=%s cpus=%d goroutines=%d model=%s docker=%v",
        c.GOOS, c.GOARCH, c.NumCPU, c.NumGoroutine, model, c.IsDocker)
}