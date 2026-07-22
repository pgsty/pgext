## 用法

来源：

- [官方 README](https://github.com/dlr-eoc/pgh3/blob/a225ac6da928cb0668780221da64061d3c4b1d54/README.md)
- [SQL API 参考](https://github.com/dlr-eoc/pgh3/blob/a225ac6da928cb0668780221da64061d3c4b1d54/doc/pgh3.md)
- [扩展 SQL](https://github.com/dlr-eoc/pgh3/blob/a225ac6da928cb0668780221da64061d3c4b1d54/sql/pgh3.sql)

`pgh3` 0.3.0 是一个已归档的 PostGIS 几何与旧版 Uber H3 C API 绑定。它用 `text` 表示 H3 索引，并提供点索引、单元几何、层级、邻域、多边形填充、压缩和平均尺寸辅助函数。

### 点与单元流程

先安装 PostGIS。坐标按 X 为经度、Y 为纬度的度数解释；包装器既不转换也不验证几何的 SRID。

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pgh3;

WITH cell AS (
    SELECT h3_geo_to_h3index(
        ST_SetSRID(ST_Point(9.4069, 52.1233), 4326),
        5
    ) AS h3index
)
SELECT h3index,
       h3_get_resolution(h3index),
       h3_h3index_to_geo(h3index),
       h3_h3index_to_geoboundary(h3index)
FROM cell;
```

调用扩展前，应把源几何转换为经纬度。返回的 PostGIS 几何不会保留 SRID，因此后续空间操作需要时要显式设置预期 SRID。

### 主要函数组

- `h3_geo_to_h3index(point, integer)` 及其 `geometry` 重载创建索引。
- `h3_h3index_to_geo(text)` 和 `h3_h3index_to_geoboundary(text)` 返回单元中心与边界几何。
- `h3_kring(text, integer)`、`h3_to_parent(text, integer)` 和 `h3_to_children(text, integer)` 遍历邻域与层级。
- `h3_compact(text[])` 和 `h3_uncompact(text[], integer)` 改变层级表示。
- `h3_hexagon_area_km2(integer)`、`h3_hexagon_area_m2(integer)`、`h3_edge_length_km(integer)` 和 `h3_edge_length_m(integer)` 返回平均尺寸。

### 多边形填充与内存

物化填充前应先估算基数，尤其是使用精细分辨率时。

```sql
WITH region AS (
    SELECT ST_GeomFromText(
        'POLYGON((9.35 52.08,9.46 52.08,9.46 52.17,9.35 52.17,9.35 52.08))',
        4326
    ) AS geom
)
SELECT h3_polyfill_estimate(geom, 8) FROM region;

WITH region AS (
    SELECT ST_GeomFromText(
        'POLYGON((9.35 52.08,9.46 52.08,9.46 52.17,9.35 52.17,9.35 52.08))',
        4326
    ) AS geom
)
SELECT h3_polyfill(geom, 8) FROM region;
```

`h3_polyfill` 支持 Polygon 和 MultiPolygon 几何，并排除孔洞。它会按估算结果预分配输出空间。`pgh3.polyfill_mem` 限制该分配，默认值为 PostgreSQL 约 1 GB 的 `MaxAllocSize`；提高它可能造成非常大的后端内存分配。与其提高上限，更应降低分辨率或先用 PostGIS 切分输入。

### 兼容边界

扩展使用 H3 v4 之前的函数名，并与构建时所用 H3 库 ABI 紧密耦合。仓库已经归档，也没有证明支持历史 PostgreSQL 范围之外的版本。应针对明确匹配的 H3 版本构建，使用已知单元 ID 和几何方向验证结果；新生产系统应优先选择仍在维护的 H3 扩展。
