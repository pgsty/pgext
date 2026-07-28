## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/italian_fts/italian_fts-1.2.1/README.rst)
- [官方扩展控制文件 (italian_fts.control)](https://api.pgxn.org/src/italian_fts/italian_fts-1.2.1/italian_fts.control)
- [官方扩展 SQL (italian_fts.sql)](https://api.pgxn.org/src/italian_fts/italian_fts-1.2.1/italian_fts.sql)

`italian_fts` — 该包可用于在 PostgreSQL 8.3 及以上版本中安装和配置 ISpell 字典。使用它来实现相应的文本搜索、解析或语言工作流。请使用上述链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION italian_fts;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.2`。
- 控制文件将扩展标记为可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行对比。
