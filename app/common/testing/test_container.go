package testing

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/yunussandikci/go-pure-api/app/common"
)

type testContainer struct {
	container testcontainers.Container
}

type TestContainer interface {
	GetHostPort(port string) string
	Terminate()
}

func NewTestContainer(image string, ports []string, env map[string]string, waitStrategy wait.Strategy) TestContainer {
	container, _ := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        fmt.Sprintf("docker.io/library/%s", image),
			ExposedPorts: ports,
			Env:          env,
			WaitingFor:   waitStrategy,
		},
		Started: true,
	})
	return &testContainer{
		container: container,
	}
}

func (t *testContainer) GetHostPort(port string) string {
	host, _ := t.container.Host(context.Background())
	mappedPort, _ := t.container.MappedPort(context.Background(), nat.Port(port))
	natPort := mappedPort.Port()
	return fmt.Sprintf("%s:%s", host, natPort)
}

func (t *testContainer) Terminate() {
	err := t.container.Terminate(context.Background())
	if err != nil {
		common.Logger.Warn(err)
	}
}
