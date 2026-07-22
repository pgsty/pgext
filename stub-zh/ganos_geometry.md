## 用法

来源：

- [ApsaraDB RDS 的 GanosBase Geometry 模型](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/geometry-model)
- [ApsaraDB RDS for PostgreSQL 支持的扩展](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)
- [RDS 扩展创建限制](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/limits-on-the-creation-of-the-pg-cron-extension)

`ganos_geometry` 是阿里云专属的 GanosBase 空间引擎，用于 ApsaraDB RDS for PostgreSQL。它提供兼容 OpenGIS 的 `geometry` 和 `geography` 数据类型、空间函数、运算符与索引，可处理二维、三维和四维矢量数据；它不是可移植的社区软件包。

### 核心工作流

在明确选定的模式中创建扩展，然后用具体几何子类型和 SRID 存储数据：

```sql
CREATE EXTENSION ganos_geometry WITH SCHEMA public CASCADE;

CREATE TABLE roads (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL,
  geom geometry(LINESTRING, 3857) NOT NULL
);

INSERT INTO roads (name, geom)
VALUES (
  'sample',
  ST_GeomFromText('LINESTRING(0 0, 100 100)', 3857)
);

CREATE INDEX roads_geom_gix ON roads USING GIST (geom);
VACUUM ANALYZE roads (geom);

SELECT id, name, ST_AsText(geom), ST_Length(geom)
FROM roads;
```

`CASCADE` 会在引擎支持时安装 GanosBase 声明的依赖。官方模型还支持 GiST 和 BRIN 索引：多维 GiST 索引使用 `gist_geometry_ops_nd`，三维或四维 BRIN 索引分别使用 `brin_geometry_inclusion_ops_3d` 或 `brin_geometry_inclusion_ops_4d`。

### 主要 SQL 对象

- 构造与序列化函数包括 `ST_GeomFromText` 和 `ST_AsText`。
- 度量函数包括 `ST_Length` 和 `ST_Area`。
- 空间关系谓词包括 `ST_Contains`、`ST_Covers` 和 `ST_Intersects`。
- 官方 SQL 参考文档覆盖范围更大的 PostGIS 兼容函数；应根据已安装的服务版本核对具体签名。

### 服务商、版本与模式边界

扩展可用性取决于 RDS 引擎版本、PostgreSQL 大版本与小版本。当前标准版表中，PostgreSQL 17 上的 `ganos_geometry` 为 7.0，PostgreSQL 12–16 上为 6.9，PostgreSQL 10–11 上为 6.3，PostgreSQL 18 上不可用；其他版本会不同。应在目标实例上查询 `pg_available_extensions`，不能将目录中的 7.0 视为所有引擎的统一版本。

当前 RDS 指南要求不要把 GanosBase 与 PostGIS 扩展安装到同一模式，因为两者都管理 `spatial_ref_sys`；应使用不同模式，或先移除冲突扩展。某些旧引擎小版本还会因安全修复而禁止新的 `CREATE EXTENSION` 操作。既有扩展可能继续存在，但新建操作会被阻止；应升级实例后重试，不应绕过服务商限制。

二进制、权限、备份、升级和支持都由阿里云控制。将自定义类型用于长期模式前，应在准确的 RDS 构建上测试迁移、SRID、空间索引、查询计划、转储恢复和跨版本行为。
