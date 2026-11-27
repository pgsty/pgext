## 用法

来源：

- [官方源码 1](https://github.com/bojan-janjic-nz/pg_monotonic_chi2/blob/d14e656b20345d3af38ffc9198c15f16efe6d041/README.md)
- [官方源码 2](https://github.com/bojan-janjic-nz/pg_monotonic_chi2/blob/d14e656b20345d3af38ffc9198c15f16efe6d041/monotonic_chi2.control)
- [官方源码 3](https://github.com/bojan-janjic-nz/pg_monotonic_chi2/blob/d14e656b20345d3af38ffc9198c15f16efe6d041/monotonic_chi2--1.0.sql)
- [官方源码 4](https://github.com/bojan-janjic-nz/pg_monotonic_chi2/blob/d14e656b20345d3af38ffc9198c15f16efe6d041/monotonic_chi2.c)

`monotonic_chi2` 1.0 根据数值特征与二元目标拟合卡方分箱。它先合并样本占比过低的箱，再执行相邻 ChiMerge，最后可选用 PAVA 使事件率单调。仓库名为 `pg_monotonic_chi2`，SQL 扩展名则是 `monotonic_chi2`。

### 启用与使用

扩展包含 C 函数与 PL/pgSQL 辅助函数。安装需要超级用户；支持迁移到其他模式，不需要预加载或重启。上游未声明支持的 PostgreSQL 主版本范围。

```sql
CREATE EXTENSION monotonic_chi2;

SELECT monotonic_chi2_cutpoints(
    ARRAY[10, 20, 30, 40]::float8[],
    ARRAY[1, 5, 10, 20]::int8[],
    ARRAY[19, 15, 10, 0]::int8[],
    alpha => 0.05, max_bins => 3, min_pct => 0.05,
    monotonic => true
);

SELECT monotonic_chi2_bin(20::float8, ARRAY[10, 20, 30]::float8[]);
```

第二条查询把数值分入第 2 箱。边界采用右闭语义：各区间为 `(previous_cut, cut]`，超过最后一个切点的数值进入末箱。

### 表级辅助函数

假设现有表包含数值列 `age` 和二元目标列 `defaulted`：

```sql
SELECT fit_chi2_bins('public.customers'::regclass, 'age', 'defaulted',
                    monotonic => false);
SELECT * FROM summarize_chi2_bins('public.customers'::regclass,
                                 'age', 'defaulted', monotonic => false);
```

`fit_chi2_bins` 先合并相等的有限特征值再预分箱，避免拆散并列值。`summarize_chi2_bins` 返回 `bin_id`、`min_val`、`max_val`、`n_rows`、`events`、`non_events` 和 `event_rate`，也可通过可选参数 `cutpoints` 复用已有切点。连续目标可指定 `target_threshold`，超过阈值的记录算作事件；未指定时，目标必须是 0/1 或布尔值。

### 参数与边界

- `alpha` 必须位于 `(0, 1)`，`max_bins` 至少为 1，`min_pct` 位于 `[0, 0.5)`，`n_prebins` 至少为 1；默认值依次为 0.05、6、0.05 和 100。
- `monotonic => false` 跳过 PAVA，适合 U 形或多峰关系。
- 拟合会排除空目标值以及空值、NaN、无穷大的特征值。拟合输入为空时返回空切点数组。使用空数组分配时落入第 1 箱，空值或 NaN 返回 NULL。汇总将空值/NaN 特征归入第 0 箱，并排除空目标值。
- 辅助函数以调用者权限读取指定表。应先在训练数据上确定并保存切点，再用于评分数据；扩展不会代替你验证预测模型。
