## 用法

来源：

- [已复核 commit 的上游手册](https://github.com/uptimejp/sql_firewall/blob/39fdec885d748975e7475e05156b91eddff9d7a7/README.sql_firewall)
- [0.8 版本 SQL 对象](https://github.com/uptimejp/sql_firewall/blob/39fdec885d748975e7475e05156b91eddff9d7a7/sql_firewall--0.8.sql)
- [扩展 control 文件](https://github.com/uptimejp/sql_firewall/blob/39fdec885d748975e7475e05156b91eddff9d7a7/sql_firewall.control)

`sql_firewall` 是一个已放弃的 PostgreSQL 9.4 查询白名单原型。它学习用户 ID 与由解析树生成的查询 ID 对，然后对未学习查询告警或拒绝。必须把它放入 `shared_preload_libraries` 并重启。

```conf
shared_preload_libraries = 'sql_firewall'
sql_firewall.firewall = 'learning'
```

在管理数据库创建扩展后，应学习具有代表性的负载，检查 `sql_firewall.sql_firewall_statements`，先切换到 `permissive`，然后才考虑 `enforcing`。钩子作用于整个实例，但管理视图和函数只存在于创建 SQL 扩展的数据库。

规则导入导出使用超级用户指定的服务器文件路径，且只能在 disabled 模式执行。查询 ID 依赖 PostgreSQL 内部解析树，因此规则文件不能跨大版本移植。该代码仅支持 PostgreSQL 9.4，未经完整移植与安全审计不得部署到当前 PostgreSQL。它只能作为纵深防御，不能替代参数化 SQL、最小权限和应用输入验证。
