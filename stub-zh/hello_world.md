## 用法

来源：

- [官方上游 README](https://github.com/andrelandgraf/russian-roulette-hello-world-postgres-extension/blob/0ab27d4357ce50a66d930c08ed5f16ba6dfbbd1b/README.md)
- [官方扩展控制文件 (hello_world.control)](https://github.com/andrelandgraf/russian-roulette-hello-world-postgres-extension/blob/0ab27d4357ce50a66d930c08ed5f16ba6dfbbd1b/hello_world.control)
- [官方扩展 SQL (hello_world--1.0.sql)](https://github.com/andrelandgraf/russian-roulette-hello-world-postgres-extension/blob/0ab27d4357ce50a66d930c08ed5f16ba6dfbbd1b/hello_world--1.0.sql)

`hello_world` — 一个简单的 PostgreSQL 扩展，仅使用 SQL 实现基本的扩展功能（无需编译 C 代码）。当需要这些特殊函数或聚合时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION hello_world;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `hello_world()` 是一个扩展函数，返回 `TEXT`。
- `hello_world(name TEXT)` 是一个扩展函数，返回 `TEXT`。
- `log_hello_world(message TEXT DEFAULT 'Hello, World!')` 是一个扩展函数，返回 `INTEGER`。
- `russian_roulette(target_user_id INTEGER, odds INTEGER DEFAULT 6)` 是一个扩展函数，返回 `TEXT`。
- `hello_world_log` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况与固定源进行比对。
