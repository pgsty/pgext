## 用法

来源：

- [官方 README](https://github.com/Zeleo/pg_natural_sort_order/blob/16d2ff812aabfaab6b83056ffbf5de200c69a5b6/README.md)
- [官方控制文件](https://github.com/Zeleo/pg_natural_sort_order/blob/16d2ff812aabfaab6b83056ffbf5de200c69a5b6/pg_natural_sort_order.control)
- [官方 PGXN 版本](https://pgxn.org/dist/pg_natural_sort_order/1.0.1/)

`pg_natural_sort_order` 将文本中的连续数字映射为固定宽度文本，使普通 PostgreSQL 排序与 B-tree 索引能够得到人类习惯的顺序，例如让 `item2` 排在 `item10` 前。上游只在 PostgreSQL 9.x 上测试过版本 `1.0.1`；用于当前服务器前必须重新构建并验证。

### 核心流程

选择足以容纳数据中所有连续数字的归一化宽度，并创建表达式索引：

```sql
CREATE EXTENSION pg_natural_sort_order;

CREATE INDEX products_name_natural_idx
  ON products (natural_sort_order(name, 75));

SELECT id, name
FROM products
ORDER BY natural_sort_order(name, 75);
```

查询表达式必须与索引表达式一致，规划器才能直接使用该索引。

### 函数行为

`natural_sort_order(text, integer)` 会把每段连续十进制数字替换为按指定宽度补零的表示。上游允许的宽度范围是 `1` 到 `150`，示例使用 `75`。宽度越大，可处理的长数字越多，但归一化值与索引也越大。

### 运维注意事项

如果输入中的连续数字比所选宽度更长，README 警告该数字会被拆分，排序很可能错误。应依据真实数据最大值选择宽度；若正确性重要，还应约束该上限，并在表达式或宽度变化后重建索引。该映射是文本归一化，不是区域感知排序，也不会按数字语义解析符号、小数或版本号。项目没有证明支持旧 PostgreSQL 9.x 之外的版本，因此需要在目标大版本上使用有代表性的 Unicode 与标点数据对比结果。
