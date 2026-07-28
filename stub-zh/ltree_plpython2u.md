## 用法

来源：

- [官方上游 README](https://github.com/vicmisael/postgres-xl-fix/blob/0bca966cd19ca7be25676db32f0f50c014349f20/contrib/README)
- [官方扩展控制文件 (ltree_plpython2u.control)](https://github.com/vicmisael/postgres-xl-fix/blob/0bca966cd19ca7be25676db32f0f50c014349f20/contrib/ltree_plpython/ltree_plpython2u.control)
- [官方扩展 SQL (ltree_plpython2u--1.0.sql)](https://github.com/vicmisael/postgres-xl-fix/blob/0bca966cd19ca7be25676db32f0f50c014349f20/contrib/ltree_plpython/ltree_plpython2u--1.0.sql)

`ltree_plpython2u` — 一个固定的 PostgreSQL 扩展，使其可以在较新的 GCC 和 glibc 版本上编译。当数据库代码必须在或与该过程语言进行交互时，请使用此扩展。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION ltree_plpython2u;
```

在目标数据库中安装扩展，运行可用的最小上游示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `ltree_to_plpython2(val internal)` 是一个扩展函数，返回 `internal`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 先安装确认的扩展依赖项：`ltree`, `plpython2u`。
- 控制文件将扩展标记为可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
