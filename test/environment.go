package test

import (
	"context"
	"os"
	"testing"

	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/Cattle0Horse/url-shortener/internal/global/database/mysql"
	"github.com/Cattle0Horse/url-shortener/tools"
	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	EnvFileName   = "docker-compose.env.yaml"
	ConfigFilName = "config.example.yaml"
)

func IsTest() bool {
	return os.Getenv("ENV") == "test"
}

func SetupEnvironment(t *testing.T) {
	t.Setenv("ENV", "test")
	compose, err := tc.NewDockerCompose(tools.SearchFile(EnvFileName))
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal))
	})
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	require.NoError(t,
		compose.WaitForService("mysql", wait.ForLog("port: 3306  MySQL Community Server")).Up(ctx, tc.Wait(true)),
	)

	config.Init(tools.SearchFile(ConfigFilName))
	mysql.Init()
}
