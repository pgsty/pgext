


## Usage

> [pg_ai_query: AI-powered SQL query generation for PostgreSQL](https://github.com/benodiwal/pg_ai_query)

`pg_ai_query` generates SQL queries from natural language descriptions using LLM providers (OpenAI, Anthropic, Google Gemini). It introspects your database schema and translates plain-English questions into SQL.

### Configuration

Create `~/.pg_ai.config`:

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

### Query Generation

```sql
SELECT generate_query('show all customers');
SELECT generate_query('monthly sales trend for the last year by category');
SELECT generate_query('customers who have not placed orders in the last 6 months');
```

### Schema Discovery

```sql
SELECT get_database_tables();
SELECT get_table_details('orders');
```

### Query Explanation

```sql
SELECT explain_query('SELECT * FROM users WHERE created_at > NOW() - INTERVAL ''7 days''');
```

Pass API credentials directly:

```sql
SELECT explain_query('SELECT * FROM products WHERE price > 100', 'your-api-key', 'anthropic');
```

### Supported Models

- **OpenAI**: gpt-4o, gpt-4o-mini
- **Anthropic**: claude-3-5-sonnet-20241022, claude-3-haiku-20240307
- **Google Gemini**: gemini-2.5-pro, gemini-2.5-flash
