## 用法

来源：

- [官方扩展控制文件](https://github.com/vibhorkum/edb_foreignkeyconstraintmanager/blob/50a89fb337426e810eecca4f3e4f1b00d7bb1023/edb_foreignkeyconstraintmanager.control)
- [官方上游文档](https://github.com/vibhorkum/edb_foreignkeyconstraintmanager/blob/50a89fb337426e810eecca4f3e4f1b00d7bb1023/README.md)

`edb_foreignkeyconstraintmanager` — 在 EDB Postgres Advanced Server 上管理分区表的触发器式外键行为

已复核目录快照记录的版本为 `1.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`refint`。

```sql
CREATE EXTENSION "edb_foreignkeyconstraintmanager";
SELECT extversion
FROM pg_extension
WHERE extname = 'edb_foreignkeyconstraintmanager';
```

该上游项目与 `EDB` 相关；应根据所链接来源核实当前支持、许可证、打包方式与部署边界。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
