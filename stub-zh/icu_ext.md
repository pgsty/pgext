

## 用法

> [icu_ext: PostgreSQL 的 ICU 扩展函数](https://github.com/dverite/icu_ext)

将 [ICU](https://icu.unicode.org/) 功能暴露给 PostgreSQL。需要 PostgreSQL 11+ 且编译时启用了 ICU（`--with-icu`）。

### 版本信息

```sql
SELECT icu_version();           -- ICU 库版本
SELECT icu_unicode_version();   -- Unicode 标准版本
```

### 区域设置函数

```sql
SELECT * FROM icu_locales_list() WHERE name LIKE 'es%' LIMIT 5;
SELECT icu_default_locale();
SELECT icu_set_default_locale('en');
```

### 排序规则属性

```sql
SELECT * FROM icu_collation_attributes('fr-u-ks-level2-kn');
```

### 字符串比较

```sql
-- 大小写敏感、重音不敏感的比较：
SELECT icu_compare('abce', 'abce', 'en-u-ks-level1-kc-true');  -- 0
SELECT icu_compare('Abce', 'abce', 'en-u-ks-level1-kc-true');  -- 1
```

### 排序键（用于唯一索引）

```sql
CREATE UNIQUE INDEX idx ON my_table((icu_sort_key(name, 'fr-u-ks-level1')));
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

### 数字拼读

```sql
SELECT icu_number_spellout(42, 'en');   -- 'forty-two'
SELECT icu_number_spellout(42, 'fr');   -- 'quarante-deux'
```

### 欺骗与混淆检测

```sql
SELECT icu_spoof_check('paypal');
SELECT icu_confusable_strings_check('google', 'gооgle');
```

### 字符信息

```sql
SELECT icu_char_name('A');
SELECT icu_char_type('A');
```
