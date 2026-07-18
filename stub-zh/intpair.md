## 用法

来源：

- [已复核 commit 的 pg_intpair README](https://github.com/citusdata/pg_intpair/blob/fa274d8aa6a73670b9cee95e94373022c2777926/README.md)
- [已复核 commit 的 intpair 0.2 安装 SQL](https://github.com/citusdata/pg_intpair/blob/fa274d8aa6a73670b9cee95e94373022c2777926/intpair--0.2.sql)
- [已复核 commit 的 intpair 回归 SQL](https://github.com/citusdata/pg_intpair/blob/fa274d8aa6a73670b9cee95e94373022c2777926/sql/intpair.sql)

0.2 版 `intpair` 提供定长 `int64pair` 类型：以 16 字节存储两个有符号 64 位整数。使用 `int64pair(a, b)` 构造值，并用下标读取从零开始的分量。它提供比较、B-tree 与哈希支持，因此可用作普通键和索引列。

```sql
CREATE EXTENSION intpair;

CREATE TABLE pair_demo (
  id int64pair PRIMARY KEY
);

INSERT INTO pair_demo VALUES
  (int64pair(-1, 1)),
  (int64pair(10, 20));

SELECT id, id[0] AS first_value, id[1] AS second_value
FROM pair_demo
ORDER BY id;
```

### 注意事项

- 下标从零开始：`p[0]` 是第一个整数，`p[1]` 是第二个整数。
- 这是紧凑标量类型，不是 PostgreSQL 复合类型。迁移结构相同的复合类型列时，需要像上游示例那样显式经文本转换，或先定义兼容类型转换。
- 索引排序遵循扩展的比较实现。用它替换复合类型或数组表示之前，应验证依赖排序的应用行为。
