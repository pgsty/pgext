## Usage

Sources:

- [Official dbdev package page](https://database.dev/supabase/dbdev)
- [Official package-installation documentation](https://github.com/supabase/dbdev/blob/8b3b966d97f87027a36c052e679ebb45f7e25736/website/content/docs/install-a-package.mdx)
- [Official project README](https://github.com/supabase/dbdev/blob/8b3b966d97f87027a36c052e679ebb45f7e25736/README.md)

`dbdev` is the in-database client for the database.dev registry of PostgreSQL Trusted Language Extensions. The catalog entry corresponds to the registry package `supabase@dbdev`; it downloads package SQL, registers it through `pg_tle`, and exposes installed packages as PostgreSQL extensions.

### Core Workflow

First install `pg_tle` and `http`, then follow the bootstrap SQL on the official package page to register and create `supabase@dbdev`. After that, install a versioned package and enable it as an extension:

```sql
CREATE EXTENSION IF NOT EXISTS http WITH SCHEMA extensions;
CREATE EXTENSION IF NOT EXISTS pg_tle;

-- Run the current bootstrap block from the official supabase@dbdev page.

SELECT dbdev.install('owner@package');
CREATE EXTENSION "owner@package"
    SCHEMA public
    VERSION '1.2.3';
```

The principal SQL API is `dbdev.install(package_name text)`. It fetches available versions and upgrade paths from the registry, calls the `pg_tle` installation functions, and registers a download event. Pin the extension version in `CREATE EXTENSION` for reproducibility.

### Operational Notes

Version `0.0.5` requires the `http` and `pg_tle` extensions plus outbound HTTPS access to database.dev. Installing a package executes SQL obtained from the remote registry inside the database, so review the selected package and version before installation and control who may call `dbdev.install`. The official documentation warns that logical restore of databases containing TLEs can fail and recommends physical backups for this workflow.
