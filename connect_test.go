package digital_ocean_db_connect_test

import (
	"log"
	"testing"

	"github.com/dragonzurfer/digital_ocean_db_connect"
)

func TestShowTables(t *testing.T) {
	db := digital_ocean_db_connect.Connect()
	if db == nil {
		t.Fatal("Failed to connect to the database")
	}
	defer db.Close()

	rows, err := db.Raw("SHOW TABLES").Rows()
	if err != nil {
		t.Fatalf("Failed to execute SHOW TABLES: %s", err)
	}
	defer rows.Close()

	log.Println("Tables in the database:")
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			t.Fatalf("Failed to scan table name: %s", err)
		}
		log.Println(tableName)
	}
}
