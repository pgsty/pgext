## 用法

来源：

- [官方扩展控制文件](https://github.com/dturon/partman_to_cstore/blob/80c6411318494837e5ef92825dd7a6c80b61fb66/partman_to_cstore.control)

`partman_to_cstore` — partman_to_cstore：分析处理类 PostgreSQL 扩展；用于将旧的 pg_partman 分区迁移到 cstore 列式存储。

已复核目录快照记录的版本为 `1.1.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`cstore_fdw`, `pg_partman`, `plpgsql`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "partman_to_cstore";
SELECT extversion
FROM pg_extension
WHERE extname = 'partman_to_cstore';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
