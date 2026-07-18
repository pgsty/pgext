## 用法

来源：

- [官方上游文档](https://michelp.github.io/pgfsm/)

`pgfsm` — pgfsm：通用类 PostgreSQL 扩展；上游说明为“使用触发器和实用函数实施有限状态机约束”。

已复核目录快照记录的版本为 `0.0.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgfsm";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgfsm';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
