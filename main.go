package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/librefitness/librefitness/internal/api"
	"github.com/librefitness/librefitness/internal/config"
	"github.com/librefitness/librefitness/internal/database"
)

func main() {
	cmd := &cobra.Command{
		Use:   "librefit",
		Short: "librefit is calorie counter and diet planner to help people achieve their goals",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(startCmd)

	if err := cmd.Execute(); err != nil {
		fmt.Printf("Error during command execution: %v", err)
		os.Exit(0)
	}
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: `Start starts the server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		conf := config.NewConfig()

		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
		cmd.PersistentFlags().StringVarP(&conf.Opts.DatabaseDSN, "db-host", "", "", "DSN string")

		cmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
			if viper.IsSet(f.Name) && viper.GetString(f.Name) != "" {
				cmd.PersistentFlags().Set(f.Name, viper.GetString(f.Name))
			}
		})

		// Initilize connection to the DB
		db, err := database.Init()
		if err != nil {
			return err
		}

		// Run migrations
		database.Migrate()

		// Creating adminUser
		err = database.InitData()
		if err != nil {
			return err
		}

		// Closing connection to the DB: https://gorm.io/docs/generic_interface.html
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}

		defer sqlDB.Close()

		api, err := api.NewHTTPServer(conf)
		if err != nil {
			return err
		}

		api.Run()

		return nil
	},
}
