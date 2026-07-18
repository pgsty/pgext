## Usage

Sources:

- [Extension control file](https://github.com/x4m/ags/blob/ceaf4c375e5142db193eae616e10d6704ad28c7e/ags.control)
- [Version 1.0 install SQL](https://github.com/x4m/ags/blob/ceaf4c375e5142db193eae616e10d6704ad28c7e/ags--1.0.sql)
- [Access-method implementation](https://github.com/x4m/ags/blob/ceaf4c375e5142db193eae616e10d6704ad28c7e/gist.c)

ags installs an experimental index access method derived from PostgreSQL GiST code. Version 1.0 exposes the access method named ags and one default operator class, ags_cube_ops, for the cube type. It is intended for cube overlap, containment, equality, and distance searches.

### Basic use

The extension does not declare its cube dependency, so install cube explicitly first:

```sql
CREATE EXTENSION cube;
CREATE EXTENSION ags;

CREATE TABLE ags_demo (
    id bigint GENERATED ALWAYS AS IDENTITY,
    point cube NOT NULL
);

INSERT INTO ags_demo (point)
VALUES (cube(ARRAY[0.0, 0.0])),
       (cube(ARRAY[1.0, 2.0])),
       (cube(ARRAY[5.0, 5.0]));

CREATE INDEX ags_demo_point_idx ON ags_demo USING ags (point);

SELECT id, point
FROM ags_demo
ORDER BY point <-> cube(ARRAY[0.0, 0.0])
LIMIT 2;
```

The operator class also registers cube overlap and containment operators. Verify plans with EXPLAIN on representative data before relying on the index.

### Caveats

- The repository has no README, release history, license file, or PostgreSQL compatibility matrix.
- Much of the implementation is a private copy of PostgreSQL GiST internals from 2019. Those internal APIs are major-version-sensitive, so build and regression-test ags separately for every PostgreSQL major version.
- Version 1.0 supplies only the cube operator class, does not declare cube as a control-file dependency, and disables parallel index scans.
- Treat the extension as research code until its correctness, crash safety, recovery behavior, and upgrade path have been validated for the target server.
