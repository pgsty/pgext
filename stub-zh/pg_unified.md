## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_unified/pg_unified.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_unified/pg_unified--1.1.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_unified/meson.build)

`pg_unified` `1.1` 提供通过统一视图查询缓冲事件与批量数据。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_unified;
CREATE TABLE unified_demo_batch(url text);
SELECT unified_register('demo', 'unified_demo_batch');
SELECT unified_ingest('demo', '{"url":"/home"}');
SELECT * FROM demo_unified;
SELECT unified_flush('demo');
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

注册的视图组合批量表行与尚未刷新的事件。unified_flush 转移队列行，unified_purge 清理符合条件的历史，unified_stream_events 暴露缓冲区。需明确安排刷新，并保持事件 JSON 与批量表兼容。应检查生成视图和底层表的权限；这是数据库本地缓冲区，不是分布式流消息系统。

此模块对应上游旧名 `alohadb_unified`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
