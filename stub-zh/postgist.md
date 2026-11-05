## 用法

来源：

- [上游来源 1](https://github.com/postgis-t/postgis-t/blob/f5ecd265fd7ae266a1038d0d1946e334f5657db0/src/postgist/postgist--0.1.0.sql)
- [上游来源 2](https://github.com/postgis-t/postgis-t/blob/f5ecd265fd7ae266a1038d0d1946e334f5657db0/src/postgist/postgist.control)

`postgist` 0.1.0 将带时间戳的测量值与几何对象组合为基于 PostGIS 的时空复合类型。

### 核心流程

`timeseries` 保存测量值、值域、时间区间及空间包络，`trajectory` 保存带时间戳的几何对象。`tst_timeseries` 与 `tst_trajectory` 聚合行，`tst_hull`、`tst_period`、`tst_range` 和 `tst_data` 读取各组成部分。

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION postgist;
SELECT tst_timeseries(value, observed_at, geom)
FROM (VALUES
  (1::numeric, timestamp '2026-01-01', ST_Point(0,0)),
  (2::numeric, timestamp '2026-01-02', ST_Point(1,1))
) AS samples(value, observed_at, geom);
```

### 运行边界

依赖 PostGIS 及创建扩展对象的权限。这是旧版 0.1.0 SQL/PLpgSQL 源码，尚无经过核实的当前 PostgreSQL 主版本支持范围。包围盒重叠辅助函数仅筛选潜在交集，并非精确时空相交判定。依赖结果前，应验证数据集的顺序、坐标系与插值行为。
