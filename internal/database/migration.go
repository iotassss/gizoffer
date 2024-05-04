package database

import "log"

func (db *GizofferDB) Migrate() {
	if err := db.AutoMigrate(models...); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
}

func (db *GizofferDB) Refresh() {
	for _, model := range models {
		if err := db.Migrator().DropTable(model); err != nil {
			log.Fatalf("failed to drop table: %v", err)
		}
	}
	db.Migrate()
}
