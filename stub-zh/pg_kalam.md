## 用法

来源：

- [已复核 commit 的 PostgreSQL 扩展 README](https://github.com/kalamdb/KalamDB/blob/e59521f548e970902d3a71b8253109ebd1fd5731/pg/README.md)
- [扩展 control 文件](https://github.com/kalamdb/KalamDB/blob/e59521f548e970902d3a71b8253109ebd1fd5731/pg/pg_kalam.control)
- [Cargo 特性矩阵](https://github.com/kalamdb/KalamDB/blob/e59521f548e970902d3a71b8253109ebd1fd5731/pg/Cargo.toml)
- [KalamDB 项目站点](https://kalamdb.org)

`pg_kalam` 通过 gRPC 将 PostgreSQL 连接到单独运行的 KalamDB 服务器。它注册 `pg_kalam` 外部数据包装器，也支持通过 `kalamdb` 表访问方法创建由 KalamDB 支持的表。

```sql
CREATE EXTENSION pg_kalam;
CREATE SERVER kalam_server
  FOREIGN DATA WRAPPER pg_kalam
  OPTIONS (
    host '127.0.0.1', port '2910',
    auth_mode 'account_login',
    login_user 'pg_bridge', login_password 'replace-me'
  );
SELECT kalam_version(), kalam_compiled_mode();
```

远端 KalamDB 服务必须事先运行。当前文档还提供旧式 `static_header` 认证模式；应优先使用账号登录，并保护包含凭据的外部服务器选项。PostgreSQL `JSON` 与 `JSONB` 映射到 KalamDB `JSON`，远端 `FILE` 类型则镜像为 `JSONB`。

目录版本为 0.5.5-rc.1，项目将该扩展标记为 preview。它通过 pgrx 特性构建支持 PostgreSQL 13–18，但远端 DDL、类型映射、JSON 运算符、认证、网络故障与事务语义仍需端到端验证后才能用于生产。
