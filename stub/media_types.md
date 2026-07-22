## Usage

Sources:

- [Official control file](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/media_types/media_types.control)
- [Official Media Types README](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/media_types/README.md)
- [Generated PostgreSQL 17 extension SQL](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/media_types/extension/usr/share/postgresql/17/extension/media_types--0.1.0.sql)

`media_types` adds the `MediaType` type for validated media-type strings such as `image/gif`. It is useful when content metadata should be parsed and constrained instead of stored as unchecked text.

### Core Workflow

```sql
CREATE EXTENSION media_types;

CREATE TABLE assets (
    path text PRIMARY KEY,
    content_type MediaType NOT NULL
);

INSERT INTO assets VALUES ('logo.gif', 'image/gif');
SELECT * FROM assets WHERE content_type = 'image/gif'::MediaType;
```

The parser checks the root type and subtype against the registry embedded in the build and preserves optional parameters. Unknown combinations are rejected.

### SQL Objects

- `MediaType` is a variable-length type whose PGRX binary protocol uses CBOR serialization in this release.
- Comparison operators `=`, `<>`, `<`, `<=`, `>`, and `>=` are available.
- `MediaType_btree_ops` and `MediaType_hash_ops` support B-tree and hash indexes.

### Operational Notes

Version 0.1.0 is non-relocatable. Its control file sets `superuser = true` and `trusted = false`, so a superuser must create it; it declares no prerequisite extension or preload requirement. The accepted subtype registry is compiled into the extension. Upstream retains some deprecated media types with guidance toward replacements, so application validation and extension upgrades should be reviewed together.
