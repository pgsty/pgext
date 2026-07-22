## 用法

来源：

- [官方控制文件](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/probabilistic/probabilistic_counter.control)
- [官方 probabilistic estimator README](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/probabilistic/README.md)

`probabilistic_counter` 将 Flajolet-Martin 概率计数器实现为 `probabilistic_estimator` 状态类型与聚合。它牺牲精确性以获得紧凑、可配置的去重计数估算。

### 核心流程

```sql
CREATE EXTENSION probabilistic_counter;

SELECT probabilistic_distinct(i, 4, 32)
FROM generate_series(1, 100000) AS s(i);
```

显式参数分别是状态字节数和盐值数量。单参数重载 `probabilistic_distinct(anyelement)` 使用 4 字节和 32 个盐值。

### API

- `probabilistic_init(int, int)` 创建 `probabilistic_estimator`；`probabilistic_size(int, int)` 返回其大小。
- `probabilistic_add_item(probabilistic_estimator, anyelement)` 更新状态，`probabilistic_get_estimate(probabilistic_estimator)` 返回估算值。
- `probabilistic_reset(probabilistic_estimator)` 清空状态，`length(probabilistic_estimator)` 返回其存储大小。
- `probabilistic_distinct(anyelement, int, int)` 是可配置聚合；`probabilistic_distinct(anyelement)` 使用默认值。

### 运维说明

1.3.3 版控制文件可重定位，未声明前置扩展或预加载要求。上游仓库已经归档并只读。更大的状态可能改善表现，但会增加内存消耗；频繁更新持久化状态也可能加剧 MVCC 膨胀。应在代表性分布上测试估算器，正确性敏感的工作仍应使用精确计数。
