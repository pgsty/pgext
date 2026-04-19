## 用法

来源：[README](https://github.com/h4kbas/pgcalendar/blob/master/README.md)，[repo](https://github.com/h4kbas/pgcalendar)

`pgcalendar` 是一个 PostgreSQL 循环日历扩展。上游 README 用四个核心对象来建模循环日程：`events`、`schedules`、`exceptions` 和生成出来的 projections。

### 创建事件和日程

```sql
CREATE EXTENSION pgcalendar;

INSERT INTO pgcalendar.events (name, description, category)
VALUES ('Daily Standup', 'Team daily standup meeting', 'meeting');

INSERT INTO pgcalendar.schedules (
    event_id, start_date, end_date, recurrence_type, recurrence_interval
) VALUES (
    1, '2024-01-01 09:00:00', '2024-01-07 23:59:59', 'daily', 1
);
```

README 展示了 `daily`、`weekly`、`monthly` 和 `yearly` 四类循环，并会根据类型使用 `recurrence_day_of_week`、`recurrence_day_of_month`、`recurrence_month` 等额外列。

### 查询投影

```sql
SELECT * FROM pgcalendar.get_event_projections(
    1, '2024-01-01'::date, '2024-01-07'::date
);

SELECT * FROM pgcalendar.get_events_detailed(
    '2024-01-01'::date, '2024-01-31'::date
);
```

README 还把 `pgcalendar.event_calendar` 视图作为快速核验结果的入口。

### 例外与日程切换

```sql
INSERT INTO pgcalendar.exceptions (
    schedule_id, exception_date, exception_type, notes
) VALUES (
    1, '2024-01-15', 'cancelled', 'Holiday - meeting cancelled'
);

SELECT pgcalendar.transition_event_schedule(
    p_event_id := 1,
    p_new_start_date := '2024-01-15 09:00:00',
    p_new_end_date := '2024-01-31 23:59:59',
    p_recurrence_type := 'weekly',
    p_recurrence_interval := 2,
    p_recurrence_day_of_week := 1,
    p_description := 'Changed to bi-weekly schedule'
);
```

如果需要在新增日程前确认日期范围不重叠，可以先调用 `pgcalendar.check_schedule_overlap(...)`。

### 注意事项

目前上游唯一公开的用户文档就是 README。它已经给出了较完整的表结构和函数示例，但没有另外维护按版本划分的用户可见 SQL 变更说明。
