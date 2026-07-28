## 用法

来源：

- [官方上游 README](https://github.com/veer66/chamkho-pg/blob/3f946dc4280dc6dc80f189d3f6f7a35f60ebc9bf/README.org)
- [官方扩展控制文件 (chamkho_parser.control)](https://github.com/veer66/chamkho-pg/blob/3f946dc4280dc6dc80f189d3f6f7a35f60ebc9bf/chamkho_parser.control)
- [官方实现源代码](https://github.com/veer66/chamkho-pg/blob/3f946dc4280dc6dc80f189d3f6f7a35f60ebc9bf/src/lib.rs)

`chamkho_parser` — ~cargo pgrx run~ — 编译并启动带有扩展的 Postgres 实例 ~cargo pgrx test~ — 运行测试套件。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION chamkho_parser;
```

在目标数据库中安装扩展，在有可用上游示例时运行最小的上游示例，并在将其集成到应用程序 SQL 之前验证安装版本和返回值。

### 重要对象

- `chamkho_parser_end` 是一个扩展函数。
- `chamkho_parser_get_token` 是一个扩展函数。
- `chamkho_parser_start` 是一个扩展函数。

### 要求与注意事项

- 该目录记录版本为 `0.6.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
