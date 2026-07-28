## Usage

Sources:

- [Official upstream README](https://github.com/dresdres/bgp240-master-thesis/blob/2a96f87df76b8db3d802824cd7e3c0bbff321036/Extensions/README.md)
- [Official extension control file (seller_ext.control)](https://github.com/dresdres/bgp240-master-thesis/blob/2a96f87df76b8db3d802824cd7e3c0bbff321036/Extensions/seller_ext/seller_ext.control)
- [Official extension SQL (seller_ext--0.0.0.sql)](https://github.com/dresdres/bgp240-master-thesis/blob/2a96f87df76b8db3d802824cd7e3c0bbff321036/Extensions/CompiledExtensions/all_ext/controlSql/seller_ext--0.0.0.sql)

`seller_ext` — Seller-service LISTEN/NOTIFY background workers for a marketplace prototype. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION seller_ext;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
