## 用法

来源：

- [官方扩展控制文件](https://github.com/alexandersamoylov/pg_prttn_tools/blob/c7a034b56cd8c2327779dbc0d8060fd2220eb010/pg_prttn_tools.control)
- [官方 0.5 版 SQL](https://github.com/alexandersamoylov/pg_prttn_tools/blob/c7a034b56cd8c2327779dbc0d8060fd2220eb010/sql/pg_prttn_tools--0.5.sql)

`pg_prttn_tools` 使用生成的检查约束、规则与插入触发器管理旧式 PostgreSQL 继承分区。它早于声明式分区，只适合确实需要其准确行为的既有继承分区设计。

### 核心流程

时间辅助函数接受 `year`、`month`、`day`、`hour` 或 `minute`，推导子表名与边界，并可选创建插入路由规则。

```sql
CREATE EXTENSION pg_prttn_tools;

CREATE TABLE public.events (
  id bigint,
  occurred_at timestamp without time zone NOT NULL,
  payload jsonb
);

SELECT *
FROM prttn_tools.part_time_add(
  'public',
  'events',
  'occurred_at',
  'month',
  '2026-07-10 12:00:00'::timestamp,
  true
);
```

### 函数组

- `create_child_table` 使用 `LIKE ... INCLUDING ALL` 克隆父表，添加继承关系、检查约束，并可选添加路由规则。
- `part_list_check`、`part_list_add` 与 `part_list_create_trigger` 处理列表值分区。
- `part_time_check`、`part_time_add` 与 `part_time_create_trigger` 处理时间戳范围。
- `part_list_time_check`、`part_list_time_add` 与 `part_list_time_create_trigger` 组合列表和时间两个维度。
- `drop_ins_trigger`、`part_merge` 与 `part_time_cleanup` 分别删除路由对象、把子表合并回父表，或删除旧的时间命名表。

### 权限与破坏性操作

control 文件要求超级用户，把扩展模式固定为 `prttn_tools`，SQL 还会把函数所有者改为 `postgres`。这些函数通过拼接文本参数构造 DDL，因此只能传入可信的标识符和值。

`part_time_cleanup` 会对指定模式中名称以给定前缀开头、且排序早于计算所得时间戳后缀的每张表执行 `DROP TABLE`。调用前必须确认准确前缀并检查候选表。这些辅助函数创建的是继承子表，而不是原生 `PARTITION BY` 分区；规划器裁剪、约束、索引和路由行为因此遵循旧式继承模型。
