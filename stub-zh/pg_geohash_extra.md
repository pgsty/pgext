## 用法

来源：

- [官方 README](https://github.com/siose-innova/pg_geohash_extra/blob/ea5fcbb319503599ef26f67bb165f99b39add763/README.md)
- [安装 SQL 构建列表](https://github.com/siose-innova/pg_geohash_extra/blob/ea5fcbb319503599ef26f67bb165f99b39add763/Makefile)
- [几何到 Geohash 的 SQL](https://github.com/siose-innova/pg_geohash_extra/blob/ea5fcbb319503599ef26f67bb165f99b39add763/sql/functions/geohashfromgeom.sql)
- [邻居实现](https://github.com/siose-innova/pg_geohash_extra/blob/ea5fcbb319503599ef26f67bb165f99b39add763/src/geohash_extra.c)

`pg_geohash_extra` 为 PostGIS 补充查找相邻 Geohash 和用 Geohash 网格覆盖几何的函数。它还安装网格辅助对象与复合类型，供另一个几何覆盖函数使用。

### 核心流程

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pg_geohash_extra;

SELECT * FROM geohash_neighbours('ezs42');

SELECT *
FROM geohashfromgeom(
    ST_MakeEnvelope(-3.72, 40.40, -3.68, 40.43, 4326),
    6
) AS hash;
```

`geohashfromgeom` 会把非 4326 几何转换为 SRID 4326，按请求精度覆盖其边界范围，并返回文本 Geohash。它是网格覆盖辅助函数，不是精确几何分解。

### 对象索引

- `geohash_neighbours(text)` 返回输入网格周围的八个相邻 Geohash。
- `geohashfromgeom(geometry, integer)` 返回覆盖几何范围的文本 Geohash。
- `ST_RegularGrid(geometry, numeric, numeric, boolean)` 以 `T_Grid` 行返回网格点或网格单元。
- `st_geohashfromgeom(geometry, integer)` 返回包含 ID、精度和几何的复合 `geohash` 行。
- `geohash` 和 `T_Grid` 是复合辅助类型。

### 运维说明

版本 `0.0.1` 依赖 `postgis`，可重定位，且未声明预加载要求。应校验精度和输入 SRID，并单独测试跨日期变更线、极区、空几何和超大几何；该实现仍处早期阶段，README 列出了大量后续工作。

应先在一次性数据库中审查安装。扩展脚本会对通用类型名 `geohash` 和 `T_Grid` 执行 `DROP TYPE IF EXISTS ... CASCADE`，并将 `ST_RegularGrid` 所有者改为硬编码角色 `postgres`。这些操作可能删除依赖对象、与现有名称冲突、需要提升权限，或在不存在 `postgres` 角色时失败。
