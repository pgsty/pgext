## Usage

Sources:

- [Official upstream README](https://github.com/dimitri/base36/blob/d1bc4175148c80b125f566c3cc577383687a6162/README.md)
- [Official extension control file (base36_gist.control)](https://github.com/dimitri/base36/blob/d1bc4175148c80b125f566c3cc577383687a6162/base36_gist.control)
- [Official extension SQL (base36_gist--1.0.sql)](https://github.com/dimitri/base36/blob/d1bc4175148c80b125f566c3cc577383687a6162/base36_gist--1.0.sql)

`base36_gist` — This data type is a quick hack and only has *demo* quality, its shy of a brick load of being *production ready*. If you need it for serious work, please consider opening an issue against this project, potentially with a patch attached (they call that a *pull request* around here). Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION base36_gist;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `base36_dist(base36, base36)` is an extension function and returns `base36`.
- `gist_base36_ops` is an extension-defined operator class.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- Install the confirmed extension dependencies first: `btree_gist`, `base36`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
