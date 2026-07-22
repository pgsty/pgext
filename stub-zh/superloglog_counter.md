## 用法

来源：

- [官方控制文件](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/superloglog/superloglog_counter.control)
- [官方 SuperLogLog README](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/superloglog/README.md)

`superloglog_counter` 将采用截断和限制规则的 LogLog 变体 SuperLogLog 实现为 `superloglog_estimator` 类型与聚合。它适用于有内存上限的近似去重计数。

### 核心流程

```sql
CREATE EXTENSION superloglog_counter;

SELECT superloglog_distinct(i, 0.01)
FROM generate_series(1, 100000) AS s(i);
```

单参数重载 `superloglog_distinct(anyelement)` 使用 2.5% 的默认误差率。应显式传入误差率以调节内存与准确度。

### API

- `superloglog_init(real)` 创建 `superloglog_estimator`，`superloglog_size(real)` 返回其大小。
- `superloglog_add_item(superloglog_estimator, anyelement)` 更新状态；`superloglog_get_estimate(superloglog_estimator)` 读取估算值。
- `superloglog_reset(superloglog_estimator)` 清空状态，`length(superloglog_estimator)` 返回其存储大小。
- `superloglog_distinct(anyelement, real)` 可配置；`superloglog_distinct(anyelement)` 使用默认误差率。

### 运维说明

1.2.3 版控制文件可重定位，未声明前置扩展或预加载要求。上游仓库已经归档并只读。估算器状态可能占用数 KB，反复更新表会造成 MVCC 膨胀。应选择可接受的最低精度、批量更新状态，并在投入运维前用真实数据验证估算值。
