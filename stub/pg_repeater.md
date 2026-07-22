## Usage

Sources:

- [Official README](https://github.com/danolivo/pg_repeater/blob/master/README.md)
- [Extension control file](https://github.com/danolivo/pg_repeater/blob/master/pg_repeater.control)
- [Hook implementation](https://github.com/danolivo/pg_repeater/blob/master/pg_repeater.c)

`pg_repeater` 0.1 is a PostgreSQL patch-and-extension prototype that forwards utility commands and serialized executor plans to a remote server. It depends on postgres_fdw and pg_execplan plus compatible patched source trees. It is research code, not a replication or distributed-transaction system.

### Core Workflow

The local and remote servers need matching object layouts and the companion execution support. Configure a postgres_fdw server named by the extension GUC, then preload the hook library and restart PostgreSQL.

```conf
shared_preload_libraries = 'pg_repeater'
repeater.fdwname = 'remoteserv'
```

```sql
CREATE EXTENSION postgres_fdw;
CREATE EXTENSION pg_execplan;
CREATE EXTENSION pg_repeater;

CREATE SERVER remoteserv
FOREIGN DATA WRAPPER postgres_fdw
OPTIONS (host '/tmp', dbname 'remote_compute');

CREATE USER MAPPING FOR CURRENT_USER
SERVER remoteserv
OPTIONS (user 'remote_executor');
```

### Execution Model

- ProcessUtility and executor hooks intercept eligible statements. Plans are serialized, object OIDs are converted to a portable representation, and the payload is sent through the configured postgres_fdw connection.
- The remote side executes the supplied plan through pg_exec_plan. Both nodes therefore need compatible PostgreSQL internals, schemas, functions, types, and data assumptions.
- The implementation excludes several utility classes, including COPY, CREATE EXTENSION, EXPLAIN, FETCH, and VACUUM.

### Critical Caveats

- A session LOAD can install hooks only in that session; preloading is required for consistent server-wide interception, and changing the preload list requires restart.
- There is no distributed commit protocol or reliable rollback coupling between local and remote execution. Errors, network loss, or topology drift can leave nodes divergent.
- Serialized PostgreSQL plan trees are private, version-sensitive internals. Never mix server builds or extension ABIs, and do not use this prototype as a production availability or consistency mechanism.
