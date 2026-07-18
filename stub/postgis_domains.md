## Usage

Sources:

- [Official extension control file](https://github.com/cahutton/postgis_domains/blob/cb9290194eebb0b63c85c7e59c406e5c0049f48e/postgis_domains.control)
- [Official upstream documentation](https://github.com/cahutton/postgis_domains/blob/cb9290194eebb0b63c85c7e59c406e5c0049f48e/README.md)

`postgis_domains` — Define reusable PostGIS geometry and geography domains constrained to valid, simple SRID 4326 values

The reviewed catalog snapshot records version `1.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `postgis`.
The curated compatibility set is `10`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "postgis_domains";
SELECT extversion
FROM pg_extension
WHERE extname = 'postgis_domains';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
