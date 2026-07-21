## Usage

Sources:

- [pg_pinyin v0.0.5 README](https://github.com/aiyou178/pg_pinyin/blob/v0.0.5/readme.md)
- [pg_pinyin v0.0.5 control file](https://github.com/aiyou178/pg_pinyin/blob/v0.0.5/pg_pinyin.control)
- [0.0.4 to 0.0.5 upgrade SQL](https://github.com/aiyou178/pg_pinyin/blob/v0.0.5/pg_pinyin--0.0.4--0.0.5.sql)

pg_pinyin romanizes Chinese text and exposes tokenizer and query helpers for search applications. Use pg_pinyin to create stable Pinyin search keys, tokenize Han text, or expand Pinyin input into a pg_search regular-expression query.

Version 0.0.5 is primarily a packaging and toolchain update; its upgrade script makes no SQL catalog changes, so the user-facing API remains compatible with 0.0.4.

### Create the Extension

    CREATE EXTENSION pg_pinyin;

The extension is relocatable and does not require shared_preload_libraries or a server restart.

### Romanize Text

Romanize character by character or use word-aware segmentation:

    SELECT pinyin_char_romanize('重庆');
    SELECT pinyin_word_romanize('重庆火锅');
    SELECT pinyin_word_romanize('重庆火锅', ' ');

Both functions accept an optional suffix inserted after each emitted Pinyin unit. Character mode is deterministic per character; word mode uses the bundled word dictionary to resolve contextual pronunciations.

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
