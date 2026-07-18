## 用法

来源：

- [官方扩展控制文件](https://github.com/yvazquezo/check_orapg/blob/d150ae5dc822c4e8f0eddf8a29e7b3d7055e562f/check_orapg.control)

`check_orapg` — 通过比较数据库对象、权限、属性、注释与行数，验证 Oracle 到 PostgreSQL 或 EDB 的迁移完整性。

已复核目录快照记录的版本为 `3.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`oracle_fdw`, `plpgsql`。
整理后的兼容版本集合为 `11,12,13`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "check_orapg";
SELECT extversion
FROM pg_extension
WHERE extname = 'check_orapg';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
