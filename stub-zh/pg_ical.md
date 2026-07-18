## 用法

来源：

- [当前上游功能与用法文档](https://github.com/thomastthai/pg_ical/blob/b68e3ce2cb90686739e831f052eeed7324f8caa7/README.md)
- [扩展控制文件](https://github.com/thomastthai/pg_ical/blob/b68e3ce2cb90686739e831f052eeed7324f8caa7/pg_ical.control)

`pg_ical` 基于 `libical` 提供 RFC 5545 重复规则函数。它能解析 RRULE 字段并展开重复事件，覆盖重复日期、排除日期、时区、时长和重复实例标识。

```sql
CREATE EXTENSION pg_ical;
SELECT get_freq('FREQ=WEEKLY;INTERVAL=1;BYDAY=MO,TH');
SELECT unnest(get_occurrences(
  'FREQ=DAILY;COUNT=3',
  '2026-01-01 09:00:00+00'::timestamptz
));
```

已复核版本为 `0.2.0`，上游明确不支持已弃用的 EXRULE。日历正确性取决于 RFC 解释、`libical` 版本、夏令时切换、边界包含语义和时间戳类型。应使用真实时区与重复上限测试，并限制不可信规则可能展开的实例数量。
