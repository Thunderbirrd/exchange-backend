package postgres

import (
	"database/sql"
	"embed"
	"log"

	sqlmigrate "github.com/rubenv/sql-migrate"
)

//go:embed migrations/*.sql
var migrationsFolder embed.FS

// MigrateUp start the migrations
func MigrateUp(dbDriver string, dbString string) error {
	migrations := &sqlmigrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationsFolder,
		Root:       "migrations",
	}

	dbMigrate, err := sql.Open(dbDriver, dbString)
	if err != nil {
		return err
	}

	n, err := sqlmigrate.Exec(dbMigrate, dbDriver, migrations, sqlmigrate.Up)
	if err != nil {
		return err
	}

	log.Printf("Applied %d migrations!\n", n)
	return nil
}

func MigrateDown(dbDriver string, dbString string) error {
	migrations := &sqlmigrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationsFolder,
		Root:       ".",
	}

	dbMigrate, err := sql.Open(dbDriver, dbString)
	if err != nil {
		return err
	}

	n, err := sqlmigrate.Exec(dbMigrate, dbDriver, migrations, sqlmigrate.Down)
	if err != nil {
		return err
	}

	log.Printf("Down %d migrations!\n", n)
	return nil
}
