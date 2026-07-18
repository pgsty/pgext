## 用法

来源：

- [官方扩展控制文件](https://github.com/goldenhelix/tsf-fdw/blob/bf8d839bf447d3300cd371b9e059b612b9330923/tsf_fdw.control)

`tsf_fdw` — Golden Helix 提供的 C 外部数据包装器，用于把一个或多个 TSF 基因组数据文件作为 PostgreSQL 外表读取。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "tsf_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'tsf_fdw';
```

该上游项目与 `Golden Helix` 相关；应根据所链接来源核实当前支持、许可证、打包方式与部署边界。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
