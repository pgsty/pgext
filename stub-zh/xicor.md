## 用法

来源：[README](https://github.com/Florents-Tselai/pgxicor/blob/main/README.md), [release 0.1.1](https://github.com/Florents-Tselai/pgxicor/releases/tag/v0.1.1)

`xicor` 将 XI（Chatterjee's xi）相关系数暴露为 PostgreSQL 聚合函数。它适合检测函数依赖，包括 Pearson `corr()` 可能遗漏的非线性关系。

```sql
CREATE EXTENSION xicor;
```

### 主要聚合

```sql
SELECT xicor(x, y) FROM xicor_test;
```

上游示例把它与抛物线数据集上的 `corr()` 对比：`corr()` 接近零，而 `xicor()` 仍然较高。

### 示例

```sql
CREATE TABLE xicor_test (x float8, y float8);
INSERT INTO xicor_test (x, y) VALUES
  (1.0, 2.0),
  (2.5, 3.5),
  (3.0, 4.0),
  (4.5, 5.5),
  (5.0, 6.0);

SELECT xicor(x, y) FROM xicor_test;
```

### 可复现性控制

对于存在 ties 的数据，上游建议开启确定性 tie 处理：

```sql
SET xicor.ties = true;
SET xicor.seed = 42;
```

### 注意事项

- `xicor()` 是双数值输入的聚合函数，不是通用统计框架。
- ties 的处理方式会影响结果；如果需要可复现行为，请启用文档中的 GUC。
