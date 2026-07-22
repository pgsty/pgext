## 用法

来源：

- [官方扩展控制文件](https://github.com/katrinaibast/constr_name_unif/blob/e69b3a79fcaa4d013e648b812807a607bd14b7ff/constr_name_unif.control)
- [官方上游文档](https://github.com/katrinaibast/constr_name_unif/blob/e69b3a79fcaa4d013e648b812807a607bd14b7ff/doc/constr_name_unif.md)

`constr_name_unif` 可按可配置的规则盘点并重命名 PostgreSQL 约束。它适合统一现有基表的约束名称，但每次重命名都是模式变更 DDL，可能影响引用约束名称的迁移和工具。

### 核心流程

安装这个纯 SQL 扩展，检查约束与可用规则，然后再执行所需的重命名函数：

```sql
CREATE EXTENSION constr_name_unif;

SELECT * FROM get_all_constraints_in_schema('public');
SELECT * FROM get_all_patterns();

BEGIN;
SELECT rename_all_constraints_in_schema(
  'public',
  'snake_case_with_short_prefix'
);
COMMIT;
```

内置 `postgresql_default` 规则使用 `pkey`、`fkey` 等后缀；内置 `snake_case_with_short_prefix` 规则使用 `pk`、`fk` 等前缀。自定义规则可以设置分隔符、前缀或后缀位置，以及各类约束的缩写。

### 重要对象

- `get_all_constraints()` 与 `get_all_constraints_in_schema(text)` 返回当前约束清单；重载版本可按约束类型筛选。
- `get_all_patterns()` 列出已配置的命名规则。
- `add_pattern(...)` 与 `delete_pattern(...)` 管理命名规则。
- `add_to_abbreviated_tables(...)`、`empty_abbreviated_tables()` 与 `get_abbreviated_tables()` 管理表名缩写。
- `rename_all_constraints(text)` 及其按模式或类型限定的变体，会为主键、外键、检查、唯一和排他约束执行 `ALTER TABLE ... RENAME CONSTRAINT` 语句。

### 运维说明

生成的标识符会限制在 PostgreSQL 的 63 字节标识符上限内，并处理名称冲突；外键名称也会包含被引用表。该扩展没有能列出全部拟议目标名称的预演 API，因此应先保存清单，在副本上测试规则，再以具备适当 DDL 权限的受控事务执行重命名。提交前检查应用迁移、监控和管理脚本中是否硬编码了约束名称。
