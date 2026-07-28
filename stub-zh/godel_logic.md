## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/fuzzy_logic/fuzzy_logic-1.1.0/README.md)
- [官方扩展控制文件 (godel_logic.control)](https://api.pgxn.org/src/fuzzy_logic/fuzzy_logic-1.1.0/godel_logic.control)
- [官方扩展 SQL (godel_logic--1.1.0.sql)](https://api.pgxn.org/src/fuzzy_logic/fuzzy_logic-1.1.0/sql/godel_logic--1.1.0.sql)

`godel_logic` 属于模糊逻辑扩展，为三种基本模糊逻辑——Łukasiewicz、Gödel 和乘积逻辑——提供合取、析取、蕴含与否定运算符；其中 Łukasiewicz 逻辑还提供弱合取和弱析取运算符。当 SQL 需要这些专用函数或聚合时可使用它。请以上方链接的固定上游修订为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION godel_logic;
```

在目标数据库中安装扩展；如果上游提供了最小示例，请运行该示例，并在集成到应用 SQL 前验证安装版本和返回值。

### 重要对象

- `godel_conjunction(a fuzzy_boolean, b fuzzy_boolean)` 是扩展函数，返回 `fuzzy_boolean`。
- `godel_disjunction(a fuzzy_boolean, b fuzzy_boolean)` 是扩展函数，返回 `fuzzy_boolean`。
- `godel_negation(a fuzzy_boolean)` 是扩展函数，返回 `fuzzy_boolean`。
- `godel_residuum(a fuzzy_boolean, b fuzzy_boolean)` 是扩展函数，返回 `fuzzy_boolean`。
- `fuzzy_boolean` 是扩展定义的域。

### 要求与注意事项

- 审阅的控制文件声明默认版本为 `1.1.0`。
- 控制文件将该扩展标记为可重定位。
- 生产使用前，请根据固定版本源码确认权限、支持的 PostgreSQL 版本、升级行为和失败情形。
