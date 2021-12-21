// Example: Fetch one row.
//
// No cancel is allowed as no context is specified in the method call Query(). If you want to capture Ctrl+C to cancel
// the query, specify the context and use QueryContext() instead. See selectmany for example.
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	sf "github.com/snowflakedb/gosnowflake"
)

// getDSN constructs a DSN based on the test connection parameters
func getDSN() (string, *sf.Config, error) {
	cfg := &sf.Config{
		Account:  os.Getenv("SNOWFLAKE_TEST_ACCOUNT"),
		User:     os.Getenv("SNOWFLAKE_TEST_USER"),
		Password: os.Getenv("SNOWFLAKE_TEST_PASSWORD"),
	}
	dsn, err := sf.DSN(cfg)
	return dsn, cfg, err
}

func main() {
	dsn, cfg, err := getDSN()
	if err != nil {
		log.Fatalf("failed to create DSN from Config: %v, err: %v", cfg, err)
	}

	db, err := sql.Open("snowflake", dsn)
	if err != nil {
		log.Fatalf("failed to connect. %v, err: %v", dsn, err)
	}
	defer db.Close()
	query := "SELECT 1"
	rows, err := db.Query(query) // no cancel is allowed
	if err != nil {
		log.Fatalf("failed to run a query. %v, err: %v", query, err)
	}
	defer rows.Close()
	var v int
	for rows.Next() {
		err := rows.Scan(&v)
		if err != nil {
			log.Fatalf("failed to get result. err: %v", err)
		}
		if v != 1 {
			log.Fatalf("failed to get 1. got: %v", v)
		}
	}
	if rows.Err() != nil {
		fmt.Printf("ERROR: %v\n", rows.Err())
		return
	}
	fmt.Printf("Congrats! You have successfully run %v with Snowflake DB!\n", query)
}
