## 用法

来源：

- [Official database.dev 包页面](https://database.dev/cem/uuidv7)

`cem-uuidv7` — UUIDv7 扩展。当 SQL 需要这些特殊函数或聚合时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION "cem-uuidv7";
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `uuid7_from_timestamp(p_timestamp timestamp with time zone)` 是一个扩展函数，返回 `uuid`。
- `uuid_generate_v7()` 是一个扩展函数，返回 `uuid`。

### 要求与注意事项

- 该目录记录版本 `1.0.2`。
- 这是一个 database.dev/pg_tle 包；在发出引用的 `CREATE EXTENSION` 之前，请先注册或生成其包迁移。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
