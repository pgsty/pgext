
## 用法

> 语法：
>
> ```sql
> CREATE EXTENSION pg_variables;
> SELECT pgv_set('vars', 'int1', 101);
> SELECT pgv_get('vars', 'int1', NULL::int);
> ```
>
> 来源：[README](https://github.com/postgrespro/pg_variables)

`pg_variables` 为 PostgreSQL 提供会话级变量。变量按 package 分组，只在当前会话中可见，并且可以配置为事务性或非事务性。

## 基本行为

默认情况下，变量不是事务性的，不会受 `BEGIN`、`COMMIT` 或 `ROLLBACK` 影响。`pgv_set()` 的可选参数 `is_transactional` 可以改变这一行为。

```sql
SELECT pgv_set('vars', 'int1', 101);
SELECT pgv_get('vars', 'int1', NULL::int);
```

事务性示例：

```sql
BEGIN;
SELECT pgv_set('vars', 'trans_int', 101, true);
SAVEPOINT sp1;
SELECT pgv_set('vars', 'trans_int', 102, true);
ROLLBACK TO sp1;
COMMIT;
SELECT pgv_get('vars', 'trans_int', NULL::int);
```

## Package

变量按 package 分组，因此可以在同一会话中并存多个命名变量，也可以一次性删除整组变量。README 说明空 package 会自动删除。

## 核心函数

### 标量与数组变量

通用 API 如下：

```sql
pgv_set(package text, name text, value anynonarray, is_transactional bool default false)
pgv_get(package text, name text, var_type anynonarray, strict bool default true)

pgv_set(package text, name text, value anyarray, is_transactional bool default false)
pgv_get(package text, name text, var_type anyarray, strict bool default true)
```

`pgv_get()` 会同时检查变量是否存在及其类型。如果 package 或变量缺失，行为取决于 `strict`。

### 记录集合

README 还记录了面向 record 的操作：

- `pgv_insert()`
- `pgv_update()`
- `pgv_delete()`
- `pgv_select()`

这些函数用于操作存放在某个 package 和变量名下的记录集合。

## 已废弃的辅助函数

项目仍保留一些旧的类型专用辅助函数，例如：

- `pgv_set_int()` / `pgv_get_int()`
- `pgv_set_text()` / `pgv_get_text()`
- `pgv_set_numeric()` / `pgv_get_numeric()`
- `pgv_set_timestamp()` / `pgv_get_timestamp()`
- `pgv_set_timestamptz()` / `pgv_get_timestamptz()`
- `pgv_set_date()` / `pgv_get_date()`
- `pgv_set_jsonb()` / `pgv_get_jsonb()`

README 将这些函数标记为已废弃，建议改用通用的 `pgv_set()` / `pgv_get()` API。
