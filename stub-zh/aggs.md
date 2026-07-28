## 用法

来源：

- [官方上游 README](https://github.com/mkindahl/pg_examples/blob/277a29ac1b31478911c75ddd886d4dea02730aa8/README.md)
- [官方扩展控制文件 (aggs.control)](https://github.com/mkindahl/pg_examples/blob/277a29ac1b31478911c75ddd886d4dea02730aa8/aggs/aggs.control)
- [官方扩展 SQL (aggs--0.1.sql)](https://github.com/mkindahl/pg_examples/blob/277a29ac1b31478911c75ddd886d4dea02730aa8/aggs/aggs--0.1.sql)

`aggs` — 测试聚合函数。当 SQL 需要这些特殊功能或聚合时使用它。在目标 PostgreSQL 构建上使用上述链接的上游固定版本作为 API 边界并进行测试。

### 核心工作流

```sql
CREATE EXTENSION aggs;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `window_agg_dropfn(internal, anynonarray)` 是一个扩展函数，返回 `internal`。
- `window_agg_dropfn(internal, anynonarray, integer)` 是一个扩展函数，返回 `internal`。
- `window_agg_finalfn(internal, anynonarray)` 是一个扩展函数，返回 `anyarray`。
- `window_agg_finalfn(internal, anynonarray, integer)` 是一个扩展函数，返回 `anyarray`。
- `window_agg_transfn(internal, anynonarray)` 是一个扩展函数，返回 `internal`。
- `window_agg_transfn(internal, anynonarray, integer)` 是一个扩展函数，返回 `internal`。
- `window_agg` 是由扩展公开的聚合函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1`。
- 控制文件标记该扩展为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
