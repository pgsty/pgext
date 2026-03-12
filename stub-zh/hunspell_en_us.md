

## 用法

> [hunspell_en_us: PostgreSQL 的英语（美国） Hunspell 词典](https://github.com/postgrespro/hunspell_dicts)

用于 PostgreSQL 全文搜索的英语（美国） Hunspell 词典和文本搜索配置。

```sql
CREATE EXTENSION hunspell_en_us;

SELECT ts_lexize('english_hunspell', 'stories');

SELECT to_tsvector('english_hunspell', 'stories');
```

该扩展提供 `english_hunspell` 词典和文本搜索配置。
