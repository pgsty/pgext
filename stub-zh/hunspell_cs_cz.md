

## 用法

> [hunspell_cs_cz: PostgreSQL 的捷克语 Hunspell 词典](https://github.com/postgrespro/hunspell_dicts)

用于 PostgreSQL 全文搜索的捷克语 Hunspell 词典和文本搜索配置。

```sql
CREATE EXTENSION hunspell_cs_cz;

SELECT ts_lexize('czech_hunspell', 'příběhy');

SELECT to_tsvector('czech_hunspell', 'příběhy');
```

该扩展提供 `czech_hunspell` 词典和文本搜索配置。
