


## 用法

来源：[PGHydro README](https://github.com/pghydro/pghydro/blob/master/README.md)，[pgh_raster SQL](https://github.com/pghydro/pghydro/blob/master/pgh_raster--6.6.sql)

`pgh_raster` 为 PgHydro 增加基于栅格的水文产品。它保存 DEM、flow path、flow direction、flow accumulation 栅格，并提供上游/下游像元追踪、排水区分析、高程查询、高程剖面和栅格裁剪函数。

### 启用

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION postgis_raster;
CREATE EXTENSION pghydro;
CREATE EXTENSION pgh_raster;
```

### 栅格表

扩展创建 `pgh_raster` schema，其中包括：

| 表 | 用途 |
|----|------|
| `pghrt_elevation` | 高程栅格数据 |
| `pghrt_flowpath` | flow path 栅格数据 |
| `pghrt_flowdirection` | flow direction 栅格数据 |
| `pghrt_flowaccumulation` | flow accumulation 栅格数据 |

可以使用 `raster2pgsql` 等 PostGIS Raster 工具加载栅格产品，并确保 SRID 与切片方式匹配排水数据集。

### 分析函数

函数组包括下游和上游像元遍历、上游面积查询、下游排水点和排水区查询、高程像元值、高程剖面，以及栅格裁剪。

```sql
SELECT nspname, proname
FROM pg_proc p
JOIN pg_namespace n ON n.oid = p.pronamespace
WHERE nspname = 'pgh_raster'
ORDER BY proname;
```

### 注意事项

`pgh_raster` 同时依赖 `pghydro` 和 `postgis_raster`。派生水文地貌指标之前，应确保栅格对齐、SRID 和 nodata 处理与矢量排水图层一致。
