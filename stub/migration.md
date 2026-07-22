## Usage

Sources:

- [Official README at the catalog revision](https://github.com/drewc/postgresql-migrations/blob/3da0fe400bfa37f92a78e8854f08522f3989079d/README.org)
- [Extension SQL at the catalog revision](https://github.com/drewc/postgresql-migrations/blob/3da0fe400bfa37f92a78e8854f08522f3989079d/migration--0.0.1.sql)

`migration` 0.0.1 is a database-resident migration runner. It registers numbered up/down SQL files, reads those files from the PostgreSQL server filesystem, executes them dynamically, and records status. It is an old administrative tool, not a safe migration service for untrusted users.

### Core Workflow

Place reviewed SQL files on the database server, choose a stable root directory, register them in order, and use one administrative session to migrate.

```sql
CREATE EXTENSION migration;

SELECT migration_set_env('ROOT_DIRECTORY', '/srv/db-migrations');
SELECT register_migration(1, '001_create.up.sql', '001_create.down.sql');
SELECT register_migration(2, '002_index.up.sql',  '002_index.down.sql');

-- Apply registered migrations through version 2.
SELECT migration_migrate_database(2);

-- Inspect the registry and recorded errors.
TABLE migration_registry;
TABLE migration_error;
```

### Object Index

- `migration_registry` stores the numeric migration, up/down file paths, and applied status.
- `migration_set_env(text, text)` and related environment helpers manage values such as the server-side root directory.
- `register_migration(...)` records file paths; it does not freeze or checksum their contents.
- `migration_run(...)` executes one direction, while `migration_migrate_database(numeric)` walks the registry toward a target version.

### Security and Operations

- File access occurs on the PostgreSQL server through server-side COPY, and file contents are executed as SQL. Restrict the extension and its functions to trusted administrators.
- Keep registered files immutable, reviewed, and backed up. The registry stores paths rather than content hashes, so replacing a file changes what a later run executes.
- Run only one migrator at a time. The implementation provides no advisory-lock protocol for concurrent runners.
- Test both up and down paths on a disposable copy and take a database backup. Error rows help diagnosis but are not a substitute for transactional review or recovery planning.
