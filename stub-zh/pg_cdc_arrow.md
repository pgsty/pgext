## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_cdc_arrow/pg_cdc_arrow.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_cdc_arrow/pg_cdc_arrow--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_cdc_arrow/meson.build)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_cdc_arrow/pg_cdc_arrow.c)

`pg_cdc_arrow` `1.0` 提供Arrow IPC 逻辑解码输出插件。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
SELECT * FROM pg_create_logical_replication_slot('arrow_demo', 'pg_cdc_arrow');
SELECT * FROM pg_logical_slot_get_binary_changes('arrow_demo', NULL, NULL);
```

创建和消费复制槽需要相应复制权限。

### 运行边界

安装 SQL 不定义 SQL 对象，解码时直接加载输出库。配置逻辑 WAL 和可用复制槽，并使用获授权的复制角色。二进制变更采用 Arrow IPC，而非 JSON 或文本。应监控保留的 WAL，消费者停止后移除不用的槽，并检查更新、删除负载所需的副本标识。

此模块对应上游旧名 `alohadb_cdc_arrow`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
