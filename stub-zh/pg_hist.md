## 用法

来源：

- [Official README](https://github.com/segasai/pg_hist/blob/3e67428bcc5017c8fda9a8635ca6a076417f30d7/README.md)
- [Extension SQL](https://github.com/segasai/pg_hist/blob/3e67428bcc5017c8fda9a8635ca6a076417f30d7/scripts/pg_hist--1.0.0.sql)
- [Histogram implementation](https://github.com/segasai/pg_hist/blob/3e67428bcc5017c8fda9a8635ca6a076417f30d7/pg_hist.c)

pg_hist 从给定 SQL 查询结果计算一到十维直方图。它只返回非空箱的坐标和计数；加权变体则返回浮点合计。当查询来源可信时，它适合临时数值汇总。

### 核心流程

箱数量、下界和上界数组长度必须相同。源查询要为每个维度返回一个数值列：

```sql
CREATE EXTENSION pg_hist;

SELECT *
FROM pg_hist_2d(
  'SELECT x, y FROM measurement',
  ARRAY[20, 10],
  ARRAY[0.0, -5.0],
  ARRAY[100.0, 5.0]
);
```

如需加权直方图，在查询最后一列追加权重，并调用对应的 `_w` 函数。

### 重要对象

- `pg_hist_1d`、`pg_hist_2d` 和 `pg_hist_3d` 返回固定形状的非空整数计数箱。
- `pg_hist` 支持任意维度，但调用方必须提供 record 结构定义。
- `pg_hist_1d_w`、`pg_hist_2d_w` 和 `pg_hist_3d_w` 返回加权合计。
- `pg_hist_w` 是任意维度的加权形式。

### 限制与安全说明

查询文本会在调用者后端中准备并执行；绝不能用不可信输入拼接它。实现会先分配稠密累加器，然后才省略空箱，最多允许 10 个维度和一亿个总箱，并以每批 1,000 行读取。超出给定范围的值或 null 输入不参与统计。大量箱组合会占用可观后端内存，因此要校验数组大小，并避开延迟敏感会话运行昂贵直方图。
