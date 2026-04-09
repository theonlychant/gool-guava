package cpuhelper

import (
    "bufio"
    "io"
    "os"
    "runtime"
    "strings"
)

// IsAndroid reports whether the current runtime appears to be Android.
// It checks runtime.GOOS and common Android filesystem markers.
func IsAndroid() bool {
    if runtime.GOOS == "android" {
        return true
    }
    if _, err := os.Stat("/system/build.prop"); err == nil {
        return true
    }
    return false
}

// GetSupportedAbis returns a best-effort list of supported ABIs on Android.
// On non-Android hosts it falls back to runtime.GOARCH.
func GetSupportedAbis() []string {
    // Try build.prop first
    if data, err := os.ReadFile("/system/build.prop"); err == nil {
        s := string(data)
        // common keys storing ABI lists
        keys := []string{"ro.product.cpu.abilist", "ro.product.cpu.abi", "ro.product.cpu.abilist32", "ro.product.cpu.abilist64"}
        for _, k := range keys {
            if v := extractProp(s, k); v != "" {
                // comma or comma-separated list
                parts := strings.Split(v, ",")
                for i := range parts {
                    parts[i] = strings.TrimSpace(parts[i])
                }
                return parts
            }
        }
    }

    // Fallback to reading /proc/cpuinfo for an 'Processor' or 'model name' line
    if info := ReadProcCpuinfo(); info != "" {
        // heuristically try to find 'model name' or 'Processor' and map to arch
        lower := strings.ToLower(info)
        if strings.Contains(lower, "armv7") || strings.Contains(lower, "armv7l") {
            return []string{"armeabi-v7a"}
        }
        if strings.Contains(lower, "aarch64") || strings.Contains(lower, "arm64") {
            return []string{"arm64-v8a"}
        }
        if strings.Contains(lower, "intel") || strings.Contains(lower, "x86_64") {
            return []string{"x86_64"}
        }
        if strings.Contains(lower, "x86") {
            return []string{"x86"}
        }
    }

    // Last resort: runtime architecture
    return []string{runtime.GOARCH}
}

func extractProp(build string, key string) string {
    rd := strings.NewReader(build)
    r := bufio.NewReader(rd)
    for {
        line, err := r.ReadString('\n')
        if err != nil && err != io.EOF {
            break
        }
        line = strings.TrimSpace(line)
        if strings.HasPrefix(line, key+"=") {
            return strings.TrimPrefix(line, key+"=")
        }
        if err == io.EOF {
            break
        }
    }
    return ""
}

// ReadProcCpuinfo returns the contents of /proc/cpuinfo if readable, otherwise
// an empty string.
func ReadProcCpuinfo() string {
    b, err := os.ReadFile("/proc/cpuinfo")
    if err != nil {
        return ""
    }
    return string(b)
}

// ParseCpuFeatures parses /proc/cpuinfo for a features/flags line and returns
// the space-separated feature tokens.
func ParseCpuFeatures() []string {
    info := ReadProcCpuinfo()
    if info == "" {
        return nil
    }
    lines := strings.Split(info, "\n")
    for _, l := range lines {
        low := strings.ToLower(l)
        if strings.HasPrefix(low, "features") || strings.HasPrefix(low, "flags") || strings.Contains(low, "cpu features") {
            idx := strings.Index(l, ":")
            if idx >= 0 && idx+1 < len(l) {
                rest := strings.TrimSpace(l[idx+1:])
                if rest == "" {
                    return nil
                }
                parts := strings.Fields(rest)
                return parts
            }
        }
    }
    return nil
}
