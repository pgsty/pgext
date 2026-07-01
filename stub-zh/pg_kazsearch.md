## 用法

来源：[README](https://github.com/darkhanakh/pg-kazsearch/blob/v2.2.0/README.md)、[v2.2.0 release](https://github.com/darkhanakh/pg-kazsearch/releases/tag/v2.2.0)、[v2.1.0 release](https://github.com/darkhanakh/pg-kazsearch/releases/tag/v2.1.0)

`pg_kazsearch` 是面向哈萨克语的 PostgreSQL full-text search 扩展。README 说明它支持 PostgreSQL 16-18，并创建可直接使用的 text search configuration `kazakh_cfg` 和 dictionary `pg_kazsearch_dict`。版本 2.2.0 在核心 stemmer 中增加拉丁字母哈萨克语支持；成功识别的拉丁与西里尔输入会收敛到 canonical Cyrillic stems。

### 快速开始

```sql
CREATE EXTENSION pg_kazsearch;

SELECT ts_lexize('pg_kazsearch_dict', 'алмаларымыздағы');
-- {алма}

SELECT to_tsvector('kazakh_cfg', 'президенттің жарлығы');
-- 'жарлық':2 'президент':1
```

### 将哈萨克语 FTS 添加到表

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

README 记录了无需重启的运行时 dictionary 调优：

```sql
ALTER TEXT SEARCH DICTIONARY pg_kazsearch_dict
  (w_deriv = 3.5, w_short_char = 100.0);
```

通过 `script_mode` 控制拉丁字母处理；默认 `auto` 模式会检测受支持的现代哈萨克拉丁正字法，并归一化为西里尔输出：

```sql
ALTER TEXT SEARCH DICTIONARY pg_kazsearch_dict
  (script_mode = cyrillic_only);
```

### Release 与打包说明

- 本项目 CSV 跟踪 package/source version `2.2.0`、extension control version `0.1.0`、`pgrx` `0.18.1`，以及 PostgreSQL versions 16-18。
- 上游 release `v2.2.0` 为核心 stemmer 增加拉丁字母哈萨克语支持。
- 上游 release `v2.0.0` 引入了当前 Rust / `pgrx` PostgreSQL extension 打包。
- 上游 release `v2.1.0` 在 PostgreSQL extension 之外增加了 Elasticsearch plugin；README 中的 PostgreSQL SQL 用法未改变。

### 注意事项

面向 PostgreSQL 的文档很简洁，重点是 stemming 和 FTS 用法。此 stub 避免推断除 `kazakh_cfg`、`pg_kazsearch_dict` 以及上面文档化示例以外的 SQL 对象。
