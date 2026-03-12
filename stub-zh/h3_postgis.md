

## 用法

> [h3_postgis: H3 的 PostGIS 集成](https://github.com/zachasme/h3-pg)

`h3_postgis` 是将 H3 六边形层次空间索引与 PostGIS 集成的桥接扩展。它实现了 H3 索引与 PostGIS 几何类型之间的转换。

```sql
CREATE EXTENSION h3_postgis CASCADE;
```

该扩展需要同时安装 `h3` 和 `postgis`。它提供了在 H3 单元格索引与 PostGIS 几何体之间转换的函数，使得可以将 H3 的六边形网格系统与 PostGIS 的空间能力结合使用。

### 主要函数

```sql
-- 将 PostGIS 点转换为 H3 单元格索引
SELECT h3_latlng_to_cell(ST_MakePoint(-73.985, 40.748)::point, 9);

-- 获取 H3 单元格边界的 PostGIS 几何体
SELECT h3_cell_to_boundary_geometry('892a1008003ffff'::h3index);

-- 将 H3 单元格转换为 PostGIS 多边形用于可视化
SELECT h3_cell_to_geometry('892a1008003ffff'::h3index);
```
