## 用法

来源：

- [PGXN biscuit 2.4.3 分发](https://pgxn.org/dist/biscuit/2.4.3/)
- [PGXN 2.4.3 元数据](https://api.pgxn.org/dist/biscuit/2.4.3/META.json)
- [PGXN 2.4.3 控制文件](https://api.pgxn.org/src/biscuit/biscuit-2.4.3/biscuit.control)
- [PGXN 2.4.3 改变日志](https://api.pgxn.org/src/biscuit/biscuit-2.4.3/CHANGELOG.md)
- [官方文档](https://biscuit.readthedocs.io/)

`biscuit` 是一个针对文本模式过滤优化的实验性 PostgreSQL 索引访问方法。它主要针对 `LIKE`, `ILIKE`, `NOT LIKE`, 和 `NOT ILIKE` 预测子，包括多列和表达式索引，同时以增加内存使用和写入工作为代价换取更快的过滤速度。

### 核心流程

```sql
CREATE EXTENSION biscuit;

CREATE INDEX message_body_biscuit_idx
ON message USING biscuit (body);

SELECT id, body
FROM message
WHERE body ILIKE '%timeout%';
```

当查询中使用相同的表达式时，表达式索引可以生效：

```sql
CREATE INDEX customer_email_biscuit_idx
ON customer USING biscuit (lower(email));

SELECT *
FROM customer
WHERE lower(email) LIKE '%@example.com';
```

对于跨越多个索引文本列的预测子，应使用多列索引，并通过在代表性数据上运行 `EXPLAIN (ANALYZE, BUFFERS)` 来确认选择的计划。

### 重要对象

- `biscuit` 是用于 `CREATE INDEX ... USING biscuit` 的索引访问方法。
- `biscuit_operators` 报告支持的操作符。
- `biscuit_version` 和 `biscuit_build_info` 暴露构建信息；`biscuit_build_info_json` 以 JSON 格式返回这些信息。
- `biscuit_status` 报告安装的构建和位图配置。
- `biscuit_index_stats` 和 `biscuit_index_memory_size` 检查索引及其内存占用情况。
- `biscuit_memory_usage` 是扩展内存使用情况视图。
- `biscuit_has_roaring` 和 `biscuit_roaring_version` 报告可选的 Roaring 位图支持。

### 限制和操作

`biscuit` 主要用于过滤，而不是有序索引扫描。它不提供正则表达式或相似性搜索功能。索引可能比 B-树更大且更昂贵；在生产使用前，请先基准测试读取选择性、摄入成本、内存使用情况以及真空行为。保留常规索引以满足排序、等值、唯一性或其他访问方法的需求。

上游项目将 Biscuit 标记为积极开发中。PGXN 发布 `2.4.3` 作为稳定分发，但该存档的改变日志仅止于 `2.4.2`，其元数据和控制文件暴露 SQL 扩展版本 `2.4.1`。将 `2.4.3` 视为分发/包刷新：未声称有额外的 SQL API 变更。材料中的 `2.4.2` 改变修复了索引缓存中的使用后释放漏洞以及编译器警告。
