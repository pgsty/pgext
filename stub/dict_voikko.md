## Usage

Sources:

- [Official README](https://github.com/Houston-Inc/dict_voikko/blob/746eeef7f40cadbe5a527ccee4250a2f8daf2a94/README.md)
- [Official extension SQL](https://github.com/Houston-Inc/dict_voikko/blob/746eeef7f40cadbe5a527ccee4250a2f8daf2a94/dict_voikko--1.0.sql)
- [Official C implementation](https://github.com/Houston-Inc/dict_voikko/blob/746eeef7f40cadbe5a527ccee4250a2f8daf2a94/dict_voikko.c)
- [Official control file](https://github.com/Houston-Inc/dict_voikko/blob/746eeef7f40cadbe5a527ccee4250a2f8daf2a94/dict_voikko.control)

`dict_voikko` adds a PostgreSQL full-text-search dictionary that uses Voikko morphological analysis to reduce Finnish words and compounds to base lexemes. It is a dictionary component, not a complete text-search configuration.

### Core Workflow

```sql
CREATE EXTENSION dict_voikko;

SELECT ts_lexize('voikko', 'kissoja');
SELECT ts_lexize('voikko', 'kerrostalollekohan');

CREATE TEXT SEARCH DICTIONARY finnish_voikko_stop (
  TEMPLATE = voikko_template,
  StopWords = finnish
);

CREATE TEXT SEARCH CONFIGURATION finnish_voikko (
  COPY = pg_catalog.finnish
);

ALTER TEXT SEARCH CONFIGURATION finnish_voikko
  ALTER MAPPING FOR asciiword, asciihword, hword_asciipart,
                    word, hword, hword_part
  WITH finnish_voikko_stop, finnish_stem;
```

The default `voikko` dictionary is installed with `voikko_template`. `StopWords` is the only accepted template option.

### Chaining Behavior

Recognized inflected words return base forms; the documented compound example produces `kerros` and `talo`. Stop words return an empty lexeme result, while an unrecognized word returns SQL NULL so the next dictionary in the mapping can try it. Keep a fallback such as `finnish_stem` after the Voikko dictionary.

The implementation uses only the first morphological analysis returned by Voikko and extracts stems with fixed regular expressions. Verify domain vocabulary, compounds, proper names, encodings, and desired normalization rather than treating it as a complete Finnish linguistic analyzer.

### Dependencies and Compatibility

The server library must link to `libvoikko` and have Finnish morphological dictionary data installed. The C source opens language `fi-x-morpho`, while the older README describes a dictionary named `mor-morpho`; confirm the language tag exposed by the installed Voikko data instead of following the name blindly.

Voikko resources are loaded in each backend that initializes the text-search dictionary. Account for their process memory and ensure the same dictionaries exist on every primary/standby or pooled server.

Version `1.0` uses PostgreSQL and Voikko APIs from an old codebase last changed in 2018 and has no current compatibility matrix. Build and test it against the exact PostgreSQL, `libvoikko`, and morphology-data versions before production indexing; changing analyzer data can change generated `tsvector` values and require reindexing/recomputation.
