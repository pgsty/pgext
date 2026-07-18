## 用法

来源：

- [官方扩展控制文件](https://github.com/benizar/pg_adm/blob/78314723ddfeb2128847018d5198d4156b527fb4/pg_adm.control)
- [官方上游文档](https://github.com/benizar/pg_adm/blob/78314723ddfeb2128847018d5198d4156b527fb4/doc/pg_adm.md)
- [官方上游 README](https://github.com/benizar/pg_adm/blob/78314723ddfeb2128847018d5198d4156b527fb4/README.md)

`pg_adm` — 提供若干 PostgreSQL 管理辅助工具。

已复核目录快照记录的版本为 `0.0.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`pg_buffercache`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_adm";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_adm';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
