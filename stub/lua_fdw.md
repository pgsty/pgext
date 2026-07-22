## Usage

Sources:

- [Official README and Lua API](https://github.com/seanpringle/lua_fdw/blob/5577a0d886b7a07cff70b7825294087c14e8272e/README.md)
- [Extension SQL](https://github.com/seanpringle/lua_fdw/blob/5577a0d886b7a07cff70b7825294087c14e8272e/sql/lua_fdw--0.0.1.sql)
- [FDW implementation](https://github.com/seanpringle/lua_fdw/blob/5577a0d886b7a07cff70b7825294087c14e8272e/src/lua_fdw.c)

`lua_fdw` version `0.0.1` lets a server-side Lua script implement a read-only PostgreSQL foreign table. Use it as an adapter for data sources that can be reached from Lua, with the important boundary that the script executes inside a PostgreSQL backend.

### Core Workflow

A minimal script returns one row:

```lua
function ScanStart()
  done = false
end

function ScanIterate()
  if not done then
    done = true
    return { data = "hello world" }
  end
end
```

Point a foreign table at the server-side file:

```sql
CREATE EXTENSION lua_fdw;
CREATE SERVER lua_srv FOREIGN DATA WRAPPER lua_fdw;

CREATE FOREIGN TABLE lua_test (data text)
SERVER lua_srv
OPTIONS (script '/srv/postgresql/lua/hello_world.lua');

SELECT * FROM lua_test;
```

### Lua API

Planning callbacks are `EstimateRowCount()`, `EstimateRowWidth()`, `EstimateStartupCost()`, and `EstimateTotalCost()`. Scan callbacks are `ScanStart()`, `ScanIterate()`, `ScanRestart()`, `ScanEnd()`, and `ScanExplain()`; missing callbacks are skipped.

The global `fdw` table exposes `fdw.table`, `fdw.columns`, `fdw.clauses`, and `fdw.ereport()`. Simple top-level WHERE clauses appear in `fdw.clauses`, but the Lua script must apply any remote filtering itself.

### Options and Security

`script` selects the Lua file. `inject` executes an additional Lua fragment after loading it; `lua_path` and `lua_cpath` extend Lua module search paths. These are code-execution capabilities, not data-only options. Lua code runs with the PostgreSQL backend's operating-system access, and `lua_cpath` can load native modules. Restrict ownership and FDW/server usage to trusted administrators, keep scripts immutable to untrusted users, and do not treat the runtime as a sandbox. The FDW supplies scan callbacks only and does not implement SQL writes.
