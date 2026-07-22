## 用法

来源：

- [官方上游仓库的归档快照](https://github.pkg.st/dmtkfs/pg-tamper-log)
- [最后已知的官方仓库地址](https://github.com/dmtkfs/pg-tamper-log)

`pg_tamperlog_rust` 是 `pg_tamperlog` 1.1 附带的可选 pgrx 辅助扩展。它为 SQL 审计链扩展提供更快的 SHA-256 计算，本身不会创建或保护审计日志。

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pg_tamperlog_rust;
CREATE EXTENSION pg_tamperlog;
```

归档的上游说明使用 `cargo pgrx` 针对特定 PostgreSQL 安装构建辅助扩展。原生 pgrx 模块绑定目标 PostgreSQL 大版本及 Rust/C ABI：应针对精确服务器做可复现构建，再用独立 SHA-256 实现逐字节核对空值、`NULL`、Unicode、大输入和并发输入的哈希结果。

哈希速度不会增强链的威胁模型。特权主体可以替换原生库或 SQL 包装器，崩溃或升级也可能改变运行行为。应限制函数与库安装，记录部署产物哈希，监控重新加载，并保留带签名的异地主机检查点。

官方 GitHub 仓库目前返回 HTTP 404；上述来源只是其 README 的归档渲染，不是完整可审计的源码归档。在恢复并审查准确的内容寻址源码之前，不要信任同名 crate 或二进制。
