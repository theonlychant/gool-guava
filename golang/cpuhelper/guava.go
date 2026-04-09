package cpuhelper

import (
    "fmt"
    "net/http"
    "strings"
)

// FetchMavenArtifactInfo checks Maven Central for the given artifact and
// returns the artifact URL and content length when available.
// groupID like "com.google.guava", artifactID like "guava", version like "32.1.2-jre".
func FetchMavenArtifactInfo(groupID, artifactID, version string) (url string, size int64, available bool, err error) {
    base := "https://repo1.maven.org/maven2/"
    path := strings.ReplaceAll(groupID, ".", "/") + "/" + artifactID + "/" + version + "/"
    jar := artifactID + "-" + version + ".jar"
    url = base + path + jar

    req, err := http.NewRequest("HEAD", url, nil)
    if err != nil {
        return url, 0, false, err
    }
    req.Header.Set("User-Agent", "cpuhelper/1.0")
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return url, 0, false, err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return url, 0, false, fmt.Errorf("artifact not found: %s -> %d", url, resp.StatusCode)
    }
    return url, resp.ContentLength, true, nil
}
