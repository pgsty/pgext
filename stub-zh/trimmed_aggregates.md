## 用法

来源：

- [上游聚合文档](https://github.com/tvondra/trimmed_aggregates/blob/bc8c1918f5d21885378cabbd070cb1b8eea02e5f/README.md)
- [扩展控制文件](https://github.com/tvondra/trimmed_aggregates/blob/bc8c1918f5d21885378cabbd070cb1b8eea02e5f/trimmed_aggregates.control)
- [PGXN 分发页](https://pgxn.org/dist/trimmed_aggregates/)

`trimmed_aggregates` 在丢弃低端与高端的可配置比例后计算平均值、方差和标准差。组合聚合 `trimmed` 共享收集状态，一次返回全部七项统计值。

```sql
CREATE EXTENSION trimmed_aggregates;
SELECT avg_trimmed(v, 0.10, 0.10)
FROM generate_series(1, 1000) AS g(v);
SELECT trimmed(v, 0.20, 0.10)
FROM generate_series(1, 1000) AS g(v);
```

这些聚合会保留并排序完整输入集，因此内存与耗时随分组规模增长；大量并发或高基数组可能耗尽后端内存。应校验裁剪比例、NULL 和小分组行为，设置语句资源限制，并压测最坏分组。目录版本 `2.0.0-dev` 属于预览构建，使用前应固定版本并进行回归测试。
