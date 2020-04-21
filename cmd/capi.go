package cmd

import (
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
)
var (
	provider string
)

var capiCmd = &cobra.Command{
	Use:   "capi",
	Short: "CAPI deploys a Cluster API (CAPI) cluster using a specific provider",
	Long:  `CAPI deploy will create an upstream CAPI management cluster, the --provider option will deploy CAPI using a specific provider`,
	Run: func(cmd *cobra.Command, args []string) {
		runCAPI()
	},
}

func init() {
	capiCmd.Flags().StringVarP(&provider, "provider", "p", "", "The provider to use (vsphere, kubevirt")
	capiCmd.MarkFlagRequired("provider")
	deployCmd.AddCommand(capiCmd)
}

func runCAPI() {
	log.Info("running CAPI engine with Provider: %v\n", provider)

}