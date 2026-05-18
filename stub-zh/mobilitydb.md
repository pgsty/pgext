## 用法

来源：[repo README](https://github.com/MobilityDB/MobilityDB), [MobilityDB 1.3 manual](https://mobilitydb.github.io/MobilityDB/master/), [v1.3.0 release](https://github.com/MobilityDB/MobilityDB/releases/tag/v1.3.0)

MobilityDB 使用 temporal 和 spatio-temporal 数据类型扩展 PostgreSQL 与 PostGIS，使车辆轨迹、传感器读数和随时间变化属性等 moving object data 能够高效存储、索引和查询。

**关键文档：**

- [MobilityDB Manual](https://mobilitydb.github.io/MobilityDB/master/)
- [Temporal Types](https://mobilitydb.github.io/MobilityDB/master/ch04.html)
- [Spatial-Temporal Types](https://mobilitydb.github.io/MobilityDB/master/ch07.html)
- [Temporal Poses](https://mobilitydb.github.io/MobilityDB/master/ch11.html)
- [Temporal Circular Buffers](https://mobilitydb.github.io/MobilityDB/master/ch13.html)
- [Indexing](https://mobilitydb.github.io/MobilityDB/master/ch10s02.html)
- [MobilityDB Workshop](https://mobilitydb.com/documentation/)
- [API Reference](https://mobilitydb.github.io/MobilityDB/master/)

### 入门

MobilityDB 要求 PostGIS。启用两个扩展：

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION mobilitydb;
```

### Temporal 类型

MobilityDB 为基础类型提供 temporal variants：

| Temporal Type | Base Type | Description |
|---------------|-----------|-------------|
| `tbool`       | `boolean` | Time-varying boolean |
| `tint`        | `integer` | Time-varying integer |
| `tfloat`      | `float`   | Time-varying float |
| `ttext`       | `text`    | Time-varying text |
| `tgeometry`   | `geometry` | Time-varying arbitrary geometry |
| `tgeography`  | `geography` | Time-varying arbitrary geography |
| `tgeompoint`  | `geometry(Point)` | Time-varying geometric point |
| `tgeogpoint`  | `geography(Point)` | Time-varying geographic point |
| `tnpoint`     | network point | Time-varying network point |
| `tcbuffer`    | circular buffer | Time-varying circular buffer |
| `tpose`       | pose | Time-varying point position and orientation |
| `trgeometry`  | rigid geometry | Time-varying rigid geometry |

MobilityDB 1.3 增加了 `tgeometry`、`tgeography`、`tcbuffer`、`tpose` 和 `trgeometry`。`tgeometry` 与 `tgeography` 支持 discrete 或 step interpolation，不支持任意几何的 linear interpolation。1.3 release notes 将 `tcbuffer`、`tpose` 和 `trgeometry` 标记为 experimental。

### Temporal 子类型

每种 temporal type 可按值如何随时间变化表示为不同子类型：

| Subtype | Description | Example |
|---------|-------------|---------|
| **Instant** | 单个时间戳上的单个值 | `'25.5@2025-01-01 08:00'` |
| **Sequence** | 一个时间区间上的连续值 | `'[25.5@08:00, 28.1@09:00, 30.0@10:00]'` |
| **SequenceSet** | 一组互不重叠的 sequences | `'{[25.5@08:00, 28.1@09:00], [30.0@11:00, 31.2@12:00]}'` |

Sequences 使用方括号表示包含 `[` 或排除 `(` 边界，就像 PostgreSQL range types 一样。

### 创建 Temporal 值

**Instant values：**

```sql
SELECT tfloat '25.5@2025-06-01 08:00:00+00';
SELECT tgeompoint 'SRID=4326;Point(2.3522 48.8566)@2025-06-01 08:00:00+00';
```

**Sequence values（continuous interpolation）：**

```sql
SELECT tfloat '[20.0@2025-06-01 08:00, 25.5@2025-06-01 09:00, 22.0@2025-06-01 10:00]';
```

**Discrete sequences（stepwise interpolation）：**

```sql
SELECT tint 'Interp=Step;[10@2025-06-01 08:00, 20@2025-06-01 09:00, 15@2025-06-01 10:00]';
```

**SequenceSet values：**

```sql
SELECT tfloat '{[20.0@08:00, 25.5@09:00], [22.0@11:00, 28.0@12:00]}';
```

**从组件构造：**

```sql
SELECT tgeompoint_inst(ST_Point(2.3522, 48.8566, 4326), '2025-06-01 08:00+00');
SELECT tgeompoint_seq(ARRAY[
    tgeompoint_inst(ST_Point(2.3522, 48.8566, 4326), '2025-06-01 08:00+00'),
    tgeompoint_inst(ST_Point(2.2945, 48.8584, 4326), '2025-06-01 08:30+00'),
    tgeompoint_inst(ST_Point(2.3364, 48.8606, 4326), '2025-06-01 09:00+00')
]);
```

### Temporal 操作

**提取指定时间的值：**

```sql
SELECT valueAtTimestamp(temp, '2025-06-01 08:30:00+00')
FROM (SELECT tfloat '[20.0@08:00, 30.0@09:00]' AS temp) t;
-- Returns 25.0 (linear interpolation)
```

**限制到某个时间段：**

```sql
SELECT atTime(trip, tstzspan '[2025-06-01 08:00, 2025-06-01 09:00]')
FROM trips;
```

**获取 temporal value 的时间跨度：**

```sql
SELECT duration(trip), startTimestamp(trip), endTimestamp(trip)
FROM trips;
```

**Temporal comparisons：**

```sql
-- Time periods when temperature exceeded 30 degrees
SELECT atValue(temperature, true)
FROM (SELECT tfloat '[20@08:00, 35@09:00, 25@10:00]' #> 30.0 AS temperature) t;
```

### Spatial-Temporal 操作

**Trajectory：提取空间路径为 geometry：**

```sql
SELECT ST_AsText(trajectory(trip))
FROM trips
WHERE vehicle_id = 42;
```

**速度计算：**

```sql
-- Speed in units per second (m/s for geographic points)
SELECT speed(trip)
FROM trips
WHERE vehicle_id = 42;
```

**轨迹长度：**

```sql
SELECT length(trip)
FROM trips
WHERE vehicle_id = 42;
```

**Space-time bounding box（stbox）：**

```sql
-- Get the space-time bounding box
SELECT stbox(trip)
FROM trips;

-- Construct an stbox for querying
SELECT stbox(
    ST_MakeEnvelope(2.2, 48.8, 2.4, 48.9, 4326),
    tstzspan '[2025-06-01, 2025-06-02]'
);
```

**空间限制：区域内的值：**

```sql
-- Portions of a trip within a polygon
SELECT atGeometry(trip, ST_Buffer(ST_Point(2.35, 48.86, 4326), 0.01))
FROM trips;
```

**两个 temporal points 之间的距离：**

```sql
SELECT distance(t1.trip, t2.trip)
FROM trips t1, trips t2
WHERE t1.vehicle_id = 1 AND t2.vehicle_id = 2;
```

**最近接近距离和时间：**

```sql
SELECT nearestApproachDistance(t1.trip, t2.trip),
       nearestApproachInstant(t1.trip, t2.trip)
FROM trips t1, trips t2
WHERE t1.vehicle_id = 1 AND t2.vehicle_id = 2;
```

### 索引

MobilityDB 支持 GiST 和 SP-GiST indexes，用于高效 temporal 与 spatio-temporal 查询。

**Temporal types（time dimension）的 SP-GiST index：**

```sql
CREATE INDEX ON measurements USING spgist(temperature);
```

**Spatio-temporal types（space + time）的 GiST index：**

```sql
CREATE INDEX ON trips USING gist(trip);
```

这些索引可加速 bounding box 查询、temporal overlap 检查和 spatial-temporal intersection：

```sql
-- Uses GiST index for space-time filtering
SELECT vehicle_id
FROM trips
WHERE trip && stbox(
    ST_MakeEnvelope(2.2, 48.8, 2.4, 48.9, 4326),
    tstzspan '[2025-06-01, 2025-06-02]'
);
```

### 示例：车辆跟踪

存储和查询车辆 GPS 轨迹的完整示例：

```sql
CREATE TABLE vehicles (
    vehicle_id  INT PRIMARY KEY,
    plate       TEXT,
    type        TEXT
);

CREATE TABLE trips (
    trip_id     BIGSERIAL PRIMARY KEY,
    vehicle_id  INT REFERENCES vehicles(vehicle_id),
    trip        tgeompoint,
    trip_date   DATE
);

CREATE INDEX ON trips USING gist(trip);

-- Insert a trip as a sequence of GPS points
INSERT INTO trips (vehicle_id, trip, trip_date) VALUES (
    1,
    tgeompoint_seq(ARRAY[
        tgeompoint_inst(ST_Point(2.3522, 48.8566, 4326), '2025-06-01 08:00+00'),
        tgeompoint_inst(ST_Point(2.2945, 48.8584, 4326), '2025-06-01 08:15+00'),
        tgeompoint_inst(ST_Point(2.3364, 48.8606, 4326), '2025-06-01 08:30+00'),
        tgeompoint_inst(ST_Point(2.3488, 48.8534, 4326), '2025-06-01 08:45+00')
    ]),
    '2025-06-01'
);

-- Where was vehicle 1 at 08:20?
SELECT valueAtTimestamp(trip, '2025-06-01 08:20+00')
FROM trips WHERE vehicle_id = 1 AND trip_date = '2025-06-01';

-- What was the average speed?
SELECT twAvg(speed(trip))
FROM trips WHERE vehicle_id = 1 AND trip_date = '2025-06-01';

-- Total distance traveled
SELECT length(trip)
FROM trips WHERE vehicle_id = 1 AND trip_date = '2025-06-01';

-- Get the full trajectory as a LineString
SELECT ST_AsGeoJSON(trajectory(trip))
FROM trips WHERE vehicle_id = 1 AND trip_date = '2025-06-01';
```

### 示例：Spatio-Temporal Intersection 查询

查找在给定时间窗口内经过特定区域的所有 trips：

```sql
-- Define area of interest: a circle around the Eiffel Tower
WITH area AS (
    SELECT ST_Buffer(ST_Point(2.2945, 48.8584, 4326)::geography, 500)::geometry AS geom
)
SELECT t.vehicle_id,
       t.trip_date,
       atGeometry(t.trip, a.geom) AS trip_in_area,
       length(atGeometry(t.trip, a.geom)) AS distance_in_area
FROM trips t, area a
WHERE t.trip && stbox(
    a.geom,
    tstzspan '[2025-06-01 07:00+00, 2025-06-01 10:00+00]'
)
  AND eIntersects(t.trip, a.geom)
ORDER BY t.trip_date;
```

### 聚合函数

MobilityDB 提供 temporal aggregates：

```sql
-- Time-weighted average of a temporal float
SELECT twAvg(temperature) FROM sensor_data WHERE sensor_id = 1;

-- Merge multiple temporal points into one
SELECT tUnion(trip) FROM trips WHERE vehicle_id = 1 AND trip_date = '2025-06-01';

-- Centroid of a set of temporal points at each timestamp
SELECT tCentroid(trip) FROM trips WHERE trip_date = '2025-06-01';
```

### 注意事项

- catalog package 和 extension 都是 `mobilitydb`，版本 `1.3.0`；打包矩阵面向 PostgreSQL 14 到 18，并要求 `postgis`。
- v1.3.0 release 增加 PostgreSQL 18 和 PostGIS 3.6 支持，但 migration note 说明二进制格式相较 MobilityDB 1.2 已改变，因此从 1.2 升级需要 backup and restore。
- 上游源码构建说明展示了在加载 MobilityDB 前设置 `shared_preload_libraries = 'postgis-3'` 和 `max_locks_per_transaction = 128`。对于未使用打包默认值的集群，请验证这些设置。
- 本地 package metadata 仍带有 curation comment `need another schema`；上游文档未确认必须使用单独 schema，因此在该备注解决前应避免给出 schema-specific guidance。
