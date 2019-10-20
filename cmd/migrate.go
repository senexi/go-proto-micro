package cmd

import (
	"fmt"

	"github.com/go-pg/migrations/v7"
	"github.com/go-pg/pg/v9"
	_ "github.com/senexi/go-proto-micro/migrations"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long: `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.`,
	Run: func(cmd *cobra.Command, args []string) {
		migrate(args)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func migrate(args []string) {
	database := viper.GetString("database.name")
	databasePort := viper.GetString("database.port")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	databaseURL := fmt.Sprintf("%s:%s", viper.GetString("database.url"), databasePort)

	db := pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
		Addr:     databaseURL,
	})
	_, err := db.Exec("SELECT 1")
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Panic("database down")
	}

	log.WithFields(log.Fields{"db": database}).Info("connected to database")
	defer db.Close()

	oldVersion, newVersion, err := migrations.Run(db, args...)
	if err != nil {
		log.Error(err)
	}
	if newVersion != oldVersion {
		log.WithFields(log.Fields{"old version": oldVersion, "new version": newVersion}).Info("migrated")
	} else {
		fmt.Printf("version is %d\n", oldVersion)
		log.WithFields(log.Fields{"current version": oldVersion}).Info("nothing done")
	}
}
