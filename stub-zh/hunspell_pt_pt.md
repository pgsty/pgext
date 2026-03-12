

## 用法

> [hunspell_pt_pt: PostgreSQL 的葡萄牙语 Hunspell 词典](https://github.com/postgrespro/hunspell_dicts)

用于 PostgreSQL 全文搜索的葡萄牙语 Hunspell 词典和文本搜索配置。

```sql
CREATE EXTENSION hunspell_pt_pt;

SELECT ts_lexize('portuguese_hunspell', 'histórias');

SELECT to_tsvector('portuguese_hunspell', 'histórias');
```

该扩展提供 `portuguese_hunspell` 词典和文本搜索配置。
