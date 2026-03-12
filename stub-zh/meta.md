


## 用法

> [meta: PostgreSQL 简化系统目录](https://github.com/aquameta/meta)

meta 提供一套规范化、易于理解的系统目录视图，使用通用的视图名和列名，构建于 `pg_catalog` 和 `information_schema` 之上。

### 系统目录视图

该扩展在 `meta` 模式下创建约 30 个视图：

```sql
-- 列出所有表
SELECT * FROM meta.table;

-- 列出所有列
SELECT * FROM meta.column;

-- 列出所有视图
SELECT * FROM meta.view;

-- 列出模式
SELECT * FROM meta.schema;

-- 列出函数
SELECT * FROM meta.function;

-- 列出扩展
SELECT * FROM meta.extension;

-- 列出触发器
SELECT * FROM meta.trigger;

-- 列出外键
SELECT * FROM meta.foreign_key;

-- 列出约束
SELECT * FROM meta.constraint_check;
SELECT * FROM meta.constraint_unique;

-- 列出类型
SELECT * FROM meta.type;

-- 列出角色
SELECT * FROM meta.role;

-- 列出序列
SELECT * FROM meta.sequence;

-- 列出运算符
SELECT * FROM meta.operator;

-- 列出策略
SELECT * FROM meta.policy;
```

### 可用视图

`cast`、`column`、`connection`、`constraint_check`、`constraint_unique`、`extension`、`foreign_column`、`foreign_data_wrapper`、`foreign_key`、`foreign_server`、`foreign_table`、`function`、`function_parameter`、`operator`、`policy`、`policy_role`、`relation`、`relation_column`、`role`、`role_inheritance`、`schema`、`sequence`、`table`、`table_privilege`、`trigger`、`type`、`view`

### 元标识符

该扩展提供复合类型作为"软"主键，通过名称（而非 OID）来标识 PostgreSQL 对象（表、列、类型等）。

### 目录触发器（可选）

通过可选的 `meta_triggers` 扩展，视图变为可更新，从而支持通过 DML 执行 DDL：

```sql
-- 通过 INSERT 创建表
INSERT INTO meta.table (schema_name, name) VALUES ('public', 'foo');
```
