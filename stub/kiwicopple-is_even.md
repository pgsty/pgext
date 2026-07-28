## Usage

Sources:

- [Official database.dev package page](https://database.dev/kiwicopple/is_even)

`kiwicopple-is_even` — Calculates if a value is even. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "kiwicopple-is_even";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `is_even(value INT)` is an extension function and returns `BOOLEAN`.

### Requirements and Caveats

- The catalog records version `0.0.1`.
- This is a database.dev/pg_tle package; register or generate its package migration before issuing the quoted `CREATE EXTENSION` identity.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
