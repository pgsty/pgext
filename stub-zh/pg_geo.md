## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_geo/pg_geo.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_geo/pg_geo--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_geo/meson.build)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_geo/geo_ops.c)

`pg_geo` `1.0` 提供二维几何与包围盒索引。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_geo;
SELECT ST_DistanceSphere(ST_Point(-157.8,21.3), ST_Point(-156.3,20.8));
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

geo_point、geo_bbox、geo_polygon 和 geo_circle 提供相应的 ST_ 谓词与 GiST 包围盒运算。这是小型二维实现，不保证兼容 PostGIS。使用查询结果前，应核对支持的几何语法、坐标假设和精确谓词；public 中的函数名可能与其他空间扩展重叠。

此模块对应上游旧名 `alohadb_geo`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
