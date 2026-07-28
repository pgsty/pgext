## Usage

Sources:

- [Official database.dev package page](https://database.dev/cem/uuidv7)

`cem-uuidv7` — UUIDv7 extension. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "cem-uuidv7";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `uuid7_from_timestamp(p_timestamp timestamp with time zone)` is an extension function and returns `uuid`.
- `uuid_generate_v7()` is an extension function and returns `uuid`.

### Requirements and Caveats

- The catalog records version `1.0.2`.
- This is a database.dev/pg_tle package; register or generate its package migration before issuing the quoted `CREATE EXTENSION` identity.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
