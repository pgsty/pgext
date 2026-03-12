

## 用法

> [xicor: PostgreSQL 的 XI 相关系数](https://github.com/Florents-Tselai/pgxicor)

提供 XI（xi）相关系数，可以检测 X 和 Y 之间的函数关系——是 `corr(X, Y)` 的更强大替代方案，适用于超越线性关系的场景。

```sql
CREATE EXTENSION xicor;
```

### 函数

| 函数 | 描述 |
|---|---|
| `xicor(x, y)` | 计算两个变量之间的 XI 相关系数 |

### 配置

对于存在并列数据的可重现结果：

```sql
SET xicor.ties = true;
SET xicor.seed = 42;
```

### 示例

```sql
CREATE TABLE xicor_test (x float8, y float8);
INSERT INTO xicor_test (x, y) VALUES
  (1.0, 2.0), (2.5, 3.5), (3.0, 4.0), (4.5, 5.5), (5.0, 6.0);

SELECT xicor(x, y) FROM xicor_test;
```
