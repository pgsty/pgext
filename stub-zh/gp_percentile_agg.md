## 用法

来源：

- [gpdb-doc/markdown/ref_guide/modules/gp_percentile_agg.html.md](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpdb-doc/markdown/ref_guide/modules/gp_percentile_agg.html.md)
- [gpcontrib/gp_percentile_agg/gp_percentile_agg--1.0.0.sql](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpcontrib/gp_percentile_agg/gp_percentile_agg--1.0.0.sql)
- [gpcontrib/gp_percentile_agg/Makefile](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpcontrib/gp_percentile_agg/Makefile)
- [LICENSE](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/LICENSE)
- [gpcontrib/gp_percentile_agg/gp_percentile_agg.control](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpcontrib/gp_percentile_agg/gp_percentile_agg.control)

`gp_percentile_agg` 提供 Greenplum GPORCA 优化百分位数与中位数查询所需的聚合函数。1.0.0 版属于 Greenplum 6 模块；应用通常继续使用内置的有序集聚合语法。

### 核心用法

```sql
CREATE EXTENSION gp_percentile_agg;
SET optimizer_enable_orderedagg = on;
CREATE TABLE measurements (value double precision);
INSERT INTO measurements VALUES (10), (20), (30);
SELECT percentile_cont(0.5) WITHIN GROUP (ORDER BY value)
FROM measurements;
```

### 运行边界

使用 GPORCA 时，在查询会话中启用 `optimizer_enable_orderedagg`。文档中的优化条件是：排序列不能是计算表达式，分位值必须是常量标量而非数组，查询不能含有 `GROUP BY`。这些是适用条件，不保证所有百分位数查询都会更快。

模块提供适用于浮点数、时间间隔、时间戳及带时区时间戳的 `gp_percentile_cont` 重载，以及多态的 `gp_percentile_disc`。这些内部辅助聚合要求有序输入和总行数与同值行组的计数；应用查询应使用普通的 `percentile_cont`、`percentile_disc` 或 `median` 接口。

在各数据库中由超级用户注册模块。C 库按需加载，无需预加载配置，扩展支持迁移模式。官方指南要求：从使用过该模块的 Greenplum 6.21.x 升级到 6.22.0 的修正版时，应删除并重新创建模块；执行维护前应检查依赖对象。这里不宣称兼容某个原生 PostgreSQL 大版本。
