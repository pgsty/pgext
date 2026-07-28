## 用法

来源：

- [官方上游 README](https://github.com/spa5k/slugify-postgres/blob/54c6a9eef1d30b3434e32e30404d65ae5f91a440/README.md)
- [官方扩展控制文件 (slugify.control)](https://github.com/spa5k/slugify-postgres/blob/54c6a9eef1d30b3434e32e30404d65ae5f91a440/slugify.control)
- [官方实现源代码](https://github.com/spa5k/slugify-postgres/blob/54c6a9eef1d30b3434e32e30404d65ae5f91a440/src/lib.rs)

`slugify` — PostgreSQL 扩展，用于从字符串生成各种变体的 Slug。当 SQL 需要这些特殊函数或聚合时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION slugify;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `slug` 是一个扩展函数。
- `slug_rand` 是一个扩展函数。
- `slug_rand_c` 是一个扩展函数。
- `slug_rand_sep` 是一个扩展函数。
- `slug_rand_sep_c` 是一个扩展函数。
- `slug_sep` 是一个扩展函数。

### 要求与注意事项

- 该目录记录版本为 `0.0.1`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
