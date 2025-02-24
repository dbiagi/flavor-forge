package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	COMPOSE_PATH       = "../docker-compose.yml"
	COMPOSE_IDENTIFIER = "g_e2e"
)

func TestSetup(t *testing.T) {
	identifier := tc.StackIdentifier(COMPOSE_IDENTIFIER)
	files := tc.WithStackFiles(COMPOSE_PATH)
	compose, err := tc.NewDockerComposeWith(files, identifier)
	require.NoError(t, err, "Failed to create docker-compose")

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	waitStrategy := wait.ForListeningPort("4566").WithStartupTimeout(30 * time.Second)
	err = compose.
		WaitForService("localstack", waitStrategy).
		Up(ctx, tc.Wait(true))

	require.NoError(t, err, "Failed to start services")

	srv := NewTestServer()

	t.Cleanup(func() {
		require.NoError(t, compose.Down(context.Background()), tc.RemoveOrphans(true), tc.RemoveImagesLocal, "Failed to stop compose")
		srv.ForceShutdown()
	})

	srv.Start()
}
