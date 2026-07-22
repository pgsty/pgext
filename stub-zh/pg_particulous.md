## 用法

来源：

- [已复核 commit 的 pg_particulous README](https://github.com/funbringer/pg_particulous/blob/25918b502ffecd0425ea7eefafc8770b2391fbf6/README.md)
- [版本 1.0 的安装 SQL](https://github.com/funbringer/pg_particulous/blob/25918b502ffecd0425ea7eefafc8770b2391fbf6/pg_particulous--1.0.sql)
- [C 辅助实现](https://github.com/funbringer/pg_particulous/blob/25918b502ffecd0425ea7eefafc8770b2391fbf6/src/pg_particulous.c)
- [上游迁移测试](https://github.com/funbringer/pg_particulous/blob/25918b502ffecd0425ea7eefafc8770b2391fbf6/sql/vanilla_to_pathman.sql)

`pg_particulous` 是 PostgreSQL 10 时代的实验性迁移代码，用于把原生声明式分区元数据转换为 `pg_pathman` 元数据。尽管 README 声称支持双向转换，版本 `1.0` 只实现原生分区到 `pg_pathman` 的方向。它会直接重写系统目录与存储元数据，因此应视为历史迁移研究代码，而不是常规扩展。

### 迁移入口

扩展依赖 `plpgsql` 和 `pg_pathman`。PostgreSQL 10 原生分区表可以隐式取得分区表达式；非原生关系则必须提供明确表达式。

```sql
CREATE EXTENSION pg_pathman;
CREATE EXTENSION pg_particulous;

SELECT migrate_to_pathman(
  'public.measurements'::regclass,
  'recorded_at'
);
```

`migrate_to_pathman(regclass, text, boolean)` 返回布尔值。对于原生分区关系，它会调用 `build_vanilla_part_key`，递归运行 `desugar_vanilla`，向原分区添加检查约束，最后调用 `add_to_pathman_config`。在已复核 SQL 中，`run_tests` 参数虽然被接收，但没有使用。

其他辅助函数检查 PostgreSQL 10 分区键与边界、构造检查条件并识别原生分区。这些函数公开的是实现机制，不是受支持的通用分区管理 API。

### 破坏性边界

`desugar_vanilla` 以排他模式锁定 `pg_class` 和 `pg_inherits`，并以 `ACCESS EXCLUSIVE` 模式锁定目标关系。它会在目标表与临时表之间交换 `relkind`、`relfilenode`、事务年龄字段和 `pg_partitioned_table` 引用，清除原生分区标记，然后删除该临时表。这些操作依赖 PostgreSQL 10 的精确内部目录布局。

### 运维注意事项

- 不能仅凭扩展安装成功就在现代 PostgreSQL 或生产关系上运行。仓库自 2017 年后未更新，目标是非稳定的内部 API。
- 迁移没有实现反向操作。函数成功也不能证明约束、索引、触发器、外键、所有权、授权、序列、复制和查询计划仍然等价。
- 应在可丢弃的物理副本上演练确切 schema，制作并验证可恢复备份，停止并发写入，并比较迁移前后的目录状态与查询结果。
- 函数需要锁定并修改关系、更新系统目录和登记 `pg_pathman` 元数据的权限。只应允许迁移专用超级用户执行；实验后不再需要时应移除扩展。
