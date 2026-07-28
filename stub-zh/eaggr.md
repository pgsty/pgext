## 用法

来源：

- [官方上游 README](https://github.com/riskaware-ltd/open-eaggr/blob/c0d6d3b1091aaca10a2baa49b018687f198b7ce7/README.md)
- [官方扩展控制文件 (eaggr.control)](https://github.com/riskaware-ltd/open-eaggr/blob/c0d6d3b1091aaca10a2baa49b018687f198b7ce7/EAGGRPostgres/eaggr.control)
- [官方扩展 SQL (eaggr--2.0.sql)](https://github.com/riskaware-ltd/open-eaggr/blob/c0d6d3b1091aaca10a2baa49b018687f198b7ce7/EAGGRPostgres/eaggr--2.0.sql)

`eaggr` — OpenEAGGR 软件库是 Discrete Global Grid System (DGGS) 的实现，用于将地球表面建模为等面积单元的网络。使用它进行相应的空间数据或地理空间工作流。在目标 PostgreSQL 构建中使用链接的上游修订版本作为 API 边界并进行测试。

### 核心工作流

```sql
CREATE EXTENSION eaggr;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `EAGGR_CellGeometry(text, text)` 是一个扩展函数，返回 `text`。
- `EAGGR_CellToPoint(text, text)` 是一个扩展函数，返回 `text`。
- `EAGGR_GetBoundingCell(text[], text)` 是一个扩展函数，返回 `text`。
- `EAGGR_ShapeComparison(text, text, text, text)` 是一个扩展函数。
- `EAGGR_ToCellArray(text)` 是一个扩展函数，返回 `text[]`。
- `EAGGR_ToCells(text, double precision, text)` 是一个扩展函数，返回 `text`。
- `EAGGR_Version()` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 审查的控制文件声明默认版本为 `2.0`。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
