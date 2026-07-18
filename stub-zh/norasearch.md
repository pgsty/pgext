## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/lemmerelassal/NoraSearch/blob/0c11af31b5aea1c58ca401e57966870b13b19ddf/README.md)
- [1.0 版本 SQL 声明](https://github.com/lemmerelassal/NoraSearch/blob/0c11af31b5aea1c58ca401e57966870b13b19ddf/norasearch--1.0.sql)
- [扩展 control 文件](https://github.com/lemmerelassal/NoraSearch/blob/0c11af31b5aea1c58ca401e57966870b13b19ddf/norasearch.control)

`norasearch` 提供一个 C 函数，在每个可能偏移位置比较两个字符串，并返回对齐字节匹配数超过阈值的位置。结果是由 `{offset,count}` 二元组构成的二维整数数组。

```sql
CREATE EXTENSION norasearch;
SELECT norasearch('abracadabra', 'abra', 2);
-- {{0,4},{7,4}}
```

函数签名为 `norasearch(hay text, needle text, minmatch int)`。偏移量从零开始，并且仅输出满足 `count > minmatch` 的记录。实现按字节而非按区域设置感知字符进行比较，因此用于面向用户的搜索前应验证多字节文本行为。

当前 README 提供 Windows/MSYS2 与 PostgreSQL 18 的构建说明。它没有索引运算符类或规划器集成：每次调用都在函数内完成匹配，因此应使用有代表性的字符串长度和行数进行基准测试。
