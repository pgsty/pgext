## 用法

来源：

- [官方扩展控制文件](https://github.com/mikeizbicki/chajda/blob/a9912d87b6f9bf8dc6a92ace5395ed5f97b3df51/pspacy.control)
- [官方上游 README](https://github.com/mikeizbicki/chajda/blob/a9912d87b6f9bf8dc6a92ace5395ed5f97b3df51/README.md)

`pspacy` — pspacy 通过 PL/Python 调用 spaCy，提供多语言词形还原、tsvector 生成和 tsquery 生成功能。

已复核目录快照记录的版本为 `1.0`、类型为 `puresql`、实现语言为 `Python`。
应先安装并验证声明的扩展依赖：`plpython3u`。
整理后的兼容版本集合为 `12,13`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pspacy";
SELECT extversion
FROM pg_extension
WHERE extname = 'pspacy';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
