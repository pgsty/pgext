## 用法

来源：

- [官方上游来源](https://gitlab.com/1on.cz/nfiesta/nfiesta_sdesign)

`nfiesta_sdesign` — 用于采样设计数据建模与操作的 PostgreSQL 扩展。

已复核目录快照记录的版本为 `1.1.17`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`, `postgis`。
整理后的兼容版本集合为 `16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "nfiesta_sdesign";
SELECT extversion
FROM pg_extension
WHERE extname = 'nfiesta_sdesign';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
