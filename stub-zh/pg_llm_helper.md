## 用法

来源：

- [官方 README](https://github.com/GeoffMontee/pg_llm_helper/blob/feae47af77a9cb735400b66b0b39f046d60cf924/README.md)
- [扩展 SQL](https://github.com/GeoffMontee/pg_llm_helper/blob/feae47af77a9cb735400b66b0b39f046d60cf924/pg_llm_helper--1.0.sql)
- [共享内存实现](https://github.com/GeoffMontee/pg_llm_helper/blob/feae47af77a9cb735400b66b0b39f046d60cf924/pg_llm_helper.c)
- [扩展控制文件](https://github.com/GeoffMontee/pg_llm_helper/blob/feae47af77a9cb735400b66b0b39f046d60cf924/pg_llm_helper.control)

`pg_llm_helper` 1.0 安装覆盖整个 postmaster 的错误钩子，以及容纳 100 条记录的共享内存环；其中保存失败查询文本、错误文本、SQLSTATE、严重级别、后端 PID 和时间戳。它还能通过 pgai 把当前后端最后捕获的错误发送给 LLM。由于历史跨会话、跨数据库，并可能含凭据或个人数据，必须按集群级诊断日志限制访问。

### 预加载与安装

把动态库加入 `shared_preload_libraries` 并重启；稍后再加载无法分配共享内存或注册钩子。

```conf
shared_preload_libraries = 'pg_llm_helper'
```

在受控管理数据库中创建扩展，并立即移除默认访问权：

```sql
CREATE EXTENSION pg_llm_helper;

REVOKE EXECUTE ON FUNCTION get_last_error() FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION get_error_history(integer) FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION clear_error_history() FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION llm_help_last_error() FROM PUBLIC;

SELECT * FROM get_last_error();
SELECT * FROM get_error_history(10);
```

上游 README 把 `vector` 称为必需依赖，但复核的控制文件没有声明 `requires`，安装的错误捕获 SQL 也不使用 vector 对象。LLM 辅助函数在调用时需要兼容的 `ai.openai_chat_complete` 函数，但 pgai 也没有声明为扩展依赖。应验证实际打包的依赖集合，而不是依赖 `CREATE EXTENSION ... CASCADE` 猜测。

### 用户函数

- `get_last_error()`：查找后端 PID 等于当前 PID 的最新环形缓冲条目。
- `get_error_history(integer)`：复制最多 100 个来自所有后端的已填充槽；数组遍历不保证按新到旧排序。
- `clear_error_history()`：清空整个集群级缓冲区，而不只是调用者的条目。
- `llm_help_last_error()`：格式化最后一条错误及其查询，通过 pgai 调用硬编码的 `gpt-4o-mini` 模型并返回响应。

编译期限制为 100 条错误、8192 字节查询文本和 1024 字节错误文本。条目既没有数据库标识也没有用户标识。PID 复用会让新后端在槽位被覆盖前找到旧进程的条目。

### 隐私与可靠性边界

钩子从 `debug_query_string` 捕获 `ERROR` 及更高级别。查询和错误消息可能含字面量、API 密钥、客户数据、对象名与安全策略细节。获准调用 `get_error_history` 的角色能读取其他会话捕获的文本，获准调用 `clear_error_history` 的角色能删除证据。应仅向已审计管理员授权，统一日志保留与披露策略，并测试多租户和 PID 复用行为。

调用 `llm_help_last_error` 会同步向外部提供方发送已捕获的查询和错误文本。授权前应落实数据分级、脱敏、出站、驻留、保留、成本、超时与提供方密钥控制。返回建议是不可信诊断文本，不能用于授权或直接执行修复。

复核的项目面向 PostgreSQL 17，并使用全局日志与共享内存钩子。应测试它与所有其他预加载模块的钩子链、启动/重启、缓冲区争用、截断、崩溃、升级和错误风暴下的性能。不要把该环当作审计日志：它容量小、会覆盖、可被全局清空，而且缺少持久排序与身份字段。
