## Usage

Sources:

- [Official README](https://github.com/k2bd/pghex/blob/0afc072f18799e36071ed6a14e44478c84db7a63/README.md)
- [SQL API implementation](https://github.com/k2bd/pghex/blob/0afc072f18799e36071ed6a14e44478c84db7a63/src/lib.rs)
- [Hex-grid algorithms](https://github.com/k2bd/pghex/blob/0afc072f18799e36071ed6a14e44478c84db7a63/src/hex_alg.rs)

`pghex` adds a `hex` coordinate type and common algorithms for hexagonal grids. It is useful for board games, map cells, field-of-view calculations, and other applications that need neighbor, distance, line, ring, range, or spiral operations inside SQL.

### Core Workflow

Coordinates are written as two-axis values `[q, r]`; the implementation derives the third cube coordinate as `s = -q - r`.

```sql
CREATE EXTENSION pghex;

CREATE TABLE units (
    name text PRIMARY KEY,
    position hex NOT NULL,
    vision_range integer NOT NULL
);

INSERT INTO units VALUES
    ('Hero', '[0, 1]', 3),
    ('Goblin', '[5, 1]', 2);

SELECT name, tile
FROM units
CROSS JOIN LATERAL hexes_in_range(position, vision_range) AS tile;

SELECT hex_distance('[0, 1]'::hex, '[5, 1]'::hex);
SELECT * FROM linedraw('[0, 1]'::hex, '[5, 1]'::hex);
```

The `+`, `-`, and `=` operators add, subtract, and compare `hex` values.

### Main Functions

- `neighbors(hex)` and `diagonals(hex)` return the six adjacent or diagonal cells.
- `hex_distance(hex, hex)` returns cube-grid distance.
- `linedraw(hex, hex)` returns the cells along a line, including its endpoints.
- `hexes_in_range(hex, integer)` returns every cell up to the requested distance.
- `ring_path(hex, integer)` returns a single ring; `spiral_path(hex, integer)` returns the center and successive rings.

The path and neighborhood functions are set-returning functions, so use them in `FROM` or with `LATERAL` when applying them per table row.

### Caveats

The reviewed package reports version `0.0.0`; upstream describes development-environment installation and does not present a stable compatibility guarantee. Its Cargo features cover PostgreSQL `12` through `17`. The extension is non-relocatable and its control file requires superuser installation. Radius and distance arguments should be validated by the application, and large ranges grow quadratically in row count. Store one consistent coordinate convention across the application; no preload or restart is required.
