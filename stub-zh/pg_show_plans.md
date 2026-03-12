


## 用法

> [pg_show_plans: 显示正在运行的 SQL 语句的查询计划](https://github.com/cybertec-postgresql/pg_show_plans)

pg_show_plans 实时显示所有正在运行的 SQL 语句的执行计划。计划存储在共享内存哈希表中。

### 查看运行中的计划

```sql
-- 查看所有正在运行的查询的计划
SELECT * FROM pg_show_plans;

-- 查看计划及其查询文本
SELECT * FROM pg_show_plans_q;
```

视图返回以下列：

| 列名 | 类型 | 描述 |
|--------|------|-------------|
| `pid` | integer | 服务器进程 ID |
| `level` | integer | 查询嵌套层级（0 = 顶层） |
| `userid` | oid | 用户 OID |
| `dbid` | oid | 数据库 OID |
| `plan` | text | 查询计划文本 |
| `query` | text | 查询字符串（仅在 `pg_show_plans_q` 中） |

### 嵌套查询

当函数调用 SQL 语句时，它们会出现在更深的嵌套层级中：

```sql
SELECT * FROM pg_show_plans;
  pid  | level | userid | dbid  |                       plan
-------+-------+--------+-------+-----------------------------------------------
 11504 |     0 |     10 | 16384 | Function Scan on print_item  (cost=...)
 11504 |     1 |     10 | 16384 | Result  (cost=0.00..0.01 rows=1 width=4)
```

### GUC 参数

| 参数 | 默认值 | 描述 |
|-----------|---------|-------------|
| `pg_show_plans.plan_format` | `text` | 输出格式：`text`、`json`、`yaml`、`xml` |
| `pg_show_plans.max_plan_length` | 16384 | 最大计划长度（字节，影响共享内存） |
| `pg_show_plans.is_enabled` | `true` | 在运行时启用或禁用扩展 |
