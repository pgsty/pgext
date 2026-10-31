## 用法

来源：

- [PostGIS 3.7.0 发布](https://postgis.net/2026/10/PostGIS-3.7.0/)
- [PostGIS 3.7 手册](https://postgis.net/docs/manual-3.7/)
- [官方发布说明](https://postgis.net/docs/release_notes.html)
- [当前手册 HTML](https://postgis.net/docs/postgis-en.html)

`postgis` 3.7.0 为 PostgreSQL 增加空间类型、索引与 SQL 函数。对用户来说，最主要的区分是：`geometry` 用于平面或投影坐标工作，`geography` 用于经纬度数据上的球面计算。此主要版本要求 PostgreSQL 14 及以上、GEOS 3.10 及以上以及 PROJ 6.1 及以上。

### 基本设置

```sql
CREATE EXTENSION postgis;
SELECT PostGIS_Full_Version();
```

### 核心类型与函数

```sql
CREATE TABLE sensors (
  id bigserial PRIMARY KEY,
  geom geometry(Point, 4326),
  geog geography(Point, 4326)
);

SELECT ST_SetSRID(ST_MakePoint(-73.985, 40.748), 4326);
SELECT ST_Intersects(a.geom, b.geom) FROM a, b;
SELECT ST_DWithin(a.geom, b.geom, 100);
SELECT ST_Distance(a.geog, b.geog);
SELECT ST_Transform(geom, 3857) FROM sensors;
```

- 构造函数：`ST_MakePoint`、`ST_GeomFromText`、`ST_GeomFromGeoJSON`
- 空间关系：`ST_Intersects`、`ST_Contains`、`ST_Within`、`ST_DWithin`
- 度量与坐标变换：`ST_Distance`、`ST_Area`、`ST_Length`、`ST_Transform`
- 几何处理：`ST_Buffer`、`ST_Intersection`、`ST_Union`

### 空间索引

```sql
CREATE INDEX idx_sensors_geom ON sensors USING GIST (geom);
```

官方手册仍将 GiST 作为通用空间索引的首选建议；对于特定数据分布与权衡，也可以考虑 BRIN 与 SP-GiST。

### 注意事项

- 进行平面距离与面积计算时，应在合适的投影 SRID 中使用 `geometry`；需要以米为单位的球面计算时，应使用 `geography`。
- `PostGIS 3.7.0` 是一个主要版本，移除了对 PostgreSQL 12/13 以及低于 3.10 的 GEOS 的支持；升级前应检查拓扑容差语义变化，以及 `address_standardizer` 和 `tiger_geocoder` 软件包迁移。安装匹配文件后，按上游升级流程的要求执行 `SELECT postgis_extensions_upgrade();`。
