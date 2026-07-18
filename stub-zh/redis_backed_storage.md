## 用法

来源：

- [上游原型设计与限制](https://github.com/iJustErikk/redis-backed-pg-storage-engine/blob/1509d3e616e4dc58a0adf809d6fcae0aa72b6032/README.md)
- [表访问方法实现](https://github.com/iJustErikk/redis-backed-pg-storage-engine/blob/1509d3e616e4dc58a0adf809d6fcae0aa72b6032/src/lib.rs)
- [Rust 包清单](https://github.com/iJustErikk/redis-backed-pg-storage-engine/blob/1509d3e616e4dc58a0adf809d6fcae0aa72b6032/Cargo.toml)

`redis_backed_storage` 是尚未完成的 pgrx 实验，尝试实现 Redis 后端的 PostgreSQL 表访问方法。原型声称的范围只有一个类字符串列，以及基本的建表、插入和扫描。

```sql
CREATE EXTENSION redis_backed_storage;
CREATE TABLE demo (payload bytea) USING redis_backed_storage;
```

上游称其为周末项目，并明确没有并发安全、WAL、事务保证或 Redis 重启持久性。已复核源码仍含占位回调，Redis 集成也不完整。不要用它保存持久数据，也不要放到有价值的集群；它只适合在一次性实验室中研究访问方法接口。
