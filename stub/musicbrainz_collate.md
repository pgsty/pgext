## Usage

Sources:

- [Official extension control file](https://github.com/metabrainz/postgresql-musicbrainz-collate/blob/6c350bcbfd3418db4c90826eeea22c24146b2d05/musicbrainz_collate.control)
- [Official upstream documentation](https://github.com/metabrainz/postgresql-musicbrainz-collate/blob/6c350bcbfd3418db4c90826eeea22c24146b2d05/README.musicbrainz_collate.md)

`musicbrainz_collate` — Provides Unicode Collation Algorithm support

The reviewed catalog snapshot records version `1.0.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "musicbrainz_collate";
SELECT extversion
FROM pg_extension
WHERE extname = 'musicbrainz_collate';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
