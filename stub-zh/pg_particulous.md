## 用法

来源：

- [已复核 commit 的 pg_particulous README](https://github.com/funbringer/pg_particulous/blob/25918b502ffecd0425ea7eefafc8770b2391fbf6/README.md)
- [已复核 commit 的 pg_particulous.control](https://github.com/funbringer/pg_particulous/blob/25918b502ffecd0425ea7eefafc8770b2391fbf6/pg_particulous.control)
- [版本 1.0 的安装 SQL](https://github.com/funbringer/pg_particulous/blob/25918b502ffecd0425ea7eefafc8770b2391fbf6/pg_particulous--1.0.sql)
- [C 辅助函数实现](https://github.com/funbringer/pg_particulous/blob/25918b502ffecd0425ea7eefafc8770b2391fbf6/src/pg_particulous.c)
- [上游迁移测试](https://github.com/funbringer/pg_particulous/blob/25918b502ffecd0425ea7eefafc8770b2391fbf6/sql/vanilla_to_pathman.sql)

`pg_particulous` 是一个尚在开发中的迁移辅助扩展，用于把 PostgreSQL 10 原生分区元数据迁移到 `pg_pathman`。当前公开操作是 `migrate_to_pathman`；`build_vanilla_part_key` 和 `desugar_vanilla` 等辅助函数会推导分区表达式、改写原生分区元数据，再向 `pg_pathman` 注册关系。

### 迁移入口

```sql
CREATE EXTENSION pg_pathman;
CREATE EXTENSION pg_particulous;

SELECT migrate_to_pathman(
  'public.measurements'::regclass,
  'recorded_at'
);
```

应将 `pg_particulous` 与 `pg_pathman` 安装到同一模式。对于 PostgreSQL 10 原生分区关系，函数会推导分区键；其他关系则要明确提供表达式。

### 注意事项

- 尽管 README 使用了“vice versa”的描述，已复核版本 `1.0` 的安装 SQL 只实现迁往 `pg_pathman` 的方向。
- `desugar_vanilla` 会取得 `ACCESS EXCLUSIVE` 锁，并直接更新 `pg_class` 和 `pg_partitioned_table`。执行失败或服务器版本不匹配可能破坏分区元数据。
- 代码面向 PostgreSQL 10 时代的内部 API，且自 2017 年后未再更新。应把它视为迁移研究代码，而不是常规生产工具；必须在一次性副本上使用已验证备份进行演练。
