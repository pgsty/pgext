## Usage

Sources:

- [Official upstream README](https://github.com/supabase/pg_netstat/blob/debce79df4737c5b5f5ea3826b3b4145df5c5c81/README.md)
- [Official extension control file (pg_netstat.control)](https://github.com/supabase/pg_netstat/blob/debce79df4737c5b5f5ea3826b3b4145df5c5c81/pg_netstat.control)
- [Official implementation source](https://github.com/supabase/pg_netstat/blob/debce79df4737c5b5f5ea3826b3b4145df5c5c81/src/lib.rs)

`pg_netstat` — pg_netstat monitors your PostgreSQL database network traffic. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_netstat;

select * from pg_netstat;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `netstat()` is an extension function.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
