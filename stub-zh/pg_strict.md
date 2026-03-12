

## 用法

> [pg_strict: 阻止不带 WHERE 子句的危险 UPDATE 和 DELETE](https://github.com/spa5k/pg_strict)

`pg_strict` 扩展阻止缺少 `WHERE` 子句的 `UPDATE` 和 `DELETE` 语句。它通过 `post_parse_analyze_hook` 在解析/分析阶段运行，为每种语句类型提供三种执行模式。

### 配置参数

| 参数 | 模式 | 描述 |
|-----------|-------|-------------|
| `pg_strict.require_where_on_update` | `on`/`warn`/`off` | 对 UPDATE 强制要求 WHERE |
| `pg_strict.require_where_on_delete` | `on`/`warn`/`off` | 对 DELETE 强制要求 WHERE |

- **`on`**：拒绝不带 WHERE 的语句（抛出错误）
- **`warn`**：允许但发出警告日志
- **`off`**：标准 PostgreSQL 行为

### 会话级配置

```sql
SET pg_strict.require_where_on_update = 'on';
SET pg_strict.require_where_on_delete = 'warn';
```

### 持久化配置

```sql
ALTER DATABASE postgres SET pg_strict.require_where_on_update = 'on';
ALTER ROLE app_service SET pg_strict.require_where_on_delete = 'on';
ALTER ROLE dba_admin SET pg_strict.require_where_on_update = 'off';
```

### 事务级覆盖

```sql
BEGIN;
SET LOCAL pg_strict.require_where_on_delete = 'off';
DELETE FROM temp_table;  -- 在此事务中允许
COMMIT;
```

### API 函数

```sql
SELECT pg_strict_version();           -- 扩展版本
SELECT pg_strict_config();            -- 所有设置及其值和描述

-- 以编程方式验证查询
SELECT pg_strict_check_where_clause('DELETE FROM t', 'DELETE');  -- 返回布尔值
SELECT pg_strict_validate_update('UPDATE t SET x=1');
SELECT pg_strict_validate_delete('DELETE FROM t');

-- 快速模式切换
SELECT pg_strict_enable_update();     -- 将 update 执行设为 'on'
SELECT pg_strict_warn_delete();       -- 将 delete 执行设为 'warn'
SELECT pg_strict_disable_update();    -- 将 update 执行设为 'off'
```

任何非空 WHERE 条件都被接受（包括 `WHERE false`）。支持 CTE 语句。
