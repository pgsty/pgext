## 用法

来源：

- [阿里云 Ganos Raster 模型](https://help.aliyun.com/en/rds/apsaradb-rds-for-postgresql/raster-model)
- [阿里云栅格数据加载指南](https://help.aliyun.com/en/rds/apsaradb-rds-for-postgresql/load-raster-data)
- [ApsaraDB RDS for PostgreSQL 扩展矩阵](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`ganos_raster` 是阿里云为 ApsaraDB RDS for PostgreSQL 托管的 GanosBase 栅格扩展。它用于存储和分析网格影像等栅格数据；这不是具有公开 SQL 源码树的可移植社区扩展。

### 核心流程

```sql
CREATE EXTENSION Ganos_Raster WITH SCHEMA public CASCADE;

CREATE TABLE raster_table (
    id integer,
    raster_obj raster
);

INSERT INTO raster_table
VALUES (1, ST_ImportFrom('chunk_table', '/home/beijing.tif'));

SELECT ST_Height(raster_obj), ST_Width(raster_obj)
FROM raster_table
WHERE id = 1;
```

`ST_ImportFrom` 也接受阿里云 OSS URL。RDS 实例必须具有对象访问授权，并且 OSS 存储桶必须位于同一区域。

### 对象索引

- `raster` 是持久化栅格类型。
- `ST_ImportFrom(chunk_table, path)` 从本地路径或 OSS 导入栅格文件。
- `ST_Height` 和 `ST_Width` 报告栅格尺寸。
- `ST_BuildPyramid`、`ST_BestPyramidLevel`、`ST_Clip` 和 `ST_ClipDimension` 支持金字塔构建与空间提取。

完整的供应商版本专属接口应以阿里云 Raster SQL 参考为准。

### 运维说明

可用性和版本取决于 ApsaraDB PostgreSQL 引擎的大版本与小版本。目录版本 `7.0` 对应一个供应商发行版，其他受支持的引擎大版本可能提供不同 Ganos Raster 版本；应在目标实例上检查 `pg_available_extension_versions` 和当前阿里云矩阵。

`CASCADE` 会安装供应商依赖。文档中的 URI 形式可以包含 AccessKey ID 与密钥，因此不要把 OSS 凭据写入可复用 SQL、日志或查询历史。导入路径由托管服务端解释，而不是客户端机器。生产使用前，应针对选定的 RDS 版本测试存储增长、金字塔维护、备份行为和访问控制。
