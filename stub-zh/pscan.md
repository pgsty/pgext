## 用法

来源：

- [固定提交的上游 README](https://github.com/MasahikoSawada/pscan/blob/fd5ff5c53cc80ea2625c27d469bf3c7a6a3eb232/README.md)
- [固定提交的扩展控制文件](https://github.com/MasahikoSawada/pscan/blob/fd5ff5c53cc80ea2625c27d469bf3c7a6a3eb232/pscan.control)

`pscan` 是 PostgreSQL 并行堆扫描测试工具。`p_tuplescan(regclass)` 用于练习 PostgreSQL 9.6 并行顺序扫描所采用的并行元组扫描路径，`p_brangescan(regclass)` 则把关系划分为连续块范围并分配给工作进程。重载形式可以显式指定工作进程数，GUC `pscan.n_workers` 提供默认值。

```sql
CREATE EXTENSION pscan;
SET pscan.n_workers = 4;

SELECT p_tuplescan('public.fact_table'::regclass);
SELECT p_brangescan('public.fact_table'::regclass, 4);
```

这些函数不是保证正确性的 SQL 扫描替代品。上游 README 明确说明它们跳过可见性检查与过滤、读取全部元组，并假设表没有垃圾元组且全部可见。块范围方法会让工作进程并发读取相距较远的范围，因此可能产生较不连续的 I/O。

版本 `1.0` 只能用于可丢弃的基准测试关系，并且必须核实准确的 PostgreSQL 构建，因为它依赖内部堆与并行工作进程 API。不要把函数授权给应用角色；不应拿其返回值替代普通查询，也不要在未控制缓存状态、表可见性、存储类型、工作进程启动和并发负载时使用计时结果。
