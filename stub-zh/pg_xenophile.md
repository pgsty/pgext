

## 用法

> [pg_xenophile: 国际化 (i18n) 与本地化 (l10n) 工具集](https://github.com/bigsmoke/pg_xenophile)

`pg_xenophile` 扩展提供 i18n/l10n 基础设施，包括国家、语言和货币的参考数据，以及自动化的本地化表管理。

```sql
CREATE EXTENSION pg_xenophile CASCADE;
```

所有对象位于 `xeno` 模式中（不可重定位）。

### 参考表

- **`xeno.country`**：ISO 3166-1 代码，包含国际电话区号和货币信息
- **`xeno.lang`**：ISO 639-1 语言代码
- **`xeno.currency`**：ISO 4217 货币代码及符号
- **`xeno.country_subdivision`**：ISO 3166-2 行政区划代码
- **`xeno.eu_country`**：欧盟成员国追踪
- **`xeno.country_postal_code_pattern`**：邮政编码验证规则

### 本地化表

该扩展自动管理翻译表。向 `xeno.l10n_table` 插入记录即可注册一个可翻译的基表：

```sql
INSERT INTO xeno.l10n_table (base_table_schema, base_table_name)
VALUES ('public', 'products');
```

这会自动创建配套的 `products_l10n` 表和特定语言的视图。

### 便捷视图

- `xeno.country_l10n_en`：英文国家名称
- `xeno.lang_l10n_en`：英文语言名称
- `xeno.country_subdivision_l10n_en`：英文行政区划名称

### 配置

```sql
SET pg_xenophile.base_lang_code = 'en';
SET pg_xenophile.user_lang_code = 'en';
SET pg_xenophile.target_lang_codes = '{nl,fr,de}';
```
