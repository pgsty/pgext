## 用法

来源：

- [官方 README](https://github.com/Abiji-2020/pg_ask/blob/7c58aeb31042b75c4f1f675b621329c689003e2c/README.md)
- [扩展 SQL](https://github.com/Abiji-2020/pg_ask/blob/7c58aeb31042b75c4f1f675b621329c689003e2c/pg/pg_ask--1.0.sql)
- [SQL 生成入口](https://github.com/Abiji-2020/pg_ask/blob/7c58aeb31042b75c4f1f675b621329c689003e2c/src/pg_ask.cpp)
- [模式探查器](https://github.com/Abiji-2020/pg_ask/blob/7c58aeb31042b75c4f1f675b621329c689003e2c/src/explorer.cpp)
- [AI 客户端实现](https://github.com/Abiji-2020/pg_ask/blob/7c58aeb31042b75c4f1f675b621329c689003e2c/src/ai_engine.cpp)

`pg_ask` 会把自然语言提示和数据库模式元数据发送到 OpenAI 兼容端点，并返回生成的 SQL。只应在严格受限的角色下把 `pg_gen_query(text)` 用作评审辅助；`pg_gen_execute(text)` 会以调用者的数据库权限执行不可信的模型输出，不适合直接暴露。

### 核心流程

启动服务器前，在 PostgreSQL 服务进程环境中设置 `PG_ASK_AI_KEY`，安装扩展，并为专用角色配置端点和模型。使用前应先撤销函数默认授予公众的访问权：

```sql
CREATE EXTENSION pg_ask;

REVOKE EXECUTE ON FUNCTION pg_gen_query(text) FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION pg_gen_execute(text) FROM PUBLIC;
GRANT EXECUTE ON FUNCTION pg_gen_query(text) TO sql_preview_role;

ALTER ROLE sql_preview_role SET pg_ask.model = 'gpt-4o';
ALTER ROLE sql_preview_role SET pg_ask.endpoint = 'https://api.openai.com/v1';

SET ROLE sql_preview_role;
SELECT pg_gen_query('count orders created today');
```

应始终先检查并验证返回文本，再通过另一个受控执行路径运行它。该扩展不会解析 SQL、实施允许列表、参数化语句，也不会强制只读。

### API 与配置

- `pg_gen_query(text) → text` 收集非系统表和列的元数据，同步请求服务商，并返回模型生成的 SQL。
- `pg_gen_execute(text) → refcursor` 调用生成器，并通过动态 SQL 打开名为 `ai_query_result` 的固定游标；使用游标需要显式事务。
- `pg_ask.model` 是用户可设置的模型名。
- `pg_ask.endpoint` 是用户可设置的服务商基础 URL。
- `PG_ASK_AI_KEY` 继承自 PostgreSQL 服务进程环境。

使用 `pg_gen_execute(text)` 时需要配合 `BEGIN`、`FETCH ALL FROM ai_query_result` 与 `COMMIT`，但授予该函数意味着模型可以发出调用者获准执行的任意语句。最好完全不授予它。

### 安全与运维边界

模式探查器会把非系统模式中的模式名、表名、列名和类型名连同用户提示一起发送出去。应将其视为元数据披露，并评审服务商的数据保留、驻留和保密要求。提示与模式详情会通过无法随事务回滚的同步网络调用离开数据库。

由于 `pg_ask.endpoint` 可由用户设置，而服务器会提供 `PG_ASK_AI_KEY`，获授权者可能重定向请求来窃取密钥或访问内部 HTTP 服务。应限制函数和配置权限、约束出站网络，且不要复用具有广泛权限的服务商凭据。

SQL 文件把 `pg_gen_query(text)` 声明为 `PARALLEL SAFE`，但实现会读取系统目录并执行网络 I/O；不要据此假定并行执行安全。README 面向 PostgreSQL 16+，而仓库其他材料曾提到更旧版本，因此采用前应核验实际打包构建，并测试失败、超时和服务商响应行为。
