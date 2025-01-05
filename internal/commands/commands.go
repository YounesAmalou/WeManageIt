package commands

import (
	"fmt"
	"net/http"
	"strings"

	"WeManageIt/internal/config"
	"WeManageIt/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const mainDescription = `WeManageIt is a project management
system written in Golang & React by DevExpert Team.`

func VersionTemplate() string {
	info := utils.Version()
	return fmt.Sprintf(`Version: %s
Commit: %s
Built: %s`, info.Version, info.CommitHash, info.BuildDate+"\n")
}

var cfgFile string

var rootCmd = &cobra.Command{
	Use:     "wemanageit",
	Long:    mainDescription,
	Version: VersionTemplate(),
}

var migration http.FileSystem

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file path")
	migration = http.Dir("./migrations")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath("./conf")
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc/wemanageit")
		viper.SetConfigName("wemanageit")
	}

	viper.SetEnvPrefix("WEMANAGEIT")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	config.InitDefaults()

	err := viper.ReadInConfig()
	if err == nil {
		return
	}
	if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		panic(err)
	}

}

// Execute the root cobra command
func Execute() {
	rootCmd.SetVersionTemplate(VersionTemplate())
	rootCmd.AddCommand(newJobCmd(), newTokenCmd(), newWebCmd(), newMigrateCmd(), newWorkerCmd(), newResetPasswordCmd(), newSeedCmd())
	rootCmd.Execute()
}
