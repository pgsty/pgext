## Usage

Sources:

- [Official PGXN distribution](https://pgxn.org/dist/pg_cld2/1.0.0/)
- [Official pg_cld2 README](https://github.com/hedges333/pg_cld2/blob/f13f0721009db68fc0333aa7671dcf27ba16c739/README.md)
- [Version 1.0.0 extension SQL](https://github.com/hedges333/pg_cld2/blob/f13f0721009db68fc0333aa7671dcf27ba16c739/sql/pg_cld2--1.0.0.sql)

`pg_cld2` exposes Google's Compact Language Detector 2 (CLD2) to PostgreSQL. It estimates the human language and writing script of text and maps candidate languages to available PostgreSQL text-search configurations. Results are statistical hints, not authoritative classifications.

### Detect language and script

The extension requires the external CLD2 library at build and run time:

```sql
CREATE EXTENSION pg_cld2;

SELECT is_reliable,
       mll_language_code,
       mll_primary_script_code,
       language_1_language_code,
       language_1_percent,
       language_1_ts_name
FROM pg_cld2_detect_language(
    'This is a sample text to detect the language.'
);
```

`pg_cld2_detect_language` returns the `pg_cld2_language_detection` composite type. Besides byte counts and UTF-8 validity, it contains a most-likely-language group (`mll_*`) and three ranked candidate groups (`language_1_*` through `language_3_*`) with language names and codes, script names and codes, percentages, normalized scores, and a matched `pg_catalog.pg_ts_config` name.

For the full call signature and defaults, run:

```sql
SELECT pg_cld2_usage();
```

Important inputs include `is_plain_text` (set false for HTML), `content_language_hint`, `tld_hint`, `cld2_language_hint`, `best_effort`, `text_encoding`, `tsconfig_language_hint`, and `locale_lang_hint`. Hints deliberately bias the result. Non-UTF-8 input can be converted when the correct PostgreSQL encoding name is supplied.

### Interpretation and caveats

Check `is_reliable` before acting on a result, and retain candidate percentages when mixed-language text matters. The default `best_effort = true` can return a guess instead of `UNKNOWN` for short input; disable it when an uncertain result is preferable to a forced answer. `valid_prefix_bytes` and `is_valid_utf8` help identify malformed input.

CLD2 is an older statistical model. Short strings, proper nouns, transliteration, mixed scripts, markup, and biased hints can produce misleading classifications. Do not use its output as an authorization, compliance, nationality, or identity decision. The PGXN metadata only states a historical PostgreSQL minimum; verify the CLD2 ABI and compile this release against the exact modern server version before deployment.
