package cmd

import (
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
		log.Fatal("destroy command is not implemented yet")
		/*
			var err error
			if specFile == "" {
				specFile = filepath.Join(specPath, defaultSpecFileName)
			}
			if !fileExists(specFile) {
				fmt.Printf("cluster spec file doesnt exist: %s\n", specFile)
				os.Exit(1)
			}
			specContents, err = ioutil.ReadFile(specFile)
			if err != nil {
				log.Fatalf("error reading config file (%s)", specFile)
			}
		*/
	},
}

func init() {
	rootCmd.AddCommand(destroyCmd)
	destroyCmd.PersistentFlags().StringVarP(&specFile, "spec-file", "f", "", "Location of cluster-spec file corresponding to the cluster, default is at ~/.cake/<cluster name>/spec.yaml")
}
