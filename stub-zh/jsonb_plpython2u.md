## 用法

来源：

- [官方文档](https://www.postgresql.org/docs/12/datatype-json.html#JSON-TRANSFORM)
- [官方扩展控制文件 (jsonb_plpython2u.control)](https://github.com/postgres/postgres/blob/REL_12_STABLE/contrib/jsonb_plpython/jsonb_plpython2u.control)
- [官方扩展 SQL (jsonb_plpython2u--1.0.sql)](https://github.com/postgres/postgres/blob/REL_12_STABLE/contrib/jsonb_plpython/jsonb_plpython2u--1.0.sql)

`jsonb_plpython2u` — 历史上的 jsonb 和不受信任的 PL/Python 2 语言之间的转换。当数据库代码必须在该过程语言中运行或与其进行交互时，请使用此扩展。上游项目已归档或不再维护。

### 核心工作流

```sql
CREATE EXTENSION jsonb_plpython2u;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `jsonb_to_plpython2(val internal)` 是一个扩展函数，返回 `internal`。
- `plpython2_to_jsonb(val internal)` 是一个扩展函数，返回 `jsonb`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 请首先安装确认的扩展依赖项：`plpython2u`。
- 控制文件将该扩展标记为可重定位。
- 该历史上的转换依赖于 `plpython2u`；当前的 PostgreSQL 使用 Python 3 转换变体。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
