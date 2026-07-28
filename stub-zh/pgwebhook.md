## 用法

来源：

- [官方上游 README](https://github.com/mansueli/tle/blob/9e1c8ce5a7d7a2f334eb21894d3e8994fb7f28ad/pgwebhook/README.md)
- [官方扩展控制文件 (pgwebhook.control)](https://github.com/mansueli/tle/blob/9e1c8ce5a7d7a2f334eb21894d3e8994fb7f28ad/pgwebhook/pgwebhook.control)
- [官方扩展 SQL (pgwebhook--0.1.1--0.1.2.sql)](https://github.com/mansueli/tle/blob/9e1c8ce5a7d7a2f334eb21894d3e8994fb7f28ad/pgwebhook/pgwebhook--0.1.1--0.1.2.sql)

`pgwebhook` — pgwebhook 是一个用于从数据库中创建和管理 webhook 的 PostgreSQL 扩展。当应用程序需要此特定的数据库功能时，请使用它。在安装扩展之前，必须先安装并验证其依赖项。

### 核心工作流

```sql
CREATE EXTENSION pgwebhook;

SELECT dbdev.install('mansueli@pgwebhook');
CREATE EXTENSION "mansueli@pgwebhook" VERSION '0.1.1';
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `hook.edge_wrapper` 是一个扩展函数。
- `hook.edgehook_trigger()` 是一个扩展函数并返回 `trigger`。
- `hook.http_request` 是一个扩展函数。
- `hook.webhook_trigger()` 是一个扩展函数并返回 `trigger`。
- `hook.migrations` 是一个由扩展安装或管理的表。
- `hook` 是一个由扩展创建的模式。

### 要求与注意事项

- 控制文件声明默认版本为 `0.1.2`。
- 请先安装并验证确认的扩展依赖项：`http`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
