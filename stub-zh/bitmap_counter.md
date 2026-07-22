## 用法

来源：

- [官方控制文件](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/bitmap/bitmap_counter.control)
- [官方 bitmap estimator README](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/bitmap/README.md)

`bitmap_counter` 使用自学习位图实现近似去重计数。它提供可在表或过程代码中保存的 `bitmap_estimator` 状态类型，以及直接查询用的 `bitmap_distinct` 聚合。

### 核心流程

```sql
CREATE EXTENSION bitmap_counter;

SELECT bitmap_distinct(i, 0.01, 100000)
FROM generate_series(1, 100000) AS s(i);
```

两个调优参数分别是目标误差和预计去重数量。单参数重载 `bitmap_distinct(anyelement)` 默认使用 0.025 和 1,000,000；预计集合更小或可接受更低精度时，应优先显式传参。

### API

- `bitmap_init(real, int)` 创建 `bitmap_estimator`，`bitmap_size(real, int)` 估算其大小。
- `bitmap_add_item(bitmap_estimator, anyelement)` 更新状态；`bitmap_get_estimate(bitmap_estimator)` 读取估算值。
- `bitmap_get_error(bitmap_estimator)` 和 `bitmap_get_ndistinct(bitmap_estimator)` 暴露已配置的状态属性。
- `bitmap_reset(bitmap_estimator)` 清空状态，`length(bitmap_estimator)` 返回其存储大小。

### 运维说明

1.3.5 版控制文件可重定位，未声明前置扩展或预加载要求。上游仓库已经归档并只读。持久化估算器状态可能占用数 KB，频繁更新行会放大 MVCC 膨胀。应保守配置状态大小、批量更新，并将结果视为估算值而不是精确基数。
