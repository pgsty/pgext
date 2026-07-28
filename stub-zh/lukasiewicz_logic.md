## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/fuzzy_logic/fuzzy_logic-1.1.0/README.md)
- [官方扩展控制文件 (lukasiewicz_logic.control)](https://api.pgxn.org/src/fuzzy_logic/fuzzy_logic-1.1.0/lukasiewicz_logic.control)
- [官方扩展 SQL (lukasiewicz_logic--1.1.0.sql)](https://api.pgxn.org/src/fuzzy_logic/fuzzy_logic-1.1.0/sql/lukasiewicz_logic--1.1.0.sql)

`lukasiewicz_logic` — 模糊逻辑 =========== 该扩展提供了基本的逻辑运算符（合取、析取、蕴含和否定）用于三种基本的模糊逻辑 - Łukasiewicz、Gödel 和产品逻辑。对于 Łukasiewicz 逻辑，还提供了弱合取和弱析取运算符。当 SQL 需要这些特殊函数或聚合时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION lukasiewicz_logic;
```

在目标数据库中安装该扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `lukasiewicz_conjunction(a fuzzy_boolean, b fuzzy_boolean)` 是一个扩展函数，返回 `fuzzy_boolean`。
- `lukasiewicz_disjunction(a fuzzy_boolean, b fuzzy_boolean)` 是一个扩展函数，返回 `fuzzy_boolean`。
- `lukasiewicz_negation(a fuzzy_boolean)` 是一个扩展函数，返回 `fuzzy_boolean`。
- `lukasiewicz_residuum(a fuzzy_boolean, b fuzzy_boolean)` 是一个扩展函数，返回 `fuzzy_boolean`。
- `lukasiewicz_weak_conjunction(a fuzzy_boolean, b fuzzy_boolean)` 是一个扩展函数，返回 `fuzzy_boolean`。
- `lukasiewicz_weak_disjunction(a fuzzy_boolean, b fuzzy_boolean)` 是一个扩展函数，返回 `fuzzy_boolean`。
- `fuzzy_boolean` 是一个扩展定义的域。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.1.0`。
- 控制文件将该扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
