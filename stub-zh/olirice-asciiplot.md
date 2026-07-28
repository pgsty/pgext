## 用法

来源：

- [Official database.dev 包页面](https://database.dev/olirice/asciiplot)

`olirice-asciiplot` — 一个玩具 ASCII 绘图库。用于相应的 SQL 或数据库实用程序工作流。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION "olirice-asciiplot";
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `scatter_internal(state scatter_state)` 是一个扩展函数，返回 `TEXT`。
- `scatter_sfunc(state scatter_state, x numeric, y numeric, title TEXT, height INTEGER, width INTEGER)` 是一个扩展函数，返回 `scatter_state`。
- `scatter` 是由扩展公开的聚合函数。
- `scatter_state` 是一个由扩展定义的类型。

### 要求与注意事项

- 该目录记录版本为 `0.0.1`。
- 这是一个 database.dev/pg_tle 包；在发出引用 `CREATE EXTENSION` 的身份之前，请先注册或生成其包迁移。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行验证。
