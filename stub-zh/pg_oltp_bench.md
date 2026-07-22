## 用法

来源：

- [官方扩展 SQL](https://github.com/postgrespro/pg_oltp_bench/blob/fbf92a58e12f6e155a4d2c09969063dc26e8a2a1/pg_oltp_bench--1.0.sql)
- [官方初始化脚本](https://github.com/postgrespro/pg_oltp_bench/blob/fbf92a58e12f6e155a4d2c09969063dc26e8a2a1/oltp_init.sql)
- [官方只读负载](https://github.com/postgrespro/pg_oltp_bench/blob/fbf92a58e12f6e155a4d2c09969063dc26e8a2a1/oltp_ro.sql)
- [官方读写负载](https://github.com/postgrespro/pg_oltp_bench/blob/fbf92a58e12f6e155a4d2c09969063dc26e8a2a1/oltp_rw.sql)

`pg_oltp_bench` 把一个随机字符串辅助函数与一组 `pgbench` 脚本组合起来，用于近似 sysbench 风格的 OLTP 负载。它是基准测试夹具，本身不是应用模式，也不是标准化分数。

### 准备测试夹具

创建扩展，然后在一次性的基准测试数据库中运行上游初始化脚本：

```sql
CREATE EXTENSION pg_oltp_bench;

SELECT sb_rand_str('order-########-@@@@');
```

在 `sb_rand_str(text)` 中，每个 ASCII `#` 会替换成随机数字，每个 ASCII `@` 会替换成随机小写字母，其他字节保持不变。随附的 `oltp_init.sql` 会删除并重建 `sbtest`，插入一千万行，并在 `k` 上建立索引。

```bash
psql "$BENCHMARK_URL" -f oltp_init.sql
pgbench "$BENCHMARK_URL" -n -c 16 -j 8 -T 300 -f oltp_ro.sql
pgbench "$BENCHMARK_URL" -n -c 16 -j 8 -T 300 -f oltp_rw.sql
```

`oltp_ro.sql` 在每个事务脚本中执行十次点查和四次范围、聚合或排序查询。`oltp_rw.sql` 还会在显式事务中执行更新、删除和补位插入。

### 基准测试纪律

- `oltp_init.sql` 以 `DROP TABLE IF EXISTS sbtest` 开头；绝不能把它指向包含需要保留的同名表的数据库。
- 脚本把 `table_size` 固定为一千万、把 `range_size` 固定为 100。若改变夹具规模，必须同步修改脚本，否则随机 ID 和范围边界将不再描述实际数据集。
- 应记录 PostgreSQL 设置、硬件、客户端数、线程数、持续时间、预热、数据规模、持久性配置以及数据是否能全部放入内存。缺少这些控制条件的结果不可比较。
- 分别运行只读和读写负载，监控饱和度与错误，并在故障后验证表状态。读写脚本会产生与负载相关的数据变动，并消耗 WAL 和存储。
- `sb_rand_str` 是易变函数，使用扩展中的简单伪随机生成。它适合填充基准数据，不适用于机密、标识符或密码学材料。
