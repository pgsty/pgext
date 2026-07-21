## 用法

来源：

- [官方 v2.3.0 README](https://github.com/darkhanakh/pg-kazsearch/blob/v2.3.0/README.md)
- [v2.3.0 发行版](https://github.com/darkhanakh/pg-kazsearch/releases/tag/v2.3.0)
- [PostgreSQL 扩展控制文件](https://github.com/darkhanakh/pg-kazsearch/blob/v2.3.0/pg_ext/pg_kazsearch.control)
- [从 v2.2.0 到 v2.3.0 的升级 SQL](https://github.com/darkhanakh/pg-kazsearch/blob/v2.3.0/pg_ext/sql/pg_kazsearch--2.2.0--2.3.0.sql)

`pg_kazsearch` 为 PostgreSQL 16 至 18 提供了哈萨克语全文词干提取。它安装了一个现成的 `kazakh_cfg` 配置和 `pg_kazsearch_dict` 字典。使用支持的现代拉丁字母书写法的西里尔文和哈萨克语会转换为标准的西里尔文词干，从而使文档和查询能够在不同书写系统之间匹配。

### 核心工作流程

```sql
CREATE EXTENSION pg_kazsearch;

SELECT ts_lexize('pg_kazsearch_dict', 'алмаларымыздағы');
-- {алма}

SELECT to_tsvector('kazakh_cfg', 'мектептеріміздегі оқушылардың');
-- 'мектеп':1 'оқушы':2
```

添加加权存储向量和 GIN 索引：

```sql
ALTER TABLE articles ADD COLUMN fts tsvector
GENERATED ALWAYS AS (
    setweight(to_tsvector('kazakh_cfg', title), 'A') ||
    setweight(to_tsvector('kazakh_cfg', body), 'B')
) STORED;

CREATE INDEX articles_fts_idx ON articles USING GIN (fts);

SELECT title
FROM articles
WHERE fts @@ websearch_to_tsquery('kazakh_cfg', 'президенттің жарлығы')
ORDER BY ts_rank_cd(
    fts,
    websearch_to_tsquery('kazakh_cfg', 'президенттің жарлығы')
) DESC;
```

### 字典调优

可以在运行时更改惩罚权重：

```sql
ALTER TEXT SEARCH DICTIONARY pg_kazsearch_dict
    (w_deriv = 3.5, w_short_char = 100.0);
```

默认的 `script_mode = auto` 会检测支持的现代哈萨克语拉丁字母书写法并返回西里尔文词干。当需要严格的西里尔文行为时，可以禁用拉丁字母处理：

```sql
ALTER TEXT SEARCH DICTIONARY pg_kazsearch_dict
    (script_mode = cyrillic_only);
```

### 升级和搜索注意事项

- 在升级到 `2.3.0` 后，重新计算存储的 `tsvector` 列或重新填充触发器维护的向量，然后对表执行 `VACUUM (ANALYZE)`。

```sql
ALTER EXTENSION pg_kazsearch UPDATE;
UPDATE articles SET title = title;
VACUUM (ANALYZE) articles;
```

- 升级前长时间运行的会话应重新连接以加载新的字典。
- 拉丁字母支持针对现代书写法。混合书写系统的输入、过时的撇号/重音符/双字符拼写以及低置信度的 ASCII 词素可能保持不变。
- `websearch_to_tsquery` 使用严格的 AND 语义处理普通词项。需要更广泛召回的应用程序应故意实现并衡量一个备选查询，而不是默默地将所有搜索都改为 OR 操作。
