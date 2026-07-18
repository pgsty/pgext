## 用法

来源：

- [上游 README](https://github.com/postgrespro/bztree/blob/a4b238720b742a336d637702ed5985b17ea31af1/README.md)
- [扩展 control 文件](https://github.com/postgrespro/bztree/blob/a4b238720b742a336d637702ed5985b17ea31af1/bztree.control)
- [SQL 安装脚本](https://github.com/postgrespro/bztree/blob/a4b238720b742a336d637702ed5985b17ea31af1/bztree--1.0.sql)
- [PostgreSQL 访问方法实现](https://github.com/postgrespro/bztree/blob/a4b238720b742a336d637702ed5985b17ea31af1/bztree_fdw.cc)

`bztree` `1.0` 版是实验性的共享内存 BzTree 索引访问方法，为若干整数、文本、二进制、网络和标识符类型提供等值运算符类。

### 示例

先把 `bztree` 加入 `shared_preload_libraries` 并重启 PostgreSQL，再执行：

```sql
CREATE EXTENSION bztree;
CREATE INDEX events_key_bz ON events USING bztree (event_key);
EXPLAIN SELECT * FROM events WHERE event_key = 42;
```

该访问方法支持等值查找，但不声明排序、唯一性、反向扫描、包含列或并行扫描能力。复核源码中仍有未完成的内存回收/删除工作，也没有说明持久恢复语义或当前 PostgreSQL 兼容性。没有在可丢弃数据上完成崩溃/重启、VACUUM、并发和损坏测试前，不要用它替代生产 B-tree。
