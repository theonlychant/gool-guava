package cpuhelper

import (
    "bufio"
    "bytes"
    "context"
    "net"
    "net/http"
    "os"
    "os/exec"
    "strings"
    "time"
)

// IsGCE attempts to detect whether the current process is running on
// Google Compute Engine by querying the metadata server. It returns true
// if the metadata server responds.
func IsGCE() bool {
    ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
    defer cancel()
    req, err := http.NewRequestWithContext(ctx, "GET", "http://169.254.169.254/computeMetadata/v1/", nil)
    if err != nil {
        return false
    }
    req.Header.Set("Metadata-Flavor", "Google")
    transport := &http.Transport{
        DialContext: (&net.Dialer{Timeout: 300 * time.Millisecond}).DialContext,
    }
    client := &http.Client{Transport: transport}
    resp, err := client.Do(req)
    if err != nil {
        return false
    }
    defer resp.Body.Close()
    return resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent
}

// GCEInstanceType returns the instance machine type string from the GCE
// metadata server (e.g. "projects/12345/zones/us-central1-a/machineTypes/n1-standard-1").
// Returns empty string if not available.
func GCEInstanceType() string {
    ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
    defer cancel()
    req, err := http.NewRequestWithContext(ctx, "GET", "http://169.254.169.254/computeMetadata/v1/instance/machine-type", nil)
    if err != nil {
        return ""
    }
    req.Header.Set("Metadata-Flavor", "Google")
    transport := &http.Transport{
        DialContext: (&net.Dialer{Timeout: 300 * time.Millisecond}).DialContext,
    }
    client := &http.Client{Transport: transport}
    resp, err := client.Do(req)
    if err != nil {
        return ""
    }
    defer resp.Body.Close()
    // read body into a buffer
    b := new(bytes.Buffer)
    _, _ = b.ReadFrom(resp.Body)
    return strings.TrimSpace(b.String())
}

// HasNvidiaGPU returns true if an NVIDIA GPU appears to be available on the host.
// It prefers to run `nvidia-smi` if present, otherwise checks for driver files.
func HasNvidiaGPU() bool {
    if _, err := exec.LookPath("nvidia-smi"); err == nil {
        out, err := exec.Command("nvidia-smi", "-L").Output()
        if err == nil && len(bytes.TrimSpace(out)) > 0 {
            return true
        }
    }
    // Fallback: check for /proc/driver/nvidia
    if fi, err := os.Stat("/proc/driver/nvidia/version"); err == nil && !fi.IsDir() {
        return true
    }
    return false
}

// NvidiaGPUInfo returns a short description of detected NVIDIA GPUs, or
// an empty string if none detected. It runs `nvidia-smi --query-gpu=name --format=csv,noheader`.
func NvidiaGPUInfo() string {
    if _, err := exec.LookPath("nvidia-smi"); err != nil {
        return ""
    }
    out, err := exec.Command("nvidia-smi", "--query-gpu=name", "--format=csv,noheader").Output()
    if err != nil {
        return ""
    }
    // return unique lines joined
    scanner := bufio.NewScanner(bytes.NewReader(out))
    var parts []string
    seen := map[string]bool{}
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" {
            continue
        }
        if !seen[line] {
            seen[line] = true
            parts = append(parts, line)
        }
    }
    return strings.Join(parts, "; ")
}
