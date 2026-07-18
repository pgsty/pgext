## 用法

来源：

- [官方上游文档](https://database.dev/basejump/supabase_test_helpers)
- [官方项目或服务商页面](https://usebasejump.com/blog/testing-on-supabase-with-pgtap)

`basejump-supabase_test_helpers` — 一组用于更方便测试 Supabase 项目的 pgTAP 辅助函数。

已复核目录快照记录的版本为 `0.0.6`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`pgtap`。

```sql
CREATE EXTENSION "basejump-supabase_test_helpers";
SELECT extversion
FROM pg_extension
WHERE extname = 'basejump-supabase_test_helpers';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
