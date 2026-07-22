## Usage

Sources:

- [Official README](https://github.com/zulip/tsearch_extras/blob/f566e9606bac00a22c4b62e9511b022c861db05c/README.md)
- [Extension SQL](https://github.com/zulip/tsearch_extras/blob/f566e9606bac00a22c4b62e9511b022c861db05c/tsearch_extras--1.0.sql)
- [Full-text implementation](https://github.com/zulip/tsearch_extras/blob/f566e9606bac00a22c4b62e9511b022c861db05c/tsearch_extras.c)

`tsearch_extras` version `1.0` exposes lower-level information from PostgreSQL full-text search: locations and lengths of query matches in source text, and the lexemes stored in a `tsvector`.

### Core Workflow

```sql
CREATE EXTENSION tsearch_extras;

SELECT ts_match_locs_array(
  'The quick brown fox jumped. The jumping continued.',
  plainto_tsquery('jump')
);

SELECT tsvector_lexemes(
  to_tsvector('The quick brown fox jumped over the lazy dog')
);
```

Use an explicit text-search configuration when results must not depend on the session default:

```sql
SELECT ts_match_locs_array(
  'english'::regconfig,
  'The quick brown fox jumped.',
  plainto_tsquery('english', 'jump')
);
```

### Objects

- `ts_match_locs_array(text, tsquery) returns int4[][2]`: match location/length pairs using the current text-search configuration.
- `ts_match_locs_array(regconfig, text, tsquery) returns int4[][2]`: the same result using the selected configuration.
- `tsvector_lexemes(tsvector) returns text[]`: lexemes stored in the vector, without positional metadata or weights.

### Compatibility Boundary

Match locations are produced through PostgreSQL's parser and headline internals, so validate offsets with the exact configuration, encoding, dictionary, and server version used by the application. The repository is archived and the C source depends on internal full-text-search headers. Verify compilation and regression behavior for the target PostgreSQL release; no preload is required.
