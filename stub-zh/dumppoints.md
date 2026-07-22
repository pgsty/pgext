## 用法

来源：

- [官方项目说明](https://github.com/nmandery/dumppoints/blob/96343c7e4eb76765da69737f79611fa7355c9e5b/README)
- [扩展控制文件](https://github.com/nmandery/dumppoints/blob/96343c7e4eb76765da69737f79611fa7355c9e5b/dumppoints.control)
- [扩展 SQL](https://github.com/nmandery/dumppoints/blob/96343c7e4eb76765da69737f79611fa7355c9e5b/dumppoints--1.0.sql)

`dumppoints` 版本 `1.0` 安装 PostGIS `ST_DumpPoints(geometry)` 函数的 C 实现。它用于把几何对象的每个顶点展开为 PostGIS `geometry_dump` 行，其中包含路径与点几何对象。

### 核心流程

PostGIS 是声明的依赖，必须先可用。

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION dumppoints;

SELECT (d).path, ST_AsText((d).geom)
FROM (
  SELECT ST_DumpPoints(
    ST_GeomFromText('POLYGON((0 0, 2 0, 2 2, 0 0))')
  ) AS d
) AS q;
```

路径数组按照 PostGIS 的 dump 约定标识组成部分、环与顶点位置；几何字段是对应的点。

### 对象

- `ST_DumpPoints(geometry) returns setof geometry_dump`：为输入中的每个顶点生成一个 `geometry_dump` 值。

### 兼容性说明

安装脚本会对 PostGIS 已暴露的同一签名执行 `CREATE OR REPLACE FUNCTION`。因此，安装 `dumppoints` 会替换目标扩展模式中的数据库函数，而不是增加一个单独命名的 API。应针对实际使用的 PostGIS 版本测试行为，并在升级或删除任一扩展前核对扩展对象归属。不需要预加载。
