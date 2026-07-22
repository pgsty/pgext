## 用法

来源：

- [官方 README](https://github.com/BRupireddy2/pg_synthesize_wal/blob/5eb6e95f7f3e615626cfc5fede3c789f8b2de852/README.md)
- [官方扩展 SQL](https://github.com/BRupireddy2/pg_synthesize_wal/blob/5eb6e95f7f3e615626cfc5fede3c789f8b2de852/pg_synthesize_wal--1.0.sql)
- [官方扩展控制文件](https://github.com/BRupireddy2/pg_synthesize_wal/blob/5eb6e95f7f3e615626cfc5fede3c789f8b2de852/pg_synthesize_wal.control)

`pg_synthesize_wal` 会有意写入带随机或调用者指定载荷的自定义 WAL 记录。它是用于测试 WAL 行为的开发、测试和沙箱工具，可以生成大到跨越多个 WAL 文件的记录；上游明确不建议在生产服务器使用。

### 启用并生成 WAL

动态库必须在 postmaster 启动阶段注册自定义资源管理器。control 文件还声明 `pg_walinspect` 为扩展依赖，尽管实现只在测试中使用它。

```sql
ALTER SYSTEM SET shared_preload_libraries = 'pg_synthesize_wal';
-- Restart PostgreSQL before continuing.

CREATE EXTENSION pg_walinspect;
CREATE EXTENSION pg_synthesize_wal;

SELECT pg_synthesize_wal_record(1024::bigint);
SELECT pg_synthesize_wal_record('test payload'::text);
SELECT pg_synthesize_wal_record(decode('deadbeef', 'hex'));
```

每个重载都返回插入记录的 `pg_lsn`。尺寸重载生成随机字节；`text` 与 `bytea` 重载把调用者提供的载荷字节写入自定义记录。

### SQL 接口

- `pg_synthesize_wal_record(bigint)` 写入指定大小的随机载荷。
- `pg_synthesize_wal_record(text)` 写入文本载荷。
- `pg_synthesize_wal_record(bytea)` 写入二进制载荷。

所有重载都是 strict 且 parallel-unsafe。上游支持具备自定义 WAL 资源管理器的 PostgreSQL `15+`。

### 安全边界

调用这些函数会消耗 WAL 带宽、归档空间、复制带宽和恢复时间；极大的输入还可能消耗大量后端内存。redo 有意不修改数据，但记录仍会经过 WAL、归档、复制和恢复流程。只能在一次性或容量受控的系统上运行，同时监控 WAL 增长和副本延迟；集群恢复正常用途前应移除预加载设置。
