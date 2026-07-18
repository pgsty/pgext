## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/accarniel/VagueGeometry/blob/96c4be4cd382b93398c56e0db2a0a01bc409ce5d/README.md)
- [1.1 版本 SQL API](https://github.com/accarniel/VagueGeometry/blob/96c4be4cd382b93398c56e0db2a0a01bc409ce5d/vasavgeom--1.1.sql)
- [扩展 control 文件](https://github.com/accarniel/VagueGeometry/blob/96c4be4cd382b93398c56e0db2a0a01bc409ce5d/vasavgeom.control)
- [项目 wiki](https://github.com/accarniel/VagueGeometry/wiki)

`vasavgeom` 在 PostGIS 之上实现 Vague Spatial Algebra。一个 `vaguegeometry` 包含确定的 kernel 与可能的 conjecture；拓扑谓词返回 `vaguebool`，面积或距离等度量则可返回 `vaguenumeric` 区间。

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION vasavgeom;
WITH v AS (
  SELECT VG_MakeVagueGeom(
    ST_GeomFromText('POLYGON((0 0,0 1,1 1,1 0,0 0))'),
    ST_GeomFromText('POLYGON((-1 -1,-1 2,2 2,2 -1,-1 -1))')
  ) AS geom
)
SELECT VG_AsText(geom) FROM v;
```

API 包含构造与序列化、kernel/conjecture 提取、集合运算、模糊拓扑谓词，以及空间度量的最小值与最大值。kernel 和 conjecture 必须具有兼容的几何类型与 SRID；应用代码应保留该代数的不确定性语义，不能过早转换成普通布尔值或数值。

1.1 是依赖 PostGIS 的旧式原生代码，没有当前兼容矩阵。必须针对准确的 PostgreSQL、PostGIS、GEOS 和编译器版本进行验证，尤其是在长期表或索引中存储自定义二进制类型之前。
