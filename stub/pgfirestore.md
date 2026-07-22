## Usage

Sources:

- [Official README](https://github.com/xwkuang5/pgfirestore/blob/bc524d1b5200218b973a157361d74fc55031c907/README.md)
- [Official control file](https://github.com/xwkuang5/pgfirestore/blob/bc524d1b5200218b973a157361d74fc55031c907/pgfirestore.control)

`pgfirestore` is an experimental PostgreSQL query engine that models part of Firestore's hierarchical document and value semantics. Documents live in one `fs_documents` table as custom `fsvalue` references and maps. Upstream describes it as an exploration with no performance expectations, and the project is abandoned.

### Core Workflow

Create typed references and maps, insert them into the document table, then query a collection group with Firestore-style comparison operators:

```sql
CREATE EXTENSION pgfirestore;

INSERT INTO fs_documents (reference, properties)
VALUES (
  fs_reference('/users/1'),
  fs_map_from_entries(ARRAY['name'], ARRAY[fs_string('Ada')])
);

SELECT reference, properties
FROM fs_collection_group('users')
WHERE properties -> 'name' #= fs_string('Ada');
```

`fs_collection(parent, collection_id)` returns one collection below a reference; `fs_collection_group(collection_id)` finds collections with that ID anywhere below the database root.

### Types, Constructors, and Operators

- `fsvalue`: CBOR-backed value type with a JSON textual representation and Firestore cross-type ordering.
- `fs_null()`, `fs_nan()`, `fs_boolean(bool)`, `fs_number_from_integer(integer)`, and `fs_number_from_double(double precision)`.
- `fs_string(text)`, `fs_reference(text)`, `fs_array(fsvalue[])`, and `fs_map_from_entries(text[], fsvalue[])`.
- `#<`, `#>`, `#<=`, `#>=`, `#=`, and `#!=`: same-type Firestore query comparisons; ordinary PostgreSQL comparison operators implement cross-type ordering.
- `->`: extracts a nested document property.

### Caveats

`fs_documents` has only its primary key and no secondary indexes, so collection scans and property filters are not designed for scale. Firestore security rules are not implemented, and the published code leaves Date and GeoPoint values unimplemented despite showing their prospective external representation.

This is not a Firestore server, synchronization layer, or compatibility guarantee. Validate every supported type and operator against the application's required semantics, enforce authorization separately, and use only disposable experimental databases.
