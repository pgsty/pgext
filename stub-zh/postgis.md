

## 用法

> [PostGIS：PostgreSQL 的空间与地理对象支持](https://github.com/postgis/postgis)

PostGIS 为 PostgreSQL 添加了地理对象支持，将其变成一个空间数据库。它实现了 [OGC 简单要素](https://www.ogc.org/standard/sfs/)规范，并提供空间索引、空间函数和坐标变换功能。

### 文档

- [PostGIS 参考手册](https://postgis.net/docs/) -- 完整文档
- [几何类型](https://postgis.net/docs/using_postgis_dbmanagement.html#RefObject) -- Point、LineString、Polygon、MultiPoint 等
- [空间关系](https://postgis.net/docs/reference.html#Spatial_Relationships) -- ST_Contains、ST_Intersects、ST_Within 等
- [度量函数](https://postgis.net/docs/reference.html#Measurement_Functions) -- ST_Distance、ST_Area、ST_Length 等
- [几何处理](https://postgis.net/docs/reference.html#Geometry_Processing) -- ST_Buffer、ST_Union、ST_Intersection 等
- [几何输入/输出](https://postgis.net/docs/reference.html#Geometry_Inputs) -- WKT、WKB、GeoJSON、KML、GML、SVG
- [坐标变换](https://postgis.net/docs/reference.html#SRS_Functions) -- ST_Transform、ST_SetSRID
- [地理类型函数](https://postgis.net/docs/using_postgis_dbmanagement.html#PostGIS_Geography) -- 真实球面距离计算
- [空间索引](https://postgis.net/docs/using_postgis_dbmanagement.html#idm2269) -- GiST 与 SP-GiST 索引

### 安装

```sql
CREATE EXTENSION postgis;
```

验证安装：

```sql
SELECT PostGIS_Full_Version();
```

--------

## 核心数据类型

PostGIS 提供两种主要的空间类型：

| 类型 | 描述 | 坐标系 |
|------|------|--------|
| `geometry` | 平面（平地球）空间类型 | 笛卡尔坐标，使用 SRID 指定投影 |
| `geography` | 球面（圆地球）空间类型 | 始终使用 WGS 84 经纬度（SRID 4326） |

### Geometry（几何类型）

`geometry` 类型在投影坐标系中工作。速度快且支持 PostGIS 的全部函数。最适合在单一投影坐标系内工作的场景（例如 UTM 分区、州平面坐标系）。

### Geography（地理类型）

`geography` 类型在球体/椭球体上进行计算。距离和面积以米为单位返回。支持的函数较少，但在大面积范围内无需选择投影即可获得精确结果。

```sql
-- 地理类型列：距离以米为单位，始终使用 WGS 84
CREATE TABLE cities (
    name text PRIMARY KEY,
    location geography(Point, 4326)
);

INSERT INTO cities VALUES
    ('New York',  ST_GeogFromText('POINT(-74.006 40.7128)')),
    ('London',    ST_GeogFromText('POINT(-0.1278 51.5074)')),
    ('Tokyo',     ST_GeogFromText('POINT(139.6917 35.6895)'));

-- 以米为单位的距离（大圆距离）
SELECT a.name, b.name, ST_Distance(a.location, b.location) / 1000 AS distance_km
FROM cities a, cities b WHERE a.name < b.name;
```

--------

## 创建空间表

空间列具有几何类型、维度（2D、3D、4D）和空间参考系统标识符（SRID）。

```sql
CREATE TABLE buildings (
    id serial PRIMARY KEY,
    name text,
    geom geometry(Polygon, 4326)
);

CREATE TABLE roads (
    id serial PRIMARY KEY,
    name text,
    geom geometry(LineString, 4326)
);

CREATE TABLE sensors (
    id serial PRIMARY KEY,
    label text,
    geom geometry(Point, 4326)
);
```

### 插入空间数据

从 WKT（Well-Known Text）格式导入：

```sql
INSERT INTO sensors (label, geom) VALUES
    ('S1', ST_GeomFromText('POINT(-73.985 40.748)', 4326)),
    ('S2', ST_GeomFromText('POINT(-73.979 40.754)', 4326));
```

使用构造函数：

```sql
INSERT INTO sensors (label, geom) VALUES
    ('S3', ST_SetSRID(ST_MakePoint(-73.990, 40.735), 4326));
```

从 GeoJSON 导入：

```sql
INSERT INTO buildings (name, geom) VALUES
    ('Plaza', ST_GeomFromGeoJSON('{"type":"Polygon","coordinates":[[[-73.98,40.74],[-73.97,40.74],[-73.97,40.75],[-73.98,40.75],[-73.98,40.74]]]}'));
```

--------

## 空间索引

GiST 索引对空间查询性能至关重要。务必在空间列上创建索引：

```sql
CREATE INDEX idx_sensors_geom ON sensors USING GIST (geom);
CREATE INDEX idx_buildings_geom ON buildings USING GIST (geom);
CREATE INDEX idx_roads_geom ON roads USING GIST (geom);
```

空间索引支持边界框运算符（`&&`、`@`、`~`），并在 WHERE 子句中被 `ST_DWithin`、`ST_Intersects` 和 `ST_Contains` 等空间函数自动使用。

对于超大数据集，可以考虑使用 BRIN 索引：

```sql
CREATE INDEX idx_sensors_geom_brin ON sensors USING BRIN (geom);
```

--------

## 核心空间函数

### 创建几何对象

```sql
-- 从坐标创建点
SELECT ST_MakePoint(-73.985, 40.748);

-- 创建带 SRID 的点
SELECT ST_SetSRID(ST_MakePoint(-73.985, 40.748), 4326);

-- 从 WKT 创建
SELECT ST_GeomFromText('LINESTRING(0 0, 1 1, 2 1)', 4326);

-- 从 GeoJSON 创建
SELECT ST_GeomFromGeoJSON('{"type":"Point","coordinates":[-73.985,40.748]}');

-- 从两个点创建线段
SELECT ST_MakeLine(
    ST_MakePoint(0, 0),
    ST_MakePoint(1, 1)
);

-- 从闭合线串创建多边形
SELECT ST_MakePolygon(
    ST_GeomFromText('LINESTRING(0 0, 1 0, 1 1, 0 1, 0 0)')
);
```

### 度量

```sql
-- 两个几何对象之间的距离（以 SRID 单位计）
SELECT ST_Distance(
    ST_GeomFromText('POINT(0 0)', 4326),
    ST_GeomFromText('POINT(1 1)', 4326)
);

-- 使用地理类型计算以米为单位的距离
SELECT ST_Distance(
    'SRID=4326;POINT(-73.985 40.748)'::geography,
    'SRID=4326;POINT(-0.1278 51.5074)'::geography
);

-- 面积（以 SRID 单位计，或地理类型以平方米计）
SELECT ST_Area(geom) FROM buildings;

-- 线串长度
SELECT ST_Length(geom) FROM roads;

-- 多边形周长
SELECT ST_Perimeter(geom) FROM buildings;
```

### 空间关系

```sql
-- A 是否包含 B？
SELECT ST_Contains(a.geom, b.geom) FROM buildings a, sensors b;

-- A 与 B 是否相交？
SELECT ST_Intersects(a.geom, b.geom) FROM buildings a, roads b;

-- B 是否在 A 的距离 D 以内？（索引友好）
SELECT ST_DWithin(a.geom, b.geom, 0.01) FROM sensors a, sensors b;

-- A 和 B 是否在距离 D 以内？（地理类型，单位为米）
SELECT ST_DWithin(a.location, b.location, 10000) FROM cities a, cities b;

-- A 是否与 B 接触？（共享边界但不共享内部）
SELECT ST_Touches(a.geom, b.geom) FROM buildings a, roads b;

-- A 是否穿越 B？
SELECT ST_Crosses(a.geom, b.geom) FROM roads a, roads b;

-- A 是否与 B 重叠？（同维度，不完全相同）
SELECT ST_Overlaps(a.geom, b.geom) FROM buildings a, buildings b;
```

### 几何处理

```sql
-- 缓冲区（按距离扩展几何对象）
SELECT ST_Buffer(geom, 0.001) FROM sensors;

-- 两个几何对象的交集
SELECT ST_Intersection(a.geom, b.geom) FROM buildings a, buildings b
WHERE ST_Intersects(a.geom, b.geom) AND a.id < b.id;

-- 几何对象的联合
SELECT ST_Union(geom) FROM buildings WHERE name LIKE 'Block%';

-- 凸包
SELECT ST_ConvexHull(ST_Collect(geom)) FROM sensors;

-- 简化几何对象（Douglas-Peucker 算法）
SELECT ST_Simplify(geom, 0.0001) FROM roads;

-- 质心
SELECT ST_Centroid(geom) FROM buildings;

-- 沃罗诺伊图
SELECT ST_VoronoiPolygons(ST_Collect(geom)) FROM sensors;
```

### 坐标变换

```sql
-- 从 WGS 84 (4326) 转换到 Web 墨卡托 (3857)
SELECT ST_Transform(geom, 3857) FROM sensors;

-- 转换到 UTM 分区以进行米制计算
SELECT ST_Area(ST_Transform(geom, 32618)) AS area_sqm FROM buildings;

-- 设置几何对象的 SRID（不会重新投影）
SELECT ST_SetSRID(geom, 4326) FROM sensors;

-- 获取当前 SRID
SELECT ST_SRID(geom) FROM sensors;
```

### 输出格式

```sql
-- 转为 GeoJSON
SELECT ST_AsGeoJSON(geom) FROM sensors;

-- 转为 WKT
SELECT ST_AsText(geom) FROM sensors;

-- 转为 KML
SELECT ST_AsKML(geom) FROM sensors;

-- 转为 SVG 路径
SELECT ST_AsSVG(geom) FROM buildings;

-- 转为 EWKT（包含 SRID）
SELECT ST_AsEWKT(geom) FROM sensors;
```

--------

## 实际示例

### 查找附近的点

查找给定位置 500 米内的所有传感器：

```sql
SELECT label, ST_Distance(
    geom::geography,
    ST_SetSRID(ST_MakePoint(-73.985, 40.748), 4326)::geography
) AS distance_m
FROM sensors
WHERE ST_DWithin(
    geom::geography,
    ST_SetSRID(ST_MakePoint(-73.985, 40.748), 4326)::geography,
    500
)
ORDER BY distance_m;
```

### 空间连接

查找每个传感器所在的建筑物：

```sql
SELECT s.label, b.name
FROM sensors s
JOIN buildings b ON ST_Contains(b.geom, s.geom);
```

### 统计多边形内的点数

```sql
SELECT b.name, COUNT(s.id) AS sensor_count
FROM buildings b
LEFT JOIN sensors s ON ST_Contains(b.geom, s.geom)
GROUP BY b.name;
```

### K 近邻查询

使用索引加速的 `<->` 运算符查找距某点最近的 5 个传感器：

```sql
SELECT label, geom <-> ST_SetSRID(ST_MakePoint(-73.985, 40.748), 4326) AS dist
FROM sensors
ORDER BY geom <-> ST_SetSRID(ST_MakePoint(-73.985, 40.748), 4326)
LIMIT 5;
```

### 聚合为单个几何对象

```sql
-- 将所有传感器点收集为 MultiPoint
SELECT ST_Collect(geom) FROM sensors;

-- 计算最小外接圆
SELECT ST_MinimumBoundingCircle(ST_Collect(geom)) FROM sensors;
```

--------

## Geography 与 Geometry 对比

| 特性 | `geometry` | `geography` |
|------|-----------|-------------|
| 坐标系 | 任意投影坐标系 | 仅 WGS 84 |
| 距离单位 | SRID 单位（度、米、英尺） | 米 |
| 大面积精度 | 需要正确的投影 | 全球范围精确 |
| 函数支持 | 完整（约 300 个函数） | 子集（约 40 个函数） |
| 索引支持 | GiST、SP-GiST、BRIN | GiST |
| 性能 | 更快 | 略慢 |

常见模式是使用 `geography` 存储数据以确保正确性，需要使用地理类型不支持的函数时再转换为 `geometry`：

```sql
SELECT ST_Area(geom::geography) AS area_sqm FROM buildings;
```
