## 用法

来源：

- [腾讯云 tencentdb_ai 函数文档](https://cloud.tencent.com/document/product/409/116227)
- [TencentDB for PostgreSQL 支持的扩展版本](https://cloud.tencent.com/document/product/409/75121)

`tencentdb_ai` 是 TencentDB for PostgreSQL 托管扩展，从 SQL 调用网络可达的模型 API。1.3 版管理模型定义与凭据，并提供通用调用以及对话、向量化、重排序和图生文辅助函数。请求会离开数据库，可能产生服务商费用、延迟、限流和数据治理义务。

### 注册并调用模型

在受支持的 TencentDB 实例上：

```sql
CREATE EXTENSION tencentdb_ai CASCADE;

SELECT tencentdb_ai.add_model(
  'hunyuan-lite',
  '2023-09-01',
  NULL,
  NULL
);

SELECT * FROM tencentdb_ai.list_models();

SELECT tencentdb_ai.chat_completions(
  'hunyuan-lite',
  'Summarize the current transaction isolation level.'
);

SELECT tencentdb_ai.get_embedding(
  'hunyuan-embedding',
  ARRAY['PostgreSQL', 'vector search']
);
```

创建时还会安装 `pgcrypto`。调用已注册模型之前，使用 `update_model_attr(model_name, attr_name, attr_value)` 配置服务商 `SecretId` 和 `SecretKey`，并按需设置 `version`、`region` 或 `json_path`。不要把凭据写入可复用 SQL、应用日志或共享 shell 历史；应严格限制函数和模型目录访问。

### 函数索引

`add_model` 登记名称、版本、区域和可选 JSONPath；该路径类似 `jsonb_path_query_first`，用于选择响应字段。`list_models`、`update_model_attr` 和 `delete_model` 管理登记。`call_model(model_name, common_params, api_params)` 暴露通用文本数组请求形式。`chat_completions`、`get_embedding`、`run_rerank` 和 `image_to_text` 提供固定场景接口。

### 服务、版本与安全边界

详细函数页面把 PostgreSQL 17 写作前提，而当前服务商扩展矩阵为 PostgreSQL 14–18 列出 `tencentdb_ai` 1.3。较旧的操作示例输出还显示扩展版本 1.0。应以实例 `pg_available_extension_versions` 与当前腾讯支持答复为准，并在安装后查询 `pg_extension`。

调用是同步外部网络操作，SQL 回滚不会赋予其事务性。提示词、文档、图片 URL 与返回内容可能包含敏感数据或不可信模型输出。在把结果用于数据库写入或自动决策前，应实施出口控制、最小权限腾讯凭据、时间/成本预算、输出验证、重试/幂等规则以及明确的同意与保留策略。
