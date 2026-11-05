## 用法

来源：

- [上游来源 1](https://github.com/justusepostgres/pg_flaro/blob/d2cc08a462e46f3072c2fdc148ebb5b80ab3b0c2/README.md)
- [上游来源 2](https://github.com/justusepostgres/pg_flaro/blob/d2cc08a462e46f3072c2fdc148ebb5b80ab3b0c2/crates/pg-cache/pg_cache_journal.control)
- [上游来源 3](https://github.com/justusepostgres/pg_flaro/blob/d2cc08a462e46f3072c2fdc148ebb5b80ab3b0c2/Makefile)

`pg_cache_journal` 0.1.0 为 Flaro 写后持久化原型增加可选的 WAL 日志队列。

### 核心流程

使用启用 `journal-prototype` 特性的构建产物，再在缓存数据库中创建两个扩展。父扩展库需要预加载和重启。批次写入 `cache_journal.segment0` 与 `cache_journal.segment1`，独立工作进程随后将其物化到 `cache_journal.entries`。

```ini
pg_cache.journal = on
```

```sql
CREATE EXTENSION pg_cache;
CREATE EXTENSION pg_cache_journal;
SELECT * FROM cache_journal.state;
```

### 运行边界

依赖 PostgreSQL 17 与 pg_cache。日志提交和物化是两个不同阶段，仅查询物化表可能落后于已提交的日志。RESP 确认与持久化之间仍是异步关系。容量有上限，持续写入压力可能耗尽容量，不能把此实验模式视为 Redis 同步持久化。
