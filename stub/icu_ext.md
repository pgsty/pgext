
## Usage

> Sources: [README](https://github.com/dverite/icu_ext/blob/master/README.md), [datetime docs](https://github.com/dverite/icu_ext/blob/master/README-datetime.md), [v1.10.0 release](https://github.com/dverite/icu_ext/releases/tag/v1.10.0)

`icu_ext` exposes [ICU](https://icu.unicode.org/) functionality to PostgreSQL. Upstream requires PostgreSQL 11+ configured with ICU (`--with-icu`); the pgext catalog tracks version `1.10.0` for PostgreSQL 14-18, with the v1.10.0 release noting PostgreSQL 18 compatibility.

### Enable the Extension

```sql
CREATE EXTENSION icu_ext;
```

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

### Sort Keys and Linguistic Search

```sql
CREATE UNIQUE INDEX idx ON my_table((icu_sort_key(name, 'fr-u-ks-level1')));

SELECT icu_strpos('Jean-Rene Dupont', 'jeanrene', 'fr-u-ks-level1-ka-shifted');
SELECT icu_replace('Jean-Rene Dupont', 'jeanrene', '{firstname}', 'fr-u-ks-level1-ka-shifted');
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

### Date and Time Localization

```sql
SET icu_ext.locale TO '@calendar=buddhist';

SELECT icu_format_date('2020-12-31'::date, '{medium}', 'en@calendar=ethiopic');
SELECT icu_parse_date('25/09/2566', 'dd/MM/yyyy');
SELECT icu_format_datetime(now(), 'GGGG dd/MMMM/yyyy HH:mm:ss.SSS z', 'fr@calendar=buddhist');
```

The datetime docs also define `icu_date`, `icu_timestamptz`, and `icu_interval`, plus the `icu_ext.locale`, `icu_ext.date_format`, and `icu_ext.timestamptz_format` settings used for localized input/output and calendar-aware arithmetic.

### Number Spellout

```sql
SELECT icu_number_spellout(42, 'en');   -- 'forty-two'
SELECT icu_number_spellout(42, 'fr');   -- 'quarante-deux'
```

### Spoof and Confusable Detection

```sql
SELECT icu_spoof_check('paypal');
SELECT icu_confusable_strings_check('google', 'gооgle');
SELECT icu_confusable_string_skeleton('phi1');
```

### Character Info

```sql
SELECT icu_char_name('A');
SELECT icu_char_type('A');
SELECT icu_char_ublock_id('A');
SELECT * FROM icu_unicode_blocks() WHERE block_name = 'Basic_Latin';
```

### Caveats

- Functions that depend on ICU collation or Unicode data can change behavior when the linked ICU library changes.
- `icu_sort_key()` is index-friendly, but indexes built on sort keys should be reviewed after ICU upgrades.
