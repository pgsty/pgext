## 用法

来源：

- [PGXN 元数据](https://api.pgxn.org/dist/pg_integrity_checksums/2.0.2/META.json)
- [官方 README](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.0.2/README.md)
- [控制文件](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.0.2/pg_integrity_checksums.control)
- [迁移指南](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.0.2/MIGRATION.md)

`pg_integrity_checksums` 为 PostgreSQL 页面、单元、元组、表、索引和数据库计算物理及逻辑校验和。它是面向 PostgreSQL 12–19 上游版本的 C 扩展；Pigsty 软件包覆盖 PostgreSQL 14–18，会直接读取关系存储，因此安装和数据库级扫描需要超级用户。

### 安装与迁移

```sql
CREATE EXTENSION pg_integrity_checksums;
```

### 2.0.2 版本

2.0.2 通过适配数组迭代器和表扫描调用，增加对 PostgreSQL 19 API 的兼容。该版本不改变 SQL 对象，2.0.1 到 2.0.2 的 SQL 升级为空操作。Pigsty 构建覆盖 PostgreSQL 14–18。

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

关系级函数要求调用者拥有关系的 `SELECT` 权限，并拒绝受行级安全策略保护的关系，因为直接存储扫描无法应用 RLS。数据库级函数只允许超级用户执行。`pg_integrity_checksums.workers` 控制并行扫描，`pg_integrity_checksums.strict` 使数据库扫描在关系无法检查时失败，而不是跳过。扩展不要求共享预加载，但应依据发行说明在相关实现变化后重新计算物理基线。
