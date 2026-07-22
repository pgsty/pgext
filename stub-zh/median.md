## 用法

来源：

- [已复核 commit 的 Median README](https://github.com/za-arthur/pg_median/blob/0a43525499153317020d8cd71f8c3258c2f34fc5/README.md)
- [已复核 commit 的聚合定义](https://github.com/za-arthur/pg_median/blob/0a43525499153317020d8cd71f8c3258c2f34fc5/median--1.0.sql)
- [已复核 commit 的 Median 实现](https://github.com/za-arthur/pg_median/blob/0a43525499153317020d8cd71f8c3258c2f34fc5/median.c)

`median` 定义了名为 `median(anyelement)` 的并行安全聚合。它忽略空值，对其余所有值排序，并返回输入类型的值。它适合编写紧凑的分析查询，但在大型或混合类型负载中，需要留意依赖类型的算术语义和每组内存用量。

### 核心工作流

```sql
CREATE EXTENSION median;

SELECT median(value)
FROM (VALUES
  (1.0::numeric),
  (9.0::numeric),
  (3.0::numeric),
  (7.0::numeric)
) AS sample(value);
```

非空输入为奇数个时，聚合返回排序后的中间值。输入为偶数个时，它把两个中间值相加，再除以通过输入类型构造的数值二。空分组或全空分组返回 `NULL`。

SQL 接口除聚合外，还公开内部的转移、终结、合并、序列化与反序列化函数。它们用于支持并行聚合，属于实现细节，而不是应用层 API。

### 类型与规模限制

- 输入类型必须可排序。偶数基数的输入还必须提供兼容的 `+` 和 `/` 操作符。
- 由于返回类型就是输入类型，整数输入会采用整数算术：中点可能被截断，两个中间整数相加也可能溢出。需要精确小数中点时，应先把整数转换为 `numeric`。
- 每个聚合状态会保留全部非空值，终结函数再对它们排序。每组状态内存大致线性增长，排序工作量为 `O(n log n)`；并行工作进程会分别维护状态，然后再合并。
- 如果更适合插值语义或标准有序集合聚合，可考虑 PostgreSQL 内置的 `percentile_cont(0.5) WITHIN GROUP`。其返回类型和查询计划需要另行确认。
- 上游版本 `1.0` 没有发布当前 PostgreSQL 兼容矩阵。应针对生产使用的服务器大版本和输入类型进行测试。
