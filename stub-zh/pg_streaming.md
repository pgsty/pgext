## 用法

来源：

- [官方扩展控制文件](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_streaming/pg_streaming.control)
- [官方上游文档](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/README.md)

`pg_streaming` — pg_streaming：在 PostgreSQL 内部运行的声明式流处理引擎。

已复核目录快照记录的版本为 `0.2.0`、类型为 `preload`、实现语言为 `Rust`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_streaming";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_streaming';
```

该上游项目与 `Matroid` 相关；应根据所链接来源核实当前支持、许可证、打包方式与部署边界。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
