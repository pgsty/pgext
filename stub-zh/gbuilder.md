## 用法

来源：

- [锁定提交的上游 README](https://github.com/geo-mathijs/gbuilder/blob/9e0df5e15885949581b43bf931abd6f4fefff2f1/README.md)
- [1.0 版安装 SQL](https://github.com/geo-mathijs/gbuilder/blob/9e0df5e15885949581b43bf931abd6f4fefff2f1/gbuilder--1.0.sql)
- [锁定提交的 C 入口](https://github.com/geo-mathijs/gbuilder/blob/9e0df5e15885949581b43bf931abd6f4fefff2f1/gbuilder.c)
- [锁定提交的凹包实现](https://github.com/geo-mathijs/gbuilder/blob/9e0df5e15885949581b43bf931abd6f4fefff2f1/hull.c)
- [锁定提交的扩展控制文件](https://github.com/geo-mathijs/gbuilder/blob/9e0df5e15885949581b43bf931abd6f4fefff2f1/gbuilder.control)

gbuilder 1.0 是一个 C 扩展，增加了一种几何构造器：针对点输入、基于 Delaunay 的凹包。它依赖 PostGIS 和 Fast Robust Predicates 库。聚合形式使用 alpha 1.0；数组形式接受自定义 alpha。

### 聚合与数组形式

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION gbuilder;

SELECT ConcaveHull(geom)
FROM demo_points;

SELECT ConcaveHull(
  ARRAY[
    ST_SetSRID(ST_Point(0, 0), 4326),
    ST_SetSRID(ST_Point(1, 0), 4326),
    ST_SetSRID(ST_Point(1, 1), 4326),
    ST_SetSRID(ST_Point(0, 1), 4326)
  ],
  1.0
);
```

聚合声明复用 PostGIS 的 geometry 累积器，并在终值函数中始终传入 alpha 1.0。如果要调整 alpha，应显式把点聚合为 geometry 数组，再调用双参数函数。实现会静默跳过 NULL 以及所有非 POINT 几何；没有可用点或没有保留多边形时会返回 NULL。

### 几何与 ABI 限制

C 代码将输出 SRID 初始化为 4326，而不是从输入推导。因此，混合 SRID、非 4326 输入、Z 或 M 维度、重复点、共线点、X 范围为零、点数过少和极端 alpha 值都需要显式验证。应在调用前对点做规范化和校验，再用 PostGIS 有效性和 SRID 函数检查输出。

gbuilder 依赖私有 liblwgeom 结构编译，并重新创建了一个不属于公开 ABI 的 PostGIS 聚合状态结构。它还直接调用 PostgreSQL 和 PostGIS 内部符号。为某一 PostgreSQL/PostGIS 组合构建的二进制文件，在另一组合上可能无法链接、崩溃或损坏内存。必须使用每个节点上实际部署的 PostgreSQL 和 PostGIS 软件包精确构建并做回归测试。

上游对 PostGIS 的性能与精度对比没有在仓库中提供可重现的基准，且最后复核的代码来自 2019 年。在选择它之前，应使用具有代表性的几何数据（包括无效和对抗性输入）与当前 PostGIS 凹包函数比较。应将该扩展视为历史原生代码，首次试用必须隔离在一个可丢弃数据库中。
