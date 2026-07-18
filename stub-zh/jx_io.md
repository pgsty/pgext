## 用法

来源：

- [官方扩展控制文件](https://github.com/asotolongo/jx_io/blob/4098d0db4d71596deed1e32a71c587bb1285c01a/jx_io.control)
- [官方 PGXN 分发页](https://pgxn.org/dist/jx_io/)
- [官方项目或服务商页面](http://anthonysotolongo.wordpress.com)

`jx_io` — jx_io：数据集成类 PostgreSQL 扩展；上游说明为“在 PostgreSQL 中导入和导出 JSON 与 XML”。

已复核目录快照记录的版本为 `0.1.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "jx_io";
SELECT extversion
FROM pg_extension
WHERE extname = 'jx_io';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
