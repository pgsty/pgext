

## 用法

> 来源：[`segasai/q3c`](https://github.com/segasai/q3c) | [ADASS 论文](http://adsabs.harvard.edu/abs/2006ASPC..351..735K) | [ASCL](https://ascl.net/1905.008)

Q3C（Quad Tree Cube）是一个用于天文星表快速天球索引的 PostgreSQL 扩展。它能够对球面坐标（赤经和赤纬）进行高效的空间查询，包括锥形搜索、椭圆搜索、多边形查询、位置交叉匹配和最近邻查找。

所有角度（ra、dec、距离）单位为**度**，自行单位为**毫角秒/年**，历元单位为**年**（如 2000.5、2010.5）。所有 Q3C 函数名以 `q3c_` 为前缀。

### 表准备

要使用 Q3C，在包含 `ra` 和 `dec` 列（单位为度）的表上创建空间索引：

```sql
CREATE INDEX ON mytable (q3c_ang2ipix(ra, dec));
```

可选择按索引对表进行聚簇，以确保大数据集上更快的查询：

```sql
CLUSTER mytable_q3c_ang2ipix_idx ON mytable;
```

或在建索引前重新排列表：

```sql
CREATE TABLE mytable1 AS SELECT * FROM mytable ORDER BY q3c_ang2ipix(ra, dec);
```

建索引后，分析表：

```sql
ANALYZE mytable;
```


## 函数

- `q3c_ang2ipix(ra, dec)` -- 返回给定 ra 和 dec 对应的 ipix 值（64 位整数像素标识符）

- `q3c_dist(ra1, dec1, ra2, dec2)` -- 返回两点之间的距离（度）

- `q3c_dist_pm(ra1, dec1, pmra1, pmdec1, cosdec_flag, epoch1, ra2, dec2, epoch2)` -- 返回考虑自行的两点距离（度）。`cosdec_flag`（0 或 1）指示自行是否包含 cos(dec) 项（1）或不包含（0）

- `q3c_join(ra1, dec1, ra2, dec2, radius)` -- 如果 (ra1, dec1) 在 (ra2, dec2) 的 `radius` 球面距离内则返回 true。需要在 `q3c_ang2ipix(ra2, dec2)` 上创建索引

- `q3c_join_pm(ra1, dec1, pmra1, pmdec1, cosdec_flag, epoch1, ra2, dec2, epoch2, max_delta_epoch, radius)` -- 类似 `q3c_join` 但考虑自行。`max_delta_epoch` 是两个表之间可能的最大历元差

- `q3c_ellipse_join(ra1, dec1, ra2, dec2, major, ratio, pa)` -- 类似 `q3c_join`，但 (ra1, dec1) 必须在以半长轴 `major`、轴比 `ratio` 和位置角 `pa`（从北向东）定义的椭圆内

- `q3c_radial_query(ra, dec, center_ra, center_dec, radius)` -- 如果 (ra, dec) 在 (center_ra, center_dec) 的 `radius` 度范围内则返回 true。锥形搜索的主要函数。需要在 `q3c_ang2ipix(ra, dec)` 上建索引

- `q3c_ellipse_query(ra, dec, center_ra, center_dec, maj_ax, axis_ratio, PA)` -- 如果 (ra, dec) 在以 (center_ra, center_dec) 为中心、由半长轴、轴比和位置角定义的椭圆内则返回 true

- `q3c_poly_query(ra, dec, poly)` -- 如果 (ra, dec) 在以 RA/DEC 值数组或 PostgreSQL polygon 类型指定的球面多边形内则返回 true。使用索引

- `q3c_ipix2ang(ipix)` -- 返回给定 ipix 对应的 (ra, dec) 两元素数组

- `q3c_pixarea(ipix, bits)` -- 返回给定 ipix 在 `bits` 指定的像素化级别对应的球面面积（1 最小，30 是立方体面）

- `q3c_ipixcenter(ra, dec, bits)` -- 返回覆盖指定 (ra, dec) 的特定像素深度的像素中心 ipix 值

- `q3c_in_poly(ra, dec, poly)` -- 返回点是否在多边形内。**不**使用 q3c 索引

- `q3c_version()` -- 返回安装的 Q3C 版本


## 示例

### 锥形搜索

查询 (ra, dec) = (11, 12) 附近 0.1 度内的所有天体：

```sql
SELECT * FROM mytable WHERE q3c_radial_query(ra, dec, 11, 12, 0.1);
```

表的列名必须在前，搜索位置在后，否则索引不会被使用。

使用 `q3c_join` 的替代锥形搜索（对小表可能更快）：

```sql
SELECT * FROM mytable WHERE q3c_join(11, 12, ra, dec, 0.1);
```

### 椭圆搜索

搜索以 (10, 20) 为中心、半长轴 1 度、轴比 0.5、位置角 10 度的椭圆内的天体：

```sql
SELECT * FROM mytable WHERE q3c_ellipse_query(ra, dec, 10, 20, 1, 0.5, 10);
```

### 多边形搜索

查询顶点为 (0,0)、(2,0)、(2,1)、(0,1) 的球面多边形内的天体：

```sql
SELECT * FROM mytable WHERE
    q3c_poly_query(ra, dec, ARRAY[0, 0, 2, 0, 2, 1, 0, 1]);
```

使用 PostgreSQL polygon 类型：

```sql
SELECT * FROM mytable WHERE
    q3c_poly_query(ra, dec, '((0, 0), (2, 0), (2, 1), (0, 1))'::polygon);
```

### 位置交叉匹配

在 0.001 度范围内交叉匹配 `table1` 和 `table2`。索引需存在于 `table2` 的 `q3c_ang2ipix(ra, dec)` 上：

```sql
SELECT * FROM table1 AS a, table2 AS b WHERE
    q3c_join(a.ra, a.dec, b.ra, b.dec, 0.001);
```

索引表的 ra/dec 列必须是第 3 和第 4 个参数。这会返回匹配距离内的**所有**配对，而不仅是最近邻。

使用逐对象误差半径：

```sql
SELECT * FROM table1 AS a, table2 AS b WHERE
    q3c_join(a.ra, a.dec, b.ra, b.dec, a.err);
```

### 椭圆交叉匹配

使用椭圆误差区域进行交叉匹配（如在星系椭圆体内匹配）：

```sql
SELECT * FROM table1 AS a, table2 AS b WHERE
    q3c_ellipse_join(a.ra, a.dec, b.ra, b.dec, a.maj_ax, a.axis_ratio, a.PA);
```

### 含自行的交叉匹配

考虑自行修正的交叉匹配。假设 `table1` 有 `pmra`、`pmdec`（毫角秒/年）和 `epoch` 列，pmra 包含 cos(dec) 因子，最大历元差为 30 年：

```sql
SELECT * FROM table1 AS a, table2 AS b WHERE
    q3c_join_pm(a.ra, a.dec, a.pmra, a.pmdec, 1,
    a.epoch, b.ra, b.dec, b.epoch, 30, 0.001);
```

### 最近邻（未匹配返回 NULL）

为每行返回最近邻，1 角秒内无匹配则返回 NULL：

```sql
SELECT t.*, ss.* FROM mytable AS t
LEFT JOIN LATERAL (
    SELECT s.*
    FROM sdssdr9.phototag AS s
    WHERE q3c_join(t.ra, t.dec, s.ra, s.dec, 1./3600)
    ORDER BY q3c_dist(t.ra, t.dec, s.ra, s.dec) ASC
    LIMIT 1
) AS ss ON true;
```

### 最近邻（仅匹配项）

仅返回有邻居的天体：

```sql
SELECT t.*, ss.* FROM mytable AS t,
LATERAL (
    SELECT s.*
    FROM sdssdr9.phototag AS s
    WHERE q3c_join(t.ra, t.dec, s.ra, s.dec, 1./3600)
    ORDER BY q3c_dist(t.ra, t.dec, s.ra, s.dec) ASC
    LIMIT 1
) AS ss;
```

### 最近邻（CTE 变体）

使用带有对象 ID 列的 CTE（需要在 ID 列上建索引）：

```sql
WITH x AS MATERIALIZED (
    SELECT *, (
        SELECT objid FROM sdssdr9.phototag AS p
        WHERE q3c_join(m.ra, m.dec, p.ra, p.dec, 1./3600)
        ORDER BY q3c_dist(m.ra, m.dec, p.ra, p.dec) ASC
        LIMIT 1
    ) AS match_objid
    FROM mytable AS m
)
SELECT * FROM x, sdssdr9.phototag AS s WHERE x.match_objid = s.objid;
```

### 密度估计

使用像素化深度 25 估计天体密度：

```sql
SELECT (q3c_ipix2ang(i))[1] AS ra,
       (q3c_ipix2ang(i))[2] AS dec,
       c,
       q3c_pixarea(i, 25) AS area
FROM (
    SELECT q3c_ipixcenter(ra, dec, 25) AS i, count(*) AS c
    FROM mytable
    GROUP BY i
) AS x;
```

注意：Q3C 的像素面积不均匀（与 HEALPIX 不同）。


## 限制

- 不支持查询直径大于约 25 度的超大多边形
- 不支持超过 100 个顶点的多边形


## 性能提示

- 确保 Q3C 函数中的参数顺序正确（如 `q3c_radial_query(ra, dec, 120, 3, 1)` 而非 `q3c_radial_query(120, 3, ra, dec, 1)`）
- 使用 `EXPLAIN` 验证查询计划使用了 Q3C 索引的位图扫描
- 如果规划器选择了不佳的计划，尝试：`SET enable_mergejoin TO off; SET enable_seqscan TO off; SET enable_hashjoin TO off;`
- 按 Q3C 索引聚簇表以获得最佳性能
- 当 `q3c_join()` 与额外的过滤条件结合使用时，使用 `MATERIALIZED` CTE 避免计划问题：

```sql
WITH x AS MATERIALIZED (SELECT * FROM t1 WHERE t1.mag < 1),
     y AS (SELECT *, t2.mag AS t2mag FROM x, t2 WHERE q3c_join(x.ra, x.dec, t2.ra, t2.dec, 1./3600))
SELECT * FROM y WHERE t2mag > 33;
```
