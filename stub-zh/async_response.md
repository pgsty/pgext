## 用法

来源：

- [官方上游 README](https://github.com/arturformella/async_response/blob/24a278e4e6eb244009efa96c1859f359a028cf70/README.md)
- [官方扩展控制文件 (async_response.control)](https://github.com/arturformella/async_response/blob/24a278e4e6eb244009efa96c1859f359a028cf70/async_response.control)
- [官方扩展 SQL (async_response--1.0.sql)](https://github.com/arturformella/async_response/blob/24a278e4e6eb244009efa96c1859f359a028cf70/async_response--1.0.sql)

`async_response` — boolean async_response(port INTEGER, channel TEXT, aspect TEXT, data TEXT) 可以在复杂的查询中立即发送到 REDIS 从 PostgreSQL。使用它来执行相应的 SQL 或数据库实用程序工作流。上游审查材料已将此功能标记为弃用。

### 核心工作流

```sql
CREATE EXTENSION async_response;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `async_response(INTEGER,TEXT,TEXT,TEXT)` 是一个扩展函数，返回 `boolean`。
- `async_response(TEXT,TEXT,TEXT,TEXT)` 是一个扩展函数，返回 `boolean`。
- `async_response_async(INTEGER,TEXT,TEXT,TEXT)` 是一个扩展函数，返回 `boolean`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 控制文件要求超级用户进行安装。
- 上游材料包含显式的弃用边界。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
