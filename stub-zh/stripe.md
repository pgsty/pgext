## 用法

来源：

- [官方项目或服务商页面](https://stripe.com)

`stripe` — 由 Stainless 生成、用于从 PostgreSQL 调用 Stripe REST API 的实验性 SQL/PLPython SDK。

已复核目录快照记录的版本为 `0.0.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpython3u`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "stripe";
SELECT extversion
FROM pg_extension
WHERE extname = 'stripe';
```

该上游项目与 `Stainless` 相关；应根据所链接来源核实当前支持、许可证、打包方式与部署边界。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
