## 用法

来源：

- [已复核 commit 的 pg_algorithms README](https://github.com/kostiantyn-nemchenko/pg_algorithms/blob/e258c1cd1710f5867d21c864a01ad90f0f04326c/README.md)
- [已复核 commit 的 pg_algorithms.control](https://github.com/kostiantyn-nemchenko/pg_algorithms/blob/e258c1cd1710f5867d21c864a01ad90f0f04326c/pg_algorithms.control)
- [版本 1.0 的安装 SQL](https://github.com/kostiantyn-nemchenko/pg_algorithms/blob/e258c1cd1710f5867d21c864a01ad90f0f04326c/pg_algorithms--1.0.sql)
- [冒泡排序实现](https://github.com/kostiantyn-nemchenko/pg_algorithms/blob/e258c1cd1710f5867d21c864a01ad90f0f04326c/bubble_sort.c)
- [快速排序实现](https://github.com/kostiantyn-nemchenko/pg_algorithms/blob/e258c1cd1710f5867d21c864a01ad90f0f04326c/quick_sort.c)

`pg_algorithms` 是一个学习项目，提供两个用于排序一维 `int[]` 的 C 函数：`bubble_sort` 和 `quick_sort`。两者都会返回新的升序数组，并拒绝包含 `NULL` 的数组。

### 排序整数数组

```sql
CREATE EXTENSION pg_algorithms;

SELECT bubble_sort(ARRAY[9, 0, -2, 9, 4]);
SELECT quick_sort(ARRAY[9, 0, -2, 9, 4]);
```

### 注意事项

- 上游明确表示不要在生产中使用该项目。已复核源码停留在 2018 年，且没有现代 PostgreSQL 兼容性矩阵。
- `bubble_sort` 的比较开销为 `O(n²)`，手写递归快速排序在不利输入下也可能退化。运维工作负载更适合使用 PostgreSQL 原生排序设施。
- 版本 `1.0` 只实现上述两个函数，尽管项目名称涵盖更广的算法范围。SQL 签名接受 PostgreSQL `integer[]`，而非任意数值或多态数组。
