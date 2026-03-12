


## 用法

> [pg_ai_query: AI 驱动的 PostgreSQL SQL 查询生成](https://github.com/benodiwal/pg_ai_query)

`pg_ai_query` 使用 LLM 提供商（OpenAI、Anthropic、Google Gemini）从自然语言描述生成 SQL 查询。它会自省数据库模式，将自然语言问题翻译为 SQL。

### 配置

创建 `~/.pg_ai.config`：

```ini
[general]
log_level = "INFO"
enable_logging = false

[query]
enforce_limit = true
default_limit = 1000

[response]
show_explanation = true
show_warnings = true

[openai]
api_key = "your-openai-api-key"
default_model = "gpt-4o"

[anthropic]
api_key = "your-anthropic-api-key"
default_model = "claude-3-5-sonnet-20241022"

[gemini]
api_key = "your-google-api-key"
default_model = "gemini-2.5-flash"
```

### 查询生成

```sql
SELECT generate_query('show all customers');
SELECT generate_query('monthly sales trend for the last year by category');
SELECT generate_query('customers who have not placed orders in the last 6 months');
```

### 模式发现

```sql
SELECT get_database_tables();
SELECT get_table_details('orders');
```

### 查询解释

```sql
SELECT explain_query('SELECT * FROM users WHERE created_at > NOW() - INTERVAL ''7 days''');
```

直接传入 API 凭证：

```sql
SELECT explain_query('SELECT * FROM products WHERE price > 100', 'your-api-key', 'anthropic');
```

### 支持的模型

- **OpenAI**: gpt-4o, gpt-4o-mini
- **Anthropic**: claude-3-5-sonnet-20241022, claude-3-haiku-20240307
- **Google Gemini**: gemini-2.5-pro, gemini-2.5-flash
