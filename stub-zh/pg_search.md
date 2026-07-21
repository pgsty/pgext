## 用法

来源：

- [pg_search v0.24.3 README](https://github.com/paradedb/paradedb/blob/v0.24.3/pg_search/README.md)
- [pg_search v0.24.3 release notes](https://github.com/paradedb/paradedb/releases/tag/v0.24.3)
- [ParadeDB 文档](https://docs.paradedb.com/)

pg_search 为 PostgreSQL 添加了 ParadeDB 的 BM25 访问方法和查询操作符，用于排名全文搜索、结构化搜索以及混合搜索。当搜索必须与 PostgreSQL 数据保持事务一致性，并支持 BM25 评分、高亮显示、过滤器、聚合和连接时，请使用此扩展。

### 安装扩展

    CREATE EXTENSION pg_search;

上游 v0.24.3 支持 PostgreSQL 15 及以上版本。请使用与具体 PostgreSQL 主版本匹配的构建包。由于该扩展参与了规划器和执行器路径，因此在生产升级前，请先测试查询计划和资源使用情况。

### 创建 BM25 索引

创建一个以稳定唯一键字段为基础的 BM25 索引：

    CREATE INDEX products_search_idx
    ON products
    USING bm25 (
      id,
      title,
      description,
      category,
      rating
    )
    WITH (key_field = 'id');

保持键字段的唯一性和非空性。仅索引用于搜索或过滤的字段；每个被索引的字段都会增加构建时间、磁盘使用和写入放大。

### 查询、排名与高亮

@@@ 操作符匹配字段或已索引行与 ParadeDB 查询：

    SELECT id,
           title,
           paradedb.score(id) AS score,
           paradedb.snippet(description) AS snippet
    FROM products
    WHERE description @@@ 'wireless keyboard'
      AND category = 'electronics'
    ORDER BY score DESC
    LIMIT 20;

当用户输入必须受到限制时，请使用字段限定的查询字符串或 ParadeDB 查询构造器。未经验证的情况下，切勿将不可信输入直接拼接到查询语法中。

对于布尔查询，paradedb.boolean() 可以组合 must、should 和 must_not 子句，并可以设置 minimum_should_match。该扩展还暴露了 index_created_at() 函数用于检查索引创建时间。

### 用户可接触的对象索引

- bm25: 用于文本和结构化字段的 BM25 索引访问方法。
- @@@: 在 WHERE 子句中使用的搜索匹配操作符。
- paradedb.score(key): 匹配行的 BM25 相关性评分。
- paradedb.snippet(field): 当前匹配的高亮摘录。
- paradedb.parse(...), paradedb.term(...), paradedb.boolean(...): 带类型查询构造器。
- paradedb.index_info(...): 索引元数据和字段配置信息。
- paradedb.index_created_at(...): 索引创建时间戳。

### 0.24.3 版本的操作变化

0.24.x 系列启用了更多聚合和连接扫描路径，并增加了 time 和 timetz 支持。版本 0.24.3 还限制了顺序扫描缓冲区，设置了索引构建工作内存上限，更早地检查可用磁盘空间，修复了 GROUP BY 分组基数路由问题，并在 Tantivy 会截断值时抛出错误。

这些防护措施减少了资源使用失控的风险，但并未消除容量规划的需求。请监控临时空间、索引大小、构建时间和查询内存。升级后重新运行代表性的 EXPLAIN ANALYZE 计划，因为计划器行为可能会发生变化。

### 兼容性与注意事项

- pg_search 使用其自己的 BM25 索引实现。不要假设由其他扩展创建的索引可以互换。
- 本地目录元数据报告了 bm25 访问方法与 pg_textsearch 和 vchord_bm25 的冲突；除非文档明确支持共存，否则请避免在同一数据库中加载竞争性实现。
- 搜索索引必须与表一起维护，并可能显著增加更新成本。
- 排名依赖于查询和语料库。请使用生产环境相似的文本和过滤器进行基准测试，而不是将示例分数视为可移植的。
