

## 用法

> [hunspell_ne_np: PostgreSQL 的尼泊尔语 Hunspell 词典](https://github.com/postgrespro/hunspell_dicts)

用于 PostgreSQL 全文搜索的尼泊尔语 Hunspell 词典和文本搜索配置。

```sql
CREATE EXTENSION hunspell_ne_np;

SELECT ts_lexize('nepali_hunspell', 'कथाहरू');

SELECT to_tsvector('nepali_hunspell', 'कथाहरू');
```

该扩展提供 `nepali_hunspell` 词典和文本搜索配置。
