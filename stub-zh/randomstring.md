## 用法

来源：

- [官方扩展控制文件](https://github.com/tvondra/randomstring/blob/ae5259706484d66eb4dac63aaaea6ba745f7ffcd/randomstring.control)

`randomstring` — randomstring：生成伪随机文本与 bytea 值

已复核目录快照记录的版本为 `1.0.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "randomstring";
SELECT extversion
FROM pg_extension
WHERE extname = 'randomstring';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
