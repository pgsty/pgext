## 用法

来源：

- [当前版本资料](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.1/README.md)
- [当前版本资料](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.1/pg_integrity_checksums.control)
- [当前版本资料](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.1/Changes)
- [当前版本资料](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.1/pg_integrity_checksums--2.9.0--2.9.1.sql)

- [2.9.0 Changes](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.0/Changes)
- [2.8.0 到 2.9.0 SQL](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.0/pg_integrity_checksums--2.8.0--2.9.0.sql)
- [2.9.0 控制文件](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.0/pg_integrity_checksums.control)
- [2.9.0 函数参考](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.0/docs/functions.md)
- [当前上游](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/Changes)
- [当前上游](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/pg_integrity_checksums--2.6.0--2.7.0.sql)

- [PGXN 元数据](https://api.pgxn.org/dist/pg_integrity_checksums/2.7.0/META.json)
- [官方 README](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/README.md)
- [控制文件](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/pg_integrity_checksums.control)
- [函数参考](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/docs/functions.md)
- [运维指南](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/docs/operations.md)
- [2.5 到 2.6 升级 SQL](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/pg_integrity_checksums--2.5.0--2.6.0.sql)
- [2.4 升级 SQL](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/pg_integrity_checksums--2.3.0--2.4.0.sql)
- [2.5 升级 SQL](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/pg_integrity_checksums--2.4.0--2.5.0.sql)
- [迁移指南](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/MIGRATION.md)
- [Upgrade and baseline guide](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/docs/installation.md)

`pg_integrity_checksums` 为 PostgreSQL 页面、单元、元组、表、索引和数据库计算物理及逻辑校验和。此 C 扩展的上游支持 PostgreSQL 12–19，目录中的 Pigsty 软件包覆盖 PostgreSQL 14–18。安装和数据库级扫描需要超级用户。

### 安装与迁移

```sql
CREATE EXTENSION pg_integrity_checksums;
```

### 2.3.0 版本与页面校验

2.3.0 新增 `pg_page_checksum_status(regclass, integer)`，将页面计算所得的 16 位校验和与已存储的 `pd_checksum` 比较。升级后仍可执行以下检查：

```sql
SHOW data_checksums;
SELECT pg_page_checksum_status('public.orders'::regclass, 0);
```

关系必须存在且包含指定块。匹配时返回 `ok`，不匹配返回 `mismatch`，未初始化页面返回 `new`，未启用存储页面校验和时返回 `checksums_disabled`。比较存储值要求 `data_checksums = on`；无效关系或块号会报错。

`pg_page_checksum()` 仅计算 PostgreSQL 的 16 位页面校验和，不与存储值比较。返回零可能表示未初始化页面，也可能是计算结果恰好为零，因此不能仅凭零值判断页面有效。状态函数直接读取关系存储，结果反映此次读取的页面，不能保证后续写入后的状态。

2.2.0 引入的 FNV-1a 32/64 位重载仍返回原始 `bytea` 摘要，并保留旧的 32 位 SQL API。这些算法与 PostgreSQL 的 16 位页面校验和格式不同。更换算法或重载后，需要使用匹配的基线；2.3.0 新增页面状态函数本身不要求重新计算所有已有基线。

### 从旧扩展名迁移

2.0.0 是旧 `pg_checksums` 扩展的重命名延续。PostgreSQL 不支持 `ALTER EXTENSION ... RENAME`；应先检查依赖对象，删除旧扩展，再创建新扩展，并在使用前修改 GUC 前缀：

```sql
DROP EXTENSION pg_checksums;
CREATE EXTENSION pg_integrity_checksums;
SET pg_integrity_checksums.workers = 4;
SET pg_integrity_checksums.strict = on;
```

旧名和新名不能共存，因为它们定义相同的 SQL 函数。PostgreSQL 核心自带的 `pg_checksums` 工具是另一套独立程序。

### 校验和层级

内容比较使用逻辑校验和，存储完整性检查使用物理校验和。逻辑表校验和不要求主键，保留重复行，并排除关系和元组标识符。逻辑索引校验和由可见表行和索引定义计算，因此在 `VACUUM`、`CLUSTER` 和 `REINDEX` 后保持稳定。

```sql
SELECT pg_table_logical_checksum('public.orders'::regclass);
SELECT pg_index_logical_checksum('public.orders_pkey'::regclass);
SELECT pg_database_logical_checksum(false, true);
```

物理校验和包含存储细节，在 `VACUUM FULL`、`CLUSTER`、`REINDEX` 或其他物理重写后可能变化。需要跟踪这些变化时，使用页面、元组和表级物理函数。

### 权限与维护

关系级函数要求表级 `SELECT` 权限，并拒绝受行级安全策略保护的关系，因为直接存储扫描无法应用 RLS；仅有列级授权并不足够。数据库级函数只允许超级用户执行。`pg_integrity_checksums.workers` 控制并行扫描，`pg_integrity_checksums.strict` 使数据库扫描在关系无法检查时失败，而不是跳过。扩展不要求共享预加载。应保留使用相同算法、兼容结构计算的可信基线，并在比较物理校验和时考虑物理重写的影响。这些校验和不提供密码学认证。

### 2.7.0 升级与基线变化

安装匹配文件，重新连接并执行 SQL 升级：

```sql
ALTER EXTENSION pg_integrity_checksums UPDATE TO '2.7.0';
SHOW pg_integrity_checksums.strict;
```

数据库级逻辑校验和现在纳入带 schema 的关系名，因此即使行内容不变，所有数据库逻辑基线也会改变。修正索引元数据屏蔽后，部分物理索引基线同样改变。应重新计算并审查这些基线；物理校验和仍会受到实际存储变化影响，包括行锁和已回滚的写入。仅凭校验和不同不能认定数据损坏。

元组和单元格校验现在检查 MVCC 可见性，对不可见行返回 NULL。索引表达式在受限上下文中以表所有者身份执行，可串行化扫描会取得必要的冲突锁。页状态检查会对磁盘不匹配进行重试；数据库扫描期间被删除的关系即使在严格模式下也会以 NOTICE 跳过。其他不可读关系仍遵循默认开启的 `pg_integrity_checksums.strict`。这些语义与先前 2.6.0 的屏蔽及基线规则不同。

### 2.9.0 算法、序列与基线

```sql
ALTER EXTENSION pg_integrity_checksums UPDATE TO '2.9.0';
```

2.9.0 为通用算法函数新增 `crc32c` 与 `xxh3`，并让 `pg_table_bucket_checksums()` 支持并行 worker。新的序列逻辑校验和会比较序列定义以及 `last_value`、`is_called`；新的关系页面状态函数可一次校验所有页面和 fork。

当 `include_system = false` 时，数据库逻辑与物理校验和现在会像排除 `pg_catalog` 一样排除 `information_schema`。由于 `information_schema` 内容随 PostgreSQL 版本变化，旧版本产生的所有非系统数据库校验和都会因该修正而改变：需要重新计算并替换这些存储基线。表、索引和序列校验值不受此修正影响，CLI 基线此前已排除 `information_schema`。


### 2.9.1 扫描变化

在 PostgreSQL 17 及以上，表、索引、bucket 和数据库扫描使用 read stream；PostgreSQL 18 可使用异步 I/O，并行 worker 处理的范围也有所扩大。2.9.1 不改变校验和值。安装匹配库并执行 `ALTER EXTENSION pg_integrity_checksums UPDATE TO '2.9.1'`；此前 2.9.0 排除 information_schema 的变更仍要求替换受影响的非系统数据库基线。
