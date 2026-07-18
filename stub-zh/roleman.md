## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/roleman/)

`roleman` — 为 PostgreSQL 提供参数化角色与权限管理函数的扩展。

已复核目录快照记录的版本为 `0.3.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "roleman";
SELECT extversion
FROM pg_extension
WHERE extname = 'roleman';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
