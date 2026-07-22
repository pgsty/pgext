## 用法

来源：

- [官方 README 与 Lua API](https://github.com/seanpringle/lua_fdw/blob/5577a0d886b7a07cff70b7825294087c14e8272e/README.md)
- [扩展 SQL](https://github.com/seanpringle/lua_fdw/blob/5577a0d886b7a07cff70b7825294087c14e8272e/sql/lua_fdw--0.0.1.sql)
- [FDW 实现](https://github.com/seanpringle/lua_fdw/blob/5577a0d886b7a07cff70b7825294087c14e8272e/src/lua_fdw.c)

`lua_fdw` 版本 `0.0.1` 允许服务端 Lua 脚本实现只读 PostgreSQL 外部表。它适合适配 Lua 可访问的数据源，但关键边界是脚本会在 PostgreSQL 后端进程内执行。

### 核心流程

下面的最小脚本返回一行：

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

让外部表指向服务端文件：

```sql
CREATE EXTENSION lua_fdw;
CREATE SERVER lua_srv FOREIGN DATA WRAPPER lua_fdw;

CREATE FOREIGN TABLE lua_test (data text)
SERVER lua_srv
OPTIONS (script '/srv/postgresql/lua/hello_world.lua');

SELECT * FROM lua_test;
```

### Lua API

规划回调包括 `EstimateRowCount()`、`EstimateRowWidth()`、`EstimateStartupCost()` 和 `EstimateTotalCost()`。扫描回调包括 `ScanStart()`、`ScanIterate()`、`ScanRestart()`、`ScanEnd()` 和 `ScanExplain()`；缺失的回调会被跳过。

全局 `fdw` 表暴露 `fdw.table`、`fdw.columns`、`fdw.clauses` 和 `fdw.ereport()`。简单的顶层 WHERE 条件会出现在 `fdw.clauses` 中，但 Lua 脚本必须自行执行远端过滤。

### 选项与安全

`script` 选择 Lua 文件。`inject` 在加载文件后执行额外 Lua 片段；`lua_path` 与 `lua_cpath` 扩展 Lua 模块搜索路径。这些是代码执行能力，而非纯数据选项。Lua 代码以 PostgreSQL 后端的操作系统访问权限运行，且 `lua_cpath` 能加载原生模块。应把所有权和 FDW/服务器使用权限制给可信管理员，确保不可信用户不能修改脚本，并且不要把该运行时视为沙箱。此 FDW 只提供扫描回调，不实现 SQL 写入。
