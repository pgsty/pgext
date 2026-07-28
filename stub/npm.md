## Usage

Sources:

- [Official upstream README](https://github.com/constructive-io/lib-count/blob/c04e5084d5fbdee6c452976d9e69532c85dcbe96/packages/npm/README.md)
- [Official extension control file (npm.control)](https://github.com/constructive-io/lib-count/blob/c04e5084d5fbdee6c452976d9e69532c85dcbe96/packages/npm/npm.control)
- [Official extension SQL (npm--0.0.1.sql)](https://github.com/constructive-io/lib-count/blob/c04e5084d5fbdee6c452976d9e69532c85dcbe96/packages/npm/sql/npm--0.0.1.sql)

`npm` — Schema and tables for tracking npm package metadata and daily download counts. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION npm;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `npm_count.set_id_from_pkg_date()` is an extension function and returns `trigger`.
- `npm_count.update_updated_at()` is an extension function and returns `trigger`.
- `npm_count.validate_download_date()` is an extension function and returns `trigger`.
- `npm_count.missing_download_dates` is an extension-defined view.
- `npm_count.category` is a table installed or managed by the extension.
- `npm_count.daily_downloads` is a table installed or managed by the extension.
- `npm_count.npm_package` is a table installed or managed by the extension.
- `npm_count.package_category` is a table installed or managed by the extension.
- `npm_count` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- Install the confirmed extension dependencies first: `btree_gist`, `plpgsql`, `uuid-ossp`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
