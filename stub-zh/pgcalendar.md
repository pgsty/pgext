
## 用法

> 语法：
>
> ```sql
> CREATE EXTENSION pgcalendar;
> INSERT INTO pgcalendar.events (name, description, category)
> VALUES ('Daily Standup', 'Team daily standup meeting', 'meeting');
> SELECT * FROM pgcalendar.get_event_projections(1, '2024-01-01'::date, '2024-01-07'::date);
> ```
>
> 来源：[README](https://github.com/h4kbas/pgcalendar)

`pgcalendar` 是一个用于 PostgreSQL 的循环日历扩展。它建模事件、日程、例外和投影，并可在任意日期范围内生成日历发生实例。

## 数据模型

README 描述了四个核心概念：

- `events`，表示会议或任务等逻辑对象
- `schedules`，表示生成投影的非重叠循环定义
- `exceptions`，表示单次发生的取消或修改
- `projections`，表示实际生成出来的日历发生实例

## 快速开始

创建事件：

```sql
INSERT INTO pgcalendar.events (name, description, category)
VALUES ('Daily Standup', 'Team daily standup meeting', 'meeting');
```

创建日程：

```sql
INSERT INTO pgcalendar.schedules (
    event_id, start_date, end_date, recurrence_type, recurrence_interval
) VALUES (
    1, '2024-01-01 09:00:00', '2024-01-07 23:59:59', 'daily', 1
);
```

获取投影：

```sql
SELECT * FROM pgcalendar.get_event_projections(
    1, '2024-01-01'::date, '2024-01-07'::date
);
```

## 循环类型

README 展示了以下日程示例：

- 每日循环
- 每周循环，带 `recurrence_day_of_week`
- 每月循环，带 `recurrence_day_of_month`
- 每年循环，带 `recurrence_month` 和 `recurrence_day_of_month`

## 例外

例外可以取消或修改某次发生：

```sql
INSERT INTO pgcalendar.exceptions (
    schedule_id, exception_date, exception_type, notes
) VALUES (
    1, '2024-01-15', 'cancelled', 'Holiday - meeting cancelled'
);
```

也可以修改日期和时间。

## 函数与视图

README 文档中包括：

- `get_event_projections(event_id, start_date, end_date)`
- `get_events_detailed(start_date, end_date)`
- `transition_event_schedule(...)`
- `check_schedule_overlap(event_id, start_date, end_date)`
- `pgcalendar.event_calendar`

其中 `transition_event_schedule(...)` 用于安全切换事件的日程配置，`check_schedule_overlap(...)` 用于验证新日程不会与现有日程冲突。
