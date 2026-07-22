## 用法

来源：

- [阿里云官方 rds_embedding 指南](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/rds-embedding)
- [ApsaraDB RDS for PostgreSQL 官方概览](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/what-is-apsaradb-rds-for-postgresql)

`rds_embedding` 是阿里云 ApsaraDB RDS for PostgreSQL 扩展，通过调用已配置的外部 HTTP 模型把文本转换为向量。版本 `1.0` 只由服务商提供，需要 PostgreSQL 14 或以上版本和 `vector`，还要求托管实例具备出站网络访问能力。

### 前置条件与设置

使用 RDS 特权账户，并在嵌入扩展前启用向量类型：

```sql
CREATE EXTENSION vector;
CREATE EXTENSION rds_embedding;
```

确切的小版本必须列出该扩展。阿里云指南还要求模型服务 API 密钥以及 NAT/SNAT 配置，因为 RDS PostgreSQL 默认无法访问互联网。

### 核心流程

使用请求与响应模板注册模型，然后请求嵌入：

```sql
SELECT rds_embedding.add_model(
  'text-embedding-v3',
  'https://dashscope-intl.aliyuncs.com/api/v1/services/embeddings/text-embedding/text-embedding',
  'Authorization: Bearer <API_KEY>',
  '{"input":{"texts":["%s"]},"model":"text-embedding-v3","parameters":{"text_type":"query"}}',
  '->''output''->''embeddings''->0->>''embedding'''
);

SELECT rds_embedding.get_embedding_by_model(
  'text-embedding-v3',
  '<API_KEY>',
  'Endless Yangtze River rolls on'
)::real[]::vector;
```

将结果保存到维度匹配的 `vector(n)` 字段中，并使用 `vector` 提供的操作符进行相似度搜索。

### 重要对象

- `rds_embedding.add_model(...)`、`rds_embedding.update_model(...)` 与 `rds_embedding.del_model(name)` 管理 `rds_embedding.models` 中的记录。
- `rds_embedding.show_models()` 列出已配置模型。
- `rds_embedding.get_embedding_by_model(model, api_key, text)` 调用服务商，并按照配置的嵌入路径提取结果。
- 服务商列出了 `rds_embedding.get_response_by_model(...)`，但文档说明该函数尚不可用。

### 运维注意事项

模型调用是数据库语句内的远程同步操作：延迟、限流、DNS/NAT 故障、服务商中断、响应结构变化与 token 成本都会直接影响 SQL。请求模板和示例包含授权材料；应限制 `rds_embedding.models` 访问权、函数执行权、查询日志与监控输出，并轮换已暴露密钥。必须针对确切模型响应验证 `%s` 替换与提取路径，强制向量维度，并避免盲目重试写入。该扩展只存在于受支持的 ApsaraDB RDS 小版本中，生命周期由托管服务决定，而不是开源发布流程。
