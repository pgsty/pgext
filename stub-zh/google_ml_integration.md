## 用法

来源：

- [Cloud SQL 模型端点管理参考](https://docs.cloud.google.com/sql/docs/postgres/model-endpoint-management-reference)
- [Cloud SQL 模型注册指南](https://docs.cloud.google.com/sql/docs/postgres/model-endpoint-register-model)

`google_ml_integration` 是 Cloud SQL PostgreSQL 的厂商托管扩展，用于注册远程模型端点、保存 Secret Manager 引用、生成嵌入和调用逐行预测。模型端点管理目前属于 Preview 功能，并要求启用 `google_ml_integration.enable_model_support` 数据库标志。

```sql
CREATE EXTENSION google_ml_integration;
SELECT * FROM google_ml.model_info_view;
SELECT google_ml.embedding(
  model_id => 'textembedding-gecko@001',
  contents => 'PostgreSQL extension metadata'
);
```

使用 `google_ml.create_model(...)` 注册端点。厂商元数据支持 Google、OpenAI、Anthropic、Hugging Face 与自定义端点；非 Google 凭据应通过 Secret Manager 引用，不要写成 SQL 明文。应把模型、秘密、嵌入和预测例程限制给受控角色。

调用会把数据库内容发送给外部模型提供方。应进行数据分级，执行驻留与出站策略，最小化提示词，清除秘密和个人数据，并考虑提供方保留、费用、限流、延迟、故障和非确定性。应验证请求/响应转换和向量维度。该功能处于 Preview 且由厂商托管，生产依赖前必须核实精确 Cloud SQL 版本类型、地域、引擎版本、扩展版本、IAM 服务代理权限和升级行为。
