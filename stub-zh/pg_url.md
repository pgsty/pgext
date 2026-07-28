## 用法

来源：

- [官方上游 README](https://github.com/semtexzv/pg_url/blob/2741caa7c7ed2a4435253f0d998aac306af7f2da/README.md)
- [官方扩展控制文件 (pg_url.control)](https://github.com/semtexzv/pg_url/blob/2741caa7c7ed2a4435253f0d998aac306af7f2da/pg_url.control)
- [官方实现源代码](https://github.com/semtexzv/pg_url/blob/2741caa7c7ed2a4435253f0d998aac306af7f2da/src/lib.rs)

`pg_url` — 实现 URL 操作方法作为 postgres 扩展。使用此扩展可以在数据库内部直接处理 URL。可以使用它在 URL 主机上创建索引。当 SQL 需要这些特殊函数或聚合时，请使用之。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_url;

CREATE INDEX tbl_url_host on tbl (url_host(url));
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `url_clear_host` 是一个扩展函数。
- `url_fragment` 是一个扩展函数。
- `url_host` 是一个扩展函数。
- `url_path` 是一个扩展函数。
- `url_query` 是一个扩展函数。
- `url_query_param` 是一个扩展函数。
- `url_scheme` 是一个扩展函数。
- `url_set_fragment` 是一个扩展函数。
- `url_set_host` 是一个扩展函数。
- `url_set_path` 是一个扩展函数。
- `url_set_query` 是一个扩展函数。
- `url_set_query_param` 是一个扩展函数。
- `url_set_scheme` 是一个扩展函数。

### 要求与注意事项

- 该目录记录版本 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的内容一致。
