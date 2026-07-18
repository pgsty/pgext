## 用法

来源：

- [官方扩展控制文件](https://github.com/mreithub/pg_recall/blob/87126883d6272a324b259b83f96436a53980dbd1/recall.control)

`recall` — recall：使用 PL/pgSQL 触发器为受管理表维护历史 _log 表，支持按时间查看行级变更。

已复核目录快照记录的版本为 `0.9.6`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`btree_gist`, `plpgsql`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "recall";
SELECT extversion
FROM pg_extension
WHERE extname = 'recall';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
