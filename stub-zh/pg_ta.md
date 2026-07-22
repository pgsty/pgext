## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/magnusp/pg_ta/blob/899a2ce4ad674543d07bd007d2d85f5594baa30a/README.md)
- [1.0 版本 SQL 函数](https://github.com/magnusp/pg_ta/blob/899a2ce4ad674543d07bd007d2d85f5594baa30a/sql/pg_ta--1.0.sql)
- [C 实现](https://github.com/magnusp/pg_ta/blob/899a2ce4ad674543d07bd007d2d85f5594baa30a/src/pg_ta.c)
- [上游测试查询](https://github.com/magnusp/pg_ta/blob/899a2ce4ad674543d07bd007d2d85f5594baa30a/test.sql)

`pg_ta` 是 2014 年在 PostgreSQL 中调用 TA-Lib 的概念验证。1.0 版本只暴露 `ta_f(double precision[])`，返回周期硬编码为 5 的简单移动平均序列。

### 核心流程

构建时需要外部 TA-Lib 头文件与库。数组顺序就是观测顺序；PostgreSQL 不会从来源查询中推断排序。由于五值移动窗口尚未形成，返回结果的前四行是 NULL。

```sql
CREATE EXTENSION pg_ta;

SELECT *
FROM ta_f(ARRAY[1, 2, 3, 4, 5, 6]::double precision[]);
-- NULL, NULL, NULL, NULL, 3, 4
```

函数拒绝包含 NULL 的数组。它没有周期或 TA-Lib 移动平均类型参数，也没有实现其他指标。

### 关键安全边界

不要在隔离测试实例之外执行此扩展。已复核 C 源码按 TA-Lib 较短的 `outNbElement` 数量分配 PostgreSQL 结果数组，却按每个输入元素写入一项。周期为 5 时会向分配区之外写入四个 `Datum`，可能破坏后端内存。它还把值为零的 `Datum` 当作 NULL 标记，因此真实移动平均值 `0` 可能被输出为 SQL NULL。

上游明确声明该代码是实验性、不稳定的，绝不能用于接近生产的场景。源码早于当前 PostgreSQL 与 TA-Lib 版本；即使示例输出看似正确，也不能证明内存安全或兼容性。
