## Usage

Sources:

- [Official extension control file](https://github.com/segasai/pg_coordconvert/blob/8acfae816cf373204918c8231eabae26c575dbc3/coordconv.control)
- [Official upstream documentation](https://github.com/segasai/pg_coordconvert/blob/8acfae816cf373204918c8231eabae26c575dbc3/README)

`coordconv` — Convert between IAU 1958 Galactic coordinates and J2000 FK5 equatorial coordinates

The reviewed catalog snapshot records version `1.0.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "coordconv";
SELECT extversion
FROM pg_extension
WHERE extname = 'coordconv';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
