package tests

import (
	"gororoba/internal"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	COMPOSE_PATH       = "../docker-compose.yml"
	COMPOSE_IDENTIFIER = "g_e2e"
)

type HttpTestSuite struct {
	suite.Suite
	internal.AppServer
}

// func BeforeTest(t *testing.T) {
// 	identifier := tc.StackIdentifier(COMPOSE_IDENTIFIER)
// 	files := tc.WithStackFiles(COMPOSE_PATH)
// 	compose, err := tc.NewDockerComposeWith(files, identifier)
// 	require.NoError(t, err, "Failed to create docker-compose")

// 	ctx, cancel := context.WithCancel(context.Background())
// 	t.Cleanup(cancel)

// 	waitStrategy := wait.ForListeningPort("4566").WithStartupTimeout(30 * time.Second)
// 	err = compose.
// 		WaitForService("localstack", waitStrategy).
// 		Up(ctx, tc.Wait(true))

// 	require.NoError(t, err, "Failed to start services")

// 	srv := NewTestServer()

// 	t.Cleanup(func() {
// 		require.NoError(t, compose.Down(context.Background()), tc.RemoveOrphans(true), tc.RemoveImagesLocal, "Failed to stop compose")
// 		srv.ForceShutdown()
// 	})

// 	srv.Start()
// }

// func TestSuiteRunner(t *testing.T) {
// 	its := IntegrationTestSuite{}
// 	suite.Run(t, &its)
// }

func TestStartIntegrationTests(t *testing.T) {
	ts := HttpTestSuite{
		AppServer: NewTestServer(),
	}

	ts.AppServer.Start()

	t.Cleanup(func() {
		ts.AppServer.ForceShutdown()
	})

	suite.Run(t, &ts)
}
