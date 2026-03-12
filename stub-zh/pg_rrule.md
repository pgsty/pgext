

## 用法

> [pg_rrule: PostgreSQL 的 iCalendar RRULE 重复规则类型](https://github.com/petropavel13/pg_rrule)

`pg_rrule` 扩展提供了 RRULE 数据类型，用于解析和展开 iCalendar 重复规则（RFC 5545）。

```sql
CREATE EXTENSION pg_rrule;
```

### 参数提取

```sql
-- 获取频率
SELECT get_freq('FREQ=WEEKLY;INTERVAL=1;WKST=MO;UNTIL=20200101T045102Z'::rrule);
-- WEEKLY

-- 获取星期几（数字数组）
SELECT get_byday('FREQ=WEEKLY;BYDAY=MO,TH,SU'::rrule);
-- {2,5,1}
```

### 生成重复事件

```sql
-- 从 RRULE 生成重复事件
SELECT unnest(get_occurrences(
    'FREQ=DAILY;INTERVAL=1;UNTIL=20200105T000000Z'::rrule,
    '2020-01-01 00:00:00'::timestamp
));
```

`get_occurrences()` 函数将 RRULE 定义展开为具体的时间戳序列，支持带时区和不带时区的时间戳参数。

### 依赖

需要 `libical` 库以符合 iCalendar 标准。
