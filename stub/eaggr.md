## Usage

Sources:

- [Official upstream README](https://github.com/riskaware-ltd/open-eaggr/blob/c0d6d3b1091aaca10a2baa49b018687f198b7ce7/README.md)
- [Official extension control file (eaggr.control)](https://github.com/riskaware-ltd/open-eaggr/blob/c0d6d3b1091aaca10a2baa49b018687f198b7ce7/EAGGRPostgres/eaggr.control)
- [Official extension SQL (eaggr--2.0.sql)](https://github.com/riskaware-ltd/open-eaggr/blob/c0d6d3b1091aaca10a2baa49b018687f198b7ce7/EAGGRPostgres/eaggr--2.0.sql)

`eaggr` — The OpenEAGGR software library is an implementation of a Discrete Global Grid System (DGGS) which models the Earth's surface as a network of equal area cells. Use it for the corresponding spatial data or geospatial workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION eaggr;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `EAGGR_CellGeometry(text, text)` is an extension function and returns `text`.
- `EAGGR_CellToPoint(text, text)` is an extension function and returns `text`.
- `EAGGR_GetBoundingCell(text[], text)` is an extension function and returns `text`.
- `EAGGR_ShapeComparison(text, text, text, text)` is an extension function.
- `EAGGR_ToCellArray(text)` is an extension function and returns `text[]`.
- `EAGGR_ToCells(text, double precision, text)` is an extension function and returns `text`.
- `EAGGR_Version()` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `2.0`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
