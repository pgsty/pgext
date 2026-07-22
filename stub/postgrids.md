## Usage

Sources:

- [Official README](https://github.com/bis-brecon/postgrids/blob/ba0a85a344cfefc8e630c7428e930e6eb8df03aa/README.md)
- [OSGB SQL implementation](https://github.com/bis-brecon/postgrids/blob/ba0a85a344cfefc8e630c7428e930e6eb8df03aa/src/osgb.rs)
- [OSI SQL implementation](https://github.com/bis-brecon/postgrids/blob/ba0a85a344cfefc8e630c7428e930e6eb8df03aa/src/osi.rs)

`postgrids` version `0.0.0` provides PostgreSQL types and helper functions for British OSGB and Irish OSI national-grid references. It parses, validates, formats, and changes grid-reference precision; it does not convert references to PostGIS geometries or other coordinate systems.

### Core Workflow

```sql
CREATE EXTENSION postgrids;

SELECT osgb_from_string('SO892437');
SELECT osgb_from_eastings_northings(389200, 243700, 100);
SELECT osgb_is_valid('SO892437');
SELECT osgb_precision(osgb_from_string('SO892437'));
SELECT osgb_recalculate(osgb_from_string('SO892437'), 1000);
```

The equivalent Irish functions use the `osi_` prefix:

```sql
SELECT osi_from_string('O892437');
SELECT osi_is_valid('O892437');
SELECT osi_recalculate(osi_from_string('O892437'), 1000);
```

### Objects

The custom types are `OSGB` and `OSI`. Each family provides `*_from_string`, `*_from_eastings_northings`, `*_is_valid`, `*_precision`, and `*_recalculate`.

Precision is expressed in metres. Supported values are `1`, `10`, `100`, `1000`, `2000`, `10000`, and `100000`; unsupported values raise an error. Parsing constructors also raise on invalid references, while the `*_is_valid` helpers return a boolean.

### Maturity Boundary

The upstream README describes the project as very much work in progress and explicitly says conversion to geospatial primitives is not yet supported. Pin the extension build and verify accepted grid-reference forms before storing these custom types durably.
