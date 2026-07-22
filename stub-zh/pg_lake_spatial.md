## 用法

来源：

- [官方空间功能文档](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/spatial.md)
- [官方 control 文件](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_spatial/pg_lake_spatial.control)

`pg_lake_spatial` 把 `pg_lake` 分析表与 PostGIS geometry、DuckDB spatial/GDAL 读取器连接起来。它可以从空间文件推断模式，直接查询远程 GeoParquet 或 GeoJSON，也能把支持的文件加载到普通分析表。它同时依赖 `pg_lake` 与 `postgis`。

### 核心流程

创建依赖链，然后定义一个从远程文件推断字段的外部表：

```sql
CREATE EXTENSION pg_lake_spatial CASCADE;

CREATE FOREIGN TABLE world ()
SERVER pg_lake
OPTIONS (path 's3://my-bucket/world.parquet');

SELECT ST_GeometryType(geom), count(*)
FROM world
GROUP BY 1;
```

如需持久化副本，可按 `pg_lake` 文档用 `load_from = '...'` 创建普通分析表，或使用 `COPY`。对于 GDAL 支持的归档格式，应设置文档规定的 `format 'gdal'`，必要时再加 `zip_path` 选项。

### 格式与执行

扩展覆盖 GeoParquet、GeoJSON/GeoJSONSeq、WKB/WKT，以及 GDAL 可读取的格式。DuckDB 能执行的空间表达式会下推到分析引擎；不支持或语义不兼容的操作会回退到 PostgreSQL 中的 PostGIS。

### 注意事项

GeoParquet 当前只支持 WKB geometry。GDAL 资源可能被下载并缓存，因此远程访问需要可用的网络/对象存储凭据与可预测的缓存空间。不支持通过 GDAL 写入。分析表可能报告 SRID `0`；QGIS 等客户端可能需要显式设置 SRID。

空间下推并不完整。回退到 PostGIS 虽然结果正确，却可能增加数据传输与执行开销，不同表类型之间的连接也可能效率较低。应为每个重要函数核对执行计划与结果，保持两个依赖版本兼容，并使用有代表性的文件大小与坐标系做基准测试。
