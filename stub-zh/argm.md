## 用法

来源：

- [已复核 commit 的 Argm README](https://github.com/bashtanov/argm/blob/b8b2db36d0a57987d413f909e82369f9677e384a/README.md)
- [已复核 commit 的 argm.control](https://github.com/bashtanov/argm/blob/b8b2db36d0a57987d413f909e82369f9677e384a/argm.control)
- [版本 1.1.2 的安装 SQL](https://github.com/bashtanov/argm/blob/b8b2db36d0a57987d413f909e82369f9677e384a/argm--1.1.2.sql)
- [PGXN 上的 Argm](https://pgxn.org/dist/argm/)

`argm` 提供 `argmax`、`argmin` 和 `anyold` 聚合。前两个函数按可排序键元组的字典序，在组内找到最大或最小的行，并返回该行对应的值。`anyold` 返回聚合过程遇到的第一个非空值，适合选择由分组键函数依赖决定的属性。

### 基本示例

```sql
CREATE EXTENSION argm;

SELECT
  argmax(label, score) AS highest_label,
  argmin(label, score) AS lowest_label,
  anyold(label) AS one_label
FROM (VALUES
  ('alpha', 10),
  ('beta', 20),
  ('gamma', NULL)
) AS sample(label, score);
```

与 `DISTINCT ON` 写法不同，这些聚合可在同一个 `GROUP BY` 操作中与其他聚合一起使用，也可能采用哈希聚合。

### 注意事项

- 如果多行具有相同的键元组，`argmax` 或 `argmin` 会任意选择其中一个对应值。
- 值可以使用任意 PostgreSQL 类型，但排序键必须可排序。上游实现对整个键元组使用统一的排序方向和排序规则。
- 当前源码与目录标识为 `1.1.2`；最新可见 Git 标签为 `1.1.1`，PGXN 仍展示 `1.0.2`。不要根据旧 PGXN 软件包推断当前行为。
