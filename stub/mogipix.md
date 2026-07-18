## Usage

Sources:

- [Pinned upstream README](https://gitlab.com/marjorielptt/mogipix/-/blob/6e134fe704b3a94a833d9ead1b2ae540b305a8d1/README.md)
- [Pinned Cargo metadata](https://gitlab.com/marjorielptt/mogipix/-/blob/6e134fe704b3a94a833d9ead1b2ae540b305a8d1/Cargo.toml)
- [Pinned HEALPix SQL interface](https://gitlab.com/marjorielptt/mogipix/-/blob/6e134fe704b3a94a833d9ead1b2ae540b305a8d1/src/lib.rs)
- [Pinned MOC implementation](https://gitlab.com/marjorielptt/mogipix/-/blob/6e134fe704b3a94a833d9ead1b2ae540b305a8d1/src/moc.rs)

mogipix version 0.0.1 exposes Rust CDS HEALPix, MOC, BMOC, and sky-region operations inside PostgreSQL. Its SQL functions use the mgx_ prefix and support cell hashing, hierarchy and neighborhood operations, MOC construction and set operations, coverage tests, and range forms suitable for indexed spatial filtering.

### Examples

The hash function accepts longitude and latitude in degrees and converts them to radians internally. The range-MOC helpers accept PostgreSQL integer ranges:

```sql
CREATE EXTENSION mogipix;

SELECT mgx_hash(13, 36.80105218, 56.78028536);

SELECT mgx_moc_to_ascii(
    mgx_create_range_moc_psql(
        29,
        ARRAY[int8range(100, 200), int8range(300, 400)]
    )
);
```

The extension also defines PostgreSQL types and operators for MOC and BMOC values. For table searches, upstream demonstrates combining a functional HEALPix index or integer multirange containment with mgx_ conversion and coverage functions.

### Validation and compatibility

The reviewed build uses pgrx 0.15.0, defaults to PostgreSQL 17, and declares build features for PostgreSQL 13 through 17. Installation is superuser-only. Version 0.0.1 is still an early interface, so test generated extension SQL and every PostgreSQL major version separately.

Depth, coordinate, hash, and range arguments cross signed PostgreSQL types into unsigned Rust library types. Some implementation paths use unchecked casts, unsafe vector representation conversion, or unchecked BMOC construction. Validate astronomical coordinate units and legal HEALPix depth/range bounds before calling them, and compare boundary and empty-set results with the underlying CDS libraries. The current GitLab project was active in 2025 and is not identified as archived.
