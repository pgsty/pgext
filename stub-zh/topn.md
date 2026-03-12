

## 用法

> [topn: PostgreSQL 的 Top-N 值近似估计](https://github.com/citusdata/postgresql-topn)

提供近似的 Top-N 值跟踪，使用近似算法维护预定数量的高频项及其计数器。支持物化、增量更新以及跨时间区间的合并。

```sql
CREATE EXTENSION topn;
```

### 数据类型

使用 `JSONB` 存储高频项及其频率。

### 聚合函数

| 函数 | 描述 |
|---|---|
| `topn_add_agg(text)` | 从文本列创建 JSONB 计数器的聚合函数 |
| `topn_union_agg(jsonb)` | 合并多个 JSONB 计数器列表的聚合函数 |

### 函数

| 函数 | 描述 |
|---|---|
| `topn(jsonb, n)` | 返回 Top-N 元素及其频率的行集 |
| `topn_add(jsonb, text)` | 向 JSONB 计数器添加一个文本值 |
| `topn_union(jsonb, jsonb)` | 合并两个 JSONB 计数器列表 |

### 配置

- `topn.number_of_counters` -- 要跟踪的计数器数量（默认：1000）

### 示例

```sql
-- 按日期物化热门产品
CREATE TABLE popular_products (
  review_date date UNIQUE,
  agg_data jsonb
);

INSERT INTO popular_products
SELECT review_date, topn_add_agg(product_id)
FROM customer_reviews GROUP BY review_date;

-- 获取每天排名第一的产品
SELECT review_date, (topn(agg_data, 1)).*
FROM popular_products ORDER BY review_date;

-- 跨时间范围的 Top 10（合并每日摘要）
SELECT (topn(topn_union_agg(agg_data), 10)).*
FROM popular_products
WHERE review_date >= '2000-01-01' AND review_date < '2000-02-01'
ORDER BY 2 DESC;
```
