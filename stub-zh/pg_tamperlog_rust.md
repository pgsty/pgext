## 用法

来源：

- [最后已知的官方上游仓库；当前不可用](https://github.com/dmtkfs/pg-tamper-log)

当前来源状态：官方仓库与 owner 均返回 HTTP 404。对 GitHub 仓库、fork、代码与 commit 搜索，raw、codeload、PGXN 和 Software Heritage 的检查都未找到当前官方来源或归档。

冻结目录历史上把 `pg_tamperlog_rust` 版本 `1.0.0` 描述为 `pg_tamperlog` 使用的 pgrx 辅助组件，用于加速 SHA-256 哈希链计算。这些仅是历史线索，仓库删除后无法重新验证。

```sql
-- Run only if matching, independently verified artifacts are recovered.
CREATE EXTENSION "pg_tamperlog_rust";
```

不要安装没有来源证明的同名 crate 或二进制。应先恢复并审计准确源码，验证 PostgreSQL 与 pgrx ABI 兼容性，复现构建，与独立实现核对哈希结果，并在把它用于审计链前测试 NULL、编码、并发、升级和崩溃行为。
