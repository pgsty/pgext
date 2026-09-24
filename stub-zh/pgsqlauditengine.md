## 用法

来源：

- [README](https://github.com/shaohuasong-fang/pgsqlauditengine/blob/4da0480d476357735f166377ee0c7179b046004c/README.md)
- [Control file / 控制文件](https://github.com/shaohuasong-fang/pgsqlauditengine/blob/4da0480d476357735f166377ee0c7179b046004c/pgsqlauditengine.control)
- [pgsqlauditengine--1.0.sql](https://github.com/shaohuasong-fang/pgsqlauditengine/blob/4da0480d476357735f166377ee0c7179b046004c/pgsqlauditengine--1.0.sql)

`pgsqlauditengine` 利用 PostgreSQL 钩子在执行前检查 SQL 规范，并通过内嵌 REST 服务提供规则和审计记录。上游声明支持 PostgreSQL 11–18。

### 启用

将库加入已有预加载列表，启用检查后重启。下例 GUC 前缀的特殊拼写来自上游，不能自行改写。允许远程访问前须配置非空 API 令牌。

```conf
shared_preload_libraries = 'pgsqlauditengine'
PGSAUDAUDITENGINE.enabled = on
PGSAUDAUDITENGINE.check_dml = on
PGSAUDAUDITENGINE.api_listen = '127.0.0.1'
PGSAUDAUDITENGINE.api_port = 8918
```

```sql
CREATE EXTENSION pgsqlauditengine;
SELECT name, setting FROM pg_settings WHERE name LIKE 'PGSAUDAUDITENGINE.%';
```

### 规则与 API

扩展 DDL 用于登记元数据，钩子由预加载激活。`PGSAUDAUDITENGINE.enabled` 默认关闭。规则覆盖 DDL、DML、权限、事务语句和可编程对象。`ERROR` 与 `WARNING` 两种规则级别都会阻止执行，`NOTICE` 才会放行。

API 提供 `/api/v1/health`、`/api/v1/rules`、`/api/v1/audit-logs` 和 `/api/v1/config`，规则可在运行时修改。审计记录位于环形缓冲区，并非持久审计档案。`PGSAUDAUDITENGINE.api_token` 为空时不鉴权，健康端点也不要求鉴权。应限制服务访问，并由管理员管理授权和配置。
