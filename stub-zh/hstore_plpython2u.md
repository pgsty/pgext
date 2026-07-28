## 用法

来源：

- [官方上游 README](https://github.com/skyrise-l/queryartisan/blob/29613aa7bbfd5108a64ac4b5115bc3f94310985b/optimization/contrib/README)
- [官方扩展控制文件 (hstore_plpython2u.control)](https://github.com/skyrise-l/queryartisan/blob/29613aa7bbfd5108a64ac4b5115bc3f94310985b/optimization/contrib/hstore_plpython/hstore_plpython2u.control)
- [官方扩展 SQL (hstore_plpython2u--1.0.sql)](https://github.com/skyrise-l/queryartisan/blob/29613aa7bbfd5108a64ac4b5115bc3f94310985b/optimization/contrib/hstore_plpython/hstore_plpython2u--1.0.sql)

`hstore_plpython2u` — 在数据库代码需要在或与此过程语言进行交互时，用于在 hstore 和 plpython2u 之间进行转换。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION hstore_plpython2u;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `hstore_to_plpython2(val internal)` 是一个扩展函数，返回 `internal`。
- `plpython2_to_hstore(val internal)` 是一个扩展函数，返回 `hstore`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 先安装确认的扩展依赖项：`hstore`, `plpython2u`。
- 控制文件将扩展标记为可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，需确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
