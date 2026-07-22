## 用法

来源：

- [阿里云扩展支持矩阵](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)
- [阿里云点云模型文档](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/point-cloud-model)

`ganos_pointcloud` 是阿里云 GanosBase 扩展，用于在 ApsaraDB RDS for PostgreSQL 中存储、压缩、检查和处理点云数据。它是云厂商专用组件，不是可移植的社区软件包；是否可用以及准确版本取决于 RDS PostgreSQL 引擎世代和版本。

### 启用扩展

在受支持的 RDS 实例上，于预期模式中创建核心扩展；需要空间转换与分析时再加入 geometry 桥接扩展：

```sql
CREATE EXTENSION ganos_pointcloud WITH SCHEMA public CASCADE;
CREATE EXTENSION ganos_pointcloud_geometry WITH SCHEMA public CASCADE;
```

`CASCADE` 可能安装服务上可用的前置依赖。创建后应确认实际扩展清单与版本，因为支持矩阵会为不同引擎版本提供不同的 Ganos 发布。

### 数据模型

点格式登记在 `pointcloud_formats` 中。每种格式通过 XML 定义维度、数据类型、缩放和偏移规则、压缩方式以及 SRID。单个点使用 `pcpoint`，点集合使用 `pcpatch`；所有值都引用格式标识符，以便正确解释二进制载荷。`ganos_pointcloud_geometry` 将这些值接入 Ganos geometry 运算。

加载数据集前应登记并验证格式，并保持维度顺序、缩放、偏移与 SRID 稳定。为值使用错误的格式标识符，可能在没有明显报错的情况下错误解释坐标或属性。

### 运维说明

部署或升级前，应在阿里云矩阵中核对目标引擎小版本。点云数据集可能很大，因此要用代表性数据测试压缩选择、写入吞吐、索引、空间转换、备份大小和查询内存。扩展创建与升级遵循阿里云的权限和维护模型；不要假定自托管构建或社区 `pointcloud` 项目的 API 可与已安装的 Ganos 发布互换。
