## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/postgrespro/lsm/blob/bbd87030ae68fd855d91b43ca9367968bdc2fa69/README.md)
- [扩展 control 文件](https://github.com/postgrespro/lsm/blob/bbd87030ae68fd855d91b43ca9367968bdc2fa69/lsm.control)
- [FDW 实现](https://github.com/postgrespro/lsm/blob/bbd87030ae68fd855d91b43ca9367968bdc2fa69/lsm_fdw.cpp)

`lsm` 将 RocksDB 嵌入 PostgreSQL 后端，并通过 `lsm_fdw` 外部数据包装器暴露。构建兼容性同时取决于 PostgreSQL 与 RocksDB；已复核上游仅声称测试过 Ubuntu 18.04、PostgreSQL 11 和 RocksDB 6.2.4。使用前把 `lsm` 加入 `shared_preload_libraries` 并重启。

```sql
CREATE EXTENSION lsm;
CREATE SERVER lsm_server FOREIGN DATA WRAPPER lsm_fdw;
CREATE FOREIGN TABLE student (
  id integer,
  name text
) SERVER lsm_server;
```

外部表第一个属性充当主键，且不支持复合主键。PostgreSQL 数据类型不会以原生形式存储，上游还说明 ACID 行为依赖存储引擎。

从 PostgreSQL 视角看，RocksDB 文件属于外部状态。在存放重要数据前，必须验证备份、恢复、崩溃恢复、复制、事务回滚、并发访问和清理行为。上游卸载示例写错了扩展名，应以已安装目录名为准，并在恢复要求确定前保留 RocksDB 目录。
