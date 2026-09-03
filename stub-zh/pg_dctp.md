## 用法

来源：

- [PGXN 0.0.1 README](https://pgxn.org/dist/pg_dctp/0.0.1/README.html)
- [PGXS Makefile](https://api.pgxn.org/src/pg_dctp/pg_dctp-0.0.1/Makefile)
- [C 模块源码](https://api.pgxn.org/src/pg_dctp/pg_dctp-0.0.1/pg_dctp.c)
- [PostgreSQL 许可证](https://api.pgxn.org/src/pg_dctp/pg_dctp-0.0.1/LICENSE)

`pg_dctp` 是无 SQL 对象的 PostgreSQL C 模块，用于拒绝 `CREATE ROLE`、`CREATE USER`、`ALTER ROLE` 与 `ALTER USER` 中的明文密码字面量。它不安装 SQL 对象，也没有控制文件，因此完全通过服务器预加载来启用。

### 启用模块

```conf
shared_preload_libraries = 'pg_dctp'
```

修改 `shared_preload_libraries` 后需要重启 PostgreSQL。模块加载后，任何内嵌明文密码的命令都会被拒绝。应使用 `psql` 的 `\password`、`createuser -P`，或由其他客户端发送预先计算的 SCRAM 校验值，而不是把密码字面量放进 SQL。

### 范围与边界

该模块只处理一种日志风险：角色 DDL 语句中嵌入的明文密码可能进入语句日志。在报告拒绝时，`pg_dctp` 会临时调整消息处理，避免错误日志再次回显密码。

它不强制密码长度、复杂度、复用、过期、认证方式或 `pg_hba.conf` 策略，也无法保护通过其他 SQL、客户端日志、shell 历史、监控或网络捕获泄露的秘密。全局启用前应测试角色创建、密码轮换、备份恢复工具与自动化流程。

上游明确把 `pg_dctp` 称为模块而不是 SQL 扩展，并报告已在 PostgreSQL 14 至 18 上验证。0.0.1 是短暂存在的 `disable_set_password` 发行包重命名后的后继版本；应使用当前模块名，不要同时预加载两个名称。
