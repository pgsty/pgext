## 用法

来源：

- [官方扩展控制文件](https://github.com/macdice/calendars/blob/22d7900c371ad142505667c14c7d9605dccf077d/calendars.control)
- [官方上游文档](https://github.com/macdice/calendars/blob/22d7900c371ad142505667c14c7d9605dccf077d/README)

`calendars` — calendars：通用类 PostgreSQL 扩展；上游说明为“日历转换函数”。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "calendars";
SELECT extversion
FROM pg_extension
WHERE extname = 'calendars';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
