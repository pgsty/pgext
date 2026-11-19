## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_model_registry/pg_model_registry.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_model_registry/pg_model_registry--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_model_registry/meson.build)

`pg_model_registry` `1.0` 提供模型版本、部署元数据与实验记录。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_model_registry;
SELECT model_register('demo', 'onnx', 'Example model');
SELECT * FROM model_list();
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

model_version_add 保存制品和指标，model_deploy、model_rollback 与 model_get_active 管理活跃版本元数据，experiment_start/end/list 记录实验。注册表中的部署操作改变元数据，推理消费者仍需使用选定制品。应为制品及部署表设置明确授权，并在备份中包含二进制内容和历史。

此模块对应上游旧名 `alohadb_model_registry`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
