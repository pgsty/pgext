## 用法

来源：

- [官方 README](https://github.com/sangli00/hash_fdw/blob/9c685317f6a08381b426706c235ac515505cac51/README)
- [官方扩展控制文件](https://github.com/sangli00/hash_fdw/blob/9c685317f6a08381b426706c235ac515505cac51/hash_fdw.control)
- [官方扩展 SQL](https://github.com/sangli00/hash_fdw/blob/9c685317f6a08381b426706c235ac515505cac51/hash_fdw--1.0.sql)
- [官方共享内存实现](https://github.com/sangli00/hash_fdw/blob/9c685317f6a08381b426706c235ac515505cac51/hash_if.c)
- [官方 FDW 实现](https://github.com/sangli00/hash_fdw/blob/9c685317f6a08381b426706c235ac515505cac51/hash_fdw.c)

`hash_fdw` 1.0 是 2016 年的内存 FDW 实验，把外部表行存入 PostgreSQL 共享内存哈希表。它可以在没有外部服务器的情况下扫描和插入键值行，但不是持久存储，也不提供普通 PostgreSQL 表语义。

### 预加载与基本流程

库会在 `_PG_init` 中分配共享内存，因此必须把它加入 `shared_preload_libraries`、设置存储大小并重启，然后再创建扩展。

```conf
shared_preload_libraries = 'hash_fdw'
hash.shmem_size = 1024
```

```sql
CREATE EXTENSION hash_fdw;

CREATE SERVER hash_server
FOREIGN DATA WRAPPER hash_fdw;

CREATE FOREIGN TABLE session_cache (
    id bigint,
    name text,
    age integer
)
SERVER hash_server
OPTIONS (key 'id', hash_idx '1');

INSERT INTO session_cache VALUES (1, 'Ada', 37);
SELECT * FROM session_cache;
```

`key` 指定要转换为哈希键的外部表列。`hash_idx` 选择一个共享哈希表；该源码中的公共辅助函数把可用索引硬限制为 1–8。

### 已安装接口

- `hash_fdw_handler()` 与 `hash_fdw_validator(text[], oid)` 实现 FDW。
- 外部扫描遍历所选哈希表的全部内容；插入根据配置键添加或替换条目。
- `hash_insert_hashdata(int, cstring, record)` 与 `hash_get_hashdata(int, cstring)` 暴露底层 C 辅助函数，但其 `record` 契约不便使用，且依赖内部元组表示。
- GUC 包括 `hash.shmem_size` 和 `hash_fdw.hash_idx`；共享内存布局在服务器启动时固定，因此不要动态修改大小/数量假设。

### 数据与安全边界

行不会写入 heap、WAL 或外部服务。重启、崩溃、故障转移或二进制不兼容都可能丢失或破坏全部条目。代码没有实现与普通表等价的 MVCC 可见性、事务回滚、约束、索引、备份、复制或并发保证。

共享哈希表是集群级的，外部表定义则属于单个数据库。跨数据库或不同行布局复用同一个 `hash_idx` 与键，可能发生碰撞，或用错误 descriptor 解释元组字节。删除回调不完整，也没有注册更新路径。只能在专用测试实例中使用唯一隔离的索引和可丢弃数据。

README 声明它针对 PostgreSQL 9.5.3 构建。服务器共享内存与 FDW API 已有大量变化，不要在现代服务器上加载未经审查的二进制。
