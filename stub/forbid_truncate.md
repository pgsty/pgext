## Usage

Sources:

- [Official README](https://github.com/NaokiNakamichi/postgres_extension/blob/47029106ccbc1fe00d0ec9a612ddfedd7abf6847/README.md)
- [Version 1.0 control file](https://github.com/NaokiNakamichi/postgres_extension/blob/47029106ccbc1fe00d0ec9a612ddfedd7abf6847/forbid_truncate.control)
- [ProcessUtility hook implementation](https://github.com/NaokiNakamichi/postgres_extension/blob/47029106ccbc1fe00d0ec9a612ddfedd7abf6847/forbid_truncate.c)

`forbid_truncate` 1.0 installs a PostgreSQL `ProcessUtility` hook that rejects every `TRUNCATE` statement seen by a backend. It is a coarse guard against accidental truncation, with no role, table, database, or maintenance exception.

### Preload and Verify

Install the shared library, add it to `shared_preload_libraries`, and restart PostgreSQL:

```conf
shared_preload_libraries = 'forbid_truncate'
```

`CREATE EXTENSION` only records the extension in a database; the versioned SQL creates no callable objects. After the restart, verify the hook in a disposable table:

```sql
CREATE EXTENSION forbid_truncate;

CREATE TEMP TABLE truncate_guard_demo (id integer);
INSERT INTO truncate_guard_demo VALUES (1);
TRUNCATE truncate_guard_demo;
```

The last statement should fail with `forbid truncate`. Verify separately in every database and connection path that matters, because sessions started without the preloaded library are unprotected.

### Coverage and Hook Compatibility

The hook blocks `TRUNCATE` for ordinary users and superusers alike. It does not block `DELETE`, `DROP TABLE`, partition detach/drop, table replacement, or other ways to remove data, so normal privilege design, backups, and change controls remain necessary.

Although the implementation saves the previously registered `ProcessUtility_hook`, it does not call it for non-TRUNCATE commands; it calls `standard_ProcessUtility` directly. This can bypass another extension's hook depending on preload order. Test the exact `shared_preload_libraries` ordering with every audit, DDL-control, or utility-hook extension in the cluster. Disabling or changing the guard requires configuration change and restart because it exposes no GUC or SQL bypass.
