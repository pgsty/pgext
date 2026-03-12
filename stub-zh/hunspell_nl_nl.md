

## 用法

> [hunspell_nl_nl: PostgreSQL 的荷兰语 Hunspell 词典](https://github.com/postgrespro/hunspell_dicts)

用于 PostgreSQL 全文搜索的荷兰语 Hunspell 词典和文本搜索配置。

```sql
CREATE EXTENSION hunspell_nl_nl;

SELECT ts_lexize('dutch_hunspell', 'verhalen');

SELECT to_tsvector('dutch_hunspell', 'verhalen');
```

该扩展提供 `dutch_hunspell` 词典和文本搜索配置。
