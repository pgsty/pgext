## Usage

Sources:

- [Official WIP notice](https://github.com/3manuek/twitter_parser/blob/1fafaae28bd7dc2c0d01617ddf992a19027f2c43/README)
- [Version 1.0 extension SQL](https://github.com/3manuek/twitter_parser/blob/1fafaae28bd7dc2c0d01617ddf992a19027f2c43/twitter_parser--1.0.sql)
- [C parser implementation](https://github.com/3manuek/twitter_parser/blob/1fafaae28bd7dc2c0d01617ddf992a19027f2c43/twitter_parser.c)

`twitter_parser` is an unfinished full-text-search parser prototype for recognizing Twitter-style `@mention` and `#hashtag` tokens. Creating the extension registers a parser named `twitterparser`; it does not create a ready-to-use text search configuration.

### Core Workflow

Use the parser introspection functions to evaluate it before considering a configuration:

```sql
CREATE EXTENSION twitter_parser;

SELECT * FROM ts_token_type('twitterparser');
SELECT * FROM ts_parse('twitterparser', '#postgres');
SELECT * FROM ts_parse('twitterparser', '@postgres');
```

The parser advertises token ID `24` as `mention` and token ID `25` as `hashtag`. A text search configuration would also need dictionary mappings for those aliases, but this prototype is not complete enough for production indexing.

### Object Index

- Text search parser `twitterparser` provides start, token, end, headline, and token-type callbacks.
- Token type `mention` (`24`) is intended for text beginning with `@`.
- Token type `hashtag` (`25`) is intended for text beginning with `#`.
- `twitterprs_start`, `twitterprs_getlexeme`, `twitterprs_end`, and `twitterprs_lextype` are low-level parser support functions, not application APIs.

### Operational Notes

Version `1.0` is relocatable and declares no dependency or preload requirement. Upstream labels it WIP. The implementation consumes a hashtag or mention only until an ASCII space; it does not validate punctuation or Unicode rules. More importantly, it does not advance over spaces or ordinary text and leaves the token type unset on such input, so mixed sentences can stop early or behave unpredictably.

Do not build durable text-search indexes with this parser without repairing and testing its state machine. Test malformed input, multiple tokens, punctuation, multibyte text, headline generation, upgrades, and reindexing in an isolated database. For production, use a maintained parser or normalize hashtags and mentions into separate indexed columns.
