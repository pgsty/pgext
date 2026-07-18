## 用法

来源：

- [上游用法文档](https://github.com/Kozea/weighted_mean/blob/4bf67038b72acadf5e83e7d58f9fff26ce202eb9/doc/weighted_mean.md)
- [C 聚合实现](https://github.com/Kozea/weighted_mean/blob/4bf67038b72acadf5e83e7d58f9fff26ce202eb9/src/weighted_mean.c)
- [SQL 聚合声明](https://github.com/Kozea/weighted_mean/blob/4bf67038b72acadf5e83e7d58f9fff26ce202eb9/sql/weighted_mean.sql)

`weighted_mean` 提供接受两个 `numeric` 参数的加权平均聚合。它用值乘权重的总和除以权重总和；没有输入行或总权重为零时返回零。

```sql
CREATE EXTENSION weighted_mean;
SELECT weighted_mean(unitprice, quantity)
FROM sales
WHERE unitprice IS NOT NULL AND quantity IS NOT NULL;
```

转移函数未声明 strict，却会解引用两个输入，因此必须显式过滤 NULL；异常 NULL 调用可能导致后端崩溃。负权重和总权重为零也可能产生误导结果。仓库已归档，聚合也缺少现代并行状态支持；应回归测试精确数值行为，实际可行时优先使用明确且受维护的 SQL 表达式。
