## 用法

来源：

- [Official extension control file](https://github.com/AjinkyaTaranekar/ai_toolkit/blob/49c3e95059e2c050e61400829004b70dc08992b2/ai_toolkit.control)
- [Official upstream README](https://github.com/AjinkyaTaranekar/ai_toolkit/blob/49c3e95059e2c050e61400829004b70dc08992b2/README.md)

`ai_toolkit` 1.0 将自然语言查询生成、查询与错误解释以及持久化提示上下文放进 PostgreSQL。最重要的边界是：`ai_toolkit.query(text)` 会把数据库上下文发送给配置的外部模型，并执行模型返回的 SQL。

### 核心流程

配置全部四项必需参数，重启 PostgreSQL，然后创建扩展。这里审阅的上游版本要求 PostgreSQL 18。

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

用 `ai_toolkit.explain_query(text)` 和 `ai_toolkit.explain_error(text)` 获取建议文本。应把 `ai_toolkit.query(text)` 当成特权执行路径，而不是只读 SQL 生成器。

### 记忆与辅助函数

- `ai_toolkit.set_memory(category, key, value, notes)` 保存供后续提示使用的上下文。
- `ai_toolkit.get_memory(category, key)`、`ai_toolkit.view_memories()` 和 `ai_toolkit.search_memory(search_term)` 用于读取这些上下文。
- `ai_toolkit.help()` 返回扩展内置帮助。

### 运维边界

调用会同步占用后端，并受供应商可用性、延迟、限流与费用影响。API 密钥保存在 PostgreSQL 配置中，而提示可能包含模式、数据、错误和已存记忆。应限制函数执行权限，使用低权限角色与事务护栏，审查供应商的数据保留策略，并且不要在缺少应用层审批边界时信任生成的 SQL。
