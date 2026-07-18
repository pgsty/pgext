## 用法

来源：

- [已复核 commit 的 pg_ollama README](https://github.com/ryicoh/pg_ollama/blob/c4b5dc429f0b9558cf6b9c4f8da5aaf59e3a9df1/README.md)
- [已复核 commit 的 ollama.control](https://github.com/ryicoh/pg_ollama/blob/c4b5dc429f0b9558cf6b9c4f8da5aaf59e3a9df1/ollama.control)
- [版本 1.0 的安装 SQL](https://github.com/ryicoh/pg_ollama/blob/c4b5dc429f0b9558cf6b9c4f8da5aaf59e3a9df1/ollama--1.0.sql)

`ollama` 是面向 Ollama 服务的 PL/pgSQL 客户端。它通过 `http` 扩展从 PostgreSQL 后端发送聊天请求，保存端点、模型、提示模板和超时设置，并把请求与响应记录到 `ollama.logs`。

SQL 接口提供使用默认设置的 `ollama(text)`，以及接收明确设置 ID 的 `ollama(integer, text)`。`ollama_insert_default_settings` 会创建指向 `http://localhost:11434` 的端点，选择 `gemma:2b`，并加入基础系统消息。

### 基本设置

```sql
CREATE EXTENSION IF NOT EXISTS http;
CREATE EXTENSION ollama;

SELECT ollama_insert_default_settings();
SELECT ollama('Explain PostgreSQL MVCC in one sentence.');
```

配置端点上的 Ollama 进程与所选模型必须事先可用。

### 注意事项

- control 文件没有声明依赖，但安装 SQL 需要 `http` 和 `plpgsql`；应在 `ollama` 之前安装这些扩展。
- 提示、请求 JSON、响应和生成结果都会持久化到 `ollama.logs`。应实施适当的访问控制与保留策略。
- 调用会从数据库后端同步发出 HTTP 请求。端点故障、模型延迟和网络策略会直接影响 SQL 调用。
- 目录与 control 文件均标识版本 `1.0`；上游没有提供广泛的 PostgreSQL 兼容矩阵。
