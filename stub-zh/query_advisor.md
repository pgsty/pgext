## 用法

来源：

- [EDB Query Advisor 官方文档](https://www.enterprisedb.com/docs/pg_extensions/query_advisor/)
- [官方配置指南](https://www.enterprisedb.com/docs/pg_extensions/query_advisor/configuring/)
- [官方使用指南](https://www.enterprisedb.com/docs/pg_extensions/query_advisor/using/)

`query_advisor` 对 workload predicate 与估算误差采样，无需预先创建全部候选对象，即可推荐有用索引与多列 extended statistic。

### 启用

产品软件包名为 EDB Query Advisor，但 canonical extension 与 preload identity 是 `query_advisor`。安装 1.2.2、预加载、重启并创建扩展：

```ini
shared_preload_libraries = 'query_advisor'
```

```sql
CREATE EXTENSION query_advisor;
```

使用 `query_advisor.sample_rate` 与 `query_advisor.exclude_schema_list` 限定采集范围。`query_advisor.max_qual_entries` 等容量参数需要重启。

### 生成建议

```sql
SELECT *
FROM query_advisor_index_recommendations();

SELECT *
FROM query_advisor_statistics_recommendations();

SELECT *
FROM query_advisor_qualstats_pretty;
```

Index recommendation 会结合 hypothetical index 对已采集 workload query 进行成本评估。Statistic recommendation 当前聚焦两列组合，并给出 weight 与受益 query ID。

### 审查边界

采集数据保存在内存中，服务器重启后丢失。建议只是候选而非自动 DDL：执行生成的 `CREATE INDEX` 或 `CREATE STATISTICS` command 前，必须验证写放大、存储、维护、冗余索引与生产执行计划变化。

