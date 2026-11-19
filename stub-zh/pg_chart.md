## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_chart/pg_chart.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_chart/pg_chart--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_chart/meson.build)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_chart/pg_chart.c)

`pg_chart` `1.0` 提供从 SQL 查询结果生成 SVG 图表。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_chart;
SELECT chart_bar('SELECT ''east'' AS region, 12 AS amount', 'Orders');
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

chart_bar、chart_line、chart_pie、chart_scatter 和 chart_area 返回 SVG 文本。查询应提供相应的标签、数值列。查询通过 SPI 以调用者权限执行；应限制查询文本的提交者，并评估大结果集对后端内存和运行时间的影响。

此模块对应上游旧名 `alohadb_chart`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
