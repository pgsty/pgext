## 用法

来源：

- [官方控制文件](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/loglog/loglog_counter.control)
- [官方 LogLog README](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/loglog/README.md)

`loglog_counter` 将 LogLog 基数估算器实现为 `loglog_estimator` 类型、状态函数和聚合。它适用于固定内存比精确性更重要的近似去重计数。

### 核心流程

```sql
CREATE EXTENSION loglog_counter;

SELECT loglog_distinct(i, 0.01)
FROM generate_series(1, 100000) AS s(i);
```

单参数重载 `loglog_distinct(anyelement)` 使用 2.5% 的默认误差率。显式重载允许查询选择不同的误差目标。

### API

- `loglog_init(real)` 创建 `loglog_estimator`，`loglog_size(real)` 返回对应大小。
- `loglog_add_item(loglog_estimator, anyelement)` 更新状态；`loglog_get_estimate(loglog_estimator)` 读取结果。
- `loglog_reset(loglog_estimator)` 清空状态，`length(loglog_estimator)` 返回其存储大小。
- `loglog_distinct(anyelement, real)` 执行可配置聚合；`loglog_distinct(anyelement)` 使用默认值。

### 运维说明

1.2.4 版控制文件可重定位，未声明前置扩展或预加载要求。上游仓库已经归档并只读。估算器状态可能占用数 KB；频繁更新持久化状态会产生 MVCC 行版本并造成膨胀。应采用可接受的最低精度、批量更新，并在依赖估算值前测量实际误差。
