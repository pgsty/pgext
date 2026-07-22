## 用法

来源：

- [上游 README](https://github.com/postgrespro/bztree/blob/a4b238720b742a336d637702ed5985b17ea31af1/README.md)
- [扩展 control 文件](https://github.com/postgrespro/bztree/blob/a4b238720b742a336d637702ed5985b17ea31af1/bztree.control)
- [SQL 安装脚本](https://github.com/postgrespro/bztree/blob/a4b238720b742a336d637702ed5985b17ea31af1/bztree--1.0.sql)
- [PostgreSQL 访问方法实现](https://github.com/postgrespro/bztree/blob/a4b238720b742a336d637702ed5985b17ea31af1/bztree_fdw.cc)

`bztree` `1.0` 版是实验性的共享内存 BzTree 索引访问方法，为若干整数、文本、二进制、网络和标识符类型提供等值运算符类。

### 一次性演示

先把 `bztree` 加入 `shared_preload_libraries` 并重启 PostgreSQL，然后趁演示表仍为空时创建索引。索引存在后插入的记录会经过已实现的插入回调：

```sql
CREATE EXTENSION bztree;
CREATE TABLE bztree_demo (event_key integer NOT NULL, payload text);
CREATE INDEX bztree_demo_key_bz ON bztree_demo USING bztree (event_key);
INSERT INTO bztree_demo VALUES (42, 'after index creation');
EXPLAIN SELECT * FROM bztree_demo WHERE event_key = 42;
```

不要在已有记录的表上构建该索引。在复核的源码中，`bztree_build()` 会扫描堆表，但其构建回调只序列化每个键，从未把键插入 BzTree；它却仍把所有扫描记录报告为索引记录。因此，规划器选用该索引时可能遗漏已有记录并返回错误结果。它也拒绝 NULL 键。

该访问方法支持等值查找，但不声明排序、唯一性、反向扫描、包含列或并行扫描能力。源码中仍有未完成的内存回收/删除工作，也没有说明持久恢复语义或当前 PostgreSQL 兼容性。只能把它用于可丢弃数据上的源码级研究；不要依赖它产生正确查询结果，也不要用它替代生产 B-tree。
