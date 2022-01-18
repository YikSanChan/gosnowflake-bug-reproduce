package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/snowflakedb/gosnowflake"
	sf "github.com/snowflakedb/gosnowflake"
)

// getDSN constructs a DSN based on the test connection parameters
func getDSN() (string, *sf.Config, error) {
	cfg := &sf.Config{
		Account:  os.Getenv("SNOWFLAKE_TEST_ACCOUNT"),
		User:     os.Getenv("SNOWFLAKE_TEST_USER"),
		Password: os.Getenv("SNOWFLAKE_TEST_PASSWORD"),
		Database: "snowflake_sample_data",
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
	query := "SELECT * FROM tpch_sf1.customer limit 1"
	rows, err := db.QueryContext(gosnowflake.WithHigherPrecision(context.Background()), query) // no cancel is allowed
	if err != nil {
		log.Fatalf("failed to run a query. %v, err: %v", query, err)
	}
	defer rows.Close()
	for rows.Next() {
		var custkey, name, address, nationkey, phone, acctbal, mktsegment, comment interface{}
		err = rows.Scan(&custkey, &name, &address, &nationkey, &phone, &acctbal, &mktsegment, &comment)
		if err != nil {
			log.Fatalf("failed to get result. err: %v", err)
		}
		// 60001 int64
		// Customer#000060001 string
		// 9Ii4zQn9cX string
		// 14 int64
		// 24-678-784-9652 string
		// +9957.56 *big.Float
		// HOUSEHOLD string
		// l theodolites boost slyly at the platelets: permanently ironic packages wake slyly pend string
		fmt.Printf("%+v %T\n", custkey, custkey)
		fmt.Printf("%+v %T\n", name, name)
		fmt.Printf("%+v %T\n", address, address)
		fmt.Printf("%+v %T\n", nationkey, nationkey)
		fmt.Printf("%+v %T\n", phone, phone)
		fmt.Printf("%+v %T\n", acctbal, acctbal)
		fmt.Printf("%+v %T\n", mktsegment, mktsegment)
		fmt.Printf("%+v %T\n", comment, comment)
	}
	if rows.Err() != nil {
		fmt.Printf("ERROR: %v\n", rows.Err())
		return
	}
}
