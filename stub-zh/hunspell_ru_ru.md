

## 用法

> [hunspell_ru_ru: PostgreSQL 的俄语 Hunspell 词典](https://github.com/postgrespro/hunspell_dicts)

用于 PostgreSQL 全文搜索的俄语 Hunspell 词典和文本搜索配置。

```sql
CREATE EXTENSION hunspell_ru_ru;

SELECT ts_lexize('russian_hunspell', 'рассказы');

SELECT to_tsvector('russian_hunspell', 'рассказы');
```

该扩展提供 `russian_hunspell` 词典和文本搜索配置。
