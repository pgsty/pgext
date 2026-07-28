## 用法

来源：

- [Official database.dev 包页面](https://database.dev/jessevent/supa_profile)

`jessevent@supa_profile` — 数据库表和列统计分析器。在收集或解释相应的 PostgreSQL 统计信息时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION "jessevent@supa_profile";
```

在目标数据库中安装扩展，当可用时运行上方的最小上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `infer_pattern` 是一个扩展函数。
- `profile_table` 是一个扩展函数。

### 要求与注意事项

- 表册记录版本 `1.0.0`。
- 这是一个 database.dev/pg_tle 包；在发出引用的 `CREATE EXTENSION` 之前，请先注册或生成其包迁移。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
