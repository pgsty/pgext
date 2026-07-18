## Usage

Sources:

- [Extension control file](https://github.com/wulczer/saio/blob/e6abbd9f0886e4067b8c42701a12b6e2589b641d/saio.control)
- [Upstream warning and configuration](https://github.com/wulczer/saio/blob/e6abbd9f0886e4067b8c42701a12b6e2589b641d/doc/saio.md)
- [Planner-hook implementation](https://github.com/wulczer/saio/blob/e6abbd9f0886e4067b8c42701a12b6e2589b641d/src/saio.c)
- [PGXN package metadata](https://github.com/wulczer/saio/blob/e6abbd9f0886e4067b8c42701a12b6e2589b641d/META.json)

`saio` version `0.0.1` is a 2011 experimental planner module that searches join orders with simulated annealing. It has no extension DDL file: after installing a compatible shared library, a session activates the planner hook with `LOAD 'saio'`. Once loaded, its `saio` switch defaults to on, while `saio_seed` and several temperature, equilibrium, move, and recomputation settings control the search.

### Inspect it only in an isolated legacy build

```sql
LOAD 'saio';
SET saio = off;
SET saio_seed = 0.5;
SET saio = on;

EXPLAIN
SELECT n.nspname, c.relname
FROM pg_namespace AS n
JOIN pg_class AS c ON c.relnamespace = n.oid;

SET saio = off;
```

The upstream documentation opens with an explicit warning not to use the module and describes its optimizer quality as very poor. Its code targets PostgreSQL internals from 2011, so modern server compatibility must not be assumed, and loading it changes planning for the whole backend session. Do not deploy it in production or load it into a session carrying valuable work. If historical research requires it, use a disposable server matching its era, keep `saio` disabled until a controlled query, capture plans and crashes, and terminate the session afterward.
