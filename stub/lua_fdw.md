## Usage

Sources:

- [Official extension control file](https://github.com/seanpringle/lua_fdw/blob/5577a0d886b7a07cff70b7825294087c14e8272e/lua_fdw.control)
- [Official upstream documentation](https://github.com/seanpringle/lua_fdw/blob/5577a0d886b7a07cff70b7825294087c14e8272e/README.md)

`lua_fdw` — Lua Foreign Data Wrapper

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "lua_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'lua_fdw';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
