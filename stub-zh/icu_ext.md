


## 用法

> 来源：[README](https://github.com/dverite/icu_ext/blob/master/README.md), [datetime docs](https://github.com/dverite/icu_ext/blob/master/README-datetime.md), [v1.10.0 release](https://github.com/dverite/icu_ext/releases/tag/v1.10.0)

`icu_ext` 将 [ICU](https://icu.unicode.org/) 功能暴露给 PostgreSQL。上游要求 PostgreSQL 11+ 且编译时启用 ICU（`--with-icu`）；pgext 目录记录的版本是 `1.10.0`，覆盖 PostgreSQL 14-18，v1.10.0 release 说明中提到 PostgreSQL 18 兼容性。

### 启用扩展

```sql
CREATE EXTENSION icu_ext;
```

### 版本信息

```sql
SELECT icu_version();           -- ICU library version
SELECT icu_unicode_version();   -- Unicode standard version
```

### Locale 函数

```sql
SELECT * FROM icu_locales_list() WHERE name LIKE 'es%' LIMIT 5;
SELECT icu_default_locale();
SELECT icu_set_default_locale('en');
```

### Collation 属性

```sql
SELECT * FROM icu_collation_attributes('fr-u-ks-level2-kn');
```

### 字符串比较

```sql
-- Case-sensitive, accent-insensitive comparison:
SELECT icu_compare('abce', 'abce', 'en-u-ks-level1-kc-true');  -- 0
SELECT icu_compare('Abce', 'abce', 'en-u-ks-level1-kc-true');  -- 1
```

### 排序键与语言学搜索

```sql
CREATE UNIQUE INDEX idx ON my_table((icu_sort_key(name, 'fr-u-ks-level1')));

SELECT icu_strpos('Jean-Rene Dupont', 'jeanrene', 'fr-u-ks-level1-ka-shifted');
SELECT icu_replace('Jean-Rene Dupont', 'jeanrene', '{firstname}', 'fr-u-ks-level1-ka-shifted');
```

### 文本边界分析

```sql
SELECT * FROM icu_character_boundaries('Hello', 'en');
SELECT * FROM icu_word_boundaries('I like books', 'en');
SELECT * FROM icu_sentence_boundaries('Mr. Smith went home. He was tired.', 'en');
SELECT * FROM icu_line_boundaries('Long text here', 'en');
```

### Unicode 规范化与转换

```sql
SELECT icu_normalize('text', 'NFC');
SELECT icu_is_normalized('text', 'NFC');
SELECT icu_transform('Hello', 'Latin-Cyrillic');
SELECT * FROM icu_transforms_list();
```

### 日期与时间本地化

```sql
SET icu_ext.locale TO '@calendar=buddhist';

SELECT icu_format_date('2020-12-31'::date, '{medium}', 'en@calendar=ethiopic');
SELECT icu_parse_date('25/09/2566', 'dd/MM/yyyy');
SELECT icu_format_datetime(now(), 'GGGG dd/MMMM/yyyy HH:mm:ss.SSS z', 'fr@calendar=buddhist');
```

datetime 文档还定义了 `icu_date`、`icu_timestamptz` 和 `icu_interval`，以及用于本地化输入/输出和感知日历算术的 `icu_ext.locale`、`icu_ext.date_format`、`icu_ext.timestamptz_format` 设置。

### 数字拼写

```sql
SELECT icu_number_spellout(42, 'en');   -- 'forty-two'
SELECT icu_number_spellout(42, 'fr');   -- 'quarante-deux'
```

### 欺骗与混淆检测

```sql
SELECT icu_spoof_check('paypal');
SELECT icu_confusable_strings_check('google', 'gооgle');
SELECT icu_confusable_string_skeleton('phi1');
```

### 字符信息

```sql
SELECT icu_char_name('A');
SELECT icu_char_type('A');
SELECT icu_char_ublock_id('A');
SELECT * FROM icu_unicode_blocks() WHERE block_name = 'Basic_Latin';
```

### 注意事项

- 依赖 ICU collation 或 Unicode 数据的函数，在链接的 ICU 库变化后行为可能改变。
- `icu_sort_key()` 适合用于索引，但基于排序键构建的索引在 ICU 升级后应复核。
