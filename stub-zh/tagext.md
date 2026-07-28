## 用法

来源：

- [官方上游 README](https://github.com/packtpublishing/learn-postgresql-second-edition/blob/9590a7604ce6656fcc6f888830a71e082c794d64/README.md)
- [官方扩展控制文件 (tagext.control)](https://github.com/packtpublishing/learn-postgresql-second-edition/blob/9590a7604ce6656fcc6f888830a71e082c794d64/CHAPTER_12/tagext/tagext.control)
- [官方扩展 SQL (tagext--1.0--1.1.sql)](https://github.com/packtpublishing/learn-postgresql-second-edition/blob/9590a7604ce6656fcc6f888830a71e082c794d64/CHAPTER_12/tagext/tagext--1.0--1.1.sql)

`tagext` — 标签编程示例扩展。当 SQL 需要这些特殊函数或聚合时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION tagext;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `tag_path(tag_to_search text)` 是一个扩展函数，返回 `TEXT`。
- `tag_path(tag_to_search text, delimiter text DEFAULT ' > ')` 是一个扩展函数，返回 `TEXT`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况与固定源代码中的信息一致。
