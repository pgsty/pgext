## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/first_last/first_last-0.1.1/README)
- [官方扩展控制文件 (first_last.control)](https://api.pgxn.org/src/first_last/first_last-0.1.1/first_last.control)
- [官方扩展 SQL (first_last--0.1.0--0.1.1.sql)](https://api.pgxn.org/src/first_last/first_last-0.1.1/first_last--0.1.0--0.1.1.sql)

`first_last` — 此扩展提供四个聚合函数：first(anyelement) first(anyelement, int4) last(anyelement) last(anyelement, int4)。当 SQL 需要这些特殊功能或聚合时，请使用此扩展。请使用上方链接的上游版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION first_last;
```

在目标数据库中安装此扩展，如果有可用示例，请运行上方的最小上游示例，并在将其集成到应用程序 SQL 之前验证安装版本和返回值。

### 重要对象

- `agg_first(IN p_state anyarray, IN p_new_element anyelement, IN p_limit int4)` 是一个扩展函数并返回 `anyarray`。
- `agg_first(IN p_state anyelement, IN p_new_element anyelement)` 是一个扩展函数并返回 `anyelement`。
- `agg_last(IN p_state anyarray, IN p_new_element anyelement, IN p_limit int4)` 是一个扩展函数并返回 `anyarray`。
- `agg_last(IN p_state anyelement, IN p_new_element anyelement)` 是一个扩展函数并返回 `anyelement`。
- `first` 是由扩展公开的聚合。
- `last` 是由扩展公开的聚合。

### 要求与注意事项

- 审查过的控制文件声明默认版本 `0.1.1`。
- 控制文件将此扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的内容一致。
