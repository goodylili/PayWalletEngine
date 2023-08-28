package db

import "log"

func (d *Database) MigrateDB() error {
	log.Println("Database Migration in Process...")

	// Use GORM AutoMigrate to migrate all the database schemas.
	err := d.Client.AutoMigrate(
	// List of database models here.
	)
	if err != nil {
		return err
	}

	log.Println("Database Migration Complete!")
	return nil
}
