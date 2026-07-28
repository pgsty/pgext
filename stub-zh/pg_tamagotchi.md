## 用法

来源：

- [官方上游 README](https://github.com/hill/pg_tamagotchi/blob/bf37d7bd2fb6d269a3f7a8202c73bf3b0ba4bb15/README.md)
- [官方扩展控制文件 (pg_tamagotchi.control)](https://github.com/hill/pg_tamagotchi/blob/bf37d7bd2fb6d269a3f7a8202c73bf3b0ba4bb15/pg_tamagotchi.control)
- [官方扩展 SQL (pg_tamagotchi--0.1.0.sql)](https://github.com/hill/pg_tamagotchi/blob/bf37d7bd2fb6d269a3f7a8202c73bf3b0ba4bb15/pg_tamagotchi--0.1.0.sql)

`pg_tamagotchi` — 一个生活在你的 Postgres 数据库中的 tamagotchi。当应用程序需要此特定数据库功能时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_tamagotchi;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `feed(food text DEFAULT NULL)` 是一个扩展函数，返回 `text`。
- `hatch(name text DEFAULT NULL)` 是一个扩展函数，返回 `text`。
- `status()` 是一个扩展函数，返回 `text`。
- `talk(message text DEFAULT NULL)` 是一个扩展函数，返回 `text`。
- `vitals` 是一个由扩展定义的视图。
- `message` 是一个由扩展安装或管理的表。
- `pet` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行对比。
