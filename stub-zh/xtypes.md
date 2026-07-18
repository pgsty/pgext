## 用法

来源：

- [官方扩展控制文件](https://github.com/mkindahl/pg_xtypes/blob/eee55dfe3fb8d02baa36926692d1ccbcee930df6/xtypes.control)

`xtypes` — 提供人类可读二进制单位输入输出、算术、比较与取整的无符号 64 位字节大小类型。

已复核目录快照记录的版本为 `0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "xtypes";
SELECT extversion
FROM pg_extension
WHERE extname = 'xtypes';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
