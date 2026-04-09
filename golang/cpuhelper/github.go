package cpuhelper

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strings"
    "time"
)

// LatestRelease contains minimal data from GitHub releases.
type LatestRelease struct {
    TagName string `json:"tag_name"`
    HTMLURL string `json:"html_url"`
    Name    string `json:"name"`
}

// GetLatestRelease queries the GitHub API for the latest release of owner/repo
// (e.g., "google/guava"). It returns a minimal release structure.
func GetLatestRelease(ownerRepo string) (LatestRelease, error) {
    parts := strings.Split(ownerRepo, "/")
    if len(parts) != 2 {
        return LatestRelease{}, fmt.Errorf("ownerRepo must be owner/repo")
    }
    url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", ownerRepo)
    client := &http.Client{Timeout: 10 * time.Second}
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return LatestRelease{}, err
    }
    req.Header.Set("User-Agent", "cpuhelper/1.0")
    resp, err := client.Do(req)
    if err != nil {
        return LatestRelease{}, err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return LatestRelease{}, fmt.Errorf("github API returned %d", resp.StatusCode)
    }
    var r LatestRelease
    if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
        return LatestRelease{}, err
    }
    return r, nil
}
