## 用法

来源：

- [阿里云 Ganos 矢量金字塔文档](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/vector-pyramid)
- [ApsaraDB RDS for PostgreSQL 支持的扩展](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`ganos_geometry_pyramid` 是阿里云 GanosBase 扩展，用于把几何表构建成多级矢量瓦片金字塔，并返回 Mapbox Vector Tile（MVT/PBF）内容。它是托管服务能力，并非可安装到任意 PostgreSQL 服务器的社区扩展。

### 核心流程

使用已启用 GanosBase 的 ApsaraDB RDS for PostgreSQL 实例，准备几何表，并在构建金字塔前创建空间索引：

```sql
CREATE EXTENSION ganos_geometry_pyramid WITH SCHEMA public CASCADE;

CREATE TABLE test (
    id bigint PRIMARY KEY,
    name text,
    geom geometry(Point, 4326)
);
CREATE INDEX test_geom_gist ON test USING gist (geom);

SELECT ST_BuildPyramid('test', 'geom', 'id', '');
SELECT ST_Tile('test', '0_0_0');
SELECT ST_DeletePyramid('test');
```

`ST_BuildPyramid(table_name, geom_col, id_col, options)` 创建金字塔。`ST_Tile(table_name, tile_id)` 获取瓦片，标识格式为 Web Mercator（`EPSG:3857`）中的 `z_x_y`。`ST_DeletePyramid(table_name)` 删除生成的金字塔数据。

### 构建选项

默认值不满足需求时，可通过 `options` 传入 JSON 文本。重要键包括 `name`、`parallel`、`tileSize`、`tileExtend`、`maxLevel` 和 `buildRules`。`tileSize` 必须是 256 的倍数，文档范围为 0–4096；`tileExtend` 范围为 0–256；`buildRules` 可为不同层级设置过滤条件、属性字段和几何合并。超过 `maxLevel` 的层级将实时生成，而不是预构建。

并行构建使用预备事务。阿里云文档要求先把 `max_prepared_transactions` 设置为至少 100 并重启托管数据库；服务允许的并行度最高为 CPU 核数的四倍。这应通过服务配置协调，而不能只修改 SQL。

### 服务与版本边界

可用性取决于 ApsaraDB 版本、PostgreSQL 大版本和引擎小版本。引用时的服务商矩阵为 PostgreSQL 17 列出 `ganos_geometry_pyramid` 7.0，并为若干较旧大版本列出 6.9，因此不能假定目录版本在所有实例上都可用。应通过实例扩展管理页面或 `pg_available_extensions` 确认。构建和刷新金字塔会消耗存储与计算资源，源表变化也可能要求维护已生成的数据。
