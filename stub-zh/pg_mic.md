## 用法

来源：

- [README.md](https://github.com/jconway/pg_mic/blob/8a8a3079e942d398deb6b4de160cc41089744f09/README.md)
- [pg_mic--1.0.sql](https://github.com/jconway/pg_mic/blob/8a8a3079e942d398deb6b4de160cc41089744f09/pg_mic--1.0.sql)
- [SECURITY.md](https://github.com/jconway/pg_mic/blob/8a8a3079e942d398deb6b4de160cc41089744f09/SECURITY.md)
- [pg_mic_compat.h](https://github.com/jconway/pg_mic/blob/8a8a3079e942d398deb6b4de160cc41089744f09/pg_mic_compat.h)
- [pg_mic.control](https://github.com/jconway/pg_mic/blob/8a8a3079e942d398deb6b4de160cc41089744f09/pg_mic.control)

`pg_mic` 为角色、关系和函数添加 Biba 完整性检查。HIGH、MEDIUM、LOW 标签限制读取低完整性数据以及写入高完整性数据。这控制的是完整性而非机密性，SQL 权限仍然有效。控制文件与 SQL 的版本为 1.0，上游安全文档仍将其描述为测试版本。

### 核心用法

```ini
shared_preload_libraries = 'pg_mic'
```

```sql
CREATE EXTENSION pg_mic;
CREATE ROLE integrity_reader;
CREATE TABLE verified_facts (id integer, fact text);
SELECT mic_promote_table('verified_facts', 'high');
GRANT SELECT ON verified_facts TO integrity_reader;
SELECT mic_get_session_level();
SELECT mic_get_table_level('verified_facts');
SET ROLE integrity_reader;
SELECT * FROM verified_facts;
RESET ROLE;
```

### 运行边界

将 `pg_mic` 加入 `shared_preload_libraries`，重启 PostgreSQL，再由超级用户在各数据库中创建扩展。上游文档列出 PostgreSQL 14–19 和开发分支，此源代码修订包含相应服务器接口的版本判断。

超级用户的完整性为 HIGH，未标记的角色或用户对象默认为 LOW。新建表和函数继承创建者的级别。SELECT 要求主体级别不高于对象，INSERT 和 DELETE 要求主体级别不低于对象，UPDATE 同时应用两条规则。内置函数和系统目录读取有豁免。因此，已有未标记对象可能无法被 HIGH 会话使用，需先核验再重新标记。

`mic_get_session_level`、`mic_get_role_level` 和 `mic_get_table_level` 用于检查标签。超级用户可通过 `mic_set_role_level`、`mic_promote_table`、`mic_promote_object`、`mic_promote_extension` 和 `mic_initialize_labels` 管理标签；`mic_labels` 和 `mic_role_labels` 是标签视图。仅提升经过核验的数据与代码。

`mic.mode` 默认为 `enforcing`；`permissive` 会告警但允许违规操作，`disabled` 则跳过检查。修改模式需要重新加载配置，`mic.debug_audit` 可开启决策日志。RLS 检查较保守，可能拒绝 PostgreSQL 原本允许的路径。函数降级会使执行计划失效，但计划缓存之外保存的函数指针可能在会话结束前保留旧行为；降级后应重新连接受影响的长会话。新增的服务器实用命令需要上游明确补充覆盖。
