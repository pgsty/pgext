## Usage

Sources:

- [Official C implementation and documentation](https://github.com/fdr/pg_connlimit/blob/5ee6221506df3cb572b5b915fe834abbad18b60a/pg_connlimit.c)
- [Official control file](https://github.com/fdr/pg_connlimit/blob/5ee6221506df3cb572b5b915fe834abbad18b60a/pg_connlimit.control)

`pg_connlimit` 1.0 is a headless preload module that limits concurrent connections per role using small files outside PostgreSQL catalogs. It was designed so primary and hot-standby servers can enforce different quotas even when their database catalogs are replicated.

### Configuration

There is no extension SQL script and no `CREATE EXTENSION` workflow. Install the shared library, create a protected directory containing one file per role, then preload the module and set the directory path.

```conf
shared_preload_libraries = 'pg_connlimit'
connlimit.directory = '/etc/postgresql/connlimit-db'
```

```text
/etc/postgresql/connlimit-db/
├── app_reader    # file content: 30
└── report_user   # file content: 5
```

Restart after changing `shared_preload_libraries`. The authentication hook reads the matching role file for every successful login and rejects the new connection when existing backends for that role are already at the integer limit.

### Exact Enforcement Rules

- Limits are cluster-wide per PostgreSQL role, not per database.
- Only role names made entirely of lowercase letters, digits, and underscore are eligible; other names are skipped to prevent path traversal.
- Missing/unreadable files, invalid integers, unknown roles, or an unset directory all fail open and do not enforce a limit.
- File changes are read by later connection attempts; changing `connlimit.directory` is a SIGHUP-level configuration change.
- The current connection is not yet counted by `CountUserBackends`, so the hook rejects when the existing count is greater than or equal to the configured value.

### Security and Operational Boundaries

Protect the directory and files from database users and unrelated operating-system accounts. A user who can remove, replace, or corrupt a file can disable enforcement; a non-positive integer can block every new connection for that role. Use atomic file replacement and validate contents before rollout.

The module intentionally fails open and is not a security boundary by itself. Keep `max_connections`, reserved administrative access, pool limits, and catalog-level `CONNECTION LIMIT` controls as appropriate. The source dates from 2013 and uses authentication/procarray hooks that require porting and load testing on modern PostgreSQL.
