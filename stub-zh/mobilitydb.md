## 用法

来源：

- [MobilityDB v1.3.1 README](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/README.md)
- [扩展控制文件](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/mobilitydb/sql/mobilitydb.in.control)
- [1.3 版迁移手册](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/doc/introduction.xml)
- [时空 API](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/doc/temporal_spatial_p1.xml)
- [1.3.1 发布与升级说明](https://github.com/MobilityDB/MobilityDB/releases/tag/v1.3.1)
- [1.3.0 至 1.3.1 SQL 迁移](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/mobilitydb/sql/mobilitydb--1.3.0--1.3.1.sql)

`mobilitydb` 1.3.1 为 PostgreSQL 与 PostGIS 增加时态值和移动对象轨迹，可存储随时间变化的属性、还原指定时刻的位置，并为时空边界建立索引。该补丁修复了可能导致后端崩溃的二进制输入漏洞，1.3.0 用户应升级。

### 启用扩展

该发布要求 PostgreSQL 14 及以上、PostGIS 3 及以上，并新增 PostgreSQL 19 构建支持。软件包可用性另行记录。上游要求加载匹配的 PostGIS 动态库，并建议采用以下锁配额：

```conf
shared_preload_libraries = 'postgis-3'
max_locks_per_transaction = 128
```

将 PostGIS 库加入现有预加载列表，重启 PostgreSQL，再使用获授权的管理角色在目标数据库启用两个扩展：

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION mobilitydb;
```

### 存储与查询轨迹

示例使用投影坐标和完整 UTC 时间戳。应用应选择合适的坐标参考系；地理坐标采用不同的距离语义。

```sql
CREATE TABLE trips (
    trip_id bigint PRIMARY KEY,
    trip tgeompoint NOT NULL
);

INSERT INTO trips VALUES (
    1,
    tgeompoint 'SRID=3857;[Point(0 0)@2026-01-01 08:00:00+00,
                         Point(1000 0)@2026-01-01 09:00:00+00]'
);

SELECT valueAtTimestamp(trip, '2026-01-01 08:30:00+00'),
       ST_AsText(trajectory(trip)),
       length(trip),
       speed(trip)
FROM trips;

CREATE INDEX trips_space_time_idx ON trips USING gist (trip);

SELECT trip_id
FROM trips
WHERE trip && stbox(
    ST_MakeEnvelope(-100, -100, 1100, 100, 3857),
    tstzspan '[2026-01-01 08:00:00+00, 2026-01-01 09:00:00+00]'
);
```

边界框操作符可提供索引过滤。若边界重叠不足以满足业务条件，还应补充相应的精确时态或空间谓词。

### 类型与函数索引

- `tbool`、`tint`、`tfloat`、`ttext`：随时间变化的标量值。
- `tgeompoint`、`tgeogpoint`：移动的几何或地理点；构建了可选类型族时，`tnpoint` 表示路网点。
- `tgeometry`、`tgeography`：任意变化的空间值，支持离散或阶梯插值。
- `tcbuffer`、`tpose`、`trgeometry`：1.3 系列的可选实验性空间类型族，不应假定所有构建都包含它们。
- 瞬时值、序列和序列集分别表示一个时间点、一个序列或多个互不重叠的序列。线性插值是否可用取决于类型。
- `valueAtTimestamp`、`startTimestamp`、`endTimestamp`、`duration`：查看时态范围及值。
- `atTime`、`atGeometry`：按时间域或几何范围裁剪值。
- `trajectory`、`length`、`speed`：查看空间路径与运动情况。
- `twAvg`、`tUnion`：时间加权汇总与时态聚合。
- GiST 与 SP-GiST 操作符类可加速其支持的时间与时空边界查询。

### 升级与安全边界

安装新动态库与 SQL 文件后，在每个数据库执行更新：

```sql
ALTER EXTENSION mobilitydb UPDATE TO '1.3.1';
SELECT extversion FROM pg_extension WHERE extname = 'mobilitydb';
```

- 1.3.1 修复 CVE-2026-102639：畸形 WKB 时态值、集合或跨度输入可能越界读取并导致后端崩溃。仅执行 SQL 迁移不能替换有漏洞的动态库；应重新连接或重启已加载旧二进制的进程。
- 迁移会移除五个同基础类型间的 `<->` 操作符及其 `set_distance` 函数，因为它们与 `btree_gist` 提供的操作符冲突。更新前应检查依赖对象，必要时使用对应的 `btree_gist` 操作符。
- 从 1.2 系列升级到 1.3 会改变时态值的二进制格式，必须按照上游流程备份与恢复。1.3.0 至 1.3.1 的原地 SQL 更新不能替代跨主系列迁移。
- 坐标系、插值、时间空洞、包含边界及单位都会影响结果。应按数据模型验证，不能将所有轨迹都视为连续的地理路径。
