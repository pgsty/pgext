## Usage

Sources:

- [Official upstream README](https://github.com/aadisanghani/pg_query_rules/blob/7c65bcd58c122cbf5b7b2b130d4cc391432c7ca4/README.md)
- [Official extension control file (pg_query_rules.control)](https://github.com/aadisanghani/pg_query_rules/blob/7c65bcd58c122cbf5b7b2b130d4cc391432c7ca4/pg_query_rules.control)
- [Official extension SQL (pg_query_rules--0.1.0.sql)](https://github.com/aadisanghani/pg_query_rules/blob/7c65bcd58c122cbf5b7b2b130d4cc391432c7ca4/pg_query_rules--0.1.0.sql)

`pg_query_rules` — pg_query_rules is a PostgreSQL C extension that rewrites or blocks SQL queries at runtime using regex-based rules. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_query_rules;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `load_query_rules_to_runtime()` is an extension function and returns `TEXT`.
- `pg_query_rules_updated_at()` is an extension function and returns `TRIGGER`.
- `pgqr_test(sql_text TEXT)` is an extension function and returns `TEXT`.
- `pgqr_version()` is an extension function and returns `TEXT`.
- `pg_query_rules` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
