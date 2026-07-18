## 用法

来源：

- [已复核 commit 的 Aiven Extras README](https://github.com/aiven/aiven-extras/blob/ed2a4a676034b1ca3edb767d5805c482262d3f9b/README.md)
- [已复核 commit 的 control 模板](https://github.com/aiven/aiven-extras/blob/ed2a4a676034b1ca3edb767d5805c482262d3f9b/aiven_extras.control.in)
- [已复核 commit 的构建版本](https://github.com/aiven/aiven-extras/blob/ed2a4a676034b1ca3edb767d5805c482262d3f9b/Makefile)
- [Aiven Extras 标签](https://github.com/aiven/aiven-extras/tags)

`aiven_extras` 是由 Aiven 维护的扩展。特权操作员完成安装后，它允许非超级用户执行一组经过筛选的受限操作。主要场景包括逻辑复制管理、会话级 `auto_explain` 配置、部分 pgaudit 设置以及受控的 dblink 执行。

逻辑复制包装函数包括 `pg_create_publication`、`pg_create_subscription`，以及订阅的启用、禁用、刷新、列举和删除操作。调用角色仍须拥有 `REPLICATION` 权限；该扩展不会把任意角色变成超级用户。

### 基本设置

```sql
CREATE EXTENSION aiven_extras;

SELECT aiven_extras.auto_explain_load();
SELECT aiven_extras.set_auto_explain_log_min_duration('2000');
SELECT aiven_extras.set_auto_explain_log_analyze('on');
```

该示例只在当前数据库会话中加载并配置 `auto_explain`。调用发布和订阅函数时，应使用完全限定的表名，并谨慎限定连接字符串。

### 注意事项

- 必须先由特权操作员安装扩展。逻辑复制操作仍要求调用角色拥有 `REPLICATION`。
- 扩展只开放经过筛选的操作，并不是通用的权限绕过机制。
- 当前 `main` 分支的 Makefile 标识版本为 `1.1.20`，与目录一致；上游最新可见标签则为 `1.1.15`。应固定并测试实际采用的源码 commit 或服务商软件包。
- 订阅连接字符串可能包含凭据。不要让它们进入日志或查询历史，并优先采用服务商支持的密钥处理方式。
