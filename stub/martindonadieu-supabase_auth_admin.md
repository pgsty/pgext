## Usage

Sources:

- [Official database.dev package page](https://database.dev/martindonadieu/supabase_auth_admin)

`martindonadieu-supabase_auth_admin` — Supabase Auth tool to check if a user is an platform admin. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "martindonadieu-supabase_auth_admin";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `public"."is_admin` is an extension function.
- `public"."set_admin` is an extension function.
- `public.is_admin()` is an extension function and returns `boolean`.

### Requirements and Caveats

- The catalog records version `0.0.1`.
- This is a database.dev/pg_tle package; register or generate its package migration before issuing the quoted `CREATE EXTENSION` identity.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
