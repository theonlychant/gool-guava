package cpuhelper

import (
    "testing"
)

func TestGCEDetection(t *testing.T) {
    // This test only verifies the helper runs without panicking; it's
    // skipped if not running on GCE.
    if !IsGCE() {
        t.Skip("not running on GCE; skipping")
    }
    _ = GCEInstanceType()
}

func TestGPUDetection(t *testing.T) {
    // GPU presence varies by environment; just exercise the API.
    _ = HasNvidiaGPU()
    _ = NvidiaGPUInfo()
}
