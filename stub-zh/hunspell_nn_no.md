

## 用法

> [hunspell_nn_no: PostgreSQL 的挪威语（尼诺斯克） Hunspell 词典](https://github.com/postgrespro/hunspell_dicts)

用于 PostgreSQL 全文搜索的挪威语（尼诺斯克） Hunspell 词典和文本搜索配置。

```sql
CREATE EXTENSION hunspell_nn_no;

SELECT ts_lexize('norwegian_hunspell', 'historier');

SELECT to_tsvector('norwegian_hunspell', 'historier');
```

该扩展提供 `norwegian_hunspell` 词典和文本搜索配置。
