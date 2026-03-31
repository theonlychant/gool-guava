package cpuhelper

import (
    "context"
    "log"
    "net"
    "net/http"
    _ "net/http/pprof"
    "os"
    "time"
)

// PprofServer holds runtime info about a running pprof HTTP server.
type PprofServer struct {
    Addr string
    stop chan struct{}
}

// StartPprof starts an HTTP server that exposes pprof endpoints on the given address
// (e.g. "localhost:6060"). It returns a *PprofServer which can be stopped by
// calling Close(). If addr is empty, "localhost:6060" is used.
func StartPprof(addr string) (*PprofServer, error) {
    if addr == "" {
        addr = "localhost:6060"
    }
    ln, err := net.Listen("tcp", addr)
    if err != nil {
        return nil, err
    }

    srv := &http.Server{}
    stop := make(chan struct{})

    go func() {
        if err := srv.Serve(ln); err != nil && err != http.ErrServerClosed {
            log.Printf("cpuhelper pprof server error: %v", err)
        }
    }()

    // support graceful shutdown
    go func() {
        <-stop
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()
        _ = srv.Shutdown(ctx)
    }()

    return &PprofServer{Addr: addr, stop: stop}, nil
}

// Close stops a running pprof server.
func (p *PprofServer) Close() {
    select {
    case <-p.stop:
        // already closed
    default:
        close(p.stop)
    }
}

// init checks CPUHELPER_PPROF_ADDR and starts a pprof server automatically
// if the variable is set (non-empty). This allows optional continuous
// profiling without changing application code.
func init() {
    addr := os.Getenv("CPUHELPER_PPROF_ADDR")
    if addr == "" {
        return
    }
    if _, err := StartPprof(addr); err != nil {
        log.Printf("failed to start cpuhelper pprof on %s: %v", addr, err)
    } else {
        log.Printf("cpuhelper pprof listening on %s", addr)
    }
}
