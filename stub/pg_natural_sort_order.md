## Usage

Sources:

- [Official README](https://github.com/Zeleo/pg_natural_sort_order/blob/16d2ff812aabfaab6b83056ffbf5de200c69a5b6/README.md)
- [Official control file](https://github.com/Zeleo/pg_natural_sort_order/blob/16d2ff812aabfaab6b83056ffbf5de200c69a5b6/pg_natural_sort_order.control)
- [Official PGXN release](https://pgxn.org/dist/pg_natural_sort_order/1.0.1/)

`pg_natural_sort_order` maps embedded digit runs to fixed-width text so ordinary PostgreSQL ordering and B-tree indexes produce human-friendly results such as `item2` before `item10`. Version `1.0.1` was tested upstream on PostgreSQL 9.x; rebuild and validate it before using it on a current server.

### Core Workflow

Create an expression index using a normalization width large enough for every numeric run in the data:

```sql
CREATE EXTENSION pg_natural_sort_order;

CREATE INDEX products_name_natural_idx
  ON products (natural_sort_order(name, 75));

SELECT id, name
FROM products
ORDER BY natural_sort_order(name, 75);
```

The query expression must match the indexed expression for the planner to use the index directly.

### Function Behavior

`natural_sort_order(text, integer)` replaces each run of decimal digits with a zero-padded representation of the requested width. The upstream range for the width is `1` through `150`, with `75` used in its examples. A larger width handles longer numbers but increases normalized values and index size.

### Operational Caveats

If an input contains a numeric run longer than the chosen width, the README warns that the run is split and ordering will probably be wrong. Choose the width from actual maximum data, enforce that bound if correctness matters, and rebuild the index if the expression or width changes. The mapping is a textual normalization, not locale-aware collation and not numeric parsing of signs, decimals, or versions. The project has no demonstrated support beyond the old PostgreSQL 9.x matrix, so compare results against representative Unicode and punctuation data on the target major version.
