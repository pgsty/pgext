## 用法

来源：

- [阿里云 rds_ai 文档](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/ai-rds-ai)
- [阿里云 RDS PostgreSQL 支持的扩展列表](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`rds_ai` 是阿里云扩展，可在 ApsaraDB RDS PostgreSQL 上通过 SQL 调用大语言模型和 embedding 模型。它提供 `rds_ai.prompt`、`rds_ai.embed`、`rds_ai.retrieve` 和 `rds_ai.rag`，分别用于生成、向量化、相似度检索与检索增强生成，还能调用已经配置、兼容 HTTP 的自定义模型。

### 安装、配置与 Prompt

特权 RDS 账号可以安装云厂商插件。`CASCADE` 还会安装其 `pgvector` 和 `pgsql-http` 依赖。

```sql
CREATE EXTENSION IF NOT EXISTS rds_ai CASCADE;

SELECT rds_ai.update_model('qwen-plus', 'token', :'model_api_key');
SELECT rds_ai.update_model(
  'qwen-plus',
  'uri',
  'https://dashscope-intl.aliyuncs.com/api/v1/services/aigc/text-generation/generation'
);

SELECT rds_ai.prompt('Summarize the purpose of PostgreSQL WAL.');
```

模型配置存储在 `rds_ai.model_list` 中。默认云厂商模型包括 Qwen prompt 模型与 `text-embedding-v3`；`rds_ai.add_model` 可以注册其他兼容 HTTP 的端点。

### 注意事项

- 这是阿里云的云厂商扩展，并非可移植的开源软件包。可用性及报告的 `1.0` 或 `1.0.0` 版本取决于 RDS PostgreSQL 主/次引擎版本和实例版本。
- 云厂商要求使用受支持的 RDS 引擎、特权数据库账号和有效的 Model Studio API 密钥。访问外部模型还要求 VPC 具备出站网络，通常需要 NAT。
- Prompt、检索上下文和 embedding 输入会经 HTTP 发送给配置的模型端点。发送数据库内容前，必须落实数据分级、驻留、日志与出站控制。
- API token 与端点模板位于 `rds_ai.model_list` 中，并通过 SQL 设置。应限制表与函数权限，避免在查询日志中出现字面秘密，并轮换已经暴露的凭据。
- 普通调用不要求预加载，但阿里云文档说明，修改 `rds_ai.default_prompt_model` 前需要把 `rds_ai` 加入 `shared_preload_libraries`。该变更必须使用托管实例参数流程。
- 稠密与稀疏 embedding 维度、默认模型和受支持次版本都可能随服务变化。应根据实时云厂商文档确定 vector 列大小，不要假定目录默认值。
