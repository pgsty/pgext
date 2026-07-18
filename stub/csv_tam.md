## Usage

Sources:

- [Official extension control file](https://github.com/InnerLife0/csv_tam/blob/5fe49b6d58c2688ae15a824069a719a95e5dd2c8/csv_tam.control)

`csv_tam` — CSV table access method for PostgreSQL that stores table tuples in CSV-formatted relation files.

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "csv_tam";
SELECT extversion
FROM pg_extension
WHERE extname = 'csv_tam';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
