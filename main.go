package main

import (
	"context"

	"github.com/spf13/cobra"

	"go-api-boilerplate/cmd"
)

func main() {
	cobra.CheckErr(cmd.NewCLI().ExecuteContext(context.Background()))
}
