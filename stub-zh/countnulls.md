## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/countnulls/countnulls-1.0.0/README)
- [官方扩展 SQL (countnulls.sql)](https://api.pgxn.org/src/countnulls/countnulls-1.0.0/countnulls.sql)

`countnulls` — 一个简单的函数用于计算 NULL 参数的数量。当 SQL 需要这些特殊函数或聚合时使用它。请使用上述固定上游版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

经过审查的分发包使用了过时的 SQL 或非控制文件安装布局，因此没有建立现代独立的 `CREATE EXTENSION` 和升级工作流。请遵循固定上游的安装机制，并在隔离数据库中验证安装的对象。

### 重要对象

- `countnulls("any")` 是一个扩展函数并返回 `int`。
- `countnulls("any", "any")` 是一个扩展函数并返回 `int`。
- `countnulls("any", "any", "any")` 是一个扩展函数并返回 `int`。
- `countnulls("any", "any", "any", "any")` 是一个扩展函数并返回 `int`。
- `countnulls("any", "any", "any", "any", "any")` 是一个扩展函数并返回 `int`。
- `countnulls("any", "any", "any", "any", "any", "any")` 是一个扩展函数并返回 `int`。
- `countnulls("any", "any", "any", "any", "any", "any", "any")` 是一个扩展函数并返回 `int`。
- `countnulls("any", "any", "any", "any", "any", "any", "any", "any")` 是一个扩展函数并返回 `int`。
- `countnulls(VARIADIC "any")` 是一个扩展函数并返回 `int`。

### 要求与注意事项

- 请确认版本记录 `1.0.0`。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以匹配固定来源。
