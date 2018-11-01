package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

func main() {
	filePath := "/home/stephendwolff/foo"

	// register file with MySQL so that it can be used with LOAD DATA LOCAL INFILE
	mysql.RegisterLocalFile(filePath)

	// prepare the database connection - it opens lazily when first used
	db, dberr := sql.Open("mysql", "user:pass@/database_name")

	if dberr != nil {
		log.Println( fmt.Errorf("wnu_csv: dberr: %s", dberr))
	} else {
		// CSV file first line:
		// "id","display_name","id","created_at"
		_, err := db.Exec("LOAD DATA LOCAL INFILE '" + filePath + "' INTO TABLE classifications " +
			"FIELDS TERMINATED BY ',' " +
			"ENCLOSED BY '\"' " +
			"IGNORE 1 LINES " +
			"(id, display_name, user_id, created_at);")
		if err != nil {
			log.Println( fmt.Errorf("wnu_csv: load_data err: %s", err))
		} else {
			log.Println( "Success" )
		}
	}

	// be a good citizen
	db.Close()
}