## 用法

来源：

- [已复核 commit 的 pg_ollama README](https://github.com/ryicoh/pg_ollama/blob/c4b5dc429f0b9558cf6b9c4f8da5aaf59e3a9df1/README.md)
- [版本 1.0 的 SQL 接口](https://github.com/ryicoh/pg_ollama/blob/c4b5dc429f0b9558cf6b9c4f8da5aaf59e3a9df1/ollama--1.0.sql)
- [已复核 commit 的 ollama.control](https://github.com/ryicoh/pg_ollama/blob/c4b5dc429f0b9558cf6b9c4f8da5aaf59e3a9df1/ollama.control)

`ollama` 是面向外部 Ollama 服务器的 PL/pgSQL 客户端。它在 `ollama` schema 中保存模型、端点、超时和系统消息设置，通过 `http` 扩展同步发送聊天请求，并记录每次请求与响应。它适合能够接受数据库后端阻塞和提示词持久化的小型 SQL 模型调用，而不是数据库内模型运行时。

### 核心工作流

先安装 control 文件中没有声明的 `http` 依赖。`ollama_insert_default_settings()` 会一次性创建使用 `http://localhost:11434`、模型 `gemma:2b`、60 秒超时和一条系统消息的默认配置。

```sql
CREATE EXTENSION IF NOT EXISTS http;
CREATE EXTENSION ollama;

SELECT ollama_insert_default_settings();
SELECT ollama('Explain PostgreSQL MVCC in one sentence.');
```

Ollama 进程必须已经能从数据库服务器访问，并且其中已经存在所配置的模型。自定义部署应直接填充 `ollama.models`、`ollama.endpoints`、`ollama.settings` 和有序的 `ollama.setting_messages` 行，而不是调用默认辅助函数。

### SQL 接口

- `ollama(text)` 使用唯一一条 `is_default` 为真的配置。
- `ollama(integer, text)` 使用明确的配置 ID。
- `ollama_insert_default_settings()` 插入示例端点、模型、配置和系统消息。它不是幂等函数，重复调用会与模型主键冲突。
- `ollama.logs` 保存输入、JSON 请求、HTTP 状态、JSON 响应、提取输出和时间戳。日志行在 HTTP 调用前插入，因此失败请求可能留下未完成的行。

返回内容通过 JSON 操作符提取，再转换为函数声明的文本结果；把结果继续拼接到文本前，应确认客户端是否会收到 JSON 风格的引号。

### 运维注意事项

- control 文件没有声明依赖，但安装 SQL 引用了 `http_response`、`http_post` 和 `http_set_curlopt`；必须先安装 `http`，再安装 `ollama`。
- 每次调用都会让一个 PostgreSQL 后端等待 `/api/chat`。端点故障、模型延迟和网络策略会直接影响查询延迟与事务持续时间。
- 提示词、系统消息、完整请求 JSON、响应和输出都持久化在普通表中。应限制 schema 访问、定义保留周期；未经批准的数据边界不得发送密钥或受监管数据。
- 端点 URL 存在数据库数据中，并控制 PostgreSQL 主机的外部访问。应限制修改配置的角色，并实施网络出口控制以降低服务端请求伪造风险。
- 目录与 control 文件都标识版本 `1.0`；上游没有发布广泛的 PostgreSQL 兼容矩阵。应在生产使用的确切服务器大版本上测试 SQL 依赖与 HTTP 行为。
