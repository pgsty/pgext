

## 用法

> [ogr_fdw: PostgreSQL 的 OGR 外部数据包装器](https://github.com/pramsey/pgsql-ogr-fdw)

OGR 是 [GDAL](http://www.gdal.org/) 空间数据访问库的**矢量**部分。它允许通过简单的 C API 访问[大量 GIS 数据格式](http://www.gdal.org/ogr_formats.html)。由于 OGR 暴露了简单的表结构，而 PostgreSQL 外部数据包装器允许访问表结构，两者的契合非常完美。

### 快速开始

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION ogr_fdw;
```

使用 `ogr_fdw_info` 工具读取 OGR 数据源并输出服务器/表定义：

```bash
ogr_fdw_info -s /tmp/test -l pt_two
```

```sql
CREATE SERVER "myserver"
  FOREIGN DATA WRAPPER ogr_fdw
  OPTIONS (
    datasource '/tmp/test',
    format 'ESRI Shapefile' );

CREATE FOREIGN TABLE "pt_two" (
  fid integer,
  "geom" geometry(Point, 4326),
  "name" varchar,
  "age" integer,
  "height" real,
  "birthdate" date )
  SERVER "myserver"
  OPTIONS (layer 'pt_two');

SELECT * FROM pt_two;
```

支持过滤下推——包括简单谓词和边界框过滤（`&&`）：

```sql
SET client_min_messages = debug1;

SELECT name, age, height
FROM pt_two
WHERE height < 5.7
AND geom && ST_MakeEnvelope(0, 0, 1, 1);
```

```
DEBUG:  OGR SQL: (height < 5.7)
DEBUG:  OGR spatial filter (0 0, 1 1)
```


## 限制

- 需要 PostgreSQL 11 或更高版本
- 仅有限的非空间查询限制会下推到 OGR（仅 `>`、`<`、`<=`、`>=`、`=`）
- 仅边界框过滤（`&&`）会下推到空间过滤
- OGR 连接每次查询都会创建（无连接池）
- 每次都会检索所有列


## 示例

### WFS（Web 要素服务）

```sql
CREATE SERVER geoserver
  FOREIGN DATA WRAPPER ogr_fdw
  OPTIONS (
    datasource 'WFS:https://demo.geo-solutions.it/geoserver/wfs',
    format 'WFS' );

CREATE FOREIGN TABLE topp_states (
  fid bigint,
  the_geom Geometry(MultiSurface,4326),
  gml_id varchar,
  state_name varchar,
  state_fips varchar,
  state_abbr varchar,
  land_km double precision,
  persons double precision )
  SERVER "geoserver"
  OPTIONS (layer 'topp:states');
```

### 文件地理数据库

```sql
CREATE SERVER fgdbtest
  FOREIGN DATA WRAPPER ogr_fdw
  OPTIONS (
    datasource '/tmp/Querying.gdb',
    format 'OpenFileGDB' );

CREATE FOREIGN TABLE cities (
  fid integer,
  geom geometry(Point, 4326),
  city_name varchar,
  state_name varchar,
  elevation integer,
  pop1990 integer )
  SERVER fgdbtest
  OPTIONS (layer 'Cities');
```


## 高级功能

### 可写表

如果 OGR 驱动支持，你可以插入/更新/删除记录。可写表需要在表定义中包含 `fid` 列。

```sql
ALTER SERVER myserver
  OPTIONS (ADD updateable 'true');
```

### 列名映射

将远程列名映射到本地列名：

```sql
CREATE FOREIGN TABLE typetest_fdw_mapped (
  fid bigint,
  supertime time OPTIONS (column_name 'clock'),
  thebestname varchar OPTIONS (column_name 'name') )
  SERVER wraparound
  OPTIONS (layer 'typetest');
```

### 自动表导入

使用 `IMPORT FOREIGN SCHEMA` 自动创建外部表定义：

```sql
CREATE SCHEMA fgdball;

-- 导入所有表
IMPORT FOREIGN SCHEMA ogr_all
  FROM SERVER fgdbtest
  INTO fgdball;

-- 导入指定表
IMPORT FOREIGN SCHEMA ogr_all
  LIMIT TO(cities)
  FROM SERVER fgdbtest
  INTO fgdball;
```

### GDAL 选项

通过配置和打开选项控制驱动行为：

```sql
CREATE SERVER myserver_latin1
  FOREIGN DATA WRAPPER ogr_fdw
  OPTIONS (
    datasource '/tmp/test',
    format 'ESRI Shapefile',
    config_options 'SHAPE_ENCODING=LATIN1' );
```

多个配置选项可以作为空格分隔的列表传递。
