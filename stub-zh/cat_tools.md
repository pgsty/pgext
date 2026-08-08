## 用法

来源：

- [cat_tools 0.3.0 README](https://github.com/Postgres-Extensions/cat_tools/blob/0.3.0/README.asc)
- [cat_tools 0.3.0 历史记录](https://github.com/Postgres-Extensions/cat_tools/blob/0.3.0/HISTORY.asc)
- [cat_tools 0.3.0 控制文件](https://github.com/Postgres-Extensions/cat_tools/blob/0.3.0/cat_tools.control)
- [cat_tools 0.3.0 安装 SQL](https://github.com/Postgres-Extensions/cat_tools/blob/0.3.0/sql/cat_tools--0.3.0.sql.in)

`cat_tools` 提供用于 PostgreSQL 目录自省的带类型视图、枚举和辅助函数。它面向需要比反复解析原始 `pg_catalog` 字段更稳定、更易读接口的数据库代码；这些视图仍会跟随 PostgreSQL 目录变化，因此每次跨大版本升级时都必须审查。

### 安装并授予访问权限

```sql
CREATE EXTENSION cat_tools;
GRANT cat_tools__usage TO app_introspection;
```

扩展安装在固定的 `cat_tools` 模式中，要求 `plpgsql`，且不可重定位。应授予 `cat_tools__usage` 角色，而不是直接暴露内部 `_cat_tools` 辅助对象。

### 检查关系与列

```sql
SELECT cat_tools.relation__kind(c.relkind::text)
FROM pg_catalog.pg_class AS c
WHERE c.oid = 'public.orders'::regclass;

SELECT cat_tools.relation__column_names('public.orders'::regclass);
SELECT cat_tools.pg_attribute__get('public.orders'::regclass, 'id');
```

常用的关系辅助函数包括 `pg_class(regclass)`、`relation__is_catalog`、`relation__is_temp`、`relation__kind` 和 `relation__relkind`。带类型的映射函数能明确表示目录中的单字符代码。

### 检查例程

版本 0.3 新增了同时覆盖函数和过程的函数与类型：

```sql
SELECT cat_tools.routine__arg_types(
  'public.calculate_total(integer, numeric)'::regprocedure
);

SELECT cat_tools.routine__parse_arg_names(
  'IN account_id integer, INOUT total numeric'
);
```

例程接口包括 `routine__parse_arg_types`、`routine__parse_arg_names`、`routine__arg_types`、`routine__arg_names`、它们的文本变体，以及用于例程种类、参数模式、易变性和并行安全性的映射。`function__arg_types` 与 `function__arg_types_text` 已弃用；请改用例程解析器。

### 版本 0.3.0 与注意事项

- 上游版本 0.3.0 支持 PostgreSQL 12-18+；当前 Pigsty 软件包覆盖 PostgreSQL 14-18。
- 该版本修正了复合类型、外部表和物化视图对应的 `c`、`f`、`m` 映射。任何曾绕过旧映射问题的代码都应重新测试。
- 内部 `_cat_tools` 辅助对象现在会撤销 `PUBLIC` 的 `EXECUTE`；调用者应继承 `cat_tools__usage` 并使用受支持的接口。
- 从 0.2.3 更新至 0.3.0 会新增枚举值，因此无法在 PostgreSQL 11 或更早版本上运行。请按照上游文档规定的顺序升级数据库大版本和扩展。
- PostgreSQL 不承诺跨大版本的目录兼容性。即使使用这些包装器，也应针对每个受支持的 PostgreSQL 大版本固定测试。
