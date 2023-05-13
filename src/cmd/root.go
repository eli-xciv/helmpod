package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	// command flags

	// static
	version = "0.0.1-beta"

	rootCmd = &cobra.Command{
		Use:   "helmpod",
		Short: "helmpod allows you to deploy helm charts on podman",
		Long: `A utility that can run helm template and create your k8s pods from
                      a helm chart.
                      Full documentation can be found at https://github.com/eli-xciv/helmpod/docs`,
		Version: version,
	}

	installCmd = &cobra.Command{
		Use:   "install",
		Short: "install a helm chart using podman",
		Run:   spawnInstall,
	}

	templateCmd = &cobra.Command{
		Use:   "template",
		Short: "print out generated k8s objects",
		Run:   spawnTemplating,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(templateCmd)
	// rootCmd.SetHelpFunc(getUsage)
}

func spawnInstall(cmd *cobra.Command, args []string) {

	fmt.Println("running helm install")

}

func spawnTemplating(cmd *cobra.Command, args []string) {
	args = append([]string{"template"}, args...)

	fmt.Println("Executing helm template")

	execCmd := exec.Command("helm", args...)
	resp, err := execCmd.Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Printf("%s", resp)
}
