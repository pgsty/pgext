## 用法

来源：

- [官方 README](https://github.com/zilder/tuple_fdw/blob/9c737a2043bfd78fe085f0966e3a58b62e38fd70/README.md)
- [FDW 实现](https://github.com/zilder/tuple_fdw/blob/9c737a2043bfd78fe085f0966e3a58b62e38fd70/tuple_fdw.c)
- [文件存储实现](https://github.com/zilder/tuple_fdw/blob/9c737a2043bfd78fe085f0966e3a58b62e38fd70/storage.c)
- [0.1 版扩展 SQL](https://github.com/zilder/tuple_fdw/blob/9c737a2043bfd78fe085f0966e3a58b62e38fd70/tuple_fdw--0.1.sql)

`tuple_fdw` 把 PostgreSQL heap tuple 存入追加式外部二进制文件，文件由带校验和、经 LZ4 压缩的数据块组成。它支持顺序读取与插入，同时绕过 PostgreSQL buffer cache 和 autovacuum。它是实验性冷存储，而不是事务表替代品。

### 核心流程

在 PostgreSQL 服务器安装 liblz4，创建扩展与服务器，再使用服务器端绝对文件路径定义外部表。如果文件不存在，validator 会以 PostgreSQL 操作系统账号创建它。

```sql
CREATE EXTENSION tuple_fdw;
CREATE SERVER tuple_store FOREIGN DATA WRAPPER tuple_fdw;

CREATE FOREIGN TABLE archive_event (
  event_id bigint,
  occurred_at timestamptz,
  payload text
)
SERVER tuple_store
OPTIONS (
  filename '/var/lib/postgresql/tuple_fdw/archive_event.bin',
  use_mmap 'false',
  lz4_acceleration '1'
);
```

按预期物理顺序批量插入，之后可以正常查询。FDW 实现了 INSERT，但没有实现 UPDATE 或 DELETE。

```sql
INSERT INTO archive_event
SELECT id, occurred_at, payload
FROM live_event
ORDER BY occurred_at, id;

SELECT *
FROM archive_event
WHERE occurred_at >= timestamptz '2026-01-01 00:00:00+00';
```

### 外部表选项

- `filename`：必需的文件路径；不存在时创建。
- `use_mmap`：以 memory-mapped 读取代替 `fread`；它可能改善高并发读取，但不会让写入具备并发安全。
- `sorted`：以空格分隔的列名，向规划器声明文件已排序。
- `lz4_acceleration`：LZ4 速度/压缩率权衡，默认 `1`；值越高越偏向速度而非压缩率。

`sorted` 选项是承诺，不是强制机制。如果物理文件实际未按这些列排序，规划器可能选择无效的顺序相关计划并返回错误结果。

### 持久性、并发与可移植性

外部文件不在 PostgreSQL WAL、MVCC、事务回滚、逻辑/物理复制和普通基础备份的管理范围内。尽管写入会 fsync 且 FDW 获取关系锁，事务中止也不会撤销文件变更；另一个数据库、外部表、进程或主机还可以绕过该锁访问相同路径。上游明确禁止并发写入，以及读写混合并发。

单行插入会反复解压和重压最后一个数据块，因此应使用批量加载。文件存储原始 heap tuple 字节，依赖外部表行布局与 PostgreSQL 二进制表示；未经验证迁移，不要用不同模式、架构或 PostgreSQL 构建挂载现有文件。必须在静默状态下通过外部流程备份与恢复文件，并在重新开放前验证校验和与行数。`mmap` 读取期间也不能在底层替换或截断文件。
