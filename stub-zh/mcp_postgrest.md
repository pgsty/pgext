## 用法

来源：

- [官方上游 README](https://github.com/cybertheory/pgmcp/blob/26e25d6beb457ac064b971b9fa856cfdea1ba046/README.md)
- [官方扩展控制文件 (mcp_postgrest.control)](https://github.com/cybertheory/pgmcp/blob/26e25d6beb457ac064b971b9fa856cfdea1ba046/mcp_postgrest.control)
- [官方扩展 SQL (mcp_postgrest--0.1.2.sql)](https://github.com/cybertheory/pgmcp/blob/26e25d6beb457ac064b971b9fa856cfdea1ba046/sql/mcp_postgrest--0.1.2.sql)

`mcp_postgrest` — The mcp_postgrest PostgreSQL 扩展将您的数据库转换为符合 Anthropic 模型上下文协议 (MCP) 的 AI 动力工具接口。使用它来实现相应的向量、模型或检索工作流。在目标 PostgreSQL 构建中使用上述链接的上游固定版本作为 API 边界并进行测试。

### 核心工作流

```sql
CREATE EXTENSION mcp_postgrest;

SELECT call_tool('tool_name', '{"arg1": "value"}'::jsonb);
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `call_tool(tool_name TEXT, args JSONB)` 是一个扩展函数，返回 `JSONB`。
- `create_crud_tools_for_table(tablename TEXT)` 是一个扩展函数，返回 `VOID`。
- `generate_ai_tool(provider TEXT, api_key TEXT, tool_name TEXT, description TEXT, table_names TEXT[])` 是一个扩展函数，返回 `TEXT`。
- `mcp_on_create_table()` 是一个扩展函数，返回 `event_trigger`。
- `mcp_tools` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1.2`。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
