## Usage

Sources:

- [pg_pinyin v0.0.8 README](https://github.com/aiyou178/pg_pinyin/blob/v0.0.8/readme.md)
- [pg_pinyin v0.0.8 control file](https://github.com/aiyou178/pg_pinyin/blob/v0.0.8/pg_pinyin.control)
- [0.0.7 to 0.0.8 upgrade SQL](https://github.com/aiyou178/pg_pinyin/blob/v0.0.8/pg_pinyin--0.0.7--0.0.8.sql)
- [0.0.8 parallel read-only regression](https://github.com/aiyou178/pg_pinyin/blob/v0.0.8/test/pgtap/04_parallel_read_only.sql)
- [0.0.8 implementation](https://github.com/aiyou178/pg_pinyin/blob/v0.0.8/src/lib.rs)

pg_pinyin romanizes Chinese text and exposes tokenizer and query helpers for search applications. Use pg_pinyin to create stable Pinyin search keys, tokenize Han text, or expand Pinyin input into a pg_search regular-expression query.

Version 0.0.8 fixes read-only SPI access for parallel romanization, tokenizer conversion and suffix-dictionary queries. It uses pgrx 0.19.3 and supports PostgreSQL 14-18 plus PostgreSQL 19 beta4. The provided 0.0.7-to-0.0.8 migration makes no SQL object changes; the runtime fix is in the shared library. After installing the new extension files, run ALTER EXTENSION pg_pinyin UPDATE TO '0.0.8'. Pigsty package metadata is maintained separately.

### Create the Extension

    CREATE EXTENSION pg_pinyin;

The extension is relocatable and does not require shared_preload_libraries or a server restart.

### Romanize Text

Romanize character by character or use word-aware segmentation:

    SELECT pinyin_char_romanize('重庆');
    SELECT pinyin_word_romanize('重庆火锅');
    SELECT pinyin_word_romanize('重庆火锅', '_custom');

The optional suffix selects custom dictionary tables; it is not an output separator. For example, _custom selects pinyin.pinyin_mapping_custom and pinyin.pinyin_words_custom, whose entries override the built-in dictionaries. Word mode uses dictionary segmentation to resolve contextual pronunciations. Clear the suffix cache after modifying these tables using public.pinyin_clear_suffix_cache('_custom').

### Use pg_search Tokenizer Input

Word romanization also accepts a pg_search tokenizer result when that extension is available:

    SELECT pinyin_word_romanize(
      description::pdb.icu::text[]
    )
    FROM documents;

The overload returns romanized text; it does not expose a row-per-token API. Use the plain-text overload when pg_search tokenization is not required.

### Build a pg_search Query

When pg_search was installed before pg_pinyin, pg_pinyin provides a typed overload that returns pdb.query:

    SELECT *
    FROM documents
    WHERE id @@@ pinyin_regex_phrase(
      'chong qing',
      slope => 1,
      max_expansions => 64,
      generated_pinyin => true
    );

If pg_search is absent, the same entry point is installed as an error-reporting stub rather than silently returning a different type. Install dependencies in the intended order and test the function signature after upgrades.

### Object Index

- pinyin_char_romanize(text [, suffix]) returns character-based Pinyin text.
- pinyin_word_romanize(text [, suffix]) returns dictionary-segmented Pinyin text.
- pinyin_word_romanize(tokenizer_input [, suffix]) accepts a pg_search tokenizer result.
- pinyin_regex_phrase(text, slope, max_expansions, generated_pinyin) constructs a pg_search Pinyin phrase query when that integration is available.
- pinyin_regex_phrase_patterns is an internal pattern-building helper; prefer the public query function.

### Operational Notes

The extension ships generated character and word dictionaries in its pinyin schema. Treat those tables as extension-managed data rather than application tables. Romanization is normalization, not translation, and ambiguous or domain-specific readings may require application-side review.
