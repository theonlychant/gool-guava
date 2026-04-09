package cpuhelper

import (
    "testing"
)

func TestGetCPUInfo(t *testing.T) {
    info := GetCPUInfo()
    if info.NumCPU <= 0 {
        t.Fatalf("expected NumCPU>0")
    }
    t.Log(info.PrettySummary())
}

func TestGuavaAndGithubHelpers(t *testing.T) {
    // Check Maven Central for a known Guava artifact (may be network-dependent)
    url, size, ok, err := FetchMavenArtifactInfo("com.google.guava", "guava", "32.1.2-jre")
    if err != nil {
        t.Logf("network error checking Maven Central: %v", err)
    } else {
        if !ok {
            t.Logf("artifact not available: %s", url)
        } else {
            t.Logf("guava jar: %s size=%d", url, size)
        }
    }

    // Query GitHub latest release for google/guava
    rel, err := GetLatestRelease("google/guava")
    if err != nil {
        t.Logf("github query failed: %v", err)
    } else {
        t.Logf("latest guava release: %s (%s)", rel.TagName, rel.HTMLURL)
    }
}
