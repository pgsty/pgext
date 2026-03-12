

## 用法

> [hunspell_de_de: PostgreSQL 的德语 Hunspell 词典](https://github.com/postgrespro/hunspell_dicts)

用于 PostgreSQL 全文搜索的德语 Hunspell 词典和文本搜索配置。

```sql
CREATE EXTENSION hunspell_de_de;

SELECT ts_lexize('german_hunspell', 'Geschichten');

SELECT to_tsvector('german_hunspell', 'Geschichten');
```

该扩展提供 `german_hunspell` 词典和文本搜索配置。
