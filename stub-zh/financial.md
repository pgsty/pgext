

## 用法

> [financial: PostgreSQL 的金融计算函数](https://github.com/intgr/pg_financial)

提供 XIRR（不规则内部收益率）聚合函数，类似于电子表格程序中的 XIRR 函数。

```sql
CREATE EXTENSION financial;
```

### 函数

| 函数 | 说明 |
|---|---|
| `xirr(amount float8, date timestamptz [, guess float8] ORDER BY date)` | 计算不规则内部收益率 |

### 示例

```sql
-- 基本 XIRR 计算
SELECT xirr(amount, time ORDER BY time) FROM transaction;
-- 0.0176201237088334

-- 使用显式初始猜测值（Excel 兼容的默认值为 0.1）
SELECT xirr(amount, time, 0.1 ORDER BY time) FROM transaction;

-- 按投资组合计算 XIRR
SELECT portfolio, xirr(amount, time ORDER BY time)
FROM transaction
GROUP BY portfolio;

-- 作为窗口函数使用
SELECT xirr(amount, time) OVER (ORDER BY time)
FROM transaction;
```

当牛顿法无法收敛时返回 NULL。提供更好的猜测值在某些情况下可能有所帮助。
