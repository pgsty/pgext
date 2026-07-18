## 用法

来源：

- [官方扩展控制文件](https://github.com/seanpringle/lua_fdw/blob/5577a0d886b7a07cff70b7825294087c14e8272e/lua_fdw.control)
- [官方上游文档](https://github.com/seanpringle/lua_fdw/blob/5577a0d886b7a07cff70b7825294087c14e8272e/README.md)

`lua_fdw` — 通过 Lua 脚本实现只读数据源访问的 PostgreSQL 外部数据包装器。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "lua_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'lua_fdw';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
