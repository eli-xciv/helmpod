package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func getUsage(cmd *cobra.Command, args []string) {
	var builder strings.Builder

	fmt.Fprintf(&builder, "help    Display this help menu\n")
	fmt.Fprintf(&builder, "install     Install a helm chart by name\n")
}
