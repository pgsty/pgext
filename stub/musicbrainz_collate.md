## Usage

Sources:

- [Official README](https://github.com/metabrainz/postgresql-musicbrainz-collate/blob/6c350bcbfd3418db4c90826eeea22c24146b2d05/README.musicbrainz_collate.md)
- [Extension SQL](https://github.com/metabrainz/postgresql-musicbrainz-collate/blob/6c350bcbfd3418db4c90826eeea22c24146b2d05/sql/musicbrainz_collate.sql)
- [ICU sort-key implementation](https://github.com/metabrainz/postgresql-musicbrainz-collate/blob/6c350bcbfd3418db4c90826eeea22c24146b2d05/musicbrainz_collate.c)

`musicbrainz_collate` converts UTF-8 text to an ICU Unicode Collation Algorithm sort key represented as `bytea`. Use the key in `ORDER BY` or an expression index when reproducing the MusicBrainz-style locale-neutral, numeric-aware ordering implemented by this archived extension.

### Core Workflow

```sql
CREATE EXTENSION musicbrainz_collate;

SELECT name
FROM artist
ORDER BY musicbrainz_collate(name);

CREATE INDEX artist_musicbrainz_name_idx
    ON artist (musicbrainz_collate(name));

SELECT name
FROM artist
ORDER BY musicbrainz_collate(name)
LIMIT 50;
```

The sole public function is `musicbrainz_collate(text) RETURNS bytea`. It is strict and marked immutable, so PostgreSQL can use it in expression indexes. The implementation opens ICU's root/default collator and enables numeric collation, making digit sequences sort by numeric value rather than character-by-character.

### Encoding and Index Maintenance

The function assumes that its input bytes are UTF-8 and is not useful in databases with another server encoding. Confirm `SHOW server_encoding` before use.

ICU version changes can change generated sort keys even though the SQL function is declared immutable. After an ICU library upgrade, run `REINDEX INDEX artist_musicbrainz_name_idx` for every expression index based on the function and revalidate application-visible ordering. Replicas and failover targets must use compatible ICU builds.

### Compatibility Boundary

The historical README requires PostgreSQL 8.3 or newer and ICU 3.8 or newer, but the repository is archived and does not establish a modern PostgreSQL/ICU compatibility matrix. Compare its ordering with built-in ICU collations available on the target PostgreSQL release before adopting legacy behavior.
