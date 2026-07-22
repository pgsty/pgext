## Usage

Sources:

- [Official README](https://github.com/obartunov/dict_roman/blob/22c2cd1acd8eb077a6d76e543e4ae6a4a0071d61/README.dict_roman)
- [Official extension control file](https://github.com/obartunov/dict_roman/blob/22c2cd1acd8eb077a6d76e543e4ae6a4a0071d61/dict_roman.control)

`dict_roman` provides the `roman` full-text-search dictionary, which normalizes a Roman-numeral token to its Arabic-integer spelling. Use it when numerals such as `XXIV` should enter a text-search vector as `24` rather than as alphabetic text.

### Core Workflow

Install the extension, then inspect the dictionary directly with `ts_lexize`:

```sql
CREATE EXTENSION dict_roman;

SELECT ts_lexize('roman', 'XXIV');  -- {24}
SELECT ts_lexize('roman', 'XLII');  -- {42}
```

To use the normalization during indexing and querying, map an appropriate token type in a text-search configuration:

```sql
CREATE TEXT SEARCH CONFIGURATION english_with_roman (COPY = english);
ALTER TEXT SEARCH CONFIGURATION english_with_roman
  ALTER MAPPING FOR asciiword
  WITH roman, english_stem;

SELECT to_tsvector('english_with_roman', 'Chapter XXIV');
```

The fallback dictionary in the mapping is important for ordinary alphabetic tokens that `roman` does not recognize.

### Operational Notes

- The documented SQL surface is the `roman` dictionary; the upstream README demonstrates `ts_lexize` but does not publish a complete grammar or accepted numeric range.
- Test zero, subtractive notation, case, malformed input, and large values against the exact packaged build before relying on the normalized value.
- Changing a text-search configuration changes future parsing only. Rebuild stored `tsvector` values and dependent indexes when a new mapping must apply to existing rows.
