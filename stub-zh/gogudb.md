## 用法

来源：

- [官方扩展控制文件](https://github.com/csudata/gogudb/blob/490294508c8df14e29f312a1720e53690ef24de4/gogudb.control)
- [官方上游文档](https://github.com/csudata/gogudb/blob/490294508c8df14e29f312a1720e53690ef24de4/Install.md)
- [官方上游 README](https://github.com/csudata/gogudb/blob/490294508c8df14e29f312a1720e53690ef24de4/README.md)

`gogudb` — 通过外部服务器实现分布式哈希与范围分区的 PostgreSQL 扩展。

已复核目录快照记录的版本为 `1.1`、类型为 `preload`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `10,11`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "gogudb";
SELECT extversion
FROM pg_extension
WHERE extname = 'gogudb';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
