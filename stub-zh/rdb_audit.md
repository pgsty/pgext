## 用法

来源：

- [docker/raodb/ext/rdb_audit/README.md](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_audit/README.md)
- [docker/raodb/ext/rdb_audit/rdb_audit--1.0.sql](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_audit/rdb_audit--1.0.sql)
- [docker/raodb/ext/rdb_audit/rdb_audit.conf](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_audit/rdb_audit.conf)
- [docker/raodb/ext/rdb_audit/rdb_audit.c](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_audit/rdb_audit.c)
- [docker/raodb/ext/rdb_audit/rdb_audit.control](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_audit/rdb_audit.control)

`rdb_audit` 1.0 是 RAO DB 仓库中的审计日志源码扩展，记录所配置的语句类别及对象信息。所核对的文件未确立完整的 PostgreSQL 兼容矩阵及再分发许可证。

### 核心用法

```ini
shared_preload_libraries = 'rdb_audit'
rdb_audit.log_scope = 'read,write,ddl,role'
```

```sql
CREATE EXTENSION rdb_audit;
SHOW rdb_audit.log_scope;
```

### 运行边界

将库追加到已有预加载列表并重启后，由超级用户创建扩展。安装 SQL 创建 DDL 与删除事件触发器以及审计表；只加载库不会安装这些对象。

`rdb_audit.log_scope` 选择读取、写入、DDL、角色等类别；`rdb_audit.log_format` 选择输出形式，`rdb_audit.log_directory` 控制日志目录。`rdb_audit.log_parameter` 记录绑定值，`rdb_audit.log_statement` 记录 SQL 文本，可能包含凭据或个人数据，应限制日志访问并设置保留周期和磁盘限制。

对象审计使用 `rdb_audit.role` 配置的审计角色及表权限。审计输出可能远大于实际数据变更；正式依赖前应在测试部署中验证开销和覆盖范围。
