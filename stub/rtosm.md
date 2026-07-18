## Usage

Sources:

- [Official extension control file](https://github.com/rtosm/rtosm/blob/fb8fbec1cfb127f6ea07abc7db3f79b9a2b2a7f7/rtosm.control)
- [Official upstream documentation](https://github.com/rtosm/rtosm/blob/fb8fbec1cfb127f6ea07abc7db3f79b9a2b2a7f7/README.md)

`rtosm` — Real-time simplification of spatial objects in an OpenStreetMap API database.

The reviewed catalog snapshot records version `0.1`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `intarray`.

```sql
CREATE EXTENSION "rtosm";
SELECT extversion
FROM pg_extension
WHERE extname = 'rtosm';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
