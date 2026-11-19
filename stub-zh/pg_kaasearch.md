## 用法

来源：

- [README.md](https://github.com/erinshek/pg-kaasearch/blob/v1.0.0/README.md)
- [pg_kaasearch.control](https://github.com/erinshek/pg-kaasearch/blob/v1.0.0/pg_kaasearch.control)
- [pg_kaasearch--1.0.sql](https://github.com/erinshek/pg-kaasearch/blob/v1.0.0/pg_kaasearch--1.0.sql)
- [LICENSE](https://github.com/erinshek/pg-kaasearch/blob/v1.0.0/LICENSE)

`pg_kaasearch` 提供卡拉卡尔帕克语全文检索，包括词干提取、停用词以及拉丁字母、西里尔字母和撇号拼写的规范化。发行版 1.0.0 安装 SQL 版本 1.0，要求 PostgreSQL 15–18 和 UTF8 数据库。

### 启用与检索

安装需要具备相应权限的扩展管理员；控制文件未将扩展标为 trusted。扩展可迁移，无需预加载。

```sql
CREATE EXTENSION pg_kaasearch;
SELECT to_tsvector('kaasearch', 'Мектепке барған балалар китап оқыйды.')
    @@ to_tsquery('kaasearch', 'mektep & kitap');
SELECT kaa_stem('mektebimizde');
```

表达式索引与查询应使用相同的全文检索配置：

```sql
CREATE TABLE articles (id bigint PRIMARY KEY, title text, body text);
CREATE INDEX articles_fts ON articles USING gin (to_tsvector('kaasearch', body));
SELECT title FROM articles
WHERE to_tsvector('kaasearch', body)
   @@ websearch_to_tsquery('kaasearch', 'mektep kitap');
```

### 对象与词典

`kaasearch` 是全文检索配置；`kaasearch_parser`、`kaasearch_template` 和 `kaasearch_dict` 提供解析器和词典。`kaa_stem` 返回词干，`kaa_normalize` 规范化拼写，`kaa_analyze` 返回规范化词、词干、分段、类别与来源。标准 `ts_rank` 和 `ts_headline` 可配合该配置使用。

停用词来自安装的 `kaasearch.stop` 文件。可通过 `ALTER TEXT SEARCH DICTIONARY kaasearch_dict (Lexicon = name)` 配置自定义词库，服务端必须存在对应的词典文件。更改分词或词典行为后，应重建受影响的全文索引。在搜索路径之外安装时，应为配置名加模式限定。
