## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/sonesuke/pg_elephantduck/blob/0427d744ab3e3b8e2fbb95328888e3551942eae7/README.md)
- [扩展 control 文件](https://github.com/sonesuke/pg_elephantduck/blob/0427d744ab3e3b8e2fbb95328888e3551942eae7/pg_elephantduck.control)
- [表访问方法实现](https://github.com/sonesuke/pg_elephantduck/blob/0427d744ab3e3b8e2fbb95328888e3551942eae7/src/tam.rs)
- [Cargo 元数据](https://github.com/sonesuke/pg_elephantduck/blob/0427d744ab3e3b8e2fbb95328888e3551942eae7/Cargo.toml)

`pg_elephantduck` 是由嵌入式 DuckDB 与 Parquet 文件支持的实验性 PostgreSQL 表访问方法，提供列式存储与有限的 `WHERE` 下推。

```sql
CREATE EXTENSION pg_elephantduck;
SET elephantduck.path = '/srv/postgresql/elephantduck';
CREATE TABLE sample USING elephantduck
AS SELECT generate_series(1, 10000) AS number;
SELECT number FROM sample WHERE number < 10;
```

存储目录默认位于 `PGDATA` 下或 `/tmp`；应明确配置由 PostgreSQL 拥有的持久路径。备份恢复、复制、访问控制和磁盘容量规划都必须覆盖该路径，不能假定它具有堆表的 WAL 语义。

0.0.0 是面向 PostgreSQL 16–17 的活跃原型。上游要求通过 issue 了解限制，并未声称其 DML、崩溃恢复、复制、vacuum、锁或事务行为与堆表完全一致。在准确构建上验证这些行为前，只能用于可丢弃数据。
