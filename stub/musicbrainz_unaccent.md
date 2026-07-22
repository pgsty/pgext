## Usage

Sources:

- [Official upstream documentation](https://github.com/metabrainz/postgresql-musicbrainz-unaccent/blob/b727896da12f03f404ed96a4d0cc8c967681fbaf/README.musicbrainz_unaccent.md)
- [Official extension SQL](https://github.com/metabrainz/postgresql-musicbrainz-unaccent/blob/b727896da12f03f404ed96a4d0cc8c967681fbaf/sql/musicbrainz_unaccent.sql)
- [Official implementation](https://github.com/metabrainz/postgresql-musicbrainz-unaccent/blob/b727896da12f03f404ed96a4d0cc8c967681fbaf/musicbrainz_unaccent.c)

`musicbrainz_unaccent` removes a fixed set of Unicode accents with a C function and a text-search dictionary. Unlike PostgreSQL's general-purpose `unaccent` function, its function is declared immutable so it can be used in expression indexes. It was built for MusicBrainz normalization and its repository is archived.

### Core Workflow

Install the extension and use either the scalar function or the text-search dictionary:

```sql
CREATE EXTENSION musicbrainz_unaccent;

SELECT musicbrainz_unaccent('ľščťžýáí');
SELECT ts_lexize('musicbrainz_unaccentdict', 'ľščťžýáí');

CREATE INDEX artist_unaccent_idx
  ON artist (musicbrainz_unaccent(name));
```

The scalar example returns `lsctzyai`. Queries intended to use the expression index must apply the identical immutable expression.

### Important Objects

- `musicbrainz_unaccent(text)` returns normalized text and is declared `IMMUTABLE` and `STRICT`.
- `musicbrainz_unaccentdict_template` is the text-search dictionary template.
- `musicbrainz_unaccentdict` is the installed dictionary; it lowercases lexemes before applying the compiled mapping.

### Compatibility and Caveats

The implementation assumes UTF-8 input regardless of the database encoding. Do not use it in a non-UTF-8 database. Its mapping is compiled into the extension binary rather than loaded from an editable dictionary rules file, so changing normalization rules requires a rebuild. Lowercasing also depends on PostgreSQL's locale behavior. Because the upstream project is archived and tied to old PostgreSQL internals, verify ABI compatibility, locale-sensitive characters, index rebuild behavior, and every language required by the application before adoption. Pin the reviewed source and mapping so normalized keys do not change unexpectedly.
