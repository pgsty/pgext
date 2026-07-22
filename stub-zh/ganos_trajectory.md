## 用法

来源：

- [阿里云轨迹介绍](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/trajectory-introduction)
- [阿里云 Ganos 基本概念](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/basic-concepts)
- [阿里云扩展创建指南](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/create-extensions)
- [阿里云 RDS PostgreSQL 扩展矩阵](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`ganos_trajectory` 是阿里云 Ganos 中存储、分析移动对象轨迹的组件。轨迹把空间点与按顺序排列的 `TIMESTAMP` 观测组合起来，并可携带点属性与事件；它是服务商托管组件，不是可移植的社区扩展。

### 启用与工作流

阿里云文档把 `ganos_spatialref` 和 `ganos_geometry` 列为前置依赖。`CASCADE` 是简洁的厂商支持安装方式；应在目标 RDS 实例上核对实际安装的依赖集合。

```sql
CREATE EXTENSION ganos_trajectory CASCADE;

SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name IN ('ganos_spatialref', 'ganos_geometry', 'ganos_trajectory');
```

其数据模型支持轨迹点、属性、事件以及 `BoxNDF` 等时空包围盒。时间辅助函数包括 `ST_UnixEpochToTS` 与 `ST_TsToUnixEpoch`。加载生产数据前，应按时间戳顺序提供观测，并根据厂商示例验证重复时间、缺失点、时区、插值与单位行为。

### 服务商边界

目录记录的是 Ganos 7.0，但厂商矩阵针对不同 PostgreSQL 引擎版本列出不同版本与可用性。RDS 版本类型及小版本还可能禁止新建扩展，而已安装副本仍可继续使用。不能假定目录版本一定可创建，应查询当前矩阵和 `pg_available_extensions`。

厂商基本概念文档指出，GanosBase 与 PostGIS 不能安装在同一模式中。启用前应规划模式，并在可丢弃实例测试依赖创建、权限、升级、备份恢复、空间索引、时间谓词、导出格式，以及迁移到不提供 Ganos 的 PostgreSQL 服务。
