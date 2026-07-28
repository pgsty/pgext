## 用法

来源：

- [官方上游 README](https://github.com/masahikosawada/gmapbench/blob/a2425a83a59454d931793febaa52fb04d8912492/source/pg-garbagemap-background/contrib/README)
- [官方扩展控制文件 (ltree_plpythonu.control)](https://github.com/masahikosawada/gmapbench/blob/a2425a83a59454d931793febaa52fb04d8912492/source/pg-garbagemap-background/contrib/ltree_plpython/ltree_plpythonu.control)
- [官方扩展 SQL (ltree_plpythonu--1.0.sql)](https://github.com/masahikosawada/gmapbench/blob/a2425a83a59454d931793febaa52fb04d8912492/source/pg-garbagemap-background/contrib/ltree_plpython/ltree_plpythonu--1.0.sql)

`ltree_plpythonu` — 在数据库代码必须在或与此过程语言进行交互时使用。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION ltree_plpythonu;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `ltree_to_plpython(val internal)` 是一个扩展函数，返回 `internal`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 先安装确认的扩展依赖项：`ltree`, `plpythonu`。
- 控制文件将扩展标记为可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用之前，需确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
