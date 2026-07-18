## 用法

来源：

- [官方扩展控制文件](https://github.com/danolivo/pg_repeater/blob/93cabed0ce5cec20d36d1cd56e200a4a0091d55d/pg_repeater.control)

`pg_repeater` — 通过 postgres_fdw 将实用语句和序列化查询计划复制到远程 PostgreSQL 节点。

已复核目录快照记录的版本为 `0.1`、类型为 `preload`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`pg_execplan`, `postgres_fdw`。

```sql
CREATE EXTENSION "pg_repeater";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_repeater';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
