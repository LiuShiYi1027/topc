package main

import (
	"fmt"
	"os"

	"github.com/LiuShiYi1027/topc/version"
	"github.com/spf13/cobra"
)

func main() {
	cmd := newTopcCommand()
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newTopcCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "topc [OPTIONS] COMMAND [ARG...]",
		Short:   "A performance monitor for container",
		Version: fmt.Sprintf("%s, build %s", version.Version, version.GitCommit),
	}

	cmd.AddCommand(newVersionCommand())
	cmd.AddCommand(newMemCommand())

	return cmd
}

func newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of TopC",
		Long:  `All software has versions. This is TopC's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("v%s -- HEAD\n", version.Version)
		},
	}
}

func newMemCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "memory",
		Short: "Print the memory info of current container",
		Run: func(cmd *cobra.Command, args []string) {
			
		},
	}
}
