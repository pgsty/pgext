## 用法

来源：

- [官方扩展控制文件](https://github.com/carped99/postfga/blob/92626b6120ab526acd00ff457cfc2071538836c6/postfga.control)
- [官方上游文档](https://github.com/carped99/postfga/blob/92626b6120ab526acd00ff457cfc2071538836c6/TESTING.md)

`postfga` — postfga 是一个需预加载的 PostgreSQL 扩展，通过 FDW、C 函数、gRPC 后台工作进程和共享内存缓存集成 OpenFGA 授权。

已复核目录快照记录的版本为 `1.0.0`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "postfga";
SELECT extversion
FROM pg_extension
WHERE extname = 'postfga';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
