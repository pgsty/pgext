## 用法

来源：

- [官方 trimmed_aggregates README](https://github.com/tvondra/trimmed_aggregates/blob/bc8c1918f5d21885378cabbd070cb1b8eea02e5f/README.md)
- [2.0.0-dev 版扩展 SQL](https://github.com/tvondra/trimmed_aggregates/blob/bc8c1918f5d21885378cabbd070cb1b8eea02e5f/sql/trimmed_aggregates--2.0.0-dev.sql)
- [2.0.0-dev 版实现](https://github.com/tvondra/trimmed_aggregates/blob/bc8c1918f5d21885378cabbd070cb1b8eea02e5f/trimmed_aggregates.c)

`trimmed_aggregates` 会对分组排序，丢弃低端和高端的可配置比例后计算统计量。目录中的 `2.0.0-dev` SQL 使用常见聚合名称的三参数重载；README 中较旧的 `avg_trimmed` 示例并不代表这个预览版 API。

### 计算截尾统计量

参数依次是输入值、低端裁剪比例和高端裁剪比例：

```sql
CREATE EXTENSION trimmed_aggregates;

SELECT avg(v, 0.10, 0.10)       AS trimmed_avg,
       stddev_samp(v, 0.10, 0.10) AS trimmed_stddev
FROM generate_series(1, 1000) AS g(v);

SELECT trimmed(v, 0.20, 0.10) AS all_statistics
FROM generate_series(1, 1000) AS g(v);
```

`2.0.0-dev` 版定义了 `avg`、`var`、`var_pop`、`var_samp`、`stddev`、`stddev_pop` 和 `stddev_samp`，签名均为 `(value, low_cut, high_cut)`。输入值可以是 `double precision`、`int`、`bigint` 或 `numeric`。

`trimmed(value, low_cut, high_cut)` 返回包含七个值的数组，顺序为：平均值、总体方差、样本方差、方差、总体标准差、样本标准差、标准差。在此实现中，`variance` 等同于样本方差结果，`stddev` 等同于样本标准差结果。

### 裁剪规则与资源消耗

两个裁剪比例都不能为 null，且必须满足 `0 <= cut < 1`，两者之和必须小于 `1`。同一分组内应为两个比例传入常量。转换状态会在首次创建分组状态时记录比例，并没有为逐行变化的比例表达式定义语义。空输入值会被跳过。

所有变体都会为每个分组保留并排序全部非空输入。版本 2 增加了合并、序列化和反序列化函数，并把聚合标记为并行安全；但并行执行只会分散状态，不能消除总体内存与排序成本。高基数分组、倾斜的大分组或大量并发查询都可能耗尽内存。

这是开发版本 API，其 SQL 与 README 的稳定版示例存在实质差异。使用时应固定精确修订版，在重载解析可能歧义时限定调用，并在上线前回归测试空分组、小保留样本、数值精度、并行计划和最坏分组规模。
