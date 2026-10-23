## 用法

来源：

- [v6.0 README](https://github.com/HexaCluster/credcheck/blob/ac9127dfe0892579a42bd2255638f9948074a151/README.md)
- [v6.0 changelog](https://github.com/HexaCluster/credcheck/blob/ac9127dfe0892579a42bd2255638f9948074a151/ChangeLog)
- [SQL objects 6.0.0](https://github.com/HexaCluster/credcheck/blob/ac9127dfe0892579a42bd2255638f9948074a151/sql/credcheck--6.0.0.sql)
- [Password history WAL implementation](https://github.com/HexaCluster/credcheck/blob/ac9127dfe0892579a42bd2255638f9948074a151/credcheck.c)
- [Login-event setup](https://github.com/HexaCluster/credcheck/blob/ac9127dfe0892579a42bd2255638f9948074a151/event_trigger.sql)

`credcheck` 在角色创建、密码修改和角色重命名时检查用户名与明文密码，还能限制密码重用、封禁多次认证失败的用户，并要求首次登录修改密码。策略由超级用户配置；默认值不构成完整的密码强度策略。

### 启用和设置策略

将库加入现有预加载列表并重启 PostgreSQL。在管理员需要视图和重置函数的各数据库中安装 SQL 对象：

```ini
shared_preload_libraries = 'credcheck'
credcheck.password_min_length = 12
credcheck.password_contain_username = on
credcheck.password_reuse_history = 2
credcheck.password_reuse_interval = 365
```

```sql
CREATE EXTENSION credcheck;
CREATE ROLE app_user LOGIN PASSWORD 'example-Strong-Pass#123';
SELECT rolename, password_date FROM pg_password_history;
```

时间间隔单位为天。安装 SQL 扩展与加载服务器范围的钩子是两个步骤。上游发布版 6.0 使用 SQL 扩展版本 6.0.0；安装升级文件后需要重启，以重新加载库。

### 策略索引

| 配置项 | 用途 |
|---|---|
| `credcheck.username_min_length`、`credcheck.username_min_special`、`credcheck.username_min_digit`、`credcheck.username_min_upper`、`credcheck.username_min_lower` | 用户名长度和字符要求 |
| `credcheck.password_min_length`、`credcheck.password_min_special`、`credcheck.password_min_digit`、`credcheck.password_min_upper`、`credcheck.password_min_lower` | 密码长度和字符要求 |
| `credcheck.username_min_repeat`、`credcheck.password_min_repeat` | 限制相邻重复次数，参数名称虽然含有 min，但实际表示最大次数 |
| `credcheck.username_contain`、`credcheck.username_not_contain`、`credcheck.password_contain`、`credcheck.password_not_contain` | 必须包含或不得包含的内容 |
| `credcheck.username_contain_password`、`credcheck.password_contain_username` | 拒绝相互包含的凭证 |
| `credcheck.username_ignore_case`、`credcheck.password_ignore_case` | 大小写处理 |
| `credcheck.password_min_length_su`、`credcheck.password_valid_until_su` | 单独的超级用户要求 |
| `credcheck.password_valid_until`、`credcheck.password_valid_max` | 密码有效期的最小和最大天数；修改密码但未指定到期日时，最小值还会自动设置到期日 |
| `credcheck.whitelist`、`credcheck.superuser_nocheck` | 明确的策略豁免 |
| `credcheck.no_password_logging` | 在策略错误日志中隐藏密码，默认启用 |

CrackLib 强度检查仅在库编译时启用相应支持且字典可用时生效。

### 密码历史与复制

历史记录保存 SHA-256 密码散列，在数据库间共享，并持久化到 `$PGDATA/pg_password_history`。应将该文件纳入备份，并保护 SQL 历史视图的访问权限；6.0 中该视图撤回 PUBLIC 权限，并向 pg_read_all_stats 授予查询权限。`credcheck.history_max_size` 改变历史容量，需要重启。

5.0 在 PostgreSQL 15 及以上通过编号 150 的自定义 WAL 资源管理器复制历史变化。重放这些 WAL 的副本和恢复服务器应预加载匹配的库。较早 PostgreSQL 版本保留文件持久化历史，但没有这项复制支持。历史重置和时间戳测试函数在恢复期间拒绝执行。

```sql
SELECT pg_password_history_reset('app_user');
```

重置历史会移除相应记录的密码重用保护，应仅由管理员操作。

### 认证与密码修改

```ini
credcheck.max_auth_failure = 3
credcheck.auth_delay_ms = 1000
credcheck.whitelist_auth_failure = 'service_user'
credcheck.password_change_first_login = true
```

```sql
SELECT * FROM pg_banned_role;
SELECT pg_banned_role_reset('app_user');
ALTER ROLE app_user SET credcheck_internal.force_change_password = true;
```

封禁持续到记录被重置，其缓存会在重启时丢失。`credcheck.reset_superuser` 提供上游说明的超级用户恢复入口；`credcheck.auth_failure_cache_size` 需要重启。

禁止修改密码的实际参数是 `credcheck.disallow_change_password`。超级用户也受其影响，除非在会话中启用 `credcheck.superuser_nocheck`。这个豁免会绕过相应的全部角色检查，应严格控制。

`credcheck.password_valid_warning` 要求 PostgreSQL 17 及以上，并需在每个相关数据库单独安装官方登录事件触发器；创建 SQL 扩展不会安装该触发器。

### 明文边界

强度与重用检查需要在修改密码时获得明文。默认拒绝已计算散列的密码，包括 psql 的 `\password` 提交的密码。启用 `credcheck.encrypted_password_allowed` 可以接受它们，但不会提供等效的明文检查。应保护密码修改连接，也不要假设扩展会追溯扫描已有凭证。创建无密码角色，或重命名没有密码的角色时，会跳过用户名检查。

### 6.0 升级与新接口

6.0 将密码历史迁移到 mmap 文件中的整数索引哈希表。首次升级启动会导入 5.x 历史，并把旧文件重命名为 pg_password_history.legacy-migrated；主库和备库都执行该迁移。升级前保留并验证历史文件备份。安装匹配库并重启后，在使用 SQL 对象的数据库执行 `ALTER EXTENSION credcheck UPDATE TO '6.0.0'`。当前 SQL 会撤回历史函数和视图的 PUBLIC 权限，并将视图查询权授予 pg_read_all_stats；复查已有定制授权。

`credcheck.password_valid_min` 为新建角色设置较短的初始密码有效期，默认 0 表示关闭；既有角色修改密码仍使用原有效期策略。新函数 `pg_check_password(username, password)` 检查候选凭证，违规时报错、通过时返回 true，不检查密码重用历史。

`credcheck.lastlog` 默认关闭，启用后可用 pg_lastlog 查看会话与服务器启动、关闭事件；可选 SQL 跟踪会记录查询文本，应明确访问权限和保留范围。PgBouncer 的失败次数和首次登录改密集成需要按官方文档配置额外 SQL 与连接池钩子，预加载扩展不会自动完成。`credcheck.no_password_logging` 现在遮蔽密码字面量而保留错误语句结构。此版还修复自助 ALTER ROLE 权限提升和内存缺陷。
