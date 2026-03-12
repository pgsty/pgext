

## 用法

> [tsm_system_rows: 基于行数的 TABLESAMPLE 采样方法](https://www.postgresql.org/docs/current/tsm-system-rows.html)

提供 `SYSTEM_ROWS` 表采样方法，精确返回指定数量的行。

```sql
CREATE EXTENSION tsm_system_rows;
```

### TABLESAMPLE 方法

`SYSTEM_ROWS(count int)` -- 最大返回行数。

### 示例

```sql
-- 精确采样 100 行
SELECT * FROM my_table TABLESAMPLE SYSTEM_ROWS(100);

-- 快速预览大表中的 10 行
SELECT * FROM large_table TABLESAMPLE SYSTEM_ROWS(10);
```

执行块级采样（小样本可能出现聚集效应）。如果表中行数少于请求数量，则返回所有行。不支持 `REPEATABLE`。
