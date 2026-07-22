## 用法

来源：

- [目录版本对应的官方 README](https://github.com/kostiantyn-nemchenko/pg_algorithms/blob/e258c1cd1710f5867d21c864a01ad90f0f04326c/README.md)
- [目录版本对应的扩展 SQL](https://github.com/kostiantyn-nemchenko/pg_algorithms/blob/e258c1cd1710f5867d21c864a01ad90f0f04326c/pg_algorithms--1.0.sql)

`pg_algorithms` 1.0 为整数数组暴露冒泡排序和快速排序的演示实现。上游将其称为个人练习项目；它适合学习或小型实验，不应作为生产排序原语。

### 核心流程

```sql
CREATE EXTENSION pg_algorithms;

SELECT bubble_sort(ARRAY[5, 1, 4, 2, 3]);
SELECT quick_sort(ARRAY[5, 1, 4, 2, 3]);

-- PostgreSQL's native relational sort is the normal production choice.
SELECT array_agg(v ORDER BY v)
FROM unnest(ARRAY[5, 1, 4, 2, 3]) AS u(v);
```

### 函数索引

- `bubble_sort(integer[])` 使用冒泡排序返回升序的一维整数数组。
- `quick_sort(integer[])` 使用首元素为枢轴的递归快速排序，返回升序的一维整数数组。
- 两个函数均为 strict 和 stable，SQL 接口只暴露 integer 数组。

### 注意事项

- C 实现会拒绝包含 NULL 元素的数组和多维数组。
- 冒泡排序时间复杂度为平方级；该快速排序的枢轴选择在已有序或对抗性输入上也会退化为平方级并产生深递归。
- 实际工作负载应优先使用 PostgreSQL 原生 ORDER BY，它具备成熟的内存、磁盘溢写、排序规则和规划器集成。
