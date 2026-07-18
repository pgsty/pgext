## Usage

Sources:

- [Upstream prototype design and limitations](https://github.com/iJustErikk/redis-backed-pg-storage-engine/blob/1509d3e616e4dc58a0adf809d6fcae0aa72b6032/README.md)
- [Table access-method implementation](https://github.com/iJustErikk/redis-backed-pg-storage-engine/blob/1509d3e616e4dc58a0adf809d6fcae0aa72b6032/src/lib.rs)
- [Rust package manifest](https://github.com/iJustErikk/redis-backed-pg-storage-engine/blob/1509d3e616e4dc58a0adf809d6fcae0aa72b6032/Cargo.toml)

`redis_backed_storage` is an unfinished pgrx experiment for a Redis-backed PostgreSQL table access method. The stated prototype scope is one string-like column with basic table creation, insertion, and scanning.

```sql
CREATE EXTENSION redis_backed_storage;
CREATE TABLE demo (payload bytea) USING redis_backed_storage;
```

Upstream calls this a weekend project and documents no concurrency safety, WAL, transactional guarantees, or Redis-restart durability. The reviewed source still contains stub callbacks and incomplete Redis integration. Do not use it for persistent data or on a valuable cluster; it is useful only for studying access-method interfaces in a disposable lab.
