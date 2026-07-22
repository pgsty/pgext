## 用法

来源：

- [已复核 commit 的 Aiven Extras README](https://github.com/aiven/aiven-extras/blob/ed2a4a676034b1ca3edb767d5805c482262d3f9b/README.md)
- [已复核 commit 的 control 模板](https://github.com/aiven/aiven-extras/blob/ed2a4a676034b1ca3edb767d5805c482262d3f9b/aiven_extras.control.in)
- [已复核 commit 的构建版本](https://github.com/aiven/aiven-extras/blob/ed2a4a676034b1ca3edb767d5805c482262d3f9b/Makefile)

`aiven_extras` 是由 Aiven 维护的一组窄权限包装接口，用于执行通常需要 PostgreSQL 高权限的操作。特权操作员安装后，符合条件的非超级用户可以管理部分逻辑复制对象、配置会话级 `auto_explain`、调整部分 pgaudit 设置，以及执行受控的 dblink 操作。它安装在固定的 `aiven_extras` schema 中，不会授予通用超级用户权限。

### 逻辑复制工作流

调用角色仍须拥有 `REPLICATION`。表名应带 schema 限定，并且只向确有需要的运维角色授予这些包装函数的执行权限。

```sql
CREATE EXTENSION aiven_extras;

SELECT aiven_extras.pg_create_publication(
  'pub1', 'INSERT,UPDATE', 'public.foo', 'public.bar'
);

SELECT * FROM aiven_extras.pg_list_all_subscriptions();
```

逻辑复制接口包括 `pg_create_publication`、`pg_create_publication_for_all_tables`、`pg_create_subscription`、`pg_drop_subscription`、`pg_alter_subscription_enable`、`pg_alter_subscription_disable`、`pg_alter_subscription_refresh_publication`、`pg_list_all_subscriptions` 和 `pg_stat_replication_list`。创建订阅仍会执行网络访问与复制槽操作；包装函数不会让这些操作变成事务性的。

### 会话诊断

`auto_explain_load` 为当前会话加载 `auto_explain`，随后可用 `set_auto_explain_*` 系列函数只修改该会话的设置。

```sql
SELECT aiven_extras.auto_explain_load();
SELECT aiven_extras.set_auto_explain_log_min_duration('2000');
SELECT aiven_extras.set_auto_explain_log_analyze('on');
```

其他公开辅助函数包括 `claim_public_schema_ownership`、`dblink_record_execute`、`dblink_slot_create_or_drop`、`explain_statement`、`session_replication_role`、`set_pgaudit_parameter` 和 `set_pgaudit_role_parameter`。

### 运维注意事项

- 当前复核的 Makefile 标识版本为 `1.1.20`，与目录一致。应固定并测试实际采用的服务商软件包或源码 commit。
- 发布与订阅包装函数属于特权接口，而不是权限绕过机制。应审计函数授权；如果部署的角色模型有此要求，应撤销 `PUBLIC` 的执行权限。
- 订阅连接字符串可能包含凭据。不要让它们进入日志、查询历史或错误报告管道，并优先采用服务商支持的密钥处理方式。
- 启用计划分析可能增加执行开销，并在日志中暴露查询文本或值。应选择性使用 `auto_explain`，尤其是在启用分析时。
