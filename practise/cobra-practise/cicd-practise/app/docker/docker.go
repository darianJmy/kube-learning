package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	containername string = "gitlab"
)

type InitDockerOptions struct {
	Image    string   `json:"image"`
	Detach   bool     `json:"detach"`
	Hostname string   `json:"hostname"`
	Publish  []string `json:"publish"`
	Name     string   `json:"name"`
	Restart  string   `json:"restart"`
	Volume   []string `json:"volume"`
	Shmsize  string   `json:"shmsize"`
}

type Config struct {
	Config           container.Config
	HostConfig       container.HostConfig
	NetworkingConfig network.NetworkingConfig
	Platform         specs.Platform
}

func NewInitDockerOptions() *InitDockerOptions {
	return &InitDockerOptions{}
}

func Docker(initOptions *InitDockerOptions) *cobra.Command {
	if initOptions == nil {
		initOptions = NewInitDockerOptions()
	}

	cmd := &cobra.Command{
		Use:   "docker",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(initOptions)
		},
	}

	AddInitDockerFlags(cmd.Flags(), initOptions)

	return cmd
}

func run(initOptions *InitDockerOptions) error {

	cli, err := NewDockerClient()
	if err != nil {
		return err
	}

	cfg, err := checkInitDockerOptions(initOptions)
	if err != nil {
		return err
	}

	ctx := context.Background()
	list, err := cli.ContainerCreate(ctx, &cfg.Config, &cfg.HostConfig,
		&cfg.NetworkingConfig, &cfg.Platform, containername)
	if err != nil {
		return err
	}
	if err := cli.ContainerStart(ctx, list.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	return nil
}

func NewDockerClient() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return cli, nil
}

func checkInitDockerOptions(init *InitDockerOptions) (*Config, error) {
	var cfg Config
	cfg.Config.Image = init.Image
	cfg.Config.Hostname = init.Hostname
	var Volume map[string]string
	for k, _ := range init.Volume {
		volumes := strings.Split(init.Volume[k], ":")
		Volume[volumes[0]] = volumes[1]
	}
	return &cfg, nil
}

func AddInitDockerFlags(flagSet *pflag.FlagSet, init *InitDockerOptions) {
	flagSet.StringVar(
		&init.Image, "image", init.Image,
		"docker image",
	)
	flagSet.BoolVar(
		&init.Detach, "detach", init.Detach,
		`docker detach`,
	)
	flagSet.StringVar(
		&init.Hostname, "hostname", init.Hostname,
		`docker hostname`,
	)
	flagSet.StringArrayVar(
		&init.Publish, "publish", init.Publish,
		`docker publish`,
	)
	flagSet.StringVar(
		&init.Name, "name", init.Name,
		`docker name`,
	)
	flagSet.StringVar(
		&init.Restart, "restart", init.Restart,
		`docker name`,
	)
	flagSet.StringArrayVar(
		&init.Volume, "volume", init.Volume,
		`docker volume`,
	)
	flagSet.StringVar(
		&init.Shmsize, "shm-size", init.Shmsize,
		`docker shm-size`,
	)
}
