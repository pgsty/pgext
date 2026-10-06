## 用法

来源：

- [README.md](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/README.md)
- [contrib/alohadb_vectorize/alohadb_vectorize--1.0.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_vectorize/alohadb_vectorize--1.0.sql)
- [contrib/alohadb_vectorize/alohadb_vectorize.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_vectorize/alohadb_vectorize.c)
- [contrib/alohadb_vectorize/alohadb_vectorize.control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_vectorize/alohadb_vectorize.control)

`alohadb_vectorize` 1.0 在 AlohaDB 中为分析 SQL 应用局部内存、并行规划与 JIT 设置。实现使用 PostgreSQL SPI，并未引入新的向量化执行引擎。

### 核心用法

```sql
CREATE EXTENSION alohadb_vectorize;
BEGIN;
SELECT * FROM vectorize_status();
SELECT * FROM vectorize_query('SELECT count(*) FROM pg_class') AS t(n bigint);
SELECT * FROM vectorize_explain('SELECT count(*) FROM pg_class');
ROLLBACK;
```

### 运行边界

`vectorize_query` 返回记录集，调用者须声明输出列；`vectorize_explain` 使用 EXPLAIN ANALYZE，因此会执行传入的查询。`vectorize_status` 显示相关设置，`vectorize_benchmark` 重复执行查询并返回耗时。

这些封装通过 `SET LOCAL` 在事务内应用设置，包括增大工作内存及调整并行工作进程。实验应在有明确边界的事务中进行，并考虑并发查询的总内存需求。不要传入不可信 SQL。安装需要超级用户，未声明预加载或重启要求。核对的发行源码属于基于 PostgreSQL 18 的 AlohaDB 分支，尚未确认原版 PostgreSQL 的软件包支持。
