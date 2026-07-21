## 用法

来源：

- [ogr_fdw v1.1.9 README](https://github.com/pramsey/pgsql-ogr-fdw/blob/v1.1.9/README.md)
- [扩展控制文件](https://github.com/pramsey/pgsql-ogr-fdw/blob/v1.1.9/ogr_fdw.control)
- [v1.1.8 到 v1.1.9 的变更对比](https://github.com/pramsey/pgsql-ogr-fdw/compare/v1.1.8...v1.1.9)
- [GDAL 矢量驱动](https://gdal.org/en/stable/drivers/vector/index.html)

`ogr_fdw` 将 GDAL/OGR 支持的矢量数据公开为 PostgreSQL 外部表。它可以通过 OGR 驱动读取文件和远程数据源；当所选驱动及数据源支持更新时，也可以写入。应先安装 PostGIS，再安装 `ogr_fdw`，以获得原生 geometry 列；否则 geometry 会以 WKB `bytea` 形式公开。

### 发现并导入图层

使用随扩展安装的辅助工具检查数据源并生成匹配的 SQL：

```console
ogr_fdw_info -s /srv/gis/cities.gpkg
ogr_fdw_info -s /srv/gis/cities.gpkg -l cities
```

等价的最小定义如下：

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION ogr_fdw;

CREATE SERVER city_source
  FOREIGN DATA WRAPPER ogr_fdw
  OPTIONS (
    datasource '/srv/gis/cities.gpkg',
    format 'GPKG'
  );

CREATE FOREIGN TABLE city (
  fid bigint,
  geom geometry,
  name text
) SERVER city_source
  OPTIONS (layer 'cities');

SELECT fid, name FROM city WHERE geom && ST_MakeEnvelope(-10, 35, 30, 60, 4326);
```

对于文件型数据源，PostgreSQL 服务器账号需要相应的文件系统权限；对于远程驱动，还需要网络和凭据访问权限。

### 导入与映射

```sql
CREATE SCHEMA gis_import;

IMPORT FOREIGN SCHEMA ogr_all
  LIMIT TO (cities)
  FROM SERVER city_source
  INTO gis_import;
```

`ogr_all` 表示全部 OGR 图层。导入通常会规范化表名和列名；需要保留远端原名时，使用 `launder_table_names` 与 `launder_column_names` 选项。外部列可以通过 `OPTIONS (column_name 'RemoteName')` 映射到不同的源字段名。

### 重要选项与对象

- 服务器选项：必需的 `datasource`，以及可选的 `format`、`updateable`、`config_options`、`open_options` 和 `character_encoding`。
- 表选项：`layer` 指定 OGR 图层，`updateable` 可用于禁用写入。
- `fid` 标识要素；可写外部表必须提供该字段。
- `ogr_fdw_info` 列出驱动与图层，并输出服务器/外部表定义。
- `ogr_fdw_version()` 返回扩展与 GDAL 版本。
- `ogr_fdw_drivers()` 列出编译进来的 OGR 驱动。

### 性能与写入边界

简单比较和边界框 `&&` 谓词可以下推，但复杂过滤条件可能在本地执行。该 FDW 会取回所有选中的源字段，并为每条查询打开两个 OGR 连接，而不会进行连接池化。请使用 `EXPLAIN`，只投影所需字段，并针对实际驱动和数据源做基准测试。

写入能力取决于驱动，并要求数据源级写权限以及 `fid`。必须保持只读的数据源应设置 `updateable = false`。1.1.9 相比 1.1.8 简化了版本字符串，没有记录 SQL 工作流变化；控制文件中的 SQL 扩展版本仍为 1.1。
