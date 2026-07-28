## Usage

Sources:

- [Official upstream README](https://github.com/constructive-io/supabase-test-suite/blob/44ceff6be0b62e1d1a60524acc0c65c1b8343726/packages/hello-world/README.md)
- [Official extension control file (hello-world.control)](https://github.com/constructive-io/supabase-test-suite/blob/44ceff6be0b62e1d1a60524acc0c65c1b8343726/packages/hello-world/hello-world.control)
- [Official extension SQL (hello-world--0.0.1.sql)](https://github.com/constructive-io/supabase-test-suite/blob/44ceff6be0b62e1d1a60524acc0c65c1b8343726/packages/hello-world/sql/hello-world--0.0.1.sql)

`hello-world` — **🛠 Built by the Constructive team — creators of modular Postgres tooling for secure, composable backends. If you like our work, contribute on GitHub.**. Use it when implementing the corresponding security, audit, or access-control workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION "hello-world";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `rls_test.update_updated_at_column()` is an extension function and returns `trigger`.
- `rls_test.pets` is a table installed or managed by the extension.
- `rls_test` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- Install the confirmed extension dependencies first: `plpgsql`, `pgcrypto`, `supabase`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
