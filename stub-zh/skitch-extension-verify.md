## 用法

来源：

- [官方扩展控制文件](https://github.com/airpage-app/pg-utils/blob/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/packages/verify/skitch-extension-verify.control)

`skitch-extension-verify` — 用于检查 PostgreSQL 对象、角色、权限、策略及其他模式属性的 PL/pgSQL 目录验证函数扩展。

已复核目录快照记录的版本为 `0.0.7`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`, `skitch-extension-utils`, `uuid-ossp`。

```sql
CREATE EXTENSION "skitch-extension-verify";
SELECT extversion
FROM pg_extension
WHERE extname = 'skitch-extension-verify';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
