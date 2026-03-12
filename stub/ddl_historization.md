


## Usage

> [ddl_historization: Track all DDL changes in a PostgreSQL database](https://github.com/rodo/pg_ddl_historization)

Records all DDL changes (CREATE, ALTER, DROP, etc.) made on a database into a historization table for auditing and tracking purposes.

### Setup

```sql
CREATE EXTENSION ddl_historization;
```

The extension installs event triggers that automatically capture DDL statements and store them in the historization table.

### Querying DDL History

After installation, all DDL changes are logged automatically. Query the history table to see what changes have been made:

```sql
SELECT * FROM ddl_history ORDER BY ddl_date DESC;
```

### Integration with pg_tle

For AWS RDS environments, the extension can be deployed via `pg_tle`:

```sql
-- Build the pg_tle deployment file
-- $ make pgtle
-- Then execute pgtle.ddl_historization-0.3.sql on your instance
```

### Notes

- DDL statements are captured via PostgreSQL event triggers
- Works with `CREATE`, `ALTER`, `DROP`, and other DDL commands
- Used as a dependency by the `schedoc` extension for schema documentation
