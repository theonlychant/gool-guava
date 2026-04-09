package cpuhelper

import (
    "testing"
)

func TestAndroidHelpersRun(t *testing.T) {
    // Methods should run safely on non-Android hosts
    _ = IsAndroid()
    abis := GetSupportedAbis()
    if len(abis) == 0 {
        t.Log("no ABIs found; this is fine on non-Android hosts")
    }
    info := ReadProcCpuinfo()
    if info == "" {
        t.Log("/proc/cpuinfo not readable or empty")
    }
    feats := ParseCpuFeatures()
    t.Logf("abis=%v feats-count=%d", abis, len(feats))
}
