## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/hexgrid/hexgrid-1.0.3/README.md)
- [官方扩展控制文件 (hexgrid.control)](https://api.pgxn.org/src/hexgrid/hexgrid-1.0.3/hexgrid.control)
- [官方扩展 SQL (hexgrid.sql)](https://api.pgxn.org/src/hexgrid/hexgrid-1.0.3/sql/hexgrid.sql)

`hexgrid` — 可配置的六边形网格在抽象表面上。用于相应的空间数据或地理空间工作流。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION hexgrid;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `hex_Orientation(name text, f float[], b float[], start_angle float)` 是一个扩展函数，返回 `hex_orientation`。
- `hex_OrientationFlat()` 是一个扩展函数，返回 `hex_orientation`。
- `hex_OrientationPointy()` 是一个扩展函数，返回 `hex_orientation`。
- `ST_Centroid(hexagon hexagon)` 是一个扩展函数，返回 `geometry`。
- `ST_Hexagon(point geometry(point), grid_id int default 1)` 是一个扩展函数，返回 `hexagon`。
- `ST_HexagonCoverage(region geometry, grid_id int default 1)` 是一个扩展函数，返回 `setof`。
- `hex_orientation` 是一个扩展定义的类型。
- `hexagon` 是一个扩展定义的类型。
- `hexgrid` 是一个扩展定义的类型。
- `hexgrids` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0.3`。
- 控制文件标记扩展为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
