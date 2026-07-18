## 用法

来源：

- [已复核 commit 的上游文档](https://github.com/JoanneBogart/postgres-cmath/blob/d2ead80fda9f41099d766190c536d630aa0c8249/README.md)
- [扩展控制文件](https://github.com/JoanneBogart/postgres-cmath/blob/d2ead80fda9f41099d766190c536d630aa0c8249/cmath.control)

`cmath` 以 `c_` 前缀提供 ISO C `<math.h>` 函数，包括单精度变体以及 PostgreSQL 传统上没有的函数。与 PostgreSQL 数值函数不同，这些 C 函数遇到定义域错误或溢出时通常返回 IEEE-754 `NaN` 或无穷，而不是抛出 SQL 错误。

```sql
CREATE EXTENSION cmath;
SELECT c_log10(100), c_exp(1e100);
SELECT c_sinf(0.1), c_cosf(0.1);
```

结果遵循平台 C 库，因此边界行为可能随操作系统和架构变化。若下游约束、JSON 序列化、索引或分析要求有限数值，调用方应显式拒绝非有限值。已复核版本为 `1.2`；依赖可重复数值结果前，应在实际目标平台测试。
