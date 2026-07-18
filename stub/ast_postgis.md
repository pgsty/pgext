## Usage

Sources:

- [AST-PostGIS README at the reviewed commit](https://github.com/lizardoluis/ast_postgis/blob/a7b71d3b350cd7c12dfb574ea886bd14c870ce6c/README.md)
- [ast_postgis.control at the reviewed commit](https://github.com/lizardoluis/ast_postgis/blob/a7b71d3b350cd7c12dfb574ea886bd14c870ce6c/ast_postgis.control)

`ast_postgis` adds richer spatial types and spatial-integrity checks on top of `postgis`. Its types include `ast_point`, `ast_polygon`, and `ast_uniline`, which wrap PostGIS geometry or raster representations with constraints suited to geo-object and geo-field models.

The extension also provides trigger procedures such as `ast_spatialrelationship`, `ast_arcnodenetwork`, and `ast_aggregation`. Validation functions such as `ast_isTopologicalRelationshipValid` can check existing data before constraints are enforced and record inconsistencies in `ast_validation_log`.

### Basic Setup

```sql
CREATE EXTENSION IF NOT EXISTS postgis;
CREATE EXTENSION ast_postgis;

CREATE TABLE landmarks (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL,
  geom ast_point
);
```

Use the advanced types as table-column types, then add the documented trigger procedures for the spatial relationships your model must enforce.

### Caveats

- The control file declares `postgis` as a dependency, so it must be available in the database.
- Upstream currently identifies version 1.1.1 and documents testing with PostgreSQL 15. Test the exact PostgreSQL/PostGIS combination used by your deployment.
- Constraint procedures accept table and column names. Reproduce the upstream trigger signatures carefully and validate existing rows before enabling enforcement on production data.
