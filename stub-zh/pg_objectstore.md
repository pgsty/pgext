## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_objectstore/pg_objectstore.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_objectstore/pg_objectstore--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_objectstore/meson.build)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_objectstore/objectstore_http.c)

`pg_objectstore` `1.0` 提供从 SQL 访问云对象存储。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_objectstore;
SELECT * FROM objectstore_credential_list();
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

objectstore_credential_register 保存凭据，objectstore_get、objectstore_put、objectstore_list 和 objectstore_delete 访问配置的服务。SQL 回滚不会撤销外部写入或删除。应只向可信操作人员授予函数执行和凭据表访问权限，核验目标并控制网络出口。使用真实凭据前需按所附传输实现检查服务支持与认证方式。

此模块对应上游旧名 `alohadb_objectstore`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
