//go:build !go1.19

package tunnel_test

//
// Psiphon not working with go1.19: TODO(https://github.com/ooni/probe/issues/2222)
//

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/apex/log"
	"github.com/ooni/probe-cli/v3/internal/engine"
	"github.com/ooni/probe-cli/v3/internal/tunnel"
)

func TestPsiphonStartStop(t *testing.T) {
	if testing.Short() {
		t.Skip("skip test in short mode")
	}
	tunnelDir, err := ioutil.TempDir("testdata", "psiphon")
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	sess, err := engine.NewSession(ctx, engine.SessionConfig{
		Logger:          log.Log,
		SoftwareName:    "miniooni",
		SoftwareVersion: "0.1.0-dev",
		TunnelDir:       tunnelDir,
	})
	if err != nil {
		t.Fatal(err)
	}
	tunnel, _, err := tunnel.Start(context.Background(), &tunnel.Config{
		Name:      "psiphon",
		Session:   sess,
		TunnelDir: tunnelDir,
	})
	if err != nil {
		t.Fatal(err)
	}
	if tunnel.SOCKS5ProxyURL() == nil {
		t.Fatal("expected non nil URL here")
	}
	if tunnel.BootstrapTime() <= 0 {
		t.Fatal("expected positive bootstrap time here")
	}
	tunnel.Stop()
}
