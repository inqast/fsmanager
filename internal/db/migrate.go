package db

import (
	"log"

	"github.com/inqast/fsmanager/internal/config"
	"github.com/pressly/goose/v3"

	_ "github.com/lib/pq"
)

func GooseExec(cfg *config.Config, args []string) {
	dir := cfg.Database.Migrations
	connString := cfg.Database.GetConnString()
	command := args[0]

	db, err := goose.OpenDBWithDriver("postgres", connString)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := make([]string, 0)
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	if err := goose.Run(command, db, dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}

func Migrate(cfg *config.Config) {
	GooseExec(cfg, []string{"up"})
}
