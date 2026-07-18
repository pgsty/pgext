## 用法

来源：

- [官方扩展控制文件](https://api.pgxn.org/src/pg_collkey/pg_collkey-0.5.1/pg_collkey.control)
- [官方上游文档](https://pgxn.org/dist/pg_collkey/0.5.1/)
- [官方 PGXN 分发页](https://pgxn.org/dist/pg_collkey/)

`pg_collkey` — pg_collkey：为 PostgreSQL 提供 ICU 排序键函数封装，可按指定 locale 生成排序键。

已复核目录快照记录的版本为 `0.5.1`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_collkey";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_collkey';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
