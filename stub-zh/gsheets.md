## 用法

来源：

- [官方扩展控制文件](https://github.com/MuhammadTahaNaveed/pg-gsheets/blob/f223b892ff2a0dcd98c5f2ac0ca8d748a3b5bb28/gsheets.control)

`gsheets` — 在 PostgreSQL 中读写 Google Sheets 数据。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "gsheets";
SELECT extversion
FROM pg_extension
WHERE extname = 'gsheets';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
