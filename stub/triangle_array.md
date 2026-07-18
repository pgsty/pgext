## Usage

Sources:

- [Project README](https://github.com/evetion/trianglearray/blob/75b6e0ec507f6bdfd444efa008efdb328c933849/README.md)
- [Extension control file](https://github.com/evetion/trianglearray/blob/75b6e0ec507f6bdfd444efa008efdb328c933849/triangle_array.control)
- [Version 0.1 SQL API](https://github.com/evetion/trianglearray/blob/75b6e0ec507f6bdfd444efa008efdb328c933849/triangle_array--0.1.sql)
- [Native decoder implementation](https://github.com/evetion/trianglearray/blob/75b6e0ec507f6bdfd444efa008efdb328c933849/triangle_array.c)

`triangle_array` 0.1 is an abandoned research prototype for decoding a custom triangle-array byte layout into WKB. Its intended SQL functions are `trianglexyz()`, `trianglez_bytea()`, and `tinz_bytea()`, with PostGIS used separately to interpret the returned geometry bytes.

### Current source is not deployable

Only inspect an existing test installation; do not assume that the repository builds:

```sql
SELECT extversion
FROM pg_extension
WHERE extname = 'triangle_array';
```

At the reviewed head, the C source declares and defines two incompatible `unstitch()` functions as though C supported overloading, includes a missing declaration terminator, and has mismatched `Link` and `Linkq` pointers. These are compile-time defects. The README also says that a source table name is hard-coded.

### Untrusted-byte and memory-safety boundary

The native functions cast caller-supplied `bytea` payloads directly to point and triangle structures and index them without validating lengths, counts, offsets, or vertex references. Malformed or merely mismatched data can read outside the input buffer and crash a PostgreSQL backend. Cross-patch lookup executes hard-coded SPI SQL against `triangle_array`, and the generated WKB uses host endianness and historical geometry type choices.

Do not expose these functions to untrusted input or load an unrepaired binary in production. Reimplementation requires a documented portable wire format, bounds and overflow checks, explicit byte order, fuzzing and sanitizers, removal of hard-coded relations, current PostGIS conversion tests, and corruption-safe errors. Prefer maintained PostGIS TIN or geometry storage instead of preserving this prototype's private binary layout.
