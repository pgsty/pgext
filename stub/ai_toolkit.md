## Usage

Sources:

- [Official extension control file](https://github.com/AjinkyaTaranekar/ai_toolkit/blob/49c3e95059e2c050e61400829004b70dc08992b2/ai_toolkit.control)
- [Official upstream README](https://github.com/AjinkyaTaranekar/ai_toolkit/blob/49c3e95059e2c050e61400829004b70dc08992b2/README.md)

`ai_toolkit` 1.0 puts natural-language query generation, query and error explanation, and persistent prompt context inside PostgreSQL. Its most important boundary is that `ai_toolkit.query(text)` sends database context to a configured external model and executes the SQL returned by that model.

### Core Workflow

Configure all four required settings, restart PostgreSQL, and create the extension. The reviewed upstream version requires PostgreSQL 18.

```conf
ai_toolkit.ai_provider = 'openai'
ai_toolkit.ai_api_key = 'sk-...'
ai_toolkit.ai_model = 'gpt-4o'
ai_toolkit.prompt_file = '/path/to/query_system_prompt.txt'
```

```sql
CREATE EXTENSION ai_toolkit;

SELECT ai_toolkit.explain_query(
  'SELECT * FROM users WHERE created_at > now() - interval ''7 days'''
);

SELECT ai_toolkit.query('show all users who signed up last week');
```

Use `ai_toolkit.explain_query(text)` and `ai_toolkit.explain_error(text)` for advisory text. Treat `ai_toolkit.query(text)` as a privileged execution path, not as a read-only SQL generator.

### Memory and Utility Functions

- `ai_toolkit.set_memory(category, key, value, notes)` stores context for later prompts.
- `ai_toolkit.get_memory(category, key)`, `ai_toolkit.view_memories()`, and `ai_toolkit.search_memory(search_term)` retrieve that context.
- `ai_toolkit.help()` returns the extension's built-in help.

### Operational Boundaries

Calls are synchronous backend operations and depend on provider availability, latency, rate limits, and cost. The API key is held in PostgreSQL configuration, while prompts can include schema, data, errors, and stored memory. Restrict function execution, use a low-privilege role and transaction guardrails, review provider retention policy, and never trust generated SQL without an application-level approval boundary.
