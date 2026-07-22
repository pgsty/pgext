## 用法

来源：

- [官方控制文件](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/pcsa/pcsa_counter.control)
- [官方 PCSA README](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/pcsa/README.md)

`pcsa_counter` 将带随机平均的概率计数实现为 `pcsa_estimator` 类型与聚合。它通过可配置的位图数量和键大小提供近似去重计数。

### 核心流程

```sql
CREATE EXTENSION pcsa_counter;

SELECT pcsa_distinct(i, 32, 3)
FROM generate_series(1, 100000) AS s(i);
```

参数分别选择位图数量和键大小。单参数重载 `pcsa_distinct(anyelement)` 默认使用 64 个位图和键大小 4；需要更小状态或不同准确度时应显式调优。

### API

- `pcsa_init(int, int)` 创建 `pcsa_estimator`；`pcsa_size(int, int)` 返回状态大小。
- `pcsa_add_item(pcsa_estimator, anyelement)` 更新状态，`pcsa_get_estimate(pcsa_estimator)` 读取估算值。
- `pcsa_reset(pcsa_estimator)` 清空状态，`length(pcsa_estimator)` 返回其存储大小。
- `pcsa_distinct(anyelement, int, int)` 是可配置聚合；`pcsa_distinct(anyelement)` 使用文档默认值。

### 运维说明

1.3.3 版控制文件可重定位，未声明前置扩展或预加载要求。上游仓库已经归档并只读。更大的配置会消耗更多内存，反复更新持久化估算器可能造成 MVCC 膨胀。应在代表性基数上评测所选参数，不要把估算值当作精确结果。
