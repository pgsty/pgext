


## 用法

来源：[仓库 README](https://github.com/MobilityDB/MobilityDB)、[MobilityDB 1.3 手册](https://mobilitydb.github.io/MobilityDB/master/)、[v1.3.0 版本](https://github.com/MobilityDB/MobilityDB/releases/tag/v1.3.0)

MobilityDB 使用时态和时空数据类型扩展 PostgreSQL 与 PostGIS，使车辆轨迹、传感器读数和随时间变化属性等移动对象数据能够高效存储、索引和查询。

**关键文档：**

- [MobilityDB 手册](https://mobilitydb.github.io/MobilityDB/master/)
- [时态类型](https://mobilitydb.github.io/MobilityDB/master/ch04.html)
- [时空类型](https://mobilitydb.github.io/MobilityDB/master/ch07.html)
- [时态姿态](https://mobilitydb.github.io/MobilityDB/master/ch11.html)
- [时态圆形缓冲区](https://mobilitydb.github.io/MobilityDB/master/ch13.html)
- [索引](https://mobilitydb.github.io/MobilityDB/master/ch10s02.html)
- [MobilityDB 教程](https://mobilitydb.com/documentation/)
- [API 参考](https://mobilitydb.github.io/MobilityDB/master/)

### 入门

MobilityDB 要求 PostGIS。启用两个扩展：

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION mobilitydb;
```

### 时态类型

MobilityDB 为基础类型提供对应的时态变体：

| 时态类型 | 基础类型 | 说明 |
|---------------|-----------|-------------|
| `tbool`       | `boolean` | 随时间变化的布尔值 |
| `tint`        | `integer` | 随时间变化的整数 |
| `tfloat`      | `float`   | 随时间变化的浮点数 |
| `ttext`       | `text`    | 随时间变化的文本 |
| `tgeometry`   | `geometry` | 随时间变化的任意几何对象 |
| `tgeography`  | `geography` | 随时间变化的任意地理对象 |
| `tgeompoint`  | `geometry(Point)` | 随时间变化的几何点 |
| `tgeogpoint`  | `geography(Point)` | 随时间变化的地理点 |
| `tnpoint`     | 网络点 | 随时间变化的网络点 |
| `tcbuffer`    | 圆形缓冲区 | 随时间变化的圆形缓冲区 |
| `tpose`       | 姿态 | 随时间变化的点位置和方向 |
| `trgeometry`  | 刚性几何对象 | 随时间变化的刚性几何对象 |

MobilityDB 1.3 增加了 `tgeometry`、`tgeography`、`tcbuffer`、`tpose` 和 `trgeometry`。`tgeometry` 与 `tgeography` 支持离散插值或阶梯插值，不支持任意几何对象的线性插值。1.3 版本说明将 `tcbuffer`、`tpose` 和 `trgeometry` 标记为实验性功能。

### 时态子类型

每种时态类型都可根据值随时间变化的方式表示为不同子类型：

| 子类型 | 说明 | 示例 |
|---------|-------------|---------|
| **Instant** | 单个时间戳上的单个值 | `'25.5@2025-01-01 08:00'` |
| **Sequence** | 一个时间区间上的连续值 | `'[25.5@08:00, 28.1@09:00, 30.0@10:00]'` |
| **SequenceSet** | 一组互不重叠的序列 | `'{[25.5@08:00, 28.1@09:00], [30.0@11:00, 31.2@12:00]}'` |

序列使用方括号表示包含边界 `[`，使用圆括号表示排除边界 `(`，与 PostgreSQL 范围类型相同。

### 创建时态值

**瞬时值：**

```sql
SELECT tfloat '25.5@2025-06-01 08:00:00+00';
SELECT tgeompoint 'SRID=4326;Point(2.3522 48.8566)@2025-06-01 08:00:00+00';
```

**序列值（连续插值）：**

```sql
SELECT tfloat '[20.0@2025-06-01 08:00, 25.5@2025-06-01 09:00, 22.0@2025-06-01 10:00]';
```

**离散序列（阶梯插值）：**

```sql
SELECT tint 'Interp=Step;[10@2025-06-01 08:00, 20@2025-06-01 09:00, 15@2025-06-01 10:00]';
```

**序列集值：**

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

### 时态操作

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

**获取时态值的时间跨度：**

```sql
SELECT duration(trip), startTimestamp(trip), endTimestamp(trip)
FROM trips;
```

**时态比较：**

```sql
-- Time periods when temperature exceeded 30 degrees
SELECT atValue(temperature, true)
FROM (SELECT tfloat '[20@08:00, 35@09:00, 25@10:00]' #> 30.0 AS temperature) t;
```

### 时空操作

**轨迹：将空间路径提取为几何对象：**

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

**时空边界框（stbox）：**

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

**两个时态点之间的距离：**

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

MobilityDB 支持 GiST 和 SP-GiST 索引，用于高效执行时态与时空查询。

**时态类型（时间维度）的 SP-GiST 索引：**

```sql
CREATE INDEX ON measurements USING spgist(temperature);
```

**时空类型（空间与时间）的 GiST 索引：**

```sql
CREATE INDEX ON trips USING gist(trip);
```

这些索引可加速边界框查询、时态重叠检查和时空相交查询：

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

### 示例：时空相交查询

查找在给定时间窗口内经过特定区域的所有行程：

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

MobilityDB 提供时态聚合函数：

```sql
-- Time-weighted average of a temporal float
SELECT twAvg(temperature) FROM sensor_data WHERE sensor_id = 1;

-- Merge multiple temporal points into one
SELECT tUnion(trip) FROM trips WHERE vehicle_id = 1 AND trip_date = '2025-06-01';

-- Centroid of a set of temporal points at each timestamp
SELECT tCentroid(trip) FROM trips WHERE trip_date = '2025-06-01';
```

### 注意事项

- 目录中的软件包名和扩展名都是 `mobilitydb`，版本为 `1.3.0`；打包矩阵面向 PostgreSQL 14 到 18，并要求 `postgis`。
- v1.3.0 增加 PostgreSQL 18 和 PostGIS 3.6 支持，但迁移说明指出二进制格式相较 MobilityDB 1.2 已改变，因此从 1.2 升级需要备份并恢复。
- 上游源码构建说明展示了在加载 MobilityDB 前设置 `shared_preload_libraries = 'postgis-3'` 和 `max_locks_per_transaction = 128`。对于未使用打包默认值的集群，请验证这些设置。
- 本地软件包元数据仍带有整理备注 `need another schema`；上游文档未确认必须使用单独模式，因此在该备注解决前应避免给出特定模式的指导。
