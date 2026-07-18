## Usage

Sources:

- [Project README](https://github.com/ekayim/non_truncate/blob/036afbb78c4cbbbbb95a739fa420cb7b58b7d021/README.md)
- [Extension control file](https://github.com/ekayim/non_truncate/blob/036afbb78c4cbbbbb95a739fa420cb7b58b7d021/non_truncate.control)
- [Utility-hook implementation](https://github.com/ekayim/non_truncate/blob/036afbb78c4cbbbbb95a739fa420cb7b58b7d021/non_truncate.c)

`non_truncate` 1.0 installs a PostgreSQL utility hook that rejects every `TRUNCATE` statement with an error. It declares support for PostgreSQL 13 or later and must be loaded at server start.

### Load and register the extension

Add the library to `shared_preload_libraries` and restart PostgreSQL:

```conf
shared_preload_libraries = 'non_truncate'
```

Then register the extension in databases where its presence should be visible in the extension catalog:

```sql
CREATE EXTENSION non_truncate;

TRUNCATE sample_table;
-- ERROR: non_truncate extension is banning TRUNCATE statement.
```

### Cluster-wide effect

The blocking behavior starts when the shared library is preloaded, not when `CREATE EXTENSION` runs. Because `shared_preload_libraries` is cluster-wide, the hook also affects databases that do not contain an extension catalog entry. The version installation SQL creates no callable SQL object.

The hook rejects all `TRUNCATE` forms without an allowlist, role exemption, table exemption, or session switch. It does not prevent equivalent data removal with `DELETE`, table replacement, or other DDL, so it is a narrow guardrail rather than an authorization policy. Test hook chaining with every other utility-hook extension before deployment and keep a restart procedure for disabling it.
