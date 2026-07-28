## Usage

Sources:

- [Official database.dev package page](https://database.dev/jessevent/supa_profile)

`jessevent@supa_profile` — Database table and column statistical profiler. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "jessevent@supa_profile";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `infer_pattern` is an extension function.
- `profile_table` is an extension function.

### Requirements and Caveats

- The catalog records version `1.0.0`.
- This is a database.dev/pg_tle package; register or generate its package migration before issuing the quoted `CREATE EXTENSION` identity.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
