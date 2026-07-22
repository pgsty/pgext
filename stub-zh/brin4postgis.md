## 用法

来源：

- [Official README and experimental warning](https://github.com/2ndquadrant-it/brin4postgis/blob/fa3259493a7223bb7ffe0daee9d9cce7abc35d6c/README.md)
- [Version 0.0.1 control file](https://github.com/2ndquadrant-it/brin4postgis/blob/fa3259493a7223bb7ffe0daee9d9cce7abc35d6c/brin4postgis.control)
- [Operator and BRIN operator-class SQL](https://github.com/2ndquadrant-it/brin4postgis/blob/fa3259493a7223bb7ffe0daee9d9cce7abc35d6c/brin4postgis--0.0.1.sql)
- [Official regression workflow](https://github.com/2ndquadrant-it/brin4postgis/blob/fa3259493a7223bb7ffe0daee9d9cce7abc35d6c/sql/brin4postgis.sql)

`brin4postgis` 0.0.1 是一个实验性的纯 SQL 操作符类，用于对 PostGIS `box3d` 值和三维几何包围盒建立 BRIN 索引。它可能降低空间有序超大表的扫描成本，但上游明确声明尚未达到生产可用状态。

### 创建并使用索引

上游声明支持 PostgreSQL 9.5 及以上和 PostGIS 2.1.8 及以上。control 文件没有把 `postgis` 声明为扩展依赖，因此必须先安装：

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION brin4postgis;

CREATE INDEX observations_geom_brin
ON observations
USING brin (box3d(geom) brin_geometry_inclusion_ops_box3d);
```

查询必须使用兼容的 `box3d` 操作符，规划器才可能选用该操作符类：

```sql
SELECT id
FROM observations
WHERE box3d(geom) &&& ST_3DMakeBox(
  ST_MakePoint(-10, -10, -10),
  ST_MakePoint( 10,  10,  10)
);
```

依赖它之前，应通过 `EXPLAIN` 确认索引被使用，并把结果与顺序扫描对比。

### 主要对象

- `brin_geometry_inclusion_ops_box3d`：用于 `box3d` 的 BRIN inclusion 操作符类。
- `&&&`、`@`、`~`、`<<`、`>>`、`<<|`、`|>>` 等空间关系操作符。
- `merge(box3d, box3d)`、`contains(box3d, box3d)`、`isempty(box3d)` 等支持函数。

### 运维边界

BRIN 摘要是有损包围盒而不是精确几何，命中的堆行仍需复查。扩展会以 SRID 为零的 `box3d` 保存几何摘要，因此不保留原始 SRID 元数据；坐标值仍采用原空间参考系的单位。没有显式转换时，不应在同一个索引工作流中混用不同 SRID。

SQL 会在所选扩展模式中安装未限定的通用函数名和操作符，并依赖 PostgreSQL 的 BRIN inclusion 支持函数。上游唯一一次发布可追溯到 2015 年。应固定准确源码，检查命名冲突及当前 PostGIS/核心替代方案，并仅在代表性正确性与性能测试通过后使用。
