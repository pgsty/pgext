## 用法

来源：

- [官方扩展控制文件](https://github.com/snaga/tablelog/blob/0981b7efe69a477c06cce53f0b2151d6b1dd38ad/tablelog.control)
- [官方上游文档](https://github.com/snaga/tablelog/blob/0981b7efe69a477c06cce53f0b2151d6b1dd38ad/README.md)

`tablelog` — 基于触发器的审计与变更数据捕获扩展：把 INSERT、UPDATE、DELETE 的行前后镜像写入集中日志表，并可生成回放 SQL。

已复核目录快照记录的版本为 `0.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`, `plv8`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "tablelog";
SELECT extversion
FROM pg_extension
WHERE extname = 'tablelog';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
