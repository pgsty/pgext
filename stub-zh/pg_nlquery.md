## 用法

来源：

- [README.md](https://github.com/schetan/pg_nlquery/blob/d15c30aec342ce1eab707391152b9595978e140a/README.md)
- [pg_nlquery--1.0.sql](https://github.com/schetan/pg_nlquery/blob/d15c30aec342ce1eab707391152b9595978e140a/pg_nlquery--1.0.sql)
- [pg_nlquery.c](https://github.com/schetan/pg_nlquery/blob/d15c30aec342ce1eab707391152b9595978e140a/pg_nlquery.c)
- [schema_introspect.c](https://github.com/schetan/pg_nlquery/blob/d15c30aec342ce1eab707391152b9595978e140a/schema_introspect.c)
- [sql_validate.c](https://github.com/schetan/pg_nlquery/blob/d15c30aec342ce1eab707391152b9595978e140a/sql_validate.c)
- [spi_exec.c](https://github.com/schetan/pg_nlquery/blob/d15c30aec342ce1eab707391152b9595978e140a/spi_exec.c)
- [pg_nlquery.control](https://github.com/schetan/pg_nlquery/blob/d15c30aec342ce1eab707391152b9595978e140a/pg_nlquery.control)

`pg_nlquery` 将问题和所选模式的元数据发送到大模型端点，校验生成的语句，并可执行该语句。`nlp_explain` 只返回 SQL 而不执行；`nlp_query` 为每条结果记录返回一个 JSON 值。安装需要超级用户。底层库使用 libcurl，无需预加载。

### 核心用法

```sql
CREATE EXTENSION pg_nlquery;
SET pg_nlquery.api_format = 'openai';
SET pg_nlquery.base_url = 'http://localhost:11434/v1';
SET pg_nlquery.model = 'qwen2.5-coder:7b';
SET pg_nlquery.api_key = 'ollama';
SET pg_nlquery.schema_filter = 'public';
SELECT nlp_explain('show the five most recent orders');
SELECT * FROM nlp_query('show the five most recent orders');
```

### 运行边界

示例使用本地兼容 OpenAI 的端点；即使采用此本地配置，也必须设置非空 API 密钥。用 pg_nlquery.schema_filter 限制描述的模式，column_hints 可补充发送给模型的说明。timeout_ms 默认为 30000，max_tokens 默认为 1024。模式元数据、提示和问题会发送到配置的端点。SQL 按调用者权限执行。校验器只接受一条顶层 SELECT，SPI 以只读模式执行；这并不是覆盖 SQL 函数和外部副作用的完整沙箱。应检查生成的 SQL，并使用权限受限的角色。结果会整体物化，扩展没有行数上限。API 密钥 GUC 设置了仅超级用户可见标志，但上下文为 USERSET，不能把它当作凭证保险库。本修订的 README 声明 PostgreSQL 13+，未提供测试主版本矩阵。
