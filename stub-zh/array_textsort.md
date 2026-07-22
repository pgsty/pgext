## 用法

来源：

- [官方 README](https://github.com/2ndQuadrant/array_textsort/blob/dd60ea1a5403e1ff938aeb6d26d0e9a901ea30f4/README.md)
- [1.1 版本扩展 SQL](https://github.com/2ndQuadrant/array_textsort/blob/dd60ea1a5403e1ff938aeb6d26d0e9a901ea30f4/array_textsort--1.1.sql)
- [C 实现](https://github.com/2ndQuadrant/array_textsort/blob/dd60ea1a5403e1ff938aeb6d26d0e9a901ea30f4/array_textsort.c)

`array_textsort` 为一维 `text[]` 值增加排序与去重函数。应用需要规范化数组顺序，或需要以数组而不是行表示不重复集合时，可以使用它。

### 核心流程

```sql
CREATE EXTENSION array_textsort;

SELECT array_textsort(ARRAY['pear', 'apple', 'banana']);
SELECT array_distinct(ARRAY['pear', 'apple', 'pear', 'banana']);
```

两个函数都按照数据库区域设置排序。`array_distinct` 随后从排序结果中删除相邻重复项。

### 对象索引

- `array_textsort(text[]) RETURNS text[]` 对一维文本数组排序。
- `array_distinct(text[]) RETURNS text[]` 对数组排序并删除重复元素。

### 运维说明

版本 `1.1` 可重定位，且未声明扩展依赖或预加载要求。两个 C 函数均为 `STRICT` 和 `IMMUTABLE`，因此 SQL `NULL` 输入会返回 SQL `NULL`。

实现会拒绝包含空元素的数组。只能将其用于一维数组：这是文档定义的接口，多维行为不是受支持的合同。比较顺序由区域设置决定，因此相同字符串在不同区域设置的数据库中可能有不同顺序。如果需要跨数据库稳定排序，应在这些函数外规范化值，或使用显式排序规则的行式查询。
