## 用法

来源：

- [官方上游文档](https://github.com/maxoodf/pg_mystem/blob/38554744c42e15c6f10132dcb314118a39bc85f1/README.md)

`pg_mystem` — 运行 Yandex Mystem 工作进程，为 PostgreSQL 查询提供俄语词形还原。

已复核目录快照记录的版本为 `1.0.1`、类型为 `preload`、实现语言为 `C++`。

```sql
CREATE EXTENSION "pg_mystem";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_mystem';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
