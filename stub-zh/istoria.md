## 用法

来源：

- [官方扩展控制文件](https://github.com/KiriakosGeorgiou/istoria/blob/e1a21fc8e6d262c1990d4ad7fbd80fb1a75d981f/istoria.control)
- [官方上游文档](https://github.com/KiriakosGeorgiou/istoria/blob/e1a21fc8e6d262c1990d4ad7fbd80fb1a75d981f/README)

`istoria` — istoria：使用 PL/pgSQL 管理表历史，支持非线性撤销与重做。

已复核目录快照记录的版本为 `1.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "istoria";
SELECT extversion
FROM pg_extension
WHERE extname = 'istoria';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
