## 用法

来源：

- [官方扩展控制文件](https://database.dev/langchain/hybrid_search)
- [官方上游 README](https://github.com/supabase/dbdev/blob/8b3b966d97f87027a36c052e679ebb45f7e25736/README.md)

`hybrid_search` — 为 LangChain 与 Supabase 提供结合 pgvector 和全文检索的混合搜索函数。

已复核目录快照记录的版本为 `1.1.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`, `vector`。

```sql
CREATE EXTENSION "hybrid_search";
SELECT extversion
FROM pg_extension
WHERE extname = 'hybrid_search';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
