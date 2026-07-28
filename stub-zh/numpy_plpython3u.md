## 用法

来源：

- [官方上游 README](https://github.com/tarkmeper/numpy_plpython/blob/acd66d925f205227acec6e87fcb5c49eae770abb/README.md)
- [官方扩展控制文件 (numpy_plpython3u.control)](https://github.com/tarkmeper/numpy_plpython/blob/acd66d925f205227acec6e87fcb5c49eae770abb/numpy_plpython3u.control)
- [官方扩展 SQL (numpy_plpython3u--1.0.sql)](https://github.com/tarkmeper/numpy_plpython/blob/acd66d925f205227acec6e87fcb5c49eae770abb/numpy_plpython3u--1.0.sql)

`numpy_plpython3u` — 用于直接将 Postgres 数组转换为 numpy 数组，而无需经过 Python 列表。当数据库代码必须在该过程语言中运行或与其进行交互时，请使用此扩展。在安装扩展及其依赖项并验证后，可以使用它。

### 核心工作流

```sql
CREATE EXTENSION numpy_plpython3u;
```

在目标数据库中安装扩展，运行可用的最小上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `numpy_to_plpython3(val internal)` 是一个扩展函数，返回 `internal`。
- `plpython3_to_numpy(val internal)` 是一个扩展函数，返回 `real[]`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 首先安装并验证确认的扩展依赖项：`plpython3u`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
