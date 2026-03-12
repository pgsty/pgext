

## 用法

> [hunspell_fr: PostgreSQL 的法语 Hunspell 词典](https://github.com/postgrespro/hunspell_dicts)

用于 PostgreSQL 全文搜索的法语 Hunspell 词典和文本搜索配置。

```sql
CREATE EXTENSION hunspell_fr;

SELECT ts_lexize('french_hunspell', 'histoires');

SELECT to_tsvector('french_hunspell', 'histoires');
```

该扩展提供 `french_hunspell` 词典和文本搜索配置。
