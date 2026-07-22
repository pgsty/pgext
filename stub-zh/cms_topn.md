## 用法

来源：

- [官方 README](https://github.com/ozturkosu/cms_topn/blob/78ce0d1e0437c0b35419d963685d5de57a87078e/README.md)
- [官方扩展控制文件](https://github.com/ozturkosu/cms_topn/blob/78ce0d1e0437c0b35419d963685d5de57a87078e/cms_topn.control)
- [官方扩展 SQL](https://github.com/ozturkosu/cms_topn/blob/78ce0d1e0437c0b35419d963685d5de57a87078e/cms_topn--1.0.0.sql)

`cms_topn` 1.0.0 把近似 Count-Min Sketch 与有限的高频候选列表存储在一起。它适合在精确聚合成本过高时进行紧凑频率估算、近似 Top-N 报告，以及合并不同分区的摘要。

### 核心流程

使用 `cms_topn_add_agg` 为每个分组构建摘要，以 `topn` 查询候选频率，再用 `cms_topn_union_agg` 合并兼容摘要。

```sql
CREATE EXTENSION cms_topn;

CREATE TABLE daily_hits (
    day date PRIMARY KEY,
    users cms_topn NOT NULL
);

INSERT INTO daily_hits (day, users)
SELECT visited_at::date,
       cms_topn_add_agg(user_id, 10, 0.001, 0.99)
FROM visits
GROUP BY visited_at::date;

SELECT day, item, frequency
FROM daily_hits
CROSS JOIN LATERAL topn(users, NULL::bigint);

SELECT *
FROM topn(
    (SELECT cms_topn_union_agg(users) FROM daily_hits),
    NULL::bigint
);
```

传给 `topn` 的有类型 null 决定其 `item` 列的 SQL 类型。将要合并的摘要应使用相同输入类型和 sketch 参数。

### 主要对象

- `cms_topn(n, error_bound, confidence)` 创建空摘要，默认值为 `0.001` 和 `0.99`。
- `cms_topn_add(summary, value)` 添加一个值；`cms_topn_add_agg` 的聚合重载从列构建摘要。
- `cms_topn_frequency(summary, value)` 返回估算频率。
- `topn(summary, typed_null)` 返回候选 `item` 及估算 `frequency` 行。
- `cms_topn_union` 与 `cms_topn_union_agg` 合并兼容摘要；`cms_topn_info` 报告 sketch 维度和元数据。

更小的误差界会增加 sketch 宽度，更高的置信度会增加深度；两者都以更多内存和 CPU 换取更好的统计保证。

### 近似与兼容性边界

Count-Min Sketch 会因哈希碰撞而高估频率，保留候选列表也不会让结果变成精确值。应使用接近生产的数据分布验证可接受误差，特别是严重倾斜、大量不同值及合并摘要场景。不要把近似计数用于计费、完整性约束或其他要求精确的决策。

源码与测试停留在 2015 年。二进制自定义类型保存依赖内部类型的数据，因此未在准确构建上测试前，不能假定 dump/restore、复制、体系结构或 PostgreSQL 大版本兼容性。应把序列化摘要视为可从源数据重建的派生数据。
