## Usage

Sources:

- [Official upstream README](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/readme.txt)
- [Official extension control file (pgrao.control)](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/vagrant/postgres17/pgrao/pgrao.control)
- [Official extension SQL (pgrao--1.0.0.sql)](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/vagrant/postgres17/pgrao/pgrao--1.0.0.sql)

`pgrao` — PGRAO Postgres DB Catalog Extension. Use it when administering or automating the database behavior described above. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pgrao;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `rao_database` is an extension-defined view.
- `rao_domain` is an extension-defined view.
- `rao_event_trigger` is an extension-defined view.
- `rao_extension` is an extension-defined view.
- `rao_fdw` is an extension-defined view.
- `rao_index` is an extension-defined view.
- `rao_language` is an extension-defined view.
- `rao_mview` is an extension-defined view.
- `rao_part_tables` is an extension-defined view.
- `rao_role` is an extension-defined view.
- `rao_routine` is an extension-defined view.
- `rao_schema` is an extension-defined view.
- `rao_sequence` is an extension-defined view.
- `rao_table` is an extension-defined view.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- Install the confirmed extension dependencies first: `plpgsql`.
- The control file marks the extension as relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
