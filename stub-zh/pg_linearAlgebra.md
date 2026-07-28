## 用法

来源：

- [官方上游 README](https://github.com/joshuajerin/pg_linearalgebra/blob/63a4ca281720241ace072df5f07b972e6a6598ac/README.md)
- [官方扩展控制文件 (pg_linearAlgebra.control)](https://github.com/joshuajerin/pg_linearalgebra/blob/63a4ca281720241ace072df5f07b972e6a6598ac/pg_linearAlgebra.control)
- [官方扩展 SQL (pg_linearAlgebra--1.0.sql)](https://github.com/joshuajerin/pg_linearalgebra/blob/63a4ca281720241ace072df5f07b972e6a6598ac/sql/pg_linearAlgebra--1.0.sql)

`pg_linearAlgebra` — 一个用于基本线性代数操作的 PostgreSQL 扩展。当 SQL 需要这些特殊函数或聚合时，请使用此扩展。请使用上述链接的上游修订版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION "pg_linearAlgebra";

SELECT mAdd('[[1.0, 2.0], [3.0, 4.0]]', '[[5.0, 6.0], [7.0, 8.0]]', 2, 2);
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `mAdd(matrix1 text, matrix2 text, rows integer, cols integer)` 是一个扩展函数，返回 `text`。
- `mMultiply(matrix1 text, matrix2 text, rows integer, cols integer)` 是一个扩展函数，返回 `text`。
- `mSubtract(matrix1 text, matrix2 text, rows integer, cols integer)` 是一个扩展函数，返回 `text`。
- `mSvd(matrix text, rows integer, cols integer)` 是一个扩展函数，返回 `text`。
- `mTranspose(matrix text, rows integer, cols integer)` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况与固定源进行比对。
