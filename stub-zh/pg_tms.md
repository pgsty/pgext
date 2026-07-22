## 用法

来源：

- [Pinned official README](https://github.com/jkeifer/pg_tms/blob/3e11a51f42095632d36bcd8b9775da41edaa8cde/README.md)
- [Pinned extension SQL](https://github.com/jkeifer/pg_tms/blob/3e11a51f42095632d36bcd8b9775da41edaa8cde/pg_tms--0.0.4.sql)

`pg_tms` 是基于 SQL/PLpgSQL 的 PostGIS 扩展，可把栅格转换为符合 Tile Map Service 坐标模型的 256 像素全局墨卡托切片。它在 PostgreSQL 内创建切片栅格和坐标辅助对象；不提供 HTTP 切片服务器。

### 核心流程

需要 PostGIS 2.1 或更高版本；依赖可用时，`CASCADE` 可以一并安装：

```sql
CREATE EXTENSION pg_tms CASCADE;

SELECT t.x, t.y, t.z, t.rast
FROM source_rasters AS s
CROSS JOIN LATERAL tms_tile_raster_to_zoom(
  s.rast,
  12,
  true,
  'NearestNeighbor'
) AS t;
```

`tms_tile_raster_to_zoom(raster, zoom, drop_blanks, algorithm)` 返回包含 `rast`、`x`、`y`、`z` 的 `tms_tile` 行。`zoom` 为 `-1` 时，扩展会根据像素大小推导原生缩放级别。`drop_blanks` 会省略所有波段都无数据的切片，`algorithm` 则传给 PostGIS 重采样。

### 重要对象

- 坐标类型：`tms_latlon`、`tms_meters`、`tms_pixels`、`tms_tilecoord`、`tms_tilecoordz` 及其范围形式。
- 切片类型：`tms_tile(rast raster, x int, y int, z int)`。
- 分辨率与边界：`tms_initialresolution`、`tms_originshift`、`tms_resolution`、`tms_zoomforpixelsize`、`tms_tilebounds_meters`、`tms_tilebounds_latlon`。
- 转换辅助：`tms_latlon2meters`、`tms_meters2latlon`、`tms_pixels2meters`、`tms_meters2pixels`、`tms_pixels2tile`、`tms_meters2tile`。
- 栅格流程：`tms_raster2tilecoord_ext`、`tms_tilecoordz_from_raster`、`tms_copy_to_tile`、`tms_has_data`、`tms_tile_raster_to_zoom`。

扩展还定义了从 `tms_tilecoordz` 到 `raster`、`geometry`、米制边界和经纬度边界的赋值转换。

### 注意事项

只实现了全局墨卡托路径；生成的切片栅格和几何对象使用 SRID 3857。切片前应重投影源数据，并检查地理参考、像素大小、波段 nodata 值和跨日期变更线行为。大栅格或高缩放级别可能生成大量消耗 CPU 与存储的切片。

项目称 0.0.4 版仍在开发中，目前已经弃用。其 SQL 面向 PostgreSQL 9.3/PostGIS 2.1 时代的 API，因此应针对部署的 PostGIS 版本验证每个使用的函数。应在 PostgreSQL 前增加 HTTP 服务或缓存来提供返回的栅格。无需预加载或重启。
