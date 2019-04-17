package utils

import (
	"fmt"

	"os/exec"
)

func MongoStart() {
	cmd := "docker run  -d -p 27017:27017 --name auction_db mongo"
	out, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		panic("Cannot start mongodb container")
	}
	fmt.Printf("%s", out)
}

func MongoKill() {
	cmd := "docker kill auction_db "
	out, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		panic("Cannot kill mongodb container")
	}
	fmt.Printf("%s", out)
}

func MongoRemove() {
	cmd := "docker rm auction_db "
	out, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		panic("Cannot kill mongodb container")
	}
	fmt.Printf("%s", out)
}

/*
func StartDockerMongo() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		panic(err)
	}

	imageName := "mongo"

	hostBinding := nat.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: "27017",
	}
	containerPort, err := nat.NewPort("tcp", "27017")
	if err != nil {
		panic("Unable to get the port")
	}
	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}

	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)

	resp, err := cli.ContainerCreate(ctx,
		&container.Config{
			Image: imageName,
		},
		&container.HostConfig{
			PortBindings: portBinding,
		},
		nil,
		"auction_db")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)
}

func StopAllContainer() {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Print("Stopping container ", container.ID[:10], "... ")
		if err := cli.ContainerStop(ctx, container.ID, nil); err != nil {
			panic(err)
		}
		fmt.Println("Success")
	}

}
*/
