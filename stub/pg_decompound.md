## Usage

Sources:

- [README.md](https://github.com/fgehrlicher/pg_decompound/blob/07487d222cd730985f289440a8032f31137e19b5/README.md)
- [sql/pg_decompound--1.0.sql](https://github.com/fgehrlicher/pg_decompound/blob/07487d222cd730985f289440a8032f31137e19b5/sql/pg_decompound--1.0.sql)
- [flake.nix](https://github.com/fgehrlicher/pg_decompound/blob/07487d222cd730985f289440a8032f31137e19b5/flake.nix)
- [pg_decompound.control](https://github.com/fgehrlicher/pg_decompound/blob/07487d222cd730985f289440a8032f31137e19b5/pg_decompound.control)

`pg_decompound` supplies the `decompound` text-search dictionary template. It splits compound words into listed constituents at indexing and query time. It was developed for German; the supplied word list determines the language and emitted lexemes.

### Core Workflow

```sql
CREATE EXTENSION pg_decompound;
CREATE TEXT SEARCH DICTIONARY german_compounds (
    TEMPLATE = decompound,
    WordList = german_compounds,
    MinPiece = 4,
    MinWord = 8,
    MinCoverage = 70
);
CREATE TEXT SEARCH CONFIGURATION german_dc (COPY = german);
ALTER TEXT SEARCH CONFIGURATION german_dc
    ALTER MAPPING FOR asciiword, word, hword_part, hword_asciipart
    WITH german_compounds, german_stem;
SELECT to_tsvector('german_dc', 'Konfigurationsdatei');
```

### Operational Boundaries

An administrator installs this native extension. Its upstream development environments target PostgreSQL 16 and 17. Provide a word list before creating the dictionary: `WordList` names a `.dict` file under `$SHAREDIR/tsearch_data`. Each line contains a surface word, optionally followed by a tab and its pre-stemmed lexeme. No word list is shipped; word-list licensing is separate from the MIT-licensed extension.

Place the compound dictionary before the stemmer. An unsplittable token falls through to later dictionaries, but emitted pieces are not passed on for stemming, so the supplied lexemes must already be stemmed. `tools/build-wordlist.sh` can generate them using an existing stemmer.

`MinPiece` defaults to 4 characters, `MinWord` to 8 and `MaxWord` to 100. `MinCoverage` defaults to 0, which emits all substring matches; a positive percentage selects a non-overlapping cover and rejects insufficient coverage. `OnlyLongest` keeps the longest match at each offset. `KeepOriginal` includes the unstemmed compound; both options default to false.

Large word lists can produce noisy substring matches. `MaxWord` bounds the quadratic matching cost. Every lexeme produced from a query token is required, so different decompositions of inflected forms can cause missed matches, especially with `KeepOriginal`. Recompute stored search vectors when changing dictionary behavior.
