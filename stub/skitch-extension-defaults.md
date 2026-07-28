## Usage

Sources:

- [Official upstream README](https://github.com/airpage-app/pg-utils/blob/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/readme.md)
- [Official extension control file (skitch-extension-defaults.control)](https://github.com/airpage-app/pg-utils/blob/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/packages/defaults/skitch-extension-defaults.control)
- [Official extension SQL (skitch-extension-defaults--0.0.7.sql)](https://github.com/airpage-app/pg-utils/blob/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/packages/defaults/sql/skitch-extension-defaults--0.0.7.sql)

`skitch-extension-defaults` — default roles. Use it when implementing the corresponding security, audit, or access-control workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION "skitch-extension-defaults";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.7`.
- Install the confirmed extension dependencies first: `plpgsql`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
