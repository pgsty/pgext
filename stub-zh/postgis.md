## 用法

来源：[Official manual](https://postgis.net/documentation/manual/)，[current manual HTML](https://postgis.net/docs/postgis-en.html)，[release notes](https://postgis.net/docs/release_notes.html)，[patch release announcement](https://postgis.net/2026/04/PostGIS-Patch-Releases/)

`postgis` 为 PostgreSQL 增加空间类型、索引与 SQL 函数。对用户来说，最主要的区分是：`geometry` 用于平面或投影坐标工作，`geography` 用于经纬度数据上的球面计算。

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
- `PostGIS 3.6.3` 是一个发布日期为 `2026-04-14` 的补丁版本。release notes 描述的是修复项与一项安全加固变更，而不是新的 stub 级使用接口，因此这次刷新主要是裁剪并对齐当前手册。
