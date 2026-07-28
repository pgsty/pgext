## 用法

来源：

- [官方上游 README](https://github.com/thomasweiser/mediatum-view/blob/27310339c24cf204325dc9d3ecdd2efdb539236d/backend/pg-extensions/snowball_bilingual/README.md)
- [官方扩展控制文件 (snowball_bilingual.control)](https://github.com/thomasweiser/mediatum-view/blob/27310339c24cf204325dc9d3ecdd2efdb539236d/backend/pg-extensions/snowball_bilingual/snowball_bilingual.control)
- [官方扩展 SQL (snowball_bilingual--1.0.sql)](https://github.com/thomasweiser/mediatum-view/blob/27310339c24cf204325dc9d3ecdd2efdb539236d/backend/pg-extensions/snowball_bilingual/snowball_bilingual--1.0.sql)

`snowball_bilingual` — snowball_bilingual 扩展提供了一个新的词典模板。它是 PostgreSQL Snowball 词典模板的一个副本。目前它为以下语言提供了词干提取算法：尼泊尔语。使用它来进行相应的文本搜索、解析或语言工作流。在目标 PostgreSQL 构建中使用上面链接的固定上游版本作为 API 边界，并对其进行测试。

### 核心工作流

```sql
CREATE EXTENSION snowball_bilingual;
```

在目标数据库中安装该扩展，当可用时运行上面的最小上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `dsnowball_bilingual_init(INTERNAL)` 是一个扩展函数并返回 `INTERNAL`。
- `dsnowball_bilingual_lexize(INTERNAL, INTERNAL, INTERNAL, INTERNAL)` 是一个扩展函数并返回 `INTERNAL`。

### 要求与注意事项

- 审核后的控制文件声明默认版本为 `1.0`。
- 控制文件将该扩展标记为可重定位。
- 在生产使用之前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
