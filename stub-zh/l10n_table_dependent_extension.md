

## 用法

> [l10n_table_dependent_extension: pg_xenophile 的本地化表依赖扩展](https://github.com/bigsmoke/pg_xenophile)

`l10n_table_dependent_extension` 是 `pg_xenophile` 的配套扩展，为依赖其本地化 (l10n) 表机制的其他扩展提供基础设施。

```sql
CREATE EXTENSION l10n_table_dependent_extension;
```

该扩展与 `pg_xenophile` 的 `l10n_table` 系统协同工作，允许依赖扩展将其表注册到本地化框架中。当基表在 `xeno.l10n_table` 中注册后，系统会自动创建配套的翻译表和特定语言的视图。

完整的本地化表管理 API 请参阅 [`pg_xenophile`](https://github.com/bigsmoke/pg_xenophile) 文档。
