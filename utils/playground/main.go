package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func main() {
	ctx := context.Background()

	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error creating Docker client: %v", err)
	}

	image := "drio/hello-server:latest"
	hostPort := "8888"

	containerPort, err := nat.NewPort("tcp", "7777")
	if err != nil {
		log.Fatalf("Error creating container port: %v", err)
	}

	containerConfig := &container.Config{
		Image: image,
		ExposedPorts: nat.PortSet{
			containerPort: struct{}{},
		},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			containerPort: []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: hostPort,
				},
			},
		},
	}

	networkConfig := &network.NetworkingConfig{}

	out, err := cli.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		log.Fatalf("Error pulling image: %v", err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)

	resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, networkConfig, nil, "my-container")
	if err != nil {
		log.Fatalf("Error creating container: %v", err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		log.Fatalf("Error starting container: %v", err)
	}

	fmt.Printf("Container %s started successfully\n", resp.ID)

	// Verify that the container is running
	inspect, err := cli.ContainerInspect(ctx, resp.ID)
	if err != nil {
		log.Fatalf("Error inspecting container: %v", err)
	}

	fmt.Printf("Container ports: %v\n", inspect.NetworkSettings.Ports)
}
