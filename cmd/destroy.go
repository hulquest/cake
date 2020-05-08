package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// destroyCmd represents the destroy command
var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy a previously deploy Cake install",
	Long: `This doesn't exist yet, but when it does, you'll
	use it the same as you do deploy, but obviously instead
	of setting up a deployment, it's going to destroy it.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Error("destroy command is not implemented yet")
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(destroyCmd)

	destroyCmd.PersistentFlags().StringVarP(&specPath, "spec-path", "p", "", "Location of cluster-spec directory cooresponding to the cluster to be destroyed, default is created at ~/.cake ")
	destroyCmd.PersistentFlags().StringVarP(&clusterName, "name", "n", "", "Name of the cluster to destroy, if specified without the -p option, will look for a spec.yml file in ~/.cake/<cluster-name>/")
	destroyCmd.MarkFlagRequired("name")
}
