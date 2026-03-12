

## 用法

> [mobilitydb: PostgreSQL 的时态与时空数据管理扩展](https://github.com/MobilityDB/MobilityDB)

MobilityDB 通过时态和时空数据类型扩展了 PostgreSQL 和 PostGIS，实现了移动对象数据（如车辆轨迹、传感器读数和时变属性）的高效存储、索引和查询。

**核心文档：**

- [MobilityDB 手册](https://docs.mobilitydb.com/MobilityDB/master/)
- [时态类型](https://docs.mobilitydb.com/MobilityDB/master/mobilitydb-manual.html#temporal-types)
- [时态操作](https://docs.mobilitydb.com/MobilityDB/master/mobilitydb-manual.html#temporal-operations)
- [时空类型](https://docs.mobilitydb.com/MobilityDB/master/mobilitydb-manual.html#spatial-temporal-types)
- [索引](https://docs.mobilitydb.com/MobilityDB/master/mobilitydb-manual.html#indexing)
- [MobilityDB 教程](https://mobilitydb.com/documentation/)
- [API 参考](https://docs.mobilitydb.com/MobilityDB/master/mobilitydb-manual.html)

### 快速开始

MobilityDB 依赖 PostGIS。需要同时启用两个扩展：

```sql
CREATE EXTENSION PostGIS;
CREATE EXTENSION MobilityDB;
```

### 时态类型

MobilityDB 提供基础类型的时态变体：

| 时态类型 | 基础类型 | 描述 |
|----------|----------|------|
| `tbool`       | `boolean` | 时变布尔值 |
| `tint`        | `integer` | 时变整数 |
| `tfloat`      | `float`   | 时变浮点数 |
| `ttext`       | `text`    | 时变文本 |
| `tgeompoint`  | `geometry(Point)` | 时变几何点 |
| `tgeogpoint`  | `geography(Point)` | 时变地理点 |

### 时态子类型

每种时态类型根据值随时间变化的方式可用不同子类型表示：

| 子类型 | 描述 | 示例 |
|--------|------|------|
| **瞬时值（Instant）** | 单个时间戳上的单个值 | `'25.5@2025-01-01 08:00'` |
| **序列（Sequence）** | 时间区间上的连续值 | `'[25.5@08:00, 28.1@09:00, 30.0@10:00]'` |
| **序列集（SequenceSet）** | 不重叠序列的集合 | `'{[25.5@08:00, 28.1@09:00], [30.0@11:00, 31.2@12:00]}'` |

序列使用方括号表示包含 `[` 或排除 `(` 的边界，与 PostgreSQL 范围类型一致。

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

**通过组件构造：**

```sql
SELECT tgeompoint_inst(ST_Point(2.3522, 48.8566, 4326), '2025-06-01 08:00+00');
SELECT tgeompoint_seq(ARRAY[
    tgeompoint_inst(ST_Point(2.3522, 48.8566, 4326), '2025-06-01 08:00+00'),
    tgeompoint_inst(ST_Point(2.2945, 48.8584, 4326), '2025-06-01 08:30+00'),
    tgeompoint_inst(ST_Point(2.3364, 48.8606, 4326), '2025-06-01 09:00+00')
]);
```

### 时态操作

**获取特定时间的值：**

```sql
SELECT valueAtTimestamp(temp, '2025-06-01 08:30:00+00')
FROM (SELECT tfloat '[20.0@08:00, 30.0@09:00]' AS temp) t;
-- 返回 25.0（线性插值）
```

**限制到特定时间段：**

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
-- 温度超过 30 度的时间段
SELECT atValue(temperature, true)
FROM (SELECT tfloat '[20@08:00, 35@09:00, 25@10:00]' #> 30.0 AS temperature) t;
```

### 时空操作

**轨迹：将空间路径提取为几何体：**

```sql
SELECT ST_AsText(trajectory(trip))
FROM trips
WHERE vehicle_id = 42;
```

**速度计算：**

```sql
-- 速度以每秒为单位（地理点为 m/s）
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
-- 获取时空边界框
SELECT stbox(trip)
FROM trips;

-- 构造用于查询的 stbox
SELECT stbox(
    ST_MakeEnvelope(2.2, 48.8, 2.4, 48.9, 4326),
    tstzspan '[2025-06-01, 2025-06-02]'
);
```

**空间限制：获取特定区域内的值：**

```sql
-- 行程中位于某多边形内的部分
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

MobilityDB 支持 GiST 和 SP-GiST 索引，用于高效的时态和时空查询。

**SP-GiST 索引用于时态类型（时间维度）：**

```sql
CREATE INDEX ON measurements USING spgist(temperature);
```

**GiST 索引用于时空类型（空间 + 时间）：**

```sql
CREATE INDEX ON trips USING gist(trip);
```

这些索引可以加速边界框查询、时态重叠检查和时空相交操作：

```sql
-- 使用 GiST 索引进行时空过滤
SELECT vehicle_id
FROM trips
WHERE trip && stbox(
    ST_MakeEnvelope(2.2, 48.8, 2.4, 48.9, 4326),
    tstzspan '[2025-06-01, 2025-06-02]'
);
```

### 示例：车辆追踪

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

-- 将行程插入为一系列 GPS 点
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

-- 车辆 1 在 08:20 的位置？
SELECT valueAtTimestamp(trip, '2025-06-01 08:20+00')
FROM trips WHERE vehicle_id = 1 AND trip_date = '2025-06-01';

-- 平均速度是多少？
SELECT twAvg(speed(trip))
FROM trips WHERE vehicle_id = 1 AND trip_date = '2025-06-01';

-- 总行驶距离
SELECT length(trip)
FROM trips WHERE vehicle_id = 1 AND trip_date = '2025-06-01';

-- 获取完整轨迹作为 LineString
SELECT ST_AsGeoJSON(trajectory(trip))
FROM trips WHERE vehicle_id = 1 AND trip_date = '2025-06-01';
```

### 示例：时空相交查询

查找在指定时间窗口内经过特定区域的所有行程：

```sql
-- 定义感兴趣区域：以埃菲尔铁塔为中心的圆形区域
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
-- 时态浮点数的时间加权平均
SELECT twAvg(temperature) FROM sensor_data WHERE sensor_id = 1;

-- 将多个时态点合并为一个
SELECT tUnion(trip) FROM trips WHERE vehicle_id = 1 AND trip_date = '2025-06-01';

-- 在每个时间戳上计算一组时态点的质心
SELECT tCentroid(trip) FROM trips WHERE trip_date = '2025-06-01';
```
