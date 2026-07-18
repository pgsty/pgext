## Usage

Sources:

- [Official extension control file](https://github.com/systemEng-Learning/postgres-redis/blob/05ec5172157932635c1f773fd49d8b61dd13a948/postgres_redis.control)
- [Official upstream documentation](https://github.com/systemEng-Learning/postgres-redis/blob/05ec5172157932635c1f773fd49d8b61dd13a948/README.md)
- [Official Rust package manifest](https://github.com/systemEng-Learning/postgres-redis/blob/05ec5172157932635c1f773fd49d8b61dd13a948/Cargo.toml)

`postgres_redis` — synchronize selected PostgreSQL table values to Redis with hooks and a background worker

The reviewed catalog snapshot records version `0.0.0`, kind `preload`, and implementation language `Rust`.
The curated compatibility set is `11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "postgres_redis";
SELECT extversion
FROM pg_extension
WHERE extname = 'postgres_redis';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
