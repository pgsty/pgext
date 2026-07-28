## 用法

来源：

- [Official database.dev 包页面](https://database.dev/bobbie/hello)

`bobbie@hello` — 一个生成问候语的扩展。当 SQL 需要这些特殊的功能或聚合时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION "bobbie@hello";
```

在目标数据库中安装该扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `greet(name text default 'world')` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 表记录版本 `0.0.1`。
- 这是一个 database.dev/pg_tle 包；在发出引用的 `CREATE EXTENSION` 之前，请先注册或生成其包迁移。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
