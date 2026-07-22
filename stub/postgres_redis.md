## Usage

Sources:

- [Official README](https://github.com/systemEng-Learning/postgres-redis/blob/05ec5172157932635c1f773fd49d8b61dd13a948/README.md)
- [Official extension source](https://github.com/systemEng-Learning/postgres-redis/blob/05ec5172157932635c1f773fd49d8b61dd13a948/src/lib.rs)
- [Official GUC definitions](https://github.com/systemEng-Learning/postgres-redis/blob/05ec5172157932635c1f773fd49d8b61dd13a948/src/gucs.rs)

`postgres_redis` is an experimental hook and background-worker extension that copies a configured table's selected values to Redis. It observes qualifying `SELECT` and `UPDATE` statements, queues a key/value pair after commit, and lets a worker issue Redis `SET` commands asynchronously.

### Configuration and Workflow

Configure one table and one key/value column pair before server start, preload the library, then restart PostgreSQL.

```conf
shared_preload_libraries = 'postgres_redis'

postgres_redis.redis_url = 'redis://127.0.0.1'
postgres_redis.table = 'users'
postgres_redis.key_column = 'first_name'
postgres_redis.value_column = 'last_name'
postgres_redis.bg_delay = 10
```

```sql
CREATE EXTENSION postgres_redis;

SELECT last_name
FROM users
WHERE first_name = 'Adebayo';

UPDATE users
SET last_name = 'Updated'
WHERE first_name = 'Adebayo';
```

### GUCs and Behavior

- `postgres_redis.redis_url` is the Redis connection URL.
- `postgres_redis.table`, `postgres_redis.key_column`, and `postgres_redis.value_column` select the single table mapping.
- `postgres_redis.bg_delay` is the worker interval in seconds and has a minimum of one.
- The query hook recognizes the configured table only when the key column appears in an equality predicate against a constant. Other predicates and commands are ignored.

### Experimental Limitations

The reviewed source uses a shared queue of 400 entries and fixed 127-character key and value arrays; a full queue discards updates with a warning. Redis delivery happens after PostgreSQL commit but is not transactionally coupled to Redis, so crashes and Redis failures can lose or delay changes. The worker opens one Redis connection and the design covers only one table mapping.

`CREATE EXTENSION` also embeds the repository's `sql/test.sql`, which creates and seeds `test` and `users` tables. Inspect or patch the packaged SQL before installing into a database that already uses those names. The control file requires superuser installation, and the only explicit SQL function is `hello_postgres_redis`; this project should be treated as a learning prototype, not a general change-data-capture system.
