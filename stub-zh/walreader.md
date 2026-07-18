## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/moritetu/walreader/blob/f54a68f9f925da149a2b33030405cfeaea8d6197/README.md)
- [1.0 版本 SQL 对象](https://github.com/moritetu/walreader/blob/f54a68f9f925da149a2b33030405cfeaea8d6197/walreader--1.0.sql)
- [WAL 读取实现](https://github.com/moritetu/walreader/blob/f54a68f9f925da149a2b33030405cfeaea8d6197/walreader.c)

`walreader` 通过 SQL 暴露 PostgreSQL WAL 记录头和资源管理器描述。它是参考 `pg_waldump` 的教学实现；上游建议使用更轻量、功能更完整的标准工具。

```sql
CREATE EXTENSION walreader;
SET walreader.read_limit = 100;
SELECT timeline, walseg, lsn, rmgr, identify, rmgr_desc
FROM read_wal_segment(pg_walfile_name(pg_current_wal_lsn()))
LIMIT 20;
```

`walreader.default_wal_directory` 可让会话指向另一个服务器目录，函数还能按段名或 LSN 读取。WAL 包含敏感运维与数据变更信息。应从 `PUBLIC` 撤销全部访问，只授予专用超级用户诊断角色，也不要接受调用者控制的路径。

WAL 格式和内部解码器依赖大版本，绝不能用不匹配二进制解析段。读取活动段末尾时，报告未完成记录可能是正常现象。应设置严格记录数和语句限制，尽可能复制归档段后分析，并在不危及主库的环境测试损坏/截断输入。优先在服务器进程外使用 `pg_waldump` 隔离崩溃与资源消耗。
