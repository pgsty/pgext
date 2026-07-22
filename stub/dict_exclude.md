## Usage

Sources:

- [Official README](https://github.com/no0p/dict_exclude/blob/2058a16ca44942a0c11eae8f7bffa9a733d7cc5a/README.md)
- [Extension SQL](https://github.com/no0p/dict_exclude/blob/2058a16ca44942a0c11eae8f7bffa9a733d7cc5a/dict_exclude--1.0.sql)
- [Dictionary implementation](https://github.com/no0p/dict_exclude/blob/2058a16ca44942a0c11eae8f7bffa9a733d7cc5a/dict_exclude.c)

`dict_exclude` is a prototype full-text-search dictionary that discards tokens matching administrator-supplied regular expressions. It is useful for filtering patterned OCR noise or application-specific token families before a normal stemming dictionary.

### Core Workflow

Install a server-side rule file named `exclude.rules` under PostgreSQL's `tsearch_data` directory. The file contains one PostgreSQL regular expression per line, for example:

```text
^abc$
^def[0-9]+$
```

Create the extension, copy an existing configuration, and put `dict_exclude` first in the mapping so nonmatching tokens continue to `english_stem`.

```sql
CREATE EXTENSION dict_exclude;

CREATE TEXT SEARCH CONFIGURATION ocr_filtered
    (COPY = pg_catalog.english);

ALTER TEXT SEARCH CONFIGURATION ocr_filtered
    ALTER MAPPING FOR asciihword, asciiword
    WITH dict_exclude, english_stem;

SELECT to_tsvector(
    'ocr_filtered',
    'fat abc cat def123 visible'
);
```

### Objects and Rule Behavior

- `dict_exclude_template` is the text-search template.
- `dict_exclude` is the dictionary created from that template.
- `dict_exclude_init` loads and compiles the rules; `dict_exclude_lexize` returns an empty lexeme set for a match and passes a nonmatch to the next dictionary.
- The filename base is fixed as `exclude`, so the resource is `exclude.rules`; there is no per-dictionary option for a different ruleset.

### Operational Notes

The rule file is read from the database server's installation tree and must be deployed by an administrator. Missing or invalid rules prevent dictionary initialization. Compiled rules are cached in backend dictionary state, so validate changes with new sessions and recreate or reload affected text-search objects as appropriate for the server version.

Every token is tested against the regex list. Broad or expensive patterns can remove legitimate terms or consume substantial CPU, and the resulting `tsvector` changes search semantics. The upstream project explicitly labels the extension a development-only prototype and documents PostgreSQL 9.4-era usage; test rules, performance, encoding, and target-version compatibility before adoption.
