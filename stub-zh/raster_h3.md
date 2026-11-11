## 用法

来源：

- [PostgreSQL 接口](https://github.com/dmuldrew/raster_h3/blob/507e980c7a91feae74bdd40607acfccb64a1bbcb/docs/engines/postgresql.md)
- [扩展控制文件](https://github.com/dmuldrew/raster_h3/blob/507e980c7a91feae74bdd40607acfccb64a1bbcb/crates/raster-h3-pg/raster_h3.control)
- [SQL 定义](https://github.com/dmuldrew/raster_h3/blob/507e980c7a91feae74bdd40607acfccb64a1bbcb/crates/raster-h3-pg/sql/raster_h3--0.7.0.sql)

`raster_h3` 将 GeoTIFF 和云优化栅格聚合为 H3 单元。`0.7.0` 版本使用 PostgreSQL 函数管理器 ABI 提供原生 Rust 库，不依赖 PGRX。支持小端 64 位系统上的 PostgreSQL 13–18；扩展库必须与服务器主版本匹配。

### 启用与聚合

安装匹配的库及扩展文件后，以超级用户启用：

```sql
CREATE EXTENSION raster_h3;

SELECT h3_hex, mean, count, min, max, sum
FROM h3_raster_continuous_aggregate('/data/elevation.tif', 8);

SELECT h3_hex, majority_class, majority_fraction, histogram
FROM h3_raster_categorical_aggregate('/data/landcover.tif', 8);
```

连续值函数返回汇总统计量；分类函数返回类别计数、占比、多样性指标及 JSON 文本直方图。两者还返回单元索引、分辨率和 WKB 几何。每次调用只计算一个分辨率，范围为 0 至 15。可选参数包括 `source_crs`、`nodata` 与 `bbox`；边界框使用经纬度范围。

```sql
SELECT h3_hex, mean
FROM h3_raster_continuous_aggregate(
    '/data/temperature.tif', h3_res => 8,
    source_crs => 'EPSG:4326', nodata => -9999,
    bbox => '-122.50,37.70,-122.35,37.85'
);
```

### 权限与辅助函数

栅格路径由数据库服务器打开，也支持 HTTP(S) 和 S3 URL。两个聚合函数撤销了 `PUBLIC` 的 `EXECUTE` 权限，并将其授予 `pg_read_server_files`。该角色还允许更广泛的服务器文件访问，因此只应授予适当的操作人员。标量辅助函数不需要栅格文件访问权限。

标量 API 包括 `h3_to_string`、`string_to_h3`、`h3_to_lat`、`h3_to_lng`、`h3_get_resolution`、`h3_is_valid`、`h3_to_wkb` 和 `h3_cell_to_parent`。可选用 PostGIS 的 `ST_GeomFromWKB` 处理返回的 WKB；PostGIS 并非安装依赖。

### 限制与维护

扩展不要求共享预加载或重启服务器。已有 `0.4.0` 及以后版本可通过 `ALTER EXTENSION raster_h3 UPDATE` 执行 SQL 升级。PostgreSQL 接口未提供项目中的 Parquet/PMTiles 写入、波段选择和分位数选项；这些操作需使用上游的其他接口。统计量按像素贡献加权，而非地表面积。解释结果时应核对源坐标系及坐标变换警告。
