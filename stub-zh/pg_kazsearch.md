## 用法

来源：[README](https://github.com/darkhanakh/pg-kazsearch/blob/main/README.md)，[releases](https://github.com/darkhanakh/pg-kazsearch/releases)

`pg_kazsearch` 是一个面向哈萨克语的 PostgreSQL 全文检索扩展。README 说明它会创建可直接使用的文本搜索配置 `kazakh_cfg` 和词典 `pg_kazsearch_dict`。

### 快速开始

```sql
CREATE EXTENSION pg_kazsearch;

SELECT ts_lexize('pg_kazsearch_dict', 'алмаларымыздағы');
-- {алма}

SELECT to_tsvector('kazakh_cfg', 'президенттің жарлығы');
-- 'жарлық':2 'президент':1
```

### 为表添加哈萨克语 FTS

```sql
ALTER TABLE articles ADD COLUMN fts tsvector
    GENERATED ALWAYS AS (
        setweight(to_tsvector('kazakh_cfg', title), 'A') ||
        setweight(to_tsvector('kazakh_cfg', body), 'B')
    ) STORED;

CREATE INDEX idx_fts ON articles USING GIN (fts);

SELECT title
FROM articles
WHERE fts @@ websearch_to_tsquery('kazakh_cfg', 'президенттің жарлығы')
ORDER BY ts_rank_cd(fts, websearch_to_tsquery('kazakh_cfg', 'президенттің жарлығы')) DESC
LIMIT 10;
```

### 调优

README 说明词典参数可以在运行时调整，无需重启：

```sql
ALTER TEXT SEARCH DICTIONARY pg_kazsearch_dict
  (w_deriv = 3.5, w_short_char = 100.0);
```

### 发布与打包说明

- 上游 `v2.0.0` 引入了当前基于 Rust / `pgrx` 的架构。
- 上游 `v2.1.0` 在 PostgreSQL 扩展之外新增了 Elasticsearch 插件，但 README 中的 PostgreSQL SQL 用法没有变化。
- 仓库 README 发布 Debian `2.x` 软件包，而本项目的 CSV 说明会单独跟踪 extension control version。

### 注意事项

面向 PostgreSQL 的文档目前较简洁，重点只覆盖词干提取与全文检索用法。这里不要推断 README 未明确列出的额外 SQL 对象，保守限定在 `kazakh_cfg`、`pg_kazsearch_dict` 和上面给出的示例。
