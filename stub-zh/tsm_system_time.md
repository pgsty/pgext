

## 用法

> [tsm_system_time: 基于时间的 TABLESAMPLE 采样方法](https://www.postgresql.org/docs/current/tsm-system-time.html)

提供 `SYSTEM_TIME` 表采样方法，在指定时间限制内返回尽可能多的行。

```sql
CREATE EXTENSION tsm_system_time;
```

### TABLESAMPLE 方法

`SYSTEM_TIME(milliseconds float)` -- 读取表的最大时间（毫秒）。

### 示例

```sql
-- 采样 1 秒内可读取的行
SELECT * FROM my_table TABLESAMPLE SYSTEM_TIME(1000);

-- 使用 500 毫秒预算从大表中采样
SELECT count(*) FROM large_table TABLESAMPLE SYSTEM_TIME(500);
```

执行块级采样（非行级）。如果整个表可以在时间限制内读完，则返回所有行。不支持 `REPEATABLE`。
