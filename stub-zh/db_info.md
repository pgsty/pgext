## 用法

来源：

- [官方扩展控制文件](https://github.com/asotolongo/db_info/blob/cfd4110846932f74e403c987449f70011ccf61d4/db_info.control)
- [官方上游文档](https://github.com/asotolongo/db_info/blob/cfd4110846932f74e403c987449f70011ccf61d4/README.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/db_info/)

`db_info` — 提供数据库大小、所有者、表空间、扩展及对象数量等信息的函数和视图。

已复核目录快照记录的版本为 `0.2.0`、类型为 `puresql`、实现语言为 `SQL`。
整理后的兼容版本集合为 `11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "db_info";
SELECT extversion
FROM pg_extension
WHERE extname = 'db_info';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
