## Usage

Sources:

- [Official extension control file](https://github.com/metabrainz/postgresql-musicbrainz-unaccent/blob/b727896da12f03f404ed96a4d0cc8c967681fbaf/musicbrainz_unaccent.control)
- [Official upstream documentation](https://github.com/metabrainz/postgresql-musicbrainz-unaccent/blob/b727896da12f03f404ed96a4d0cc8c967681fbaf/README.musicbrainz_unaccent.md)

`musicbrainz_unaccent` — Removes accents from Unicode data

The reviewed catalog snapshot records version `1.0.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "musicbrainz_unaccent";
SELECT extversion
FROM pg_extension
WHERE extname = 'musicbrainz_unaccent';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
