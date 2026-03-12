

## 用法

> [pgpointcloud/pointcloud: 用于存储点云（LIDAR）数据的 PostgreSQL 扩展](https://github.com/pgpointcloud/pointcloud)
>
> [文档](https://pgpointcloud.github.io/pointcloud/)

PostgreSQL Pointcloud 在 PostgreSQL 中存储点云（LIDAR）数据。它引入了两种新的数据类型：`PcPoint` 表示单个点，`PcPatch` 表示点的集合。数据通过 schema 文档来组织，描述每个点的维度和编码方式。

```sql
CREATE EXTENSION pointcloud;
-- 用于 PostGIS 集成：
CREATE EXTENSION pointcloud_postgis;
```


## 核心概念

### Schema

PostgreSQL Pointcloud 使用"schema 文档"来描述特定 LIDAR 点的内容。每个点包含多个维度，每个维度可以是任意数据类型，并可应用缩放和/或偏移量来转换实际值与数据库存储值。schema 文档格式与 [PDAL](https://pdal.io/) 库使用的格式相同。

以下是一个简单的 4 维 schema 文档：

```sql
INSERT INTO pointcloud_formats (pcid, srid, schema) VALUES (1, 4326,
'<?xml version="1.0" encoding="UTF-8"?>
<pc:PointCloudSchema xmlns:pc="http://pointcloud.org/schemas/PC/1.1"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <pc:dimension>
    <pc:position>1</pc:position>
    <pc:size>4</pc:size>
    <pc:description>X 坐标，以长整数表示。需要使用头部的
                    缩放和偏移信息来确定双精度值。</pc:description>
    <pc:name>X</pc:name>
    <pc:interpretation>int32_t</pc:interpretation>
    <pc:scale>0.01</pc:scale>
  </pc:dimension>
  <pc:dimension>
    <pc:position>2</pc:position>
    <pc:size>4</pc:size>
    <pc:description>Y 坐标，以长整数表示。需要使用头部的
                    缩放和偏移信息来确定双精度值。</pc:description>
    <pc:name>Y</pc:name>
    <pc:interpretation>int32_t</pc:interpretation>
    <pc:scale>0.01</pc:scale>
  </pc:dimension>
  <pc:dimension>
    <pc:position>3</pc:position>
    <pc:size>4</pc:size>
    <pc:description>Z 坐标，以长整数表示。需要使用头部的
                    缩放和偏移信息来确定双精度值。</pc:description>
    <pc:name>Z</pc:name>
    <pc:interpretation>int32_t</pc:interpretation>
    <pc:scale>0.01</pc:scale>
  </pc:dimension>
  <pc:dimension>
    <pc:position>4</pc:position>
    <pc:size>2</pc:size>
    <pc:description>强度值是脉冲回波幅度的整数表示。
                    该值是可选的且与系统相关。但如果可用，
                    应始终包含在内。</pc:description>
    <pc:name>Intensity</pc:name>
    <pc:interpretation>uint16_t</pc:interpretation>
    <pc:scale>1</pc:scale>
  </pc:dimension>
  <pc:metadata>
    <Metadata name="compression">dimensional</Metadata>
  </pc:metadata>
</pc:PointCloudSchema>');
```

Schema 文档存储在 `pointcloud_formats` 表中，与 `pcid`（点云标识符）关联。每个对象只有一个 `pcid`，用于在 `pointcloud_formats` 中查找 schema。这类似于 PostGIS 中通过 `srid` 解析空间参考系统的方式。

### 点云对象

**PcPoint**：基本点云类型。每个点包含大量维度，但至少有 X 和 Y 坐标。可通过 `PC_AsText(pcpoint)` 渲染为 JSON 格式：

```json
{
    "pcid" : 1,
      "pt" : [0.01, 0.02, 0.03, 4]
}
```

**PcPatch**：`PcPoint` 的集合。LIDAR 数据不是存储数十亿条独立的点记录，而是表示为较少数量的 `PcPatch` 记录。Patch 通过 `PC_AsText(pcpatch)` 渲染：

```json
{
    "pcid" : 1,
     "pts" : [
              [0.02, 0.03, 0.05, 6],
              [0.02, 0.03, 0.05, 8]
             ]
}
```

### 表

```sql
-- 需要上面的 schema 条目，使 pcid==1 存在。

-- 点表
CREATE TABLE points (
    id SERIAL PRIMARY KEY,
    pt PCPOINT(1)
);

-- Patch 表
CREATE TABLE patches (
    id SERIAL PRIMARY KEY,
    pa PCPATCH(1)
);
```

两个系统提供的表：

- `pointcloud_formats`：保存所有 pcid 条目和 schema 文档
- `pointcloud_columns`：显示所有包含点云对象的列的视图

```sql
SELECT * FROM pointcloud_columns;

 schema |    table    | column | pcid | srid |  type
--------+-------------+--------+------+------+---------
 public | points      | pt     |    1 | 4326 | pcpoint
 public | patches     | pa     |    1 | 4326 | pcpatch
```

### 压缩

压缩在 schema 文档的 `<pc:metadata>` 块中声明：

```xml
<pc:metadata>
  <Metadata name="compression">dimensional</Metadata>
</pc:metadata>
```

支持三种压缩方法：

- **None**：使用 schema 文档中描述的类型和格式，以字节数组存储点和 patch。
- **Dimensional**：将 patch 存储为维度数据数组的集合，并应用"适当的"压缩。适合较小的 patch 尺寸。
- **LAZ**（LASZip）：需要 Pointcloud 构建时支持 laz-perf。

如果未声明压缩方式，默认使用 `none`。

维度压缩在内部使用三种方案：游程编码（低变异性）、公共位移除（窄位范围变异性）、以及基于 zlib 的原始 deflate 压缩。


## 函数：Schema

### PC_SchemaGetNDims

`PC_SchemaGetNDims(pcid integer) returns integer` -- 返回维度数量。

```sql
SELECT PC_SchemaGetNDims(1);
-- 18
```

### PC_SchemaIsValid

`PC_SchemaIsValid(xml text) returns boolean` -- 如果点云 schema 有效则返回 true。

```sql
SELECT PC_SchemaIsValid(schema) FROM pointcloud_formats LIMIT 1;
-- t
```


## 函数：PcPoint

### PC_MakePoint

`PC_MakePoint(pcid integer, vals float8[]) returns pcpoint` -- 从 pcid 和双精度数组构造新的 pcpoint。

```sql
SELECT PC_MakePoint(1, ARRAY[-127, 45, 124.0, 4.0]);

-- 010100000064CEFFFF94110000703000000400
```

插入测试值：

```sql
INSERT INTO points (pt)
SELECT PC_MakePoint(1, ARRAY[x,y,z,intensity])
FROM (
  SELECT
  -127+a/100.0 AS x,
    45+a/100.0 AS y,
         1.0*a AS z,
          a/10 AS intensity
  FROM generate_series(1,100) AS a
) AS values;
```

### PC_AsText（点）

`PC_AsText(p pcpoint) returns text` -- 返回点数据的 JSON 版本。

```sql
SELECT PC_AsText('010100000064CEFFFF94110000703000000400'::pcpoint);
-- {"pcid":1,"pt":[-127,45,124,4]}
```

### PC_Get

`PC_Get(pt pcpoint) returns float8[]` -- 返回所有维度值的数组。

```sql
SELECT PC_Get('010100000064CEFFFF94110000703000000400'::pcpoint);
-- {-127,45,124,4}
```

`PC_Get(pt pcpoint, dimname text) returns numeric` -- 返回指定维度的值。

```sql
SELECT PC_Get('010100000064CEFFFF94110000703000000400'::pcpoint, 'Intensity');
-- 4
```

### PC_MemSize（点）

`PC_MemSize(pt pcpoint) returns int4` -- 返回 pcpoint 的内存大小。

```sql
SELECT PC_MemSize(PC_MakePoint(1, ARRAY[-127, 45, 124.0, 4.0]));
-- 25
```

### PC_PCId（点）

`PC_PCId(p pcpoint) returns integer` -- 返回该点的 pcid schema 编号。

```sql
SELECT PC_PCId('010100000064CEFFFF94110000703000000400'::pcpoint);
-- 1
```


## 函数：PcPatch

### PC_Patch

`PC_Patch(pts pcpoint[]) returns pcpatch` -- 将 pcpoint 值聚合为 pcpatch 的聚合函数。

```sql
INSERT INTO patches (pa)
SELECT PC_Patch(pt) FROM points GROUP BY id/10;
```

### PC_MakePatch

`PC_MakePatch(pcid integer, vals float8[]) returns pcpatch` -- 从 pcid 和双精度数组构造新的 pcpatch（数组大小必须是维度数的倍数）。

```sql
SELECT PC_AsText(PC_MakePatch(1, ARRAY[-126.99,45.01,1,0, -126.98,45.02,2,0, -126.97,45.03,3,0]));

-- {"pcid":1,"pts":[
--  [-126.99,45.01,1,0],[-126.98,45.02,2,0],[-126.97,45.03,3,0]
-- ]}
```

### PC_AsText（patch）

`PC_AsText(p pcpatch) returns text` -- 返回 patch 数据的 JSON 版本。

```sql
SELECT PC_AsText(pa) FROM patches LIMIT 1;

-- {"pcid":1,"pts":[
--  [-126.99,45.01,1,0],[-126.98,45.02,2,0],[-126.97,45.03,3,0],
--  [-126.96,45.04,4,0],[-126.95,45.05,5,0],[-126.94,45.06,6,0],
--  [-126.93,45.07,7,0],[-126.92,45.08,8,0],[-126.91,45.09,9,0]
-- ]}
```

### PC_Summary

`PC_Summary(p pcpatch) returns text` -- 返回 patch 数据的 JSON 格式摘要。

```sql
SELECT PC_Summary(pa) FROM patches LIMIT 1;

-- {"pcid":1, "npts":9, "srid":4326, "compr":"dimensional",
--  "dims":[{"pos":0,"name":"X","size":4,"type":"int32_t","compr":"sigbits",
--           "stats":{"min":-126.99,"max":-126.91,"avg":-126.95}}, ...]}
```

### PC_NumPoints

`PC_NumPoints(p pcpatch) returns integer` -- 返回 patch 中的点数。

```sql
SELECT PC_NumPoints(pa) FROM patches LIMIT 1;
-- 9
```

### PC_PCId（patch）

`PC_PCId(p pcpatch) returns integer` -- 返回 patch 的 pcid schema 编号。

### PC_MemSize（patch）

`PC_MemSize(p pcpatch) returns int4` -- 返回 pcpatch 的内存大小。

### PC_Explode

`PC_Explode(p pcpatch) returns SetOf[pcpoint]` -- 将 patch 转换为独立点记录的集合返回函数。

```sql
SELECT PC_AsText(PC_Explode(pa)), id
FROM patches WHERE id = 7;

              pc_astext               | id
--------------------------------------+----
 {"pcid":1,"pt":[-126.5,45.5,50,5]}   |  7
 {"pcid":1,"pt":[-126.49,45.51,51,5]} |  7
 {"pcid":1,"pt":[-126.48,45.52,52,5]} |  7
 ...
```

### PC_PointN

`PC_PointN(p pcpatch, n int4) returns pcpoint` -- 返回第 n 个点（从 1 开始）。负数 n 从末尾计数。

### PC_Range

`PC_Range(p pcpatch, start int4, n int4) returns pcpatch` -- 返回从第 start 个点开始包含 n 个点的 patch。

### PC_Union

`PC_Union(p pcpatch[]) returns pcpatch` -- 将多个 pcpatch 合并为单个 pcpatch 的聚合函数。

```sql
SELECT PC_NumPoints(PC_Union(pa)) FROM patches;
-- 100
```

### PC_Intersects（patch-patch）

`PC_Intersects(p1 pcpatch, p2 pcpatch) returns boolean` -- 如果 p1 的边界与 p2 的边界相交则返回 true。

### PC_PatchAvg

`PC_PatchAvg(p pcpatch, dimname text) returns numeric` -- 返回所有点中指定维度的平均值。

```sql
SELECT PC_PatchAvg(pa, 'intensity') FROM patches WHERE id = 7;
-- 5.0000000000000000
```

`PC_PatchAvg(p pcpatch) returns pcpoint` -- 返回每个维度平均值的 PcPoint。

```sql
SELECT PC_AsText(PC_PatchAvg(pa)) FROM patches WHERE id = 7;
-- {"pcid":1,"pt":[-126.46,45.54,54.5,5]}
```

### PC_PatchMin

`PC_PatchMin(p pcpatch, dimname text) returns numeric` -- 返回指定维度的最小值。

`PC_PatchMin(p pcpatch) returns pcpoint` -- 返回每个维度最小值的 PcPoint。

### PC_PatchMax

`PC_PatchMax(p pcpatch, dimname text) returns numeric` -- 返回指定维度的最大值。

`PC_PatchMax(p pcpatch) returns pcpoint` -- 返回每个维度最大值的 PcPoint。

### PC_FilterGreaterThan

`PC_FilterGreaterThan(p pcpatch, dimname text, float8 value) returns pcpatch` -- 过滤值大于给定值的点。

```sql
SELECT PC_AsText(PC_FilterGreaterThan(pa, 'y', 45.57))
FROM patches WHERE id = 7;

-- {"pcid":1,"pts":[[-126.42,45.58,58,5],[-126.41,45.59,59,5]]}
```

### PC_FilterLessThan

`PC_FilterLessThan(p pcpatch, dimname text, float8 value) returns pcpatch` -- 过滤值小于给定值的点。

### PC_FilterEquals

`PC_FilterEquals(p pcpatch, dimname text, float8 value) returns pcpatch` -- 过滤值等于给定值的点。

### PC_FilterBetween

`PC_FilterBetween(p pcpatch, dimname text, float8 value1, float8 value2) returns pcpatch` -- 过滤值在 value1 和 value2 之间（不含）的点。

### PC_Sort

`PC_Sort(p pcpatch, dimnames text[]) returns pcpatch` -- 返回按给定维度字典排序的 patch 副本。

### PC_IsSorted

`PC_IsSorted(p pcpatch, dimnames text[], strict boolean default true) returns boolean` -- 检查 pcpatch 是否按字典顺序排序。`strict` 选项检查是否无重复。

### PC_SetPCId

`PC_SetPCId(p pcpatch, pcid int4, def float8 default 0.0) returns pcpatch` -- 设置 PcPatch 的 schema。对于新 schema 中有但旧 schema 中没有的维度，使用值 `def`（默认 `0.0`）。

### PC_Transform

`PC_Transform(p pcpatch, pcid int4, def float8 default 0.0) returns pcpatch` -- 返回基于目标 schema 转换数据的新 patch。与 `PC_SetPCId` 不同，如果解释方式、缩放或偏移不同，可能会改变 patch 数据。

### PC_Compress

`PC_Compress(p pcpatch, global_compression_scheme text, compression_config text) returns pcpatch` -- 使用手动指定的方案压缩 patch。

允许的全局压缩方案：`auto`、`laz`、`dimensional`。对于 dimensional，配置是每个维度压缩方式的逗号分隔列表：`auto`、`zlib`、`sigbits`、`rle`。

### PC_Uncompress

`PC_Uncompress(p pcpatch) returns pcpatch` -- 返回未压缩版本（压缩类型 `none`）。必须作为查询中的最外层函数才能在线路上返回未压缩数据。


## 函数：WKB

### PC_AsBinary（点）

`PC_AsBinary(p pcpoint) returns bytea` -- 返回点的 OGC "well-known binary" 格式。

### PC_EnvelopeAsBinary

`PC_EnvelopeAsBinary(p pcpatch) returns bytea` -- 返回 patch 二维边界的 OGC WKB。

注意：`PC_Envelope` 是 `PC_EnvelopeAsBinary` 的已弃用别名。

### PC_BoundingDiagonalAsBinary

`PC_BoundingDiagonalAsBinary(p pcpatch) returns bytea` -- 返回 patch 边界对角线的 OGC WKB。


## 函数：PostGIS 集成

`pointcloud_postgis` 扩展添加了将 Pointcloud 与 PostGIS 配合使用的函数，可将 `PcPoint` 和 `PcPatch` 转换为 Geometry 并进行空间过滤。

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pointcloud;
CREATE EXTENSION pointcloud_postgis;
```

### Geometry 转换

`Geometry(pcpoint) returns geometry` / `pcpoint::geometry` -- 将 PcPoint 转换为 PostGIS geometry，映射 x/y/z/m。

```sql
SELECT ST_AsText(PC_MakePoint(1, ARRAY[-127, 45, 124.0, 4.0])::geometry);
-- POINT Z (-127 45 124)
```

### PC_EnvelopeGeometry

`PC_EnvelopeGeometry(pcpatch) returns geometry` -- 返回 2D 边界的 PostGIS 二维多边形。

```sql
SELECT ST_AsText(PC_EnvelopeGeometry(pa)) FROM patches LIMIT 1;
-- POLYGON((-126.99 45.01,-126.99 45.09,-126.91 45.09,-126.91 45.01,-126.99 45.01))
```

适合创建索引：

```sql
CREATE INDEX ON patches USING GIST(PC_EnvelopeGeometry(patch));
```

### PC_BoundingDiagonalGeometry

`PC_BoundingDiagonalGeometry(pcpatch) returns geometry` -- 返回边界对角线的 LineString（根据可用维度为 2D/Z/M/ZM）。

```sql
SELECT ST_AsText(PC_BoundingDiagonalGeometry(pa)) FROM patches;
-- LINESTRING Z (-126.99 45.01 1,-126.91 45.09 9)
```

适合创建 N 维索引：

```sql
CREATE INDEX ON patches USING GIST(PC_BoundingDiagonalGeometry(patch) gist_geometry_ops_nd);
```

### PC_Intersection

`PC_Intersection(pcpatch, geometry) returns pcpatch` -- 返回仅包含与几何体相交的点的 PcPatch。

```sql
SELECT PC_AsText(PC_Explode(PC_Intersection(
      pa,
      'SRID=4326;POLYGON((-126.451 45.552, -126.42 47.55, -126.40 45.552, -126.451 45.552))'::geometry
)))
FROM patches WHERE id = 7;

             pc_astext
--------------------------------------
 {"pcid":1,"pt":[-126.44,45.56,56,5]}
 {"pcid":1,"pt":[-126.43,45.57,57,5]}
 {"pcid":1,"pt":[-126.42,45.58,58,5]}
 {"pcid":1,"pt":[-126.41,45.59,59,5]}
```

### PC_Intersects（patch-geometry）

`PC_Intersects(p pcpatch, g geometry) returns boolean` / `PC_Intersects(g geometry, p pcpatch) returns boolean` -- 如果 patch 边界与几何体相交则返回 true。

```sql
SELECT PC_Intersects('SRID=4326;POINT(-126.451 45.552)'::geometry, pa)
FROM patches WHERE id = 7;
-- t
```


## 函数：工具

### PC_Version / PC_Lib_Version / PC_Script_Version

```sql
SELECT PC_Version();           -- 1.2.5
SELECT PC_Lib_Version();       -- 1.2.5 2346cc2
SELECT PC_Script_Version();    -- 1.2.5
SELECT PC_Full_Version();
-- POINTCLOUD="1.2.5 2346cc2" PGSQL="170" LIBXML2="2.14.3 LAZPERF enabled=false
```

### PC_Lazperf_Enabled

`PC_Lazperf_Enabled() returns boolean` -- 如果 LAZperf 压缩支持可用则返回 true。


## 使用 PDAL 加载数据

[PDAL](https://pdal.io/) 用于将 LIDAR 文件加载到 PostgreSQL Pointcloud 中。PDAL 管线是声明 reader、filter 和 writer 的 JSON 文件。

加载 LAS 文件的示例管线：

```json
{
  "pipeline":[
    {
      "type":"readers.las",
      "filename":"/home/lidar/st-helens-small.las",
      "spatialreference":"EPSG:26910"
    },
    {
      "type":"filters.chipper",
      "capacity":400
    },
    {
      "type":"writers.pgpointcloud",
      "connection":"host='localhost' dbname='pc' user='lidar' password='lidar' port='5432'",
      "table":"sthsm",
      "compression":"dimensional",
      "srid":"26910"
    }
  ]
}
```

执行：

```bash
pdal pipeline --input pipelinefile.json
```

`filters.chipper` 将无序点分组为紧凑的 patch，以便高效存储。

### PDAL Writer 选项

- **connection**：PostgreSQL 连接字符串
- **table**：写入 patch 的目标表
- **schema**：创建表的 schema（可选）
- **column**：patch 列名（默认：`pa`）
- **compression**：patch 压缩格式（默认：`dimensional`）
- **overwrite**：替换已有表（默认：`true`）
- **srid**：空间参考 ID（默认：`4326`）
- **pcid**：使用已有的 PCID（可选）
- **pre_sql** / **post_sql**：管线前/后执行的 SQL（可选）

### PDAL Reader 选项

- **connection**：PostgreSQL 连接字符串
- **table**：读取 patch 的源表
- **schema**：读取的 schema（可选）
- **column**：读取的列名（默认：`pa`）
- **where**：约束查询的 SQL where 子句（可选）
- **spatialreference**：覆盖数据库 SRID（可选）

读取并导出的示例管线：

```json
{
  "pipeline":[
    {
      "type":"readers.pgpointcloud",
      "connection":"host='localhost' dbname='pc' user='lidar' password='lidar' port='5432'",
      "table":"sthsm",
      "column":"pa",
      "spatialreference":"EPSG:26910",
      "where":"PC_Intersects(pa, ST_MakeEnvelope(560037.36, 5114846.45, 562667.31, 5118943.24, 26910))"
    },
    {
      "type":"writers.text",
      "filename":"/home/lidar/st-helens-small-out.txt"
    }
  ]
}
```
