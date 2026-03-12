

## 用法

> [periods: PostgreSQL 的时间段与系统版本控制](https://github.com/xocolatl/periods)

本扩展实现了 SQL:2016 标准（最初于 SQL:2011 引入）中关于时间段（Period）和 `SYSTEM VERSIONING` 系统版本控制的行为。

### 什么是时间段？

时间段是表上的一个定义，指定一个名称和两个列。时间段的值包含起始值，但不包含结束值。

```sql
-- 标准 SQL 语法
CREATE TABLE example (
    id bigint,
    start_date date,
    end_date date,
    PERIOD FOR validity (start_date, end_date)
);
```

由于扩展无法修改 PostgreSQL 的语法，我们通过函数、视图和触发器来尽可能模拟相同的行为：

```sql
CREATE TABLE example (
    id bigint,
    start_date date,
    end_date date
);
SELECT periods.add_period('example', 'validity', 'start_date', 'end_date');
```

定义时间段后，会约束两列：起始值必须严格小于结束值，且两列均不能为空。

## 唯一约束

时间段可以作为 `PRIMARY KEY` 和 `UNIQUE` 约束的一部分：

```sql
CREATE TABLE example (
    id bigint,
    start_date date,
    end_date date
);
SELECT periods.add_period('example', 'validity', 'start_date', 'end_date');
SELECT periods.add_unique_key('example', ARRAY['id'], 'validity');
```

扩展会创建一个覆盖所有指定列和时间段两列的唯一约束，并通过 GiST 排他约束实现 `WITHOUT OVERLAPS` 语义。

## 外键

有了带时间段的唯一键，就可以创建指向它们的外键：

```sql
SELECT periods.add_foreign_key('example2', 'ARRAY[ex_id]', 'validity', 'example_id_validity');
```

## 部分操作

SQL 标准允许对时间段的一部分进行更新或删除，扩展通过 `INSTEAD OF` 触发器的视图来实现：

```sql
UPDATE example__for_portion_of_validity
SET ...,
    start_date = ...,
    end_date = ...
WHERE ...;
```

使用此功能前，表必须有主键。

## 谓词函数

SQL 标准定义了多种时间段谓词，以内联函数形式实现：

```sql
-- "t" 和 "u" 是具有时间段 "p" 和 "q" 的表
-- 两个时间段的底层列为 "s" 和 "e"

WHERE periods.contains(t.s, t.e, 42)            -- t.p CONTAINS 42
WHERE periods.contains(t.s, t.e, u.s, u.e)      -- t.p CONTAINS u.q
WHERE periods.equals(t.s, t.e, u.s, u.e)        -- t.p EQUALS u.q
WHERE periods.overlaps(t.s, t.e, u.s, u.e)      -- t.p OVERLAPS u.q
WHERE periods.precedes(t.s, t.e, u.s, u.e)      -- t.p PRECEDES u.q
WHERE periods.succeeds(t.s, t.e, u.s, u.e)      -- t.p SUCCEEDS u.q
WHERE periods.immediately_precedes(t.s, t.e, u.s, u.e)  -- t.p IMMEDIATELY PRECEDES u.q
WHERE periods.immediately_succeeds(t.s, t.e, u.s, u.e)  -- t.p IMMEDIATELY SUCCEEDS u.q
```


## 系统版本控制表

### SYSTEM_TIME

如果时间段名为 `SYSTEM_TIME`，则适用特殊规则。列类型必须是 `date`、`timestamp without time zone` 或 `timestamp with time zone`，且用户不可修改。扩展使用触发器将起始列设为 `transaction_timestamp()`，结束列始终为 `'infinity'`。

***注意：*** 一般建议使用 `timestamp with time zone`，因为时区配置参数或夏令时变更可能导致历史记录失真。

```sql
CREATE TABLE example (
    id bigint PRIMARY KEY,
    value text
);
SELECT periods.add_system_time_period('example', 'row_start', 'row_end');
```

这些列无需预先存在——扩展会自动创建。

### 排除列

可以阻止某些列的更新触发 `SYSTEM_TIME` 值变化：

```sql
SELECT periods.add_system_time_period(
            'example',
            excluded_column_names => ARRAY['foo', 'bar']);
```

### 启用系统版本控制

```sql
SELECT periods.add_system_time_period('example', 'row_start', 'row_end');
SELECT periods.add_system_versioning('example');
```

系统会将所有变更记录到一张独立的历史表中。你也可以自行创建历史表（例如添加分区）并指定扩展使用它。

### 时态查询

SQL 标准扩展了 `FROM` 和 `JOIN` 子句以支持时间点或时间范围查询，扩展通过内联函数实现：

```sql
SELECT * FROM t__as_of('...');                       -- FOR system_time AS OF '...'
SELECT * FROM t__from_to('...', '...');              -- FOR system_time FROM '...' TO '...'
SELECT * FROM t__between('...', '...');              -- FOR system_time BETWEEN '...' AND '...'
SELECT * FROM t__between_symmetric('...', '...');    -- FOR system_time BETWEEN SYMMETRIC '...' AND '...'
```

### 访问控制

历史表及辅助函数遵循基表的所有权和访问权限。历史数据为只读。如需清理旧数据，必须先暂停系统版本控制：

```sql
BEGIN;
SELECT periods.drop_system_versioning('t');
GRANT DELETE ON TABLE t TO CURRENT_USER;
DELETE FROM t_history WHERE system_time_end < now() - interval '1 year';
SELECT periods.add_system_versioning('t');
COMMIT;
```
