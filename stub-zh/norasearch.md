## 用法

来源：

- [已复核 commit 的官方 README](https://github.com/lemmerelassal/NoraSearch/blob/0c11af31b5aea1c58ca401e57966870b13b19ddf/README.md)
- [1.0 版本 SQL 声明](https://github.com/lemmerelassal/NoraSearch/blob/0c11af31b5aea1c58ca401e57966870b13b19ddf/norasearch--1.0.sql)
- [匹配实现](https://github.com/lemmerelassal/NoraSearch/blob/0c11af31b5aea1c58ca401e57966870b13b19ddf/norasearch.c)
- [扩展 control 文件](https://github.com/lemmerelassal/NoraSearch/blob/0c11af31b5aea1c58ca401e57966870b13b19ddf/norasearch.control)

`norasearch` 提供一个面向字节的相似度函数。它把 needle 与 haystack 的每个非负位置对齐，统计相等的字节位置，再返回计数严格大于阈值的对齐结果。这更像互相关评分，并不是普通的连续子串搜索。

### 核心流程

```sql
CREATE EXTENSION norasearch;

SELECT norasearch('abracadabra', 'abra', 2);
-- {{0,4},{7,4}}
```

函数签名是 `norasearch(hay text, needle text, minmatch int) RETURNS int[][]`。每个结果行是 `{offset,count}`。`offset` 从零开始，只考虑从 haystack 起点向后的偏移。负数 `minmatch` 会按零处理；匹配采用严格条件 `count > minmatch`。

### 结果与性能边界

函数对 SQL NULL 输入是 strict。空字符串、没有合格对齐或其他零结果情况返回 `NULL`，不是空数组。实现比较原始字节，因此偏移是字节位置，多字节 UTF-8 字符可能被拆开；不能把结果理解为字符下标。

它没有索引操作符类或规划器集成。实现会让每个 haystack 位置与 needle 比较，每次调用的最坏工作量与两者长度都相关，并分配结果数组。在许多行上应用 `norasearch` 前，应使用大型或重复值进行基准测试。
