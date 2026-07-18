## 用法

来源：

- [官方扩展控制文件](https://github.com/rongfengliang/nvl-pg-extension/blob/b2c153a2ff86e5fb2acb4b094448a8835700f031/nvlfunc.control)
- [官方上游文档](https://github.com/rongfengliang/nvl-pg-extension/blob/b2c153a2ff86e5fb2acb4b094448a8835700f031/README.md)

`nvlfunc` — 提供 Oracle 兼容 NVL 函数的 PostgreSQL 扩展。

已复核目录快照记录的版本为 `1.0`、类型为 `puresql`、实现语言为 `SQL`。

```sql
CREATE EXTENSION "nvlfunc";
SELECT extversion
FROM pg_extension
WHERE extname = 'nvlfunc';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
