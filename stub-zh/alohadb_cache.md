## 用法

来源：

- [官方 alohadb_cache.control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_cache/alohadb_cache.control)
- [官方 alohadb_cache--1.0.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_cache/alohadb_cache--1.0.sql)
- [官方 alohadb_cache.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_cache/alohadb_cache.c)
- [官方 cache_store.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_cache/cache_store.c)
- [官方 alohadb_cache.h](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_cache/alohadb_cache.h)

`alohadb_cache` 1.0 是 AlohaDB 源码树中的共享内存 JSONB 键值缓存，支持过期时间与 LRU 淘汰。本文描述该分支的实现，不宣称它兼容原版 PostgreSQL。

### 启用与使用

先预加载动态库并重启，再以超级用户创建扩展：

```conf
shared_preload_libraries = 'alohadb_cache'
alohadb.cache_max_entries = 1000
```

```sql
CREATE EXTENSION alohadb_cache;
SELECT public.cache_set('demo:item', '{"value":42}'::jsonb, interval '1 minute');
SELECT public.cache_get('demo:item');
SELECT * FROM public.cache_stats();
SELECT public.cache_delete('demo:item');
```

control 文件将对象固定在 `public`。`cache_set` 替换值及其 TTL；`cache_get` 对不存在或过期的条目返回 NULL；`cache_delete` 返回条目是否存在；`cache_flush` 清空整个缓存并返回数量。

### 容量与隔离

`alohadb.cache_max_entries` 仅能在启动时设置，默认 1000，范围 16–100000。键必须放入含终止符的 256 字节缓冲区，JSONB 文本必须放入 8192 字节值缓冲区。访问时清理过期条目，容量不足时淘汰最久未使用的条目。

此缓存内容易失，不写 WAL，也不会随事务回滚。键中没有数据库或角色标识，不同数据库的调用者共用集群级命名空间。SQL 脚本未撤销默认的 PUBLIC 函数执行权限。使用前应限制函数权限并划分键前缀，不能将缓存当作授权边界或持久数据源。
