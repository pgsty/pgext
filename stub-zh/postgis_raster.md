

## 用法

> [PostGIS Raster：PostGIS 的栅格数据支持](https://github.com/postgis/postgis)

PostGIS Raster 为 PostGIS 扩展了栅格（网格）数据支持，可直接在 PostgreSQL 中存储栅格数据。它支持在 SQL 中进行栅格分析、栅格/矢量交互和地图代数运算。

- [栅格参考手册](https://postgis.net/docs/RT_reference.html)

### 安装

```sql
CREATE EXTENSION postgis_raster;
```

--------

## 加载栅格数据

`raster2pgsql` 命令行工具可将栅格文件（GeoTIFF 等）导入 PostgreSQL：

```bash
# 将 GeoTIFF 以 100x100 瓦片导入，创建空间索引，使用 COPY
raster2pgsql -s 4326 -t 100x100 -I -C -M elevation.tif public.dem | psql mydb

# 追加到已有表
raster2pgsql -s 4326 -t 100x100 -a more_data.tif public.dem | psql mydb
```

关键参数：
- `-s <srid>` -- 设置 SRID
- `-t <width>x<height>` -- 将栅格切分为瓦片
- `-I` -- 创建 GiST 空间索引
- `-C` -- 应用栅格约束
- `-M` -- 加载后执行 vacuum analyze

--------

## 查询栅格数据

### 栅格元数据

```sql
-- 获取栅格尺寸和像素大小
SELECT rid,
    ST_Width(rast) AS width,
    ST_Height(rast) AS height,
    ST_ScaleX(rast) AS pixel_size_x,
    ST_ScaleY(rast) AS pixel_size_y,
    ST_NumBands(rast) AS bands,
    ST_SRID(rast) AS srid
FROM dem LIMIT 5;
```

### 像素值

```sql
-- 获取指定点的值
SELECT ST_Value(rast, ST_SetSRID(ST_MakePoint(-73.985, 40.748), 4326)) AS elevation
FROM dem
WHERE ST_Intersects(rast, ST_SetSRID(ST_MakePoint(-73.985, 40.748), 4326));

-- 获取列/行位置的值（波段 1）
SELECT ST_Value(rast, 1, 10, 20) FROM dem WHERE rid = 1;
```

### 波段统计

```sql
SELECT (ST_SummaryStats(rast)).*
FROM dem WHERE rid = 1;
-- 返回：count、sum、mean、stddev、min、max
```

--------

## 栅格处理

### 按矢量几何裁剪栅格

```sql
-- 将栅格裁剪到多边形边界
SELECT ST_Clip(rast, geom) AS clipped_rast
FROM dem, boundaries
WHERE ST_Intersects(rast, geom);
```

### 地图代数

逐像素运算：

```sql
-- 单栅格地图代数：高程分类
SELECT ST_MapAlgebra(rast, 1, NULL,
    'CASE WHEN [rast] > 100 THEN 1 WHEN [rast] > 50 THEN 2 ELSE 3 END') AS classified
FROM dem;

-- 双栅格地图代数：两个 DEM 的差值
SELECT ST_MapAlgebra(a.rast, 1, b.rast, 1, '[rast1] - [rast2]') AS diff
FROM dem_old a, dem_new b
WHERE ST_Intersects(a.rast, b.rast);
```

### 栅格/矢量交互

```sql
-- 将栅格像素转换为矢量点
SELECT (ST_PixelAsPoints(rast)).*
FROM dem WHERE rid = 1;

-- 将栅格转换为多边形（每个唯一值一个）
SELECT (ST_DumpAsPolygons(rast)).*
FROM dem WHERE rid = 1;

-- 栅格与矢量相交并获取值
SELECT p.name, ST_Value(d.rast, p.geom) AS elevation
FROM dem d, points p
WHERE ST_Intersects(d.rast, p.geom);
```

### 重采样与重投影

```sql
-- 重采样为不同的像素大小
SELECT ST_Rescale(rast, 0.001, -0.001) FROM dem;

-- 重投影到不同的 SRID
SELECT ST_Transform(rast, 3857) FROM dem;
```

--------

## 导出栅格

```sql
-- 导出为 GeoTIFF（二进制）
SELECT ST_AsTIFF(rast) FROM dem WHERE rid = 1;

-- 导出为 PNG
SELECT ST_AsPNG(rast) FROM dem WHERE rid = 1;

-- 导出为 JPEG
SELECT ST_AsJPEG(rast) FROM dem WHERE rid = 1;
```
