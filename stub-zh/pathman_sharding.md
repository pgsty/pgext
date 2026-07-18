## 用法

来源：

- [官方扩展控制文件](https://github.com/funbringer/pathman_sharding/blob/0c1bb0cfbb512f6850eb87de382e235c1db7b967/pathman_sharding.control)

`pathman_sharding` — pathman_sharding：分析处理类 PostgreSQL 扩展；为 pg_pathman 提供改进的分片能力。

已复核目录快照记录的版本为 `0.1`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`pg_pathman`, `plpgsql`, `postgres_fdw`。
整理后的兼容版本集合为 `10`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pathman_sharding";
SELECT extversion
FROM pg_extension
WHERE extname = 'pathman_sharding';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
