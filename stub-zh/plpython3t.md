## 用法

来源：

- [官方扩展控制文件](https://github.com/pykello/plpythont/blob/3b05e9ffbe405d7768bec8fc3a0764642ceff125/plpython3t.control)

`plpython3t` — plpython3t 是一个纯 SQL 包装扩展，把 PostgreSQL PL/Python 3 处理器注册为名为 plpython3t 的可信过程语言。

已复核目录快照记录的版本为 `1.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpython3u`。

```sql
CREATE EXTENSION "plpython3t";
SELECT extversion
FROM pg_extension
WHERE extname = 'plpython3t';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
