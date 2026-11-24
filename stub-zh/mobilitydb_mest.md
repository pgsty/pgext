## 用法

来源：

- [MobilityDB MEST README](https://github.com/MobilityDB/mest/blob/26503db3d1e6a019094da2316eb2862a8215f053/contrib/mobilitydb/README.md)
- [Control](https://github.com/MobilityDB/mest/blob/26503db3d1e6a019094da2316eb2862a8215f053/contrib/mobilitydb/mobilitydb_mest.control)
- [SQL 1.0](https://github.com/MobilityDB/mest/blob/26503db3d1e6a019094da2316eb2862a8215f053/contrib/mobilitydb/mobilitydb_mest--1.0.sql)

`mobilitydb_mest` 1.0 为 MobilityDB 的跨度集合、时态值与时态点提供多条目 R 树、四叉树及 k-d 树操作符类。

### 前置条件

已核对的源码要求 PostgreSQL 17 或 18、MobilityDB/MEOS 1.4、`mest`，编译时还需 json-c 开发头文件。控制文件中的扩展依赖为 `mobilitydb` 和 `mest`。

保留原有条目，加入已安装的 PostGIS 与 MobilityDB 库并重启 PostgreSQL。上游示例为：

```ini
shared_preload_libraries = 'postgis-3,libMobilityDB-1.4'
```

### 基本流程

使用具有相应权限的会话启用扩展。下面的表用于演示时态跨度集合索引：

```sql
CREATE EXTENSION mobilitydb_mest CASCADE;
CREATE TABLE recorded_spans (id bigint, t tstzspanset);
CREATE INDEX recorded_spans_mrtree ON recorded_spans
USING mgist (t tstzspanset_mrtree_equisplit_ops (num_spans = 3));
```

对于包含时态点列的表：

```sql
CREATE TABLE recorded_trips (id bigint, trip tgeompoint);
CREATE INDEX recorded_trips_mquadtree ON recorded_trips
USING mspgist (trip tgeompoint_mquadtree_equisplit_ops (num_boxes = 3));
```

### 操作符类与维护

`mgist`、`mspgist` 由 `mest` 提供。本扩展增加特定类型的操作符类，如 `tstzspanset_mrtree_equisplit_ops`、`tgeompoint_mquadtree_equisplit_ops`；安装 SQL 还定义了时态数值和时态地理点的拆分策略。

`num_spans`、`num_boxes` 控制拆分，影响索引大小。应根据代表性查询选择操作符类和拆分策略，不能假定一种配置始终最优。已核对的源码要求 MobilityDB 1.4；发行渠道提供的更早 MobilityDB 版本不能视为已验证的替代品。
