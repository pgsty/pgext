## Usage

Sources:

- [Official README](https://github.com/za-arthur/dict_translate/blob/c95381026d760bdfae6a8cb2bab94c7ca9c06575/README.md)
- [Official extension SQL](https://github.com/za-arthur/dict_translate/blob/c95381026d760bdfae6a8cb2bab94c7ca9c06575/dict_translate--1.0.sql)

`dict_translate` 1.0 provides the `translate` full-text-search dictionary template. It normalizes an input token with another dictionary and expands it to the translation group stored in a server-side `.trn` file, making translated terms searchable as equivalent lexemes.

### Core Workflow

Create a translation file such as `$SHAREDIR/tsearch_data/test_trn.trn` on every database server that can execute the configuration:

```text
forest wald forst holz
home haus heim
```

Install the extension and build a dictionary and configuration around that file:

```sql
CREATE EXTENSION dict_translate;

CREATE TEXT SEARCH DICTIONARY test_trn (
  TEMPLATE = translate,
  DictFile = test_trn,
  InputDict = pg_catalog.english_stem
);

CREATE TEXT SEARCH CONFIGURATION test_cfg (COPY = simple);

ALTER TEXT SEARCH CONFIGURATION test_cfg
  ALTER MAPPING FOR asciiword, asciihword, hword_asciipart,
                    word, hword, hword_part
  WITH test_trn, english_stem;

SELECT to_tsvector('test_cfg', 'homes forest haus');
SELECT to_tsvector('test_cfg', 'homes forest haus')
       @@ to_tsquery('test_cfg', 'home');
```

`DictFile` names the translation file without its `.trn` suffix. Each line starts with a source lexeme followed by the lexemes that should be emitted with it. `InputDict` identifies the dictionary used to normalize the input before translation lookup.

### Important Objects

- `translate` is the text search template installed by the extension.
- `dtrn_init(internal)` initializes a dictionary from `DictFile` and `InputDict` options.
- `dtrn_lexize(internal, internal, internal, internal)` performs normalization and translation expansion for the text-search subsystem.

### Operational Notes

The `.trn` file is filesystem configuration, not database data; deploy and version it consistently across primaries, standbys, and replacement hosts. Changing the file does not rewrite existing `tsvector` values, so stored vectors or indexes must be refreshed when translation rules change. Use a trusted normalization dictionary and test stemming behavior in both indexing and query configurations.
