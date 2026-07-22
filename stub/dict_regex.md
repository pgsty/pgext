## Usage

Sources:

- [Official README](https://github.com/obartunov/dict_regex/blob/master/README.dict_regex)
- [Official extension SQL](https://github.com/obartunov/dict_regex/blob/master/dict_regex--1.0.sql)
- [Official example rules](https://github.com/obartunov/dict_regex/blob/master/data/dict_regex.rules)
- [Official control file](https://github.com/obartunov/dict_regex/blob/master/dict_regex.control)

`dict_regex` is a PostgreSQL full-text-search dictionary that rewrites one or more input tokens through PCRE-compatible regular-expression rules. It combines synonym expansion, thesaurus-style multiword reduction, capture-group substitution, and stop-word removal in one dictionary.

### Core Workflow

Place a reviewed rule file where the PostgreSQL server process can read it, create the extension, and point the installed `regex` dictionary at that file:

```text
# synonym expansion
catalogue catalogue catalog

# collapse a phrase
regular\sexpressions? regex

# reorder two captured groups
(\d\d\d\d)(\d\d\d\d) $2$1
```

```sql
CREATE EXTENSION dict_regex;
ALTER TEXT SEARCH DICTIONARY regex
  (RULES = '/srv/postgresql/dict_regex.rules');

SELECT ts_lexize('regex', 'catalogue');
```

Attach `regex` to selected token types in a text-search configuration with `ALTER TEXT SEARCH CONFIGURATION ... ALTER MAPPING ...` when moving beyond direct `ts_lexize` tests.

### Rule Semantics

Each non-comment line contains a whitespace-separated pattern and output recipe. An empty recipe makes the match a stop word. The dictionary anchors patterns automatically, chooses the longest matching token sequence, and breaks equal-length ties by file order. Multiple output words are returned at the same position as synonyms; `$1` through `$9` substitute captured groups and output is lowercased.

The partial-matching engine restricts some repeated single characters/metasequences. For example, use `(\d)+` rather than `\d+`, and `(a){2,4}` rather than `a{2,4}`.

### Caveats

The extension requires PCRE and reads a server-side rules file. Upstream explicitly marks version `1.0` insecure because it does not adequately restrict the file path or contents: a user able to alter dictionary options can point it at arbitrary files readable by the PostgreSQL server and may reconstruct their contents. Do not expose dictionary alteration to untrusted roles or use it in a security-critical environment without hardening the source.

Rules are order-sensitive and multi-token matching changes lexeme positions. Test both indexing and query configurations with representative text, and reload/recreate sessions as required after changing the file.
