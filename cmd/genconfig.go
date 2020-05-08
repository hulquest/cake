package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/dustinkirkland/golang-petname"
	"github.com/gookit/color"
	"github.com/netapp/cake/pkg/config/types"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	red         = color.New(color.FgRed)
	blue        = color.New(color.FgBlue)
	green       = color.New(color.FgGreen)
	yellow      = color.New(color.FgYellow)
	specPath    string
	clusterSpec string
	clusterName string
)

// genconfigCmd represents the genconfig command
var genconfigCmd = &cobra.Command{
	Use:   "genconfig",
	Short: "Create a cluster-spec via command line prompts",
	Long: `genconfig provides an option to interactively create a cluster-spec file.  By default the 
	spec file is placed in ~/.cake/<cluster-name>/spec.yml.  If no name is provides an auto generated
	name will be used.
  For example:
	    "cake genconfig -n my-demo"
  Will take interactive input from the user and create the file "~/.cake/my-demo/spe.yml"`,
	Run: func(cmd *cobra.Command, args []string) {
		runEasyConfig()
	},
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	genconfigCmd.PersistentFlags().StringVarP(&specPath, "spec-path", "p", "", "Root directory to create cluster files in, default is ~/.cake")
	genconfigCmd.PersistentFlags().StringVarP(&clusterName, "name", "n", "", "Name to assign to the cluster being created, if omitted a random name will be generated.")
	cobra.OnInitialize(initSpecFile)
	rootCmd.AddCommand(genconfigCmd)
}

func initSpecFile() {
	if clusterName == "" {
		clusterName = petname.Generate(2, "-")
		fmt.Printf("generated a cluster name:  %s\n", clusterName)
	}
	initSpecDir()

}
func fileExists(fn string) bool {
	_, err := os.Stat(fn)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func initSpecDir() {
	if specPath == "" {
		basePath := cakeBaseDirPath()
		specPath = fmt.Sprintf("%s/%s", basePath, clusterName)
	}

	if _, err := os.Stat(specPath); os.IsNotExist(err) {
		fmt.Printf("creating config directory: %s\n", specPath)
		os.MkdirAll(specPath, 0700)
	}
	clusterSpec = fmt.Sprintf("%s/spec.yaml", specPath)
	if fileExists(clusterSpec) {
		fmt.Println("A cluster spec file already exists for the cluster-name specified, please use another name or delete the existing spec.yml file")
		os.Exit(1)
	}

}

func runEasyConfig() {
	fmt.Printf("creating spec file based on user input: %s\n", clusterSpec)
	var spec = &types.ConfigSpec{}
	configure(spec)
	writeSpec(spec)
}

func cakeBaseDirPath() string {
	if runtime.GOOS == "windows" {
		return fmt.Sprintf("%s/.cake", os.Getenv("USERPROFILE"))
	}

	return fmt.Sprintf("%s/.cake", os.Getenv("HOME"))
}

func writeSpec(spec *types.ConfigSpec) {
	var configOut []byte
	var err error

	writeSpec := *spec

	if configOut, err = yaml.Marshal(writeSpec); err != nil {
		log.Fatalln(err)
	}

	err = writeFile(clusterSpec, configOut, 0644)
	if err != nil {
		log.Println(fmt.Sprintf("Unable to save cluster spec file (%s), %s", clusterSpec, err.Error()))
		return
	}
}

func writeFile(specFile string, contents []byte, permissionCode os.FileMode) error {
	if permissionCode == 0 {
		permissionCode = 0644
	}

	if err := ioutil.WriteFile(specFile, contents, permissionCode); err != nil {
		return fmt.Errorf("unable to write config file, %v", err)
	}

	return nil
}

func createConfigDirectory(directoryName string) error {
	basePath := cakeBaseDirPath()
	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		err = os.Mkdir(basePath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	fullPath := fmt.Sprintf("%s/%s", basePath, directoryName)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		err = os.Mkdir(fullPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func configure(spec *types.ConfigSpec) {
	// fail fast if we can't connect to specified vSphere
	if err := collectVsphereInformation(spec); err != nil {
		log.Fatalln(err)
	}

	collectNetworkInformation(spec)
	collectAdditionalConfiguration(spec)
	writeSpec(spec)
}
