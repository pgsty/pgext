## 用法

- 来源: [Codeberg 仓库](https://codeberg.org/kop/pg_isok), [文档首页](https://codeberg.org/kop/pg_isok/src/branch/main/doc_src/index.html), [文档源文件](https://codeberg.org/kop/pg_isok/src/branch/main/doc_src/isok.xml)
- Isok 是面向 PostgreSQL 的查询中心型监控扩展。它关注的是相较于先前已见过的可疑数据模式的变化，而不仅仅是某些行是否存在。

```sql
CREATE SCHEMA isok;
CREATE EXTENSION pg_isok SCHEMA isok;
```

## 核心流程

该扩展围绕两张表展开：

- `ISOK_QUERIES` 存储监控查询
- `ISOK_RESULTS` 存储发现的问题及其处理状态

通过 `run_isok_queries()` 运行监控：

```sql
SELECT * FROM run_isok_queries();
SELECT * FROM run_isok_queries($$VALUES ('new_countries')$$) AS problems;
```

`ISOK_RESULTS` 中的行可以标记为已解决或延后，因此后续运行不再将其报告为活动问题。

## 备注

文档将 Isok 描述为一种“软触发器”风格的数据清理与完整性审查工具。它可在 PostgreSQL 10 及以上版本安装，并且在受管环境中可按纯 SQL 方式构建。
