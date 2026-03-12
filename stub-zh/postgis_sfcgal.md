

## 用法

> [PostGIS SFCGAL：基于 SFCGAL 的三维几何与高级操作](https://github.com/postgis/postgis)

PostGIS SFCGAL 通过封装 [SFCGAL](https://sfcgal.gitlab.io/SFCGAL/) 库，提供高级的二维和三维空间操作。它增加了对三维几何运算、体积计算、拉伸、三角剖分等核心 PostGIS GEOS 后端不支持的功能。

- [SFCGAL 函数参考手册](https://postgis.net/docs/reference_sfcgal.html)

### 安装

```sql
CREATE EXTENSION postgis_sfcgal;
```

--------

## 三维操作

### 三维交集与差集

```sql
-- 两个实体的三维交集
SELECT ST_3DIntersection(
    ST_GeomFromText('POLYHEDRALSURFACE Z(((0 0 0,1 0 0,1 1 0,0 1 0,0 0 0)),((0 0 1,1 0 1,1 1 1,0 1 1,0 0 1)),((0 0 0,0 1 0,0 1 1,0 0 1,0 0 0)),((1 0 0,1 1 0,1 1 1,1 0 1,1 0 0)),((0 0 0,1 0 0,1 0 1,0 0 1,0 0 0)),((0 1 0,1 1 0,1 1 1,0 1 1,0 1 0)))'),
    ST_GeomFromText('POLYHEDRALSURFACE Z(((0.5 0.5 0.5,1.5 0.5 0.5,1.5 1.5 0.5,0.5 1.5 0.5,0.5 0.5 0.5)),((0.5 0.5 1.5,1.5 0.5 1.5,1.5 1.5 1.5,0.5 1.5 1.5,0.5 0.5 1.5)),((0.5 0.5 0.5,0.5 1.5 0.5,0.5 1.5 1.5,0.5 0.5 1.5,0.5 0.5 0.5)),((1.5 0.5 0.5,1.5 1.5 0.5,1.5 1.5 1.5,1.5 0.5 1.5,1.5 0.5 0.5)),((0.5 0.5 0.5,1.5 0.5 0.5,1.5 0.5 1.5,0.5 0.5 1.5,0.5 0.5 0.5)),((0.5 1.5 0.5,1.5 1.5 0.5,1.5 1.5 1.5,0.5 1.5 1.5,0.5 1.5 0.5)))')
);

-- 三维差集
SELECT ST_3DDifference(solid_a, solid_b) FROM solids;

-- 三维并集
SELECT ST_3DUnion(solid_a, solid_b) FROM solids;
```

### 三维度量

```sql
-- 曲面的三维面积
SELECT ST_3DArea(geom) FROM surfaces;

-- 实体的体积
SELECT ST_Volume(geom) FROM solids;
```

### 拉伸

```sql
-- 将二维多边形拉伸为三维实体
SELECT ST_Extrude(
    ST_GeomFromText('POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))'),
    0, 0, 10  -- dx, dy, dz
);
```

### 细分与三角剖分

```sql
-- 将多边形细分为三角形
SELECT ST_Tesselate(
    ST_GeomFromText('POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))')
);

-- 约束 Delaunay 三角剖分
SELECT ST_ConstrainedDelaunayTriangles(
    ST_GeomFromText('POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))')
);
```

### 其他函数

```sql
-- 多边形的直骨架
SELECT ST_StraightSkeleton(
    ST_GeomFromText('POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))')
);

-- 近似中轴线
SELECT ST_ApproximateMedialAxis(
    ST_GeomFromText('POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))')
);

-- 闵可夫斯基和
SELECT ST_MinkowskiSum(
    ST_GeomFromText('LINESTRING(0 0, 4 0)'),
    ST_GeomFromText('POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))')
);

-- 检查曲面是否为平面
SELECT ST_IsPlanar(geom) FROM surfaces;
```
