## 用法

来源：

- [上游函数文档](https://github.com/uptimejp/pg_part/blob/6a5c96db670f75fa486006536e659d5c5f4b876d/README.md)
- [0.1.0 版本 SQL 实现](https://github.com/uptimejp/pg_part/blob/6a5c96db670f75fa486006536e659d5c5f4b876d/pg_part--0.1.0.sql)
- [PGXN 分发页](https://pgxn.org/dist/pg_part/)

`pg_part` 在固定的 `pgpart` schema 中提供 PL/pgSQL 辅助函数，用于传统继承式分区。它可以添加、合并、挂接、分离和列出子表。

```sql
CREATE EXTENSION pg_part;
SELECT pgpart.show_partition('public', 'orders');
SELECT pgpart.attach_partition('public', 'orders', 'orders_2025',
                               'order_date >= ''2025-01-01''');
```

添加与合并路径会构造动态 DDL，并通过服务端 `COPY` 文件搬移行。标识符、检查表达式和文件路径被直接拼入 SQL，因此只能授权可信管理员执行，绝不能传入用户可控文本。它早于声明式分区；处理重要数据前，应在副本上演练锁、约束、索引、权限、失败回滚、磁盘容量和清理，新系统应优先使用内核声明式分区。
