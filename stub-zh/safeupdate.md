

## 用法

> [safeupdate: 要求 UPDATE 和 DELETE 必须带条件](https://github.com/eradman/pg-safeupdate)

`safeupdate` 扩展通过在执行不带 `WHERE` 子句的 `UPDATE` 或 `DELETE` 语句时抛出错误来防止意外的批量数据变更。

### 激活

```sql
-- 会话级别
LOAD 'safeupdate';

-- 数据库级别（持久化）
ALTER DATABASE mydb SET session_preload_libraries = 'safeupdate';

-- 全局（所有数据库，需要重启）
-- shared_preload_libraries = 'safeupdate'   -- 在 postgresql.conf 中
```

### 行为

```sql
-- 阻止：不带 WHERE 的 UPDATE
UPDATE rack SET fan_speed = 70;
-- ERROR: UPDATE requires a WHERE clause

-- 阻止：不带 WHERE 的 DELETE
DELETE FROM rack;
-- ERROR: DELETE requires a WHERE clause

-- 允许：带 WHERE 子句
UPDATE rack SET fan_speed = 90 WHERE fan_speed = 70;

-- 变通方法：显式的恒真条件
UPDATE rack SET fan_speed = 90 WHERE 1 = 1;
```

### 管理员覆盖

```sql
-- 在当前会话中临时禁用保护
SET safeupdate.enabled = 0;
```

基于 CTE 的不带 WHERE 条件的修改也会被阻止。该扩展对于 PostgREST 或其他提供直接数据库写入访问的系统特别有用。
