## 用法

来源：

- [Official database.dev 包页面](https://database.dev/jessevent/ai_sql)

`jessevent@ai_sql` — Supabase 的 AI 推断 SQL 函数。通过 ai.deploy() 自动部署其边缘函数。使用上游链接的固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION "jessevent@ai_sql";
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `ai._edge_function_source` 是一个扩展函数。
- `ai.deploy(project_ref text)` 是一个扩展函数并返回 `jsonb`。
- `ai.status()` 是一个扩展函数并返回 `jsonb`。
- `public._ai_call` 是一个扩展函数。
- `public._ai_extract_agg_ffunc` 是一个扩展函数。
- `public._ai_summarize_agg_ffunc` 是一个扩展函数。
- `public._ai_text_accum` 是一个扩展函数。
- `public.ai_classify(input_text text, categories text[] default array['positive','negative','neutral'], model_name text default null, provider text default 'openai')` 是一个扩展函数并返回 `text`。
- `public.ai_complete2(input_text text, system_text text default null, model_name text default null, provider text default 'openai')` 是一个扩展函数并返回 `text`。
- `public.ai_embed(input_text text, model_name text default 'text-embedding-3-small', provider text default 'openai')` 是一个扩展函数并返回 `real[]`。
- `public.ai_extract(input_text text, schema_hint text default null, model_name text default null, provider text default 'openai')` 是一个扩展函数并返回 `jsonb`。
- `public.ai_redact(input_text text, entity_types text[] default array['email','phone','ssn','name'], model_name text default null, provider text default 'openai')` 是一个扩展函数并返回 `text`。
- `public.ai_sentiment(input_text text, model_name text default null, provider text default 'openai')` 是一个扩展函数并返回 `text`。
- `public.ai_translate(input_text text, target_language text default 'Spanish', model_name text default null, provider text default 'openai')` 是一个扩展函数并返回 `text`。

### 要求与注意事项

- 该目录记录版本 `1.0.0`。
- 这是一个 database.dev/pg_tle 包；在发出引用的 `CREATE EXTENSION` 身份之前，请先注册或生成其包迁移。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行确认。
