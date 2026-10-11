## 用法

来源：

- [gpdb-doc/markdown/ref_guide/modules/gp_array_agg.html.md](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpdb-doc/markdown/ref_guide/modules/gp_array_agg.html.md)
- [gpcontrib/gp_array_agg/gp_array_agg--1.0.0.sql](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpcontrib/gp_array_agg/gp_array_agg--1.0.0.sql)
- [gpcontrib/gp_array_agg/Makefile](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpcontrib/gp_array_agg/Makefile)
- [LICENSE](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/LICENSE)
- [gpcontrib/gp_array_agg/gp_array_agg.control](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpcontrib/gp_array_agg/gp_array_agg.control)

`gp_array_agg` 提供可合并 Greenplum 各工作节点部分结果的数组聚合函数。这里核验的 1.0.0 模块属于 OpenGPDB 维护的 Greenplum 6 发行版。

### 核心用法

```sql
CREATE EXTENSION gp_array_agg;
SELECT gp_array_agg(v ORDER BY v)
FROM (VALUES (3), (1), (NULL::integer), (2)) AS input(v);
```

### 运行边界

`gp_array_agg(anyelement)` 将输入值收集为数组，也保留空值；没有输入行时返回空值。除非在聚合内部指定 `ORDER BY`，否则元素顺序不确定。配套的 `array_agg_combine`、`array_agg_serialize` 和 `array_agg_deserialize` 服务于聚合实现，不是独立的应用工作流。

在每个需要使用它的数据库中，由超级用户注册发行版自带的模块。C 库按需加载，无需预加载配置，并支持安装到指定模式。本条描述 Greenplum 内核模块，不代表支持某个原生 PostgreSQL 大版本，也不代表已有 Pigsty 软件包。
