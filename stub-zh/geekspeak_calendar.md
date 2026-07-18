## 用法

来源：

- [官方扩展控制文件](https://github.com/ttfkam/pg_geekspeak_calendar/blob/34dcd6de80a4bf3116a19d21d0d222337c1f1d7f/geekspeak_calendar.control)
- [官方上游文档](https://github.com/ttfkam/pg_geekspeak_calendar/blob/34dcd6de80a4bf3116a19d21d0d222337c1f1d7f/README.md)

`geekspeak_calendar` — 为 GeekSpeak 播客录制日程与节目提供日历及 iCalendar 集成。

已复核目录快照记录的版本为 `1.0.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`geekspeak`。

```sql
CREATE EXTENSION "geekspeak_calendar";
SELECT extversion
FROM pg_extension
WHERE extname = 'geekspeak_calendar';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
