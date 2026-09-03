## Usage

Sources:

- [Official documentation](https://github.com/rogerwelin/pg_fk_indexer/blob/5ab97eaff8e456d58bda8d300efd64083247b534/README.md)
- [Extension control file](https://github.com/rogerwelin/pg_fk_indexer/blob/5ab97eaff8e456d58bda8d300efd64083247b534/pg_fk_indexer.control)
- [Official repository](https://github.com/rogerwelin/pg_fk_indexer)

`pg_fk_indexer` Automatically creates indexes for newly defined foreign-key columns.

### Enablement

Merge `pg_fk_indexer` into the existing preload list, restart PostgreSQL, and then create `pg_fk_indexer` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'pg_fk_indexer'
```

```sql
CREATE EXTENSION pg_fk_indexer;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
-- Foreign keys are now auto-indexed
CREATE TABLE users (id int PRIMARY KEY, username text);
CREATE TABLE orders (user_id int REFERENCES users(id));
--  index on orders(user_id) is created automatically
```

### Main Objects

The official sources define the extension surface dynamically or through provider tooling; inspect the installed version before granting access.

### Operations and Boundaries

- Preloading `pg_fk_indexer` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
