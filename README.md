# gool -- CPU helpers and Guava build tooling

[![Go CI](https://github.com/theonlychant/gool-guava/actions/workflows/go-ci.yml/badge.svg)](https://github.com/theonlychant/gool-guava/actions/workflows/go-ci.yml) [![Java CI](https://github.com/theonlychant/gool-guava/actions/workflows/java-ci.yml/badge.svg)](https://github.com/theonlychant/gool-guava/actions/workflows/java-ci.yml)

This repository collects small, practical helpers and example projects for
inspecting and profiling CPU/GPU environments in Go and Java, plus a
centralized Guava build/test area and CI workflows.

Overview
--
- Go CPU helpers: `golang/cpuhelper` contains runtime and hardware helpers,
  pprof utilities, Android-friendly detection helpers, and cloud/GPU probes.
- Java examples: `java/cpuhelper` contains a Java CPU helper and examples.
- Guava build: `src/guava-build` contains a small Gradle project that pulls
  Guava and runs example tests. A copied subset of Guava tests is available
  under `src/guava-build/guava-tests` for convenience.
- Integration tests and examples for cross-language experimentation are under
  `barmix/`.

Repository layout
--
- `golang/` - Go modules and examples.
  - `golang/cpuhelper` - main Go module with helpers and tests.
- `java/` - Java example projects.
  - `java/cpuhelper` - Java CPU helper project.
- `src/guava-build` - standalone Gradle project for Guava examples and tests.
- `barmix/` - combined example test suites and profiling examples.
- `.github/workflows` - GitHub Actions workflows for Go and Java CI.

Quick start - Go
--
Run tests for a single module:

```bash
cd golang/cpuhelper
go test ./... -v
```

Run tests for every Go module in the repo (CI-style):

```bash
for mod in $(find . -name 'go.mod' -not -path './vendor/*' -exec dirname {} \; | sort -u); do
  (cd "$mod" && go test ./... -v)
done
```

Quick start - Java/Gradle
--
If you have Gradle installed you can run per-project tests. Examples:

```bash
gradle -p java/cpuhelper test --no-daemon --console=plain
gradle -p src/guava-build test --no-daemon --console=plain
gradle -p barmix/java test --no-daemon --console=plain
```

To make builds reproducible in CI, consider adding the Gradle wrapper to each
Java project and invoking `./gradlew test` instead of a system `gradle`.

Notes on Android support
--
- The Go module includes Android-friendly helpers that attempt to detect
  runtime/ABI information without requiring Android SDK APIs at build time.
- For full Android builds or Android-specific instrumentation prefer native
  Android/Kotlin helpers and the Android toolchain.

CI
--
- The repository includes two GitHub Actions workflows:
  - `.github/workflows/go-ci.yml` -- runs Go tests per module using Go 1.22.
  - `.github/workflows/java-ci.yml` -- discovers Gradle projects and runs
    `gradle test` for each project found (uses JDK 17).

Development notes
--
- Tests that query external services (Maven Central, GitHub API) are
  best-effort and tolerant to network failures; CI may report different
  results depending on network availability.
- Large imported folders such as `src/guava-build/guava-tests` are included
  to aid local experimentation; consider converting them to a submodule or
  pruning unneeded files for a smaller repository.

Contributing
--
- Pull requests are welcome. Please run the relevant tests locally before
  submitting. If you add Java projects, include the Gradle wrapper to avoid
  CI version mismatches.

License
--
This repository does not include a license file. Add a license if you want to
clarify reuse terms.

Contact
--
If you want me to make further structural changes (add README sections,
add Gradle wrappers, clean up large folders, or add CI linting), tell me which
task to do next.
