## 用法

来源：

- [目录版本对应的官方 README](https://github.com/jaiminpan/pg_scws/blob/338d1ebe911372165acad331264bbf48afa56a9e/README.md)
- [目录版本对应的扩展 SQL](https://github.com/jaiminpan/pg_scws/blob/338d1ebe911372165acad331264bbf48afa56a9e/pg_scws--1.0.sql)
- [目录版本对应的解析器实现](https://github.com/jaiminpan/pg_scws/blob/338d1ebe911372165acad331264bbf48afa56a9e/pg_scws.c)

`pg_scws` 1.0 将 SCWS 中文分词器集成为 PostgreSQL 全文检索解析器。它安装 scws 解析器和 scwscfg 文本搜索配置，并把若干词法类别映射到 simple 词典。

### 核心流程

```sql
CREATE EXTENSION pg_scws;

SELECT to_tsvector(
    'scwscfg',
    '小明硕士毕业于中国科学院计算所，后在日本京都大学深造'
);

CREATE INDEX article_search_idx
ON article
USING gin (to_tsvector('scwscfg', body));

SELECT id
FROM article
WHERE to_tsvector('scwscfg', body)
      @@ to_tsquery('scwscfg', '计算 & 专家');
```

索引表达式和所有匹配查询必须使用同一文本搜索配置。

### 对象与设置

- 扩展创建 scws 文本搜索解析器和 scwscfg 配置。
- 设置项控制词典常驻内存、字符集、规则文件、附加词典、标点处理、二元分词和复合分词模式。
- 词典和规则文件从 PostgreSQL 文本搜索数据目录解析；所有可能执行查询的服务器都需要相同文件与配置。

```sql
SHOW scws.dict_in_memory;
SHOW scws.charset;
SHOW scws.rules;
SHOW scws.extra_dicts;
SHOW scws.punctuation_ignore;
SHOW scws.seg_with_duality;
SHOW scws.multi_mode;
```

### 兼容性与运维

- 上游只报告在 PostgreSQL 9.4 上测试，并仅预期后续 9.x 可以工作。解析器使用可能跨大版本变化的服务器 API，必须在每个确切目标大版本上构建和测试。
- 源码捆绑 SCWS、默认 UTF-8 词典和规则。重建受影响索引前，应把自定义词典原子地部署到主库、副本和故障转移候选节点。
- 分词或词典变化会改变 lexeme。计划内修改规则后应 REINDEX 表达式索引，并用有代表性的中文文本验证召回率和准确率。
