## 用法

来源：

- [官方上游 README](https://github.com/eatonphil/pgweb/blob/0944c98ac7b09741830920db413d40cf633538f2/README.md)
- [官方扩展控制文件 (pgweb.control)](https://github.com/eatonphil/pgweb/blob/0944c98ac7b09741830920db413d40cf633538f2/pgweb.control)
- [官方扩展 SQL (pgweb--0.0.1.sql)](https://github.com/eatonphil/pgweb/blob/0944c98ac7b09741830920db413d40cf633538f2/pgweb--0.0.1.sql)

`pgweb` — 创建一个新的 Postgres 数据库并运行测试脚本来定义一个 web 服务并启动服务器：当应用程序需要此特定数据库功能时使用它。在目标 PostgreSQL 构建中使用上述链接的上游固定版本作为 API 边界并进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgweb;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `handle_hello_world(params JSON)` 是一个扩展函数，返回 `TEXT`。
- `pgweb.register_get(TEXT, TEXT)` 是一个扩展函数，返回 `VOID`。
- `pgweb.serve(TEXT, INT)` 是一个扩展函数，返回 `VOID`。
- `pgweb` 是由扩展创建的一个模式。

### 要求与注意事项

- 控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况与固定源进行比对。
