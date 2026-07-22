## Usage

Sources:

- [Extension control file](https://github.com/JonM0/hyperspace/blob/99ebb600ef1b5e606a54e79978a6824a33492d3f/hyperspace.control)
- [Type definitions](https://github.com/JonM0/hyperspace/blob/99ebb600ef1b5e606a54e79978a6824a33492d3f/sql/type.sql)
- [Operators and SP-GiST class](https://github.com/JonM0/hyperspace/blob/99ebb600ef1b5e606a54e79978a6824a33492d3f/sql/kd.sql)

`hyperspace` version `1.0` adds four-dimensional geometric types and indexes. Use it for points that need distance ordering or containment tests in a four-axis space.

### Core Workflow

```sql
CREATE EXTENSION hyperspace;
CREATE TABLE samples (p point4d);
CREATE INDEX samples_p_spgist ON samples USING spgist (p);

INSERT INTO samples VALUES ('(1,2,3,4)'), ('(5,6,7,8)');
SELECT p, p <-> '(0,0,0,0)'::point4d AS distance
FROM samples
WHERE p <@ '((-10,-10,-10,-10),(10,10,10,10))'::box4d
ORDER BY p <-> '(0,0,0,0)'::point4d
LIMIT 5;
```

The default `point4d_kd` SP-GiST operator class supports equality, box or circle containment, and nearest-neighbor ordering with `<->`.

### Objects and Caveats

The principal types are `point4d`, `box4d`, and `circle4d`. Constructors with the same names are available. Points support equality, addition, `sum`, and Euclidean distance. Containment uses `<@` with the commutator `@>`.

`point4d` also has a default lexical B-tree class and an alternate magnitude-based `point4d_abs_ops` class. Text input is `(x,y,z,w)` for a point, two points for a box, and a point plus radius for a circle. These are custom C types; validate binary compatibility and query plans on the exact PostgreSQL build before storing durable data.
