## 用法

来源：

- [官方 README](https://github.com/ttfkam/pg_geekspeak_calendar/blob/34dcd6de80a4bf3116a19d21d0d222337c1f1d7f/README.md)
- [官方扩展控制文件](https://github.com/ttfkam/pg_geekspeak_calendar/blob/34dcd6de80a4bf3116a19d21d0d222337c1f1d7f/geekspeak_calendar.control)
- [官方扩展 SQL](https://github.com/ttfkam/pg_geekspeak_calendar/blob/34dcd6de80a4bf3116a19d21d0d222337c1f1d7f/geekspeak_calendar--1.0.0.sql)

`geekspeak_calendar` 1.0.0 是面向 GeekSpeak 播客数据库的专用 SQL 扩展。它保存周期性录制时间窗，并通过滚动的 `calendar` 视图，将日程与必需的 `geekspeak` 扩展提供的节目、参与者、人员和地点数据组合起来。

### 核心流程

在预期的 GeekSpeak schema 中安装依赖与扩展，然后添加录制日程并查询生成的日历：

```sql
CREATE EXTENSION geekspeak;
CREATE EXTENSION geekspeak_calendar WITH SCHEMA gs;
SET search_path = gs, public;

INSERT INTO recording_schedules ("start", "end", location, cancellations)
VALUES (
  '2026-07-25 09:00:00+00',
  '2026-12-26 10:00:00+00',
  NULL,
  ARRAY['2026-08-01'::date]
);

SELECT title, description, participants, location, "start", "end"
FROM calendar
ORDER BY "start";
```

`start` 与 `end` 的时间差定义每次录制的时长，两者的日期限定周期日程的起止范围；`cancellations` 排除指定日期，`location` 提供默认地点。

### 重要对象

- `recording_schedules` 保存每周日程的首次开始时间、最终结束时间、默认地点和取消日期。
- `sessions_similarity_gist` 是排除约束，用来拒绝星期相同且日内时间范围重叠的录制时间窗。
- `calendar` 展开当前日期附近的日程，连接待录制节目数据，并返回标题、描述、参与者、地点、几何位置、修改时间、开始与结束时间。

### 运维说明

这不是通用的 iCalendar 解析器或导出器。该视图与 `geekspeak` schema 中的 `episodes`、`participants`、`people` 和 `locations` 对象耦合，控制文件也将 `geekspeak` 声明为必需扩展。滚动视图使用 `now()` 和固定的 `generate_series(0, 50)` 范围，因此结果会随时间变化，覆盖大约未来 50 个每周周期及此前一周。部署前应针对准确的 `geekspeak` schema 复核上游 SQL。
