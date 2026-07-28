## 用法

来源：

- [官方上游 README](https://github.com/jonatas/pgscraper/blob/f17ddc9de0f00e9308948bb188c0416264c08e25/README.md)
- [官方扩展控制文件 (pgscraper.control)](https://github.com/jonatas/pgscraper/blob/f17ddc9de0f00e9308948bb188c0416264c08e25/pgscraper.control)
- [官方实现源代码](https://github.com/jonatas/pgscraper/blob/f17ddc9de0f00e9308948bb188c0416264c08e25/src/lib.rs)

`pgscraper` — 这是一个小型扩展，允许您直接从您的 PostgreSQL 数据库抓取数据。在将相应数据从 PostgreSQL 移动、转换或集成时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgscraper;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `html_select` 是一个扩展函数。
- `html_select_text` 是一个扩展函数。
- `http` 是一个扩展函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
