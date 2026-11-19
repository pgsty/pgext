## 用法

来源：

- [Official source](https://github.com/noizu-labs-ml/elixir-mcp-lib/blob/f3401ad75ba2299648a8a0a7f80c559a9fe51ae9/pg/README.md)
- [Official source](https://github.com/noizu-labs-ml/elixir-mcp-lib/blob/f3401ad75ba2299648a8a0a7f80c559a9fe51ae9/pg/pg_mcp/Cargo.toml)
- [Official source](https://github.com/noizu-labs-ml/elixir-mcp-lib/blob/f3401ad75ba2299648a8a0a7f80c559a9fe51ae9/pg/pg_mcp/pg_mcp.control)
- [Official source](https://github.com/noizu-labs-ml/elixir-mcp-lib/blob/f3401ad75ba2299648a8a0a7f80c559a9fe51ae9/pg/pg_mcp/sql/pg_mcp--0.3.0.sql)
- [Official source](https://github.com/noizu-labs-ml/elixir-mcp-lib/blob/f3401ad75ba2299648a8a0a7f80c559a9fe51ae9/pg/pg_mcp/src/session.rs)
- [Official source](https://github.com/noizu-labs-ml/elixir-mcp-lib/blob/f3401ad75ba2299648a8a0a7f80c559a9fe51ae9/pg/pg_mcp/src/lib.rs)
- [Official source](https://github.com/noizu-labs-ml/elixir-mcp-lib/blob/f3401ad75ba2299648a8a0a7f80c559a9fe51ae9/docs/pg-mcp-install.md)
- [Official source](https://github.com/noizu-labs-ml/elixir-mcp-lib/blob/f3401ad75ba2299648a8a0a7f80c559a9fe51ae9/LICENSE)

`pg_mcp` 0.3.0 是 Noizu MCP 仓库中的 Rust PostgreSQL 扩展，提供 `mcp_fdw`、MCP 调用函数和生成的 SQL 对象。Cargo 功能矩阵覆盖 PostgreSQL 16–18，使用 pgrx 0.19.2。安装需要超级用户，扩展自行创建 `mcp` 模式、不可迁移，且无需共享预加载。

### 连接 MCP 服务器

使用下例前，先准备可达的 MCP Streamable HTTP 服务及引用的秘密表。秘密查询以调用者的数据库权限执行，应按对应源码协议配置令牌行和授权。

```sql
CREATE EXTENSION pg_mcp;
CREATE SERVER engine FOREIGN DATA WRAPPER mcp_fdw
  OPTIONS (url 'https://engine.internal/mcp', timeout_ms '15000');
CREATE USER MAPPING FOR CURRENT_USER SERVER engine
  OPTIONS (token_secret 'mcp_secrets.engine');
CREATE SCHEMA mcp_tools;
SELECT mcp.import('engine', 'mcp_tools');
SELECT * FROM mcp.generate_functions('engine', 'mcp_tools');
SELECT mcp.refresh('engine');
```

control 文件要求超级用户，优先于运维指南中较宽松的描述。不要另外创建扩展拥有的 `mcp` 模式。仓库保存的安装 SQL 含开发测试包装函数；应使用与生产库匹配生成的 SQL，不能将这份启用测试的快照与任意二进制搭配。

### SQL 接口

`mcp.call_tool` 返回 JSON，`mcp.call_tool_text` 返回文本。`mcp.get_prompt`、`mcp.read_resource` 与 `mcp.complete` 提供相应的 MCP 调用。通用模式导入的目录和读取表展示工具与资源；向 `tool_calls` 插入一行会调用一次外部工具。生成的工具表、函数和视图记录在 `mcp.generated` 中；重新生成仅用 RESTRICT 替换已登记对象，因此依赖它们的用户对象可能阻止操作。

可选的引擎布局使用 `all_upstreams` 和 `per_upstream_schema`。工具可见性取决于远端身份，而非仅由本地表权限决定。`mcp.refresh` 丢弃当前后端的会话和目录缓存，不会自动重新生成 SQL 对象。

### 权限与外部影响

为各角色配置独立 USER MAPPING，并明确授予外部服务器、模式和对象权限。共享 PUBLIC 映射会合并远端身份。迁移文件中优先使用秘密引用而非直接写令牌。非回环地址必须使用 HTTPS，匿名认证也只允许回环地址。超时限制单次交换，远端错误或服务不可达可能使 SQL 调用失败。

远端工具执行会产生外部影响，不能假设它会随 PostgreSQL 事务回滚。生成的 SELECT 包装函数可能调用工具，应保留只读注解检查并审核任何覆盖设置。协议中的 `sql/*` 系列仍属实验性功能。升级和恢复时保留外部管理的秘密及审计对象，并使用匹配库文件的 SQL 升级链。
