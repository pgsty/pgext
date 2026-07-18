## 用法

来源：

- [官方扩展控制文件](https://github.com/mansueli/tle/blob/3825e5a6dd3b3d0e48eea1a01eba89655788b746/email_guard/email_guard.control)
- [官方上游文档](https://database.dev/mansueli/email_guard)
- [官方上游 README](https://github.com/mansueli/tle/blob/3825e5a6dd3b3d0e48eea1a01eba89655788b746/README.md)

`email_guard` — 用于 Supabase Auth 注册流程的邮箱防护扩展，可拦截一次性邮箱域名并规范化 Gmail 地址。

已复核目录快照记录的版本为 `0.35.0`、类型为 `puresql`、实现语言为 `SQL`。

```sql
CREATE EXTENSION "email_guard";
SELECT extversion
FROM pg_extension
WHERE extname = 'email_guard';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
