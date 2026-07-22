## 用法

来源：

- [上游 README](https://github.com/tomhebbron/pg_bitstring_helpers/blob/43011651dd5c600e7b9cb413d0ed3ea32683a1bd/README.rst)
- [1.0 版安装 SQL](https://github.com/tomhebbron/pg_bitstring_helpers/blob/43011651dd5c600e7b9cb413d0ed3ea32683a1bd/bitstring_helpers--1.0.sql)
- [C 实现](https://github.com/tomhebbron/pg_bitstring_helpers/blob/43011651dd5c600e7b9cb413d0ed3ea32683a1bd/bitstring_helpers.c)

bitstring_helpers 1.0 增加位计数、汉明距离、单个位邻居、拼接、转换和整数洗牌辅助函数。

### 位串操作

```sql
CREATE EXTENSION bitstring_helpers;
SELECT popcount(B'101101'::bit varying);
SELECT hamming_distance(B'101101'::bit varying, B'100001'::bit varying);
SELECT * FROM neighbours(B'101'::bit varying);
```

计算汉明距离时应使用等长操作数。neighbours 函数会为每一种单个位翻转返回一个值。

### 转换与聚合

`text2bitstring(text)` 将仅由零和一组成的字符串转换为 `bit varying`，`bit2text(bit varying)` 则将其转换回 `text`。扩展还创建了两个方向的显式类型转换。

安装脚本提供的聚合对象包括：将输入值收集为数组的 `aconcat(anyelement)`、拼接文本值的 `concat(text)`，以及拼接位值的 `concat(bit)`。这些声明使用 PostgreSQL 的旧式聚合语法，因此在依赖它们前，应确认 `CREATE EXTENSION` 能在目标大版本上成功执行。

### 注意事项

- shuffled_ints 实际生成从零到 limit 减一的数值，尽管 README 将范围描述为从一开始。
- 负数洗牌上限会在验证前参与内存分配，过大的上限也可能耗尽后端内存。绝不能传入不可信或无上限的数值。
- 洗牌使用基于取模的进程全局随机生成器，且存在偏差；不适用于密码学、安全、抽样或公平性敏感工作。
- 这份 2013 年代码没有当前 PostgreSQL 兼容矩阵、测试或已发布许可证。其旧式多态聚合声明在现代 PostgreSQL 版本上可能需要修改。
