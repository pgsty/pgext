## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/session_hash_tools/session_hash_tools-1.0.0/README)

`session_hash_tools` — Start the $dbname createlang plperl $dbname psql -c "CREATE SCHEMA tools" $dbname psql -f $PGDATA/share/contrib/tools.sql $dbname. Use it when SQL needs these specialized functions or aggregates. Its extension dependencies must be installed and validated first.

### Core Workflow

This component has no confirmed standalone `CREATE EXTENSION` workflow in the reviewed source. Build, load, or enable it only through the exact upstream mechanism, then verify the resulting server behavior in an isolated database.

### Requirements and Caveats

- The catalog records version `1.0.0`.
- Install the confirmed extension dependencies first: `plperl`, `plperlu`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
