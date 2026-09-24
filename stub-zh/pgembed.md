## 用法

来源：

- [README](https://github.com/Tecnisys-OSS/pgembed/blob/1aed2a0f7e7c7be729cb9558c031133453aeeee3/README.md)
- [Control file / 控制文件](https://github.com/Tecnisys-OSS/pgembed/blob/1aed2a0f7e7c7be729cb9558c031133453aeeee3/pgembed.control)
- [sql/pgembed--0.1.0.sql](https://github.com/Tecnisys-OSS/pgembed/blob/1aed2a0f7e7c7be729cb9558c031133453aeeee3/sql/pgembed--0.1.0.sql)
- [LICENSE](https://github.com/Tecnisys-OSS/pgembed/blob/1aed2a0f7e7c7be729cb9558c031133453aeeee3/LICENSE)

`pgembed` 通过 Ollama、OpenAI 或自定义 HTTP 端点生成 pgvector 向量。扩展使用非受信 PL/Python 实现，推理由外部服务完成。

### 核心工作流

须由超级用户提供 `plpython3u` 和 `vector`，查询前还需启动兼容的 Ollama 服务并准备指定模型。

```sql
CREATE EXTENSION plpython3u;
CREATE EXTENSION vector;
CREATE EXTENSION pgembed;
SELECT pgembed.embed_ollama('PostgreSQL extensions', 'embeddinggemma');
SELECT * FROM pgembed.embed_batch_ollama(
  ARRAY['PostgreSQL extensions', 'Vector similarity'], 'embeddinggemma');
```

### 接口与配置

`pgembed` 模式中的 `embed_ollama`、`embed_openai` 和 `embed_custom` 返回向量，批量形式处理文本数组。目标向量列维度必须匹配模型。`pgembed.url_allowlist` 控制允许的主机；`pgembed.max_retries`、`pgembed.initial_backoff_ms`、`pgembed.max_backoff_ms` 控制重试；`pgembed.circuit_breaker_threshold` 和 `pgembed.circuit_breaker_reset_timeout_s` 控制各后端独立的熔断器。

README 声明要求 PostgreSQL 13 或更新版本。扩展本身不需要预加载或重启。HTTP 调用完成前会占用数据库后端，应明确控制原文和凭据的访问权限。可配置主机白名单不能替代数据库权限和网络策略；不要把真实 API 密钥写入会记录日志的 SQL。Tecnisys Community License 1.0 限制竞争性托管服务用途。
