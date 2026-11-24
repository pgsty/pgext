## Usage

Sources:

- [README 0.2.0](https://github.com/Vinnu124/pg_anyFK/blob/v0.2.0/README.md)
- [Control](https://github.com/Vinnu124/pg_anyFK/blob/v0.2.0/ext/pg_anyfk.control)
- [Cargo](https://github.com/Vinnu124/pg_anyFK/blob/v0.2.0/ext/Cargo.toml)
- [Release with generated SQL](https://github.com/Vinnu124/pg_anyFK/releases/tag/v0.2.0)

`pg_anyfk` 0.2.0 enforces references from individual array elements to a parent key using triggers. It is intended for existing array-based schemas; a junction table remains an ordinary alternative.

### Core Workflow

Upstream supports PostgreSQL 15–18. A superuser installs this untrusted, non-relocatable extension; no preload or restart is required.

```sql
CREATE EXTENSION pg_anyfk;
CREATE TABLE tags (id bigint PRIMARY KEY);
CREATE TABLE posts (
    id bigserial PRIMARY KEY,
    tag_ids bigint[] NOT NULL DEFAULT '{}'
);
CREATE INDEX ON posts USING gin (tag_ids);
SELECT anyfk.add('posts(tag_ids)', 'tags(id)');
SELECT * FROM anyfk.constraints;
```

The caller must own both tables. The array element type must match the parent column type, and the parent needs a single-column unique index. A GIN index avoids scanning all child rows for every parent deletion.

### Validation and Deletion

`anyfk.add` validates existing data by default. Inspect exceptions before adoption:

```sql
SELECT * FROM anyfk.violations('posts(tag_ids)', 'tags(id)');
```

The options include `validate`, `is_deferrable`, `conname`, and `on_delete`. The default `restrict` rejects deleting referenced parent rows. Select `scrub` when deletion should remove that element from child arrays; the deleting role then needs UPDATE on the child table. `scrub` does not permit updating referenced parent keys.

`is_deferrable => true` defers checking to transaction completion. Remove a registered constraint with `anyfk.drop_constraint`, using its name from `anyfk.constraints`.

### Boundaries

Null arrays, empty arrays, null elements and duplicate elements are accepted. Parent truncation is rejected while nonempty references remain, but truncating parent and child together is supported. Row-key locks protect checks against concurrent parent changes.

These are triggers, not native foreign-key entries in `pg_constraint`; disabling triggers disables enforcement. The extension tracks table identities and uses event triggers for renames and dependent-object cleanup.
