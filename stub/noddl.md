## Usage

Sources:

- [Official documentation](https://github.com/bonesmoses/noddl/blob/7c9f1c6558bb8379ce176ce599542d123129f841/README.md)
- [Extension control file](https://github.com/bonesmoses/noddl/blob/7c9f1c6558bb8379ce176ce599542d123129f841/noddl.control)
- [Official repository](https://github.com/bonesmoses/noddl)

`noddl` Cluster-wide DDL blocker with role and database exceptions.

### Enablement

Merge `noddl` into the existing preload list, restart PostgreSQL, and then create `noddl` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'noddl'
```

```sql
CREATE EXTENSION noddl;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
ALTER SYSTEM SET noddl.enable TO true;
SELECT pg_reload_conf();
```

### Main Objects

The official sources define the extension surface dynamically or through provider tooling; inspect the installed version before granting access.

### Operations and Boundaries

- Preloading `noddl` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
