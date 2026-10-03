## 用法

来源：

- [docs/integrations/postgres-extension.md](https://github.com/pgquery-narrative/pgquerynarrative/blob/b0d91876f50ea2f3519ff12ee27c01c5e253e7a3/docs/integrations/postgres-extension.md)
- [infra/postgres-extension/pgquerynarrative.control](https://github.com/pgquery-narrative/pgquerynarrative/blob/b0d91876f50ea2f3519ff12ee27c01c5e253e7a3/infra/postgres-extension/pgquerynarrative.control)
- [infra/postgres-extension/pgquerynarrative--1.1.sql](https://github.com/pgquery-narrative/pgquerynarrative/blob/b0d91876f50ea2f3519ff12ee27c01c5e253e7a3/infra/postgres-extension/pgquerynarrative--1.1.sql)

`pgquerynarrative` 1.1 通过 SQL 调用正在运行的 PgQueryNarrative 服务。它与在数据库内执行调查的 `pqn` 扩展相互独立。

### 核心工作流

由超级用户先安装 HTTP 依赖，再安装封装扩展，并向已有角色授予访问权限。配置可访问且可信的服务端点。

```sql
CREATE EXTENSION http;
CREATE EXTENSION pgquerynarrative;
SELECT pgquerynarrative_set_api_url('http://localhost:8080');
SELECT pgquerynarrative_grant_access('app_reader');
```

### API 与边界

`pgquerynarrative_set_api_key` 为每个调用者设置会话凭据。`pgquerynarrative_run_query` 通过服务执行只读查询；`pgquerynarrative_generate_report` 请求报告，`pgquerynarrative_list_saved` 列出已保存查询。`pgquerynarrative_revoke_access` 撤销访问权限。

若安装时没有 `http`，扩展会创建返回等待状态的函数，而非真正发出 HTTP 请求；之后补装依赖不会自动改写这些函数。1.1 版本撤销了 `PUBLIC` 的函数执行权限，存储的 API URL 由所有者控制。包含 API 密钥字面量的语句可能进入查询日志，SQL 也会被发送给外部服务。仍须遵守服务端授权与结果限制。服务文档覆盖 PostgreSQL 16–18；这些 SQL 封装无需预加载。
