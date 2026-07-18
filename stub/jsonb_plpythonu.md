## Usage

Sources:

- [Upstream README](https://github.com/postgrespro/jsonb_plpython/blob/dff6c22394f97ed0618ba2c8e350ad2e1779ca64/README.md)
- [Extension control file](https://github.com/postgrespro/jsonb_plpython/blob/dff6c22394f97ed0618ba2c8e350ad2e1779ca64/jsonb_plpythonu.control)
- [SQL installation script](https://github.com/postgrespro/jsonb_plpython/blob/dff6c22394f97ed0618ba2c8e350ad2e1779ca64/jsonb_plpythonu--1.0.sql)
- [PL/Python 3 regression example](https://github.com/postgrespro/jsonb_plpython/blob/dff6c22394f97ed0618ba2c8e350ad2e1779ca64/sql/jsonb_plpython3u.sql)

`jsonb_plpythonu` version `1.0` defines a PostgreSQL transform between `jsonb` and PL/Python values. Functions that explicitly request `TRANSFORM FOR TYPE jsonb` receive Python dictionaries, lists, strings, numbers, booleans, or `None` instead of the default textual representation and can return those values as `jsonb`.

### Transform use

```sql
CREATE EXTENSION jsonb_plpythonu CASCADE;
CREATE FUNCTION jsonb_key_count(value jsonb) RETURNS integer
LANGUAGE plpythonu
TRANSFORM FOR TYPE jsonb
AS $$
return len(value)
$$;
SELECT jsonb_key_count('{"a":1,"b":2}'::jsonb);
```

This exact extension depends on the legacy unversioned `plpythonu` language and its control/module metadata references the Python-2-era library name. Modern PostgreSQL installations generally use the sibling `jsonb_plpython3u` variant instead. PL/Python is untrusted, so only superusers can create its functions. The repository has not changed since 2017; verify Python/PostgreSQL ABI compatibility and numeric/container round trips, and do not load the legacy variant merely to run current Python code.
