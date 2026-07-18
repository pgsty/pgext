## 用法

来源：

- [最后已知的官方上游仓库；当前不可用](https://github.com/dmtkfs/pg-tamper-log)

当前来源状态：官方仓库与 owner 均返回 HTTP 404。对 GitHub 仓库、fork、代码与 commit 搜索，raw、codeload、PGXN 和 Software Heritage 的检查都未找到当前官方来源或归档。

冻结目录历史上把 `pg_tamperlog` 版本 `1.1` 描述为使用哈希链的纯 SQL 防篡改审计日志，并依赖 `pg_tamperlog_rust` 与 `pgcrypto`。这些仅是历史线索，仓库删除后无法重新验证。

```sql
-- Run only if matching, independently verified artifacts are recovered.
CREATE EXTENSION "pg_tamperlog";
```

不要从未经验证的 fork 获取替代二进制或 SQL。任何使用前都必须恢复可信的签名或内容寻址源码，审查全部安装对象与权限，确认哈希链威胁模型，并独立测试删除、更新、并发、密钥管理、备份、恢复和校验行为。
