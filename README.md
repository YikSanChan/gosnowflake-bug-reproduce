FIX: use WithHigherPrecision context to ensure Query returns natural type rather than string.

See https://github.com/snowflakedb/gosnowflake/issues/517#issuecomment-1014967682

---

Reproducible code for [gosnowflake issue](https://github.com/snowflakedb/gosnowflake/issues/517).

Run `go build .` to build. Then run `./gosnowflake-bug-reproduce` to reproduce the bug.

Note: make sure env vars `SNOWFLAKE_TEST_ACCOUNT`, `SNOWFLAKE_TEST_USER` and `SNOWFLAKE_TEST_PASSWORD` are configured correctly.