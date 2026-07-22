## 用法

来源：

- [官方 README](https://github.com/edgarogh/postgres-ical/blob/8a2a9b7fa32002c53298b723de86045163812f74/README.md)
- [扩展控制文件](https://github.com/edgarogh/postgres-ical/blob/8a2a9b7fa32002c53298b723de86045163812f74/postgres_ical.control)
- [SQL 可见的 Rust 实现](https://github.com/edgarogh/postgres-ical/blob/8a2a9b7fa32002c53298b723de86045163812f74/src/lib.rs)
- [iCalendar 事件解析器](https://github.com/edgarogh/postgres-ical/blob/8a2a9b7fa32002c53298b723de86045163812f74/postgres-ical-parser/src/parser.rs)

`postgres_ical` 0.1.0 将 RFC 5545 iCalendar `VEVENT` 组件解析为 PostgreSQL 行。它既接受内存日历文本，也可接受由数据库后端获取的 URL，适合显式查询或严格受控的物化视图刷新。

### 解析内存日历文本

应显式选择列，因为上游保留增加输出列而不将其视为破坏性变更的权利：

```sql
CREATE EXTENSION postgres_ical;

SELECT uid, summary, dt_start, dt_start_naive
FROM pg_ical($ical$
BEGIN:VCALENDAR
BEGIN:VEVENT
UID:event-42@example.org
DTSTART:20260722T090000Z
SUMMARY:Review
END:VEVENT
END:VCALENDAR
$ical$);
```

UTC 或带时区值会填充 `timestamptz` 列，例如 `dt_start`；浮动本地值会填充对应的 `timestamp` 列，例如 `dt_start_naive`。

### 获取日历 URL

```sql
SELECT uid, summary, dt_start, dt_start_naive
FROM pg_ical_curl('https://example.com/calendar.ical');
```

`pg_ical_curl(url text)` 通过 libcurl 流式读取响应，返回与 `pg_ical(calendar text)` 相同的行结构。

### 解析表面

当前解析器要求 `UID` 和 `DTSTART`。它识别 `CREATED`、`DESCRIPTION`、`DTSTAMP`、`DTEND`、`LAST-MODIFIED`、`LOCATION`、`SEQUENCE` 和 `SUMMARY`；`SEQUENCE` 默认为零。未知属性和非 `VEVENT` 组件会被跳过。

返回行类型暴露了更多面向 RFC 5545 的列，但所审计实现目前会把 attachment、categories、class、comments、completed、due、duration、geography、percent-complete、priority、resources 和 status 留为 NULL 或空值。每行返回的 `component_type` 都是 `VEVENT`。

### 安全与兼容边界

`pg_ical_curl` 从 PostgreSQL 后端发起出站请求，所审计代码没有设置显式协议允许列表、响应大小限制或请求超时。应撤销不受信任角色的执行权并约束出口网络，以避免服务器端请求伪造和后端资源耗尽。如果 URL 由用户控制，应优先在数据库外获取并验证远程数据。

畸形事件、缺少必需属性、无效时区和传输失败会被 Rust 代码直接 unwrap，并表现为 PostgreSQL 错误。上游 README 与 Cargo 清单声明项目未授权/`UNLICENSED`，因此源码再分发和生产采用都需要单独的法律审查。
