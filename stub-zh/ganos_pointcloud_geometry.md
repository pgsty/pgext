## 用法

来源：

- [阿里云点云模型指南](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/point-cloud-model)
- [阿里云 RDS PostgreSQL 扩展矩阵](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`ganos_pointcloud_geometry` 是阿里云 Ganos 在点云值与空间几何之间的集成。厂商文档要求配合 `ganos_pointcloud`，其 `pcpoint` 与 `pcpatch` 类型由 `pointcloud_formats` 中的行描述。

```sql
CREATE EXTENSION ganos_pointcloud WITH SCHEMA public CASCADE;
CREATE EXTENSION ganos_pointcloud_geometry WITH SCHEMA public CASCADE;
SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name IN ('ganos_pointcloud', 'ganos_pointcloud_geometry');
```

存储数据前应按厂商指南定义点云模式。模式 ID、SRID、维度类型、缩放/偏移和压缩设置决定二进制点的解释；未经迁移直接修改会破坏语义。大型数据集应使用 `pcpatch` 和压缩，但必须先基准精度与查询行为。

这是厂商托管的 Ganos 7.0 软件，不是可移植的公开扩展。可用性和允许操作可能随 RDS 引擎版本、版本类型、地域与策略变化。应在可丢弃实例验证 `ST_MakePoint`、patch 统计、几何转换、空间索引、备份恢复和导出可移植性，并为目标环境不存在该组件的迁移场景做好设计。
