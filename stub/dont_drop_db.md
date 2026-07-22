## Usage

Sources:

- [Official README](https://github.com/s-hironobu/dont_drop_db/blob/8610f0b4b37af5634135414614dad1b3080bc8de/README.md)
- [Official control file](https://github.com/s-hironobu/dont_drop_db/blob/8610f0b4b37af5634135414614dad1b3080bc8de/dont_drop_db.control)

`dont_drop_db` prevents `DROP DATABASE` for database names listed in a server configuration parameter. It is a shared-preload safety guard, not an authorization boundary or a substitute for backups: a privileged operator can change the setting, omit the module at restart, or remove data by other means.

### Core Workflow

Configure the module before starting PostgreSQL:

```ini
shared_preload_libraries = 'dont_drop_db'
dont_drop_db.list = 'postgres,template0,template1,production'
```

Restart the server, optionally register the extension object in an administrative database, and verify both protected and unprotected names in a disposable cluster:

```sql
CREATE EXTENSION dont_drop_db;

SHOW dont_drop_db.list;
```

The default list is `postgres,template0,template1`. The special value `all` blocks every database name, while a value containing only spaces disables protection.

### Configuration Semantics

`dont_drop_db.list` is a comma-separated database-name list evaluated by the preloaded module. Review exact spelling whenever databases are renamed or cloned. Because the behavior is loaded at server start, changing only the extension object in one database does not establish cluster-wide protection.

### Operational Notes

Treat configuration management and restart verification as part of the control. After deployment, confirm `shared_preload_libraries`, inspect the server log for module-load failures, and execute a negative test against a throwaway database name that appears in the list.

The guard covers `DROP DATABASE`, not table drops, schema drops, filesystem deletion, storage failure, or a superuser bypassing the configuration. Preserve tested backups and restrict configuration-file and service-control access. Upstream reports checks through PostgreSQL 18 RC1, but each packaged binary and target major version still needs a startup and command-hook test.
