## Usage

Sources:

- [Official documentation](https://github.com/jasonmassie01/pg-sats/blob/3c1a55bf8d550f6e34fc000dcd3f6bd6480d6651/README.md)
- [Extension control file](https://github.com/jasonmassie01/pg-sats/blob/3c1a55bf8d550f6e34fc000dcd3f6bd6480d6651/pg_sats.control)
- [Official repository](https://github.com/jasonmassie01/pg-sats)

`pg_sats` Per-role query metering and optional enforcement using satoshi-denominated balances.

### Enablement

Merge `pg_sats` into the existing preload list, restart PostgreSQL, and then create `pg_sats` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'pg_sats'
```

```sql
CREATE EXTENSION pg_sats;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
-- Fund an account with 1000 sats
SELECT pg_sats.deposit('alice', 1000);

-- Check balance
SELECT pg_sats.balance('alice');
-- Returns: 1000

-- Alice runs queries... each costs 1 sat (default)
-- After 10 queries:
SELECT pg_sats.balance('alice');
-- Returns: 990

-- View spending history
SELECT * FROM pg_sats.history('alice');
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `pg_sats.balance` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_sats.deposit` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_sats.history` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_sats.balances` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `pg_sats.query_history` | TABLE | Extension-owned table; account for its data in backup and upgrades. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 14, 15, 16, 17, 18; do not infer unlisted majors.
- Preloading `pg_sats` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
- The extension fixes or creates schema objects under `pg_sats`; include them in privilege and backup review.
