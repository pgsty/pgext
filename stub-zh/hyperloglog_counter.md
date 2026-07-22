## 用法

来源：

- [官方控制文件](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/hyperloglog/hyperloglog_counter.control)
- [官方 HyperLogLog README](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/hyperloglog/README.md)

`hyperloglog_counter` 使用 HyperLogLog 算法提供近似去重计数。它包含 `hyperloglog_estimator` 状态类型以及聚合和状态管理函数；适合有内存上限的基数估算，不适合精确计数。

### 核心流程

```sql
CREATE EXTENSION hyperloglog_counter;

SELECT hyperloglog_distinct(i, 0.01)
FROM generate_series(1, 100000) AS s(i);
```

单参数重载 `hyperloglog_distinct(anyelement)` 使用 2% 的默认误差率。应显式指定误差率以控制内存与准确度之间的权衡。

### API

- `hyperloglog_init(real)` 创建 `hyperloglog_estimator`；`hyperloglog_size(real)` 返回指定误差率所需的大小。
- `hyperloglog_add_item(hyperloglog_estimator, anyelement)` 更新状态，`hyperloglog_get_estimate(hyperloglog_estimator)` 读取估算值。
- `hyperloglog_reset(hyperloglog_estimator)` 清空状态，`length(hyperloglog_estimator)` 返回其存储大小。
- `hyperloglog_distinct(anyelement, real)` 是可配置聚合；`hyperloglog_distinct(anyelement)` 使用默认误差率。

### 运维说明

1.2.6 版控制文件可重定位，未声明前置扩展或预加载要求。上游仓库已经归档并只读。持久化估算器可能占用数 KB，反复更新会造成 MVCC 膨胀。应选择内存占用最低且可接受的误差率，并在工作负载的真实基数分布上验证估算值。
