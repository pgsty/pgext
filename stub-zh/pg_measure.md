## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_measure/pg_measure.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_measure/pg_measure--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_measure/meson.build)

`pg_measure` `1.0` 提供附带单位、不确定度与溯源信息的测量值。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_measure;
SELECT m_convert(measurement(1, 0.01, 'm'), 'cm');
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

measurement 类型提供单位转换、不确定度传播、算术及聚合。m_wmean 计算加权平均，prov_track 与 prov_lineage 记录和检查行级溯源。不确定度采用一阶传播，并不构成完整的统计依赖模型。共享结果时应检查溯源表策略与授权，并保留追溯所需的历史。

此模块对应上游旧名 `alohadb_measure`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
