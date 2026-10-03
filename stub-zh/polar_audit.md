## 用法

来源：

- [Implementation (polar_audit.c)](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/aabf18dcf39abee55de91f84630cfa3143ea7431/external/polar_audit/polar_audit.c)
- [Extension control file](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/aabf18dcf39abee55de91f84630cfa3143ea7431/external/polar_audit/polar_audit.control)
- [Installation SQL](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/aabf18dcf39abee55de91f84630cfa3143ea7431/external/polar_audit/polar_audit--1.0.sql)

`polar_audit` 使用 PolarDB 专有审计钩子记录语句类别和数据库对象。本条目依据上游基于 PostgreSQL 15 的 PolarDB 分支核验，不以竞赛分叉或原生 PostgreSQL 构建为依据。

### 启用

将该库追加到已有预加载列表并重启 PolarDB。应显式选择审计类别，默认不记录任何类别。版本化扩展 SQL 不创建 SQL 对象，仅注册扩展不会启用审计。

```conf
shared_preload_libraries = 'polar_audit'
polar_audit.log = 'read,write,ddl'
```

### 设置与边界

`polar_audit.log` 选择逗号分隔的类别，类别前加 `-` 表示排除。`polar_audit.log_catalog` 控制仅涉及系统目录的活动，`polar_audit.log_relation` 生成逐关系条目，`polar_audit.log_statement` / `polar_audit.log_parameter` 控制 SQL 文本和参数记录。`polar_audit.role` 选择审计角色。这些设置需由具备特权的管理员管理。日志可能包含应用数据并快速增长，应设置保留和访问控制策略。源码调用 PolarDB 审计函数，并拒绝在服务器预加载之外加载，因此不声明支持原生 PostgreSQL。
