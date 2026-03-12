


## Usage

> [icu_ext: ICU extension functions for PostgreSQL](https://github.com/dverite/icu_ext)

Exposes [ICU](https://icu.unicode.org/) functionality to PostgreSQL. Requires PostgreSQL 11+ configured with ICU (`--with-icu`).

### Version Info

```sql
SELECT icu_version();           -- ICU library version
SELECT icu_unicode_version();   -- Unicode standard version
```

### Locale Functions

```sql
SELECT * FROM icu_locales_list() WHERE name LIKE 'es%' LIMIT 5;
SELECT icu_default_locale();
SELECT icu_set_default_locale('en');
```

### Collation Attributes

```sql
SELECT * FROM icu_collation_attributes('fr-u-ks-level2-kn');
```

### String Comparison

```sql
-- Case-sensitive, accent-insensitive comparison:
SELECT icu_compare('abce', 'abce', 'en-u-ks-level1-kc-true');  -- 0
SELECT icu_compare('Abce', 'abce', 'en-u-ks-level1-kc-true');  -- 1
```

### Sort Keys (for unique indexes)

```sql
CREATE UNIQUE INDEX idx ON my_table((icu_sort_key(name, 'fr-u-ks-level1')));
```

### Text Boundary Analysis

```sql
SELECT * FROM icu_character_boundaries('Hello', 'en');
SELECT * FROM icu_word_boundaries('I like books', 'en');
SELECT * FROM icu_sentence_boundaries('Mr. Smith went home. He was tired.', 'en');
SELECT * FROM icu_line_boundaries('Long text here', 'en');
```

### Unicode Normalization and Transformation

```sql
SELECT icu_normalize('text', 'NFC');
SELECT icu_is_normalized('text', 'NFC');
SELECT icu_transform('Hello', 'Latin-Cyrillic');
SELECT * FROM icu_transforms_list();
```

### Number Spellout

```sql
SELECT icu_number_spellout(42, 'en');   -- 'forty-two'
SELECT icu_number_spellout(42, 'fr');   -- 'quarante-deux'
```

### Spoof and Confusable Detection

```sql
SELECT icu_spoof_check('paypal');
SELECT icu_confusable_strings_check('google', 'gооgle');
```

### Character Info

```sql
SELECT icu_char_name('A');
SELECT icu_char_type('A');
```
