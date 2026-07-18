## 用法

来源：

- [上游 README 与函数参考](https://github.com/decision-labs/a5pg/blob/main/README.md)
- [扩展 control 文件](https://github.com/decision-labs/a5pg/blob/main/a5pg.control)
- [Cargo 包元数据](https://github.com/decision-labs/a5pg/blob/main/Cargo.toml)

`a5pg` 在 PostgreSQL 中提供等面积 [A5](https://a5geo.org/) 离散全球网格。它可以把经纬度坐标转换为 64 位单元格标识符、还原单元格中心或边界，并在单元格层级间导航。核心 API 与 DuckDB A5 扩展兼容。

### 核心操作

```sql
CREATE EXTENSION a5pg;

SELECT a5_lonlat_to_cell(-73.9857, 40.7580, 10);
SELECT a5_cell_to_lonlat(2742822465196523520);
SELECT a5_get_resolution(2742822465196523520);
SELECT a5_cell_to_parent(2742822465196523520, 8);
SELECT a5_cell_to_children(2742822465196523520, 12);
```

`a5_cell_to_boundary()` 以二维 `double precision` 数组返回边界。如果数据库中已有 PostGIS `geometry` 类型，还可使用 `a5_point_to_cell(geometry, resolution)` 便利封装；核心经纬度函数并不依赖 PostGIS。

上游包与本目录记录的版本均为 `0.6.1`。Cargo feature 覆盖 PostgreSQL 13 至 18。该扩展无需预加载。
