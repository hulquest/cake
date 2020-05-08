package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type settings struct {
	config                   string
	vCenterURL               string
	vCenterUser              string
	vCenterPassword          string
	managementClusterPodCIDR string
	managementClusterCIDR    string
	disableCleanup           bool
	disablePreflight         bool
	logLevel                 string
}

var (
	cliSettings settings
	envSettings settings
	cfgFile     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cake",
	Short: "Deploy a kubernetes management cluster on vSphere",
	Long: `Cake provides a mechanism to deploy either a CAPv or Rancher
	Management Cluster on your vSphere easily with a single go binary.
	For example: "cake deploy -d rke --config my-cluster-config.yaml"
	Will create VMs and install RKE and Rancher Server`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Errorf("error executing command: %v", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix("cake")
}
