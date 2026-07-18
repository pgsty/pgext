## 用法

来源：

- [官方扩展控制文件](https://github.com/GeoffMontee/pg_llm_helper/blob/feae47af77a9cb735400b66b0b39f046d60cf924/pg_llm_helper.control)

`pg_llm_helper` — 在共享内存中捕获 PostgreSQL 错误，提供错误历史查询及可选的 pgai 大模型故障分析函数。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`plpgsql`, `vector`。
整理后的兼容版本集合为 `17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_llm_helper";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_llm_helper';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
