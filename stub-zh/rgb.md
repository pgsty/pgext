## 用法

来源：

- [官方上游 README](https://github.com/nixgeekdev/pg_rgb/blob/0674066b0bec753a4c6b3e1c903cf74400ea809b/README.md)
- [官方扩展控制文件 (rgb.control)](https://github.com/nixgeekdev/pg_rgb/blob/0674066b0bec753a4c6b3e1c903cf74400ea809b/rgb.control)
- [官方实现源代码](https://github.com/nixgeekdev/pg_rgb/blob/0674066b0bec753a4c6b3e1c903cf74400ea809b/src/rgb.c)

`rgb` — 该库包含一个单一的 PostgreSQL 扩展，即 RGB 颜色数据类型，以及用于构建、转换和索引 RGB 颜色的便利函数。当应用程序数据需要此类型、域或其操作符时，请使用此扩展。在目标 PostgreSQL 构建中测试链接的上游修订版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION rgb;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以匹配固定源代码。
