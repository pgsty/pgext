## 用法

来源：

- [官方上游 README](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/extensions/endpoint/README.md)
- [官方扩展控制文件 (endpoint.control)](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/extensions/endpoint/endpoint.control)

`endpoint` — 此扩展自身不强制任何安全约束。在应用程序需要此特定数据库功能时使用它。上游将其描述为一个正在进行中的项目。

### 核心工作流

```sql
CREATE EXTENSION endpoint;

select endpoint.request(
    '0.3',                              -- version
	'GET',                              -- method
	'/endpoint/0.3/row/{meta.row_id}',  -- path
	'{"key": "val"}',                   -- query string as JSON
	'{"key": "val"}'                    -- post args as JSON
);
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.5.0`。
- 首先安装确认的扩展依赖项：`meta`。
- 控制文件标记此扩展为不可重定位。
- 上游将该项目描述为一个正在进行中的项目。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
