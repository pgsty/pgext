## 用法

来源：

- [官方扩展控制文件](https://github.com/gciolli/pg2podg/blob/039dedf753e8c4aeb05564b4bf9082d80aa8b417/pg2podg.control)
- [官方上游文档](https://pgxn.org/dist/pg2podg/)
- [官方上游 README](https://github.com/gciolli/pg2podg/blob/039dedf753e8c4aeb05564b4bf9082d80aa8b417/README.md)

`pg2podg` — pg2podg：模拟/游戏类 PostgreSQL 扩展；提供两人、完全信息、确定性游戏的通用工具。

已复核目录快照记录的版本为 `0.1.3`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg2podg";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg2podg';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
