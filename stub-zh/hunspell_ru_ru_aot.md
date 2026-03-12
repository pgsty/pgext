

## 用法

> [hunspell_ru_ru_aot: PostgreSQL 的俄语（AOT）Hunspell 词典](https://github.com/postgrespro/hunspell_dicts)

用于 PostgreSQL 全文搜索的俄语（AOT 变体）Hunspell 词典和文本搜索配置。

```sql
CREATE EXTENSION hunspell_ru_ru_aot;

SELECT ts_lexize('russian_aot_hunspell', 'рассказы');

SELECT to_tsvector('russian_aot_hunspell', 'рассказы');
```

该扩展提供 `russian_aot_hunspell` 词典和文本搜索配置。
