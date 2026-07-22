## Usage

Sources:

- [Official README](https://github.com/za-arthur/pg_textparser/blob/master/README.md)
- [Extension control file](https://github.com/za-arthur/pg_textparser/blob/master/pg_textparser.control)
- [Version 1.0 extension SQL](https://github.com/za-arthur/pg_textparser/blob/master/pg_textparser--1.0.sql)
- [Version 1.0 C implementation](https://github.com/za-arthur/pg_textparser/blob/master/pg_textparser.c)
- [Version 1.0 scanner source](https://github.com/za-arthur/pg_textparser/blob/master/textscan.l)

pg_textparser is an abandoned, incomplete prototype for a PostgreSQL full-text-search parser. The cataloged 1.0 sources do not provide a usable parser and should be treated as source material for development, not as a deployable extension.

### Upstream State

The extension SQL declares parser callback functions, but its parser definition references a headline callback that the SQL does not create and omits the declared lexical-type callback from the parser definition. As shipped, extension creation is therefore incomplete.

The C callbacks for start, next-token, end, lexical types, and headline all return null instead of maintaining parser state or emitting tokens. The separate scanner prototype recognizes only digits and returns an end-like token value, and it is not integrated into the callbacks.

### Intended Verification Surface

After a maintainer completes the SQL declarations and callback implementation, PostgreSQL's parser inspection functions should be the first acceptance gate:

```sql
SELECT * FROM ts_token_type('textparser');
SELECT * FROM ts_parse('textparser', 'alpha 123 beta');
```

A real implementation must return a stable token-type catalog, consume input without looping, preserve correct byte lengths and multibyte boundaries, and produce tokens that can be mapped into a text-search configuration. Headline behavior also needs independent tests.

### Deployment Boundary

Do not create production indexes or search configurations that depend on this source snapshot. A successful compile alone would not establish parser correctness, and local changes that merely make the SQL installable would still require regression tests for tokenization, empty input, whitespace, punctuation, encoding, long values, and headline generation.

The control file marks the extension relocatable and declares no preload or restart requirement, but that does not change its incomplete lifecycle status. Keep it quarantined until an authoritative maintained release supplies a coherent SQL surface, functional callbacks, and tests.
