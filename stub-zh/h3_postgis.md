## 用法

来源：

- [4.5.0 PostGIS API](https://github.com/postgis/h3-pg/blob/v4.5.0/docs/api.md)
- [Dependencies and extension definition](https://github.com/postgis/h3-pg/blob/v4.5.0/h3_postgis/CMakeLists.txt)
- [4.5.0 migration SQL](https://github.com/postgis/h3-pg/blob/v4.5.0/h3_postgis/sql/updates/h3_postgis--4.2.3--4.5.0.sql)
- [4.5.0 release](https://github.com/postgis/h3-pg/releases/tag/v4.5.0)

`h3_postgis` 将 H3 单元格与 PostGIS 几何、地理及栅格数据连接起来。它依赖 `h3`、`postgis` 和 `postgis_raster`，即便仅处理几何也需要这些依赖。输入几何必须使用 SRID 4326，坐标顺序为经度、纬度；这些函数不会自动重新投影输入。

### 点与单元格转换

```sql
CREATE EXTENSION h3_postgis CASCADE;
SET h3.strict = true;

SELECT h3_latlng_to_cell(
    ST_SetSRID(ST_MakePoint(-122.0553238, 37.3615593), 4326), 9
);
SELECT h3_cell_to_geometry('85283473fffffff'::h3index);
SELECT h3_cell_to_boundary_geometry('85283473fffffff'::h3index);
```

其他坐标系应先通过 PostGIS 转换到 SRID 4326，再调用 H3 函数。仅设置 SRID 标签不会转换坐标。

### 核心接口

| 任务 | 函数 |
| --- | --- |
| 点转单元格 | `h3_latlng_to_cell(geometry, integer)`、`h3_latlng_to_cell(geography, integer)` |
| 单元格中心 | `h3_cell_to_geometry`、`h3_cell_to_geography` |
| 单元格边界 | `h3_cell_to_boundary_geometry`、`h3_cell_to_boundary_geography` |
| 多边形覆盖 | `h3_polygon_to_cells`、`h3_cells_to_multi_polygon_geometry`、`h3_cells_to_multi_polygon_geography` |
| 连续栅格统计 | `h3_raster_summary`、`h3_raster_summary_stats_agg` |
| 分类栅格统计 | `h3_raster_class_summary`、`h3_raster_class_summary_item_agg` |

几何与分辨率之间的 `@` 运算符也能将位置映射到 H3 单元格。使用 `ST_IsValid()` 检查多边形；`ST_MakeValid()` 修复可能改变拓扑并产生几何集合，应在覆盖计算前提取并检查多边形部分。无效多边形的行为未定义。

### 汇总栅格数据

```sql
SELECT (summary).h3,
       (h3_raster_summary_stats_agg((summary).stats)).*
FROM (
    SELECT h3_raster_summary(rast, 8) AS summary
    FROM rasters
) AS r
GROUP BY (summary).h3;
```

默认汇总函数会自动选择方法，也可显式使用裁剪、像素中心或子像素变体控制像素如何分配到单元格。应结合栅格分辨率和目标 H3 分辨率检查所选方法。

### 升级与边界

```sql
ALTER EXTENSION h3 UPDATE TO '4.5.0';
ALTER EXTENSION h3_postgis UPDATE TO '4.5.0';
```

基础扩展更新会重建受影响的 btree 索引并刷新距离依赖对象，应先安排维护窗口，再更新伴随扩展。4.5.0 修复 PostgreSQL 17+ 受限搜索路径下的维护、表达式索引导出恢复，以及多项几何和多边形生成错误。两个扩展版本应保持一致。更新任一扩展前，先安装匹配的 4.5.0 软件包文件。

平面叠加运算通常应保持 `h3.extend_antimeridian` 为 false。两个扩展都可重定位；执行未限定模式的 SQL 时，应确保 H3 和 PostGIS 所在模式可见。两者均无需共享预加载。
