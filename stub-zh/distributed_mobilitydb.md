## 用法

来源：

- [README.md](https://github.com/mbakli/DistributedMobilityDB/blob/f13cd2f1b4e9aeb8c83a132914560a53aa5c99f9/README.md)
- [sql/distributed_mobilitydb--1.0.sql](https://github.com/mbakli/DistributedMobilityDB/blob/f13cd2f1b4e9aeb8c83a132914560a53aa5c99f9/sql/distributed_mobilitydb--1.0.sql)
- [sql/partitioning/tiling.sql](https://github.com/mbakli/DistributedMobilityDB/blob/f13cd2f1b4e9aeb8c83a132914560a53aa5c99f9/sql/partitioning/tiling.sql)
- [CMakeLists.txt](https://github.com/mbakli/DistributedMobilityDB/blob/f13cd2f1b4e9aeb8c83a132914560a53aa5c99f9/CMakeLists.txt)
- [src/general/shared_library_init.c](https://github.com/mbakli/DistributedMobilityDB/blob/f13cd2f1b4e9aeb8c83a132914560a53aa5c99f9/src/general/shared_library_init.c)
- [src/distributed_mobilitydb.control](https://github.com/mbakli/DistributedMobilityDB/blob/f13cd2f1b4e9aeb8c83a132914560a53aa5c99f9/src/distributed_mobilitydb.control)

`distributed_mobilitydb` 通过 Citus 分布 PostGIS 和 MobilityDB 数据，提供空间/时间切分、分布元数据和规划器钩子，上游标记为开发中。

### 核心用法

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION mobilitydb;
CREATE EXTENSION citus;
CREATE EXTENSION distributed_mobilitydb;
```

### 运行边界

先准备正常工作的 Citus 集群以及兼容的 PostGIS 3+、MobilityDB 1.1+ 和 Citus 10+，再由超级用户创建扩展。`create_spatiotemporal_distributed_table` 根据输入表创建切分后的输出表，表名、分片数、方法、几何列和分段选项决定物理布局。引用表选项复制整表而不切分，并拒绝非默认分片数。

README 要求 PostgreSQL 13+，构建元数据指定 13–14；目录采用这一较窄的构建范围，不据此断言更高版本的兼容性。部分扩展元数据和类型放在 pg_catalog 中。转换真实表前应规划数据重分布、磁盘用量和权限，并协调备份分布数据与元数据。Citus 自身有集群预加载和配置要求；源码未声明本库还需单独预加载。
