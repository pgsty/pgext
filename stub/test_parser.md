## Usage

Sources:

- [Official repository](https://github.com/r888888888/test_parser)
- [Extension control file at the catalog source revision](https://github.com/r888888888/test_parser/blob/356c676d16a987f9226b3ea419b4e18612c61fef/test_parser.control)
- [Version 1.0 extension SQL](https://github.com/r888888888/test_parser/blob/356c676d16a987f9226b3ea419b4e18612c61fef/test_parser--1.0.sql)
- [Version 1.0 C implementation](https://github.com/r888888888/test_parser/blob/356c676d16a987f9226b3ea419b4e18612c61fef/test_parser.c)
- [PostgreSQL text-search parser testing documentation](https://www.postgresql.org/docs/current/textsearch-debugging.html)

test_parser is a minimal example of a custom PostgreSQL full-text-search parser. The extension name contains an underscore, while the installed parser is named testparser; its implementation only separates runs of literal ASCII spaces from all other byte sequences.

### Inspect the Parser

Install it in a development database and inspect the two token types and resulting token stream before creating a text-search configuration.

```sql
CREATE EXTENSION test_parser;

SELECT * FROM ts_token_type('testparser');
SELECT * FROM ts_parse('testparser', 'alpha  beta,123');
```

Non-space runs are reported as the word token type and consecutive spaces as the blank token type. Punctuation, digits, tabs, newlines, Unicode whitespace, and linguistic word boundaries receive no special classification.

### Minimal Configuration

Map the word token to a dictionary; blank tokens normally remain unmapped but are retained by the parser so the built-in headline callback can preserve spacing.

```sql
CREATE TEXT SEARCH CONFIGURATION testcfg (PARSER = testparser);
ALTER TEXT SEARCH CONFIGURATION testcfg
    ADD MAPPING FOR word WITH simple;

SELECT to_tsvector('testcfg', 'alpha beta alpha');
```

### Intended Scope and Compatibility

This code is an educational parser, not a production tokenizer. It has only two token types, no configuration knobs, and no language-aware segmentation. Use it to study the parser callback interface or as a starting point for a separately maintained implementation, not as a claim of useful search semantics.

The C code uses PostgreSQL's parser callback ABI and reuses the built-in headline function, so build and test it against the exact PostgreSQL major version. Acceptance tests should cover multibyte text, embedded null assumptions, long input, all whitespace forms, token offsets, configuration mappings, and headline output. The extension is relocatable and declares no preload or restart requirement.
