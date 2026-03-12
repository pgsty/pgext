

## 用法

> [vasco: PostgreSQL 的最大信息系数（MIC）扩展](https://github.com/Florents-Tselai/vasco)

使用最大信息系数（MIC）和 MINE 系列统计量发现数据中隐藏的相关性。

```sql
CREATE EXTENSION vasco;
```

### 聚合函数

| 函数 | 描述 |
|---|---|
| `mic(x, y)` | 最大信息系数——检测任何关系 |
| `mas(x, y)` | 最大不对称分数——偏离单调性的程度 |
| `mev(x, y)` | 最大边值——连续函数采样的程度 |
| `mcn(x, y)` | 最小单元数——关联的复杂度 |
| `mcn_general(x, y)` | MCN，其中 `eps = 1 - MIC` |
| `tic(x, y)` | 总信息系数 |
| `gmic(x, y)` | 广义平均信息系数 |

### 工具函数

| 函数 | 描述 |
|---|---|
| `vasco_corr_matrix(table_name, output_table)` | 计算所有列对的 MIC 并存储为相关矩阵表 |

### 配置

```sql
SET vasco.mic_estimator = 'ApproxMIC';  -- 或 'MIC_e'
SET vasco.mine_c = ...;
SET vasco.mine_alpha = ...;
```

### 示例

```sql
-- 计算列对之间的 MIC
SELECT mic(x, cubic), mic(x, periodic), mic(x, rand_y)
FROM vasco_data;
-- 1, 1, 0.15

-- 创建完整的相关矩阵
SELECT vasco_corr_matrix('my_table', 'mic_my_table');
SELECT * FROM mic_my_table;
```
