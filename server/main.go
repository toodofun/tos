package main

import (
	"fmt"
	"github.com/MR5356/tos/config"
	"github.com/MR5356/tos/constant"
	_ "github.com/MR5356/tos/log"
	"github.com/MR5356/tos/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"runtime"
)

var (
	port  int
	debug bool
	root  string
)

var cmd = &cobra.Command{
	Version: constant.Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.Current(
			config.WithDebug(debug),
			config.WithPort(port),
			config.WithStorageRoot(root),
		)

		svc, err := server.New(cfg)
		if err != nil {
			return err
		}

		return svc.Run()
	},
}

func init() {
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true

	cmd.SetVersionTemplate(fmt.Sprintf("TOS %s[%s] %s %s with %s %s\n", constant.Version, constant.Commit, runtime.GOOS, runtime.GOARCH, runtime.Version(), constant.BuildTime))
	cmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debug mode")
	cmd.PersistentFlags().IntVarP(&port, "port", "p", 61296, "server port")
	cmd.PersistentFlags().StringVarP(&root, "root", "r", "/tmp/tos/data", "root path")
}

func main() {
	if err := cmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
