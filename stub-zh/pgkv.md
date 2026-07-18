## 用法

来源：

- [当前 database.dev 软件包文档](https://database.dev/jacobfvanzyl/pgkv)

`pgkv` 是由 PostgreSQL 表和 `JSONB` 支持、仿照 Redis 的纯 SQL 键值层。它支持字符串与计数器、哈希、列表、集合、有序集合、多键操作、glob 风格键查找和可选存活时间，同时保留 PostgreSQL 事务能力。

```sql
CREATE EXTENSION pgkv;
SELECT pgkv.set('greeting', 'hello');
SELECT pgkv.get('greeting');
SELECT pgkv.set('session:42', 'active', 60);
SELECT pgkv.ttl('session:42');
```

值保存在扩展共享存储中，而不是外部 Redis 服务器。过期项在读取时惰性清理；若需及时回收过期行，应调度 `pgkv.cleanup_expired()`。`pgkv.keys()` 可能扫描整个存储，大型集合与有序集合操作也在 PostgreSQL 内执行，因此应评估争用、WAL 量与表增长。

软件包还提供 `pgkv.flushall()` 等破坏性管理操作。应按角色限制函数执行权，尤其是删除、清理、模式扫描和全局清空。0.1.0 是较新的供应商分发 SQL；将其作为应用持久化契约前，应审查已安装脚本并测试升级行为。
