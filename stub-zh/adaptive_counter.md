## 用法

来源：

- [官方控制文件](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/adaptive/adaptive_counter.control)
- [官方 adaptive estimator README](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/adaptive/README.md)

`adaptive_counter` 提供用于自适应采样近似去重计数的 `adaptive_estimator` 类型、函数与聚合。它适合允许有界内存估算的场景；需要精确结果时，不能用它替代精确的 `count(DISTINCT ...)`。

### 核心流程

创建扩展，然后为聚合指定误差目标与预计最大基数：

```sql
CREATE EXTENSION adaptive_counter;

SELECT adaptive_distinct(i, 0.01, 100000)
FROM generate_series(1, 100000) AS s(i);
```

单参数重载 `adaptive_distinct(anyelement)` 使用 0.025 的误差目标和 1,000,000 的预计基数。如果这些默认值会分配超出工作负载所需的估算器，应显式传参。

### API

- `adaptive_init(error real, ndistinct int)` 创建 `adaptive_estimator`；`adaptive_size(error real, ndistinct int)` 返回对应大小。
- `adaptive_add_item(adaptive_estimator, anyelement)` 更新状态，`adaptive_get_estimate(adaptive_estimator)` 读取当前估算值。
- `adaptive_get_error(adaptive_estimator)`、`adaptive_get_ndistinct(adaptive_estimator)` 和 `adaptive_get_item_size(adaptive_estimator)` 暴露状态参数。
- `adaptive_merge(adaptive_estimator, adaptive_estimator)` 合并兼容状态；`adaptive_reset(adaptive_estimator)` 清空状态，`length(adaptive_estimator)` 返回其存储大小。

### 运维说明

1.3.3 版控制文件将扩展标记为可重定位，未声明前置扩展或预加载要求。上游仓库已经归档并只读。估算器值可能占用数 KB；在表中反复更新会按 MVCC 生成新行版本并可能造成膨胀。应采用可接受的最低精度和现实的基数上限，批量更新状态，并用代表性数据验证误差表现。
