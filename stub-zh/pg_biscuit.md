## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_biscuit/pg_biscuit-1.0.0/README.md)
- [官方扩展控制文件 (pg_biscuit.control)](https://api.pgxn.org/src/pg_biscuit/pg_biscuit-1.0.0/pg_biscuit.control)
- [官方扩展 SQL (pg_biscuit--1.0.sql)](https://api.pgxn.org/src/pg_biscuit/pg_biscuit-1.0.0/sql/pg_biscuit--1.0.sql)
- [当前上游继续开发仓库](https://github.com/CrystallineCore/Biscuit)

`pg_biscuit` — 一种 PostgreSQL 索引访问方法（IAM），用于高性能的文本列模式匹配。Biscuit 索引特别设计用于加速带有任意通配符的 LIKE 查询。请在相应的文本搜索、解析或语言工作流中使用它。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION pg_biscuit;
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `biscuit_handler(internal)` 是一个扩展函数，返回 `index_am_handler`。
- `biscuit_index_stats(oid)` 是一个扩展函数，返回 `text`。
- `biscuit_like_support(internal)` 是一个扩展函数。
- `biscuit_multicolumn_enabled()` 是一个扩展函数，返回 `boolean`。
- `biscuit_indexes` 是一个扩展定义的视图。
- `biscuit_indexes_detailed` 是一个扩展定义的视图。
- `biscuit_version` 是一个由扩展安装或管理的表。
- `products` 是一个由扩展安装或管理的表。
- `biscuit` 是一个扩展定义的访问方法。
- `biscuit_bpchar_ops` 是一个扩展定义的操作类。
- `biscuit_text_ops` 是一个扩展定义的操作类。
- `biscuit_varchar_ops` 是一个扩展定义的操作类。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `1.0`。
- 先安装确认的扩展依赖项：`plpgsql`。
- 控制文件将扩展标记为不可重定位。
- 上述 PGXN 源代码记录了 `pg_biscuit` 的身份；当前上游仓库发布了重命名的 `biscuit` 扩展，请在替换之前审查其迁移边界。
- 在生产使用之前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
