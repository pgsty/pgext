## Usage

Sources:

- [Official upstream README](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/extensions/filesystem/README.md)
- [Official extension control file (filesystem.control)](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/extensions/filesystem/filesystem.control)

`filesystem` — Filesystem Foreign Data Wrapper =============================== Exposes the filesystem to PostgreSQL, allowing files and directories to be read via SQL commands. Use it when PostgreSQL must access the corresponding external data source through a foreign-data interface. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION filesystem;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.4.0`.
- Install the confirmed extension dependencies first: `meta`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
