## 用法

来源：

- [Source snapshot README](https://github.com/maludb/maludb-core/blob/71cf34de6986281b36f8bec5b2c924907da97cb7/README.md)
- [Control file](https://github.com/maludb/maludb-core/blob/71cf34de6986281b36f8bec5b2c924907da97cb7/maludb_core.control)
- [Version 0.106.0 SQL](https://github.com/maludb/maludb-core/blob/71cf34de6986281b36f8bec5b2c924907da97cb7/sql/extension/maludb_core--0.106.0.sql)
- [Principals and scopes](https://github.com/maludb/maludb-core/blob/71cf34de6986281b36f8bec5b2c924907da97cb7/docs/principal-scoping.md)
- [Changelog](https://github.com/maludb/maludb-core/blob/71cf34de6986281b36f8bec5b2c924907da97cb7/CHANGELOG.md)
- [Ingest and retrieval example](https://github.com/maludb/maludb-core/blob/71cf34de6986281b36f8bec5b2c924907da97cb7/examples/01-ingest-to-replay.sql)

`maludb_core` 将机构记忆保存为文档、主体、主张、事实、事件及图关系，并提供文本和向量检索。本页描述声明版本为 0.106.0、面向 PostgreSQL 17 的固定源码快照。变更日志将其标为尚未发布；最新带标签的发行版为 v4.5.0，其中 SQL 扩展版本为 0.100.0。

### 启用租户模式

安装扩展文件后，以管理员身份在目标数据库中启用。依赖包括 `vector`、`btree_gist`、`pg_trgm`、`pgcrypto` 和 PL/pgSQL。扩展使用 C 共享库，但不要求共享预加载。核心模式固定，应用模式则必须逐个启用。

```sql
CREATE EXTENSION maludb_core CASCADE;
CREATE ROLE app LOGIN;
GRANT maludb_user TO app;
CREATE SCHEMA app AUTHORIZATION app;
ALTER ROLE app SET search_path TO app, maludb_core, public;
SELECT * FROM maludb_core.enable_memory_schema('app');

SET ROLE app;
SET search_path TO app, maludb_core, public;
SELECT * FROM maludb_subject;
RESET ROLE;
```

远程登录前需单独配置认证。`SET ROLE` 不会应用角色的登录默认值，因此示例显式设置路径。常规应用连接使用租户本地视图与函数。`maludb_user` 用于应用访问，只读访问使用 `maludb_read`。

### 摄取、检索与维护记忆

租户 API 包括用于图摄取的 `maludb_memory_ingest_edge` 和 `maludb_memory_ingest_extraction`，用于检索的 `maludb_memory_search` 和 `maludb_vector_search`，以及用于有界节点边导入的 `maludb_graph_import`。扩展保存模型输出；抽取和嵌入工作由独立客户端执行。官方摄取到回放示例展示来源登记、主张、已核实事实、文本检索、事件回放及通过替代关系纠错的过程。

`maludb_forget_document` 和 `maludb_forget_chunk` 删除选定材料。0.106.0 的三种检索路径均会过滤已标记删除的片段。处于法律保留状态的来源会拒绝删除，仍被引用的来源会保留并在结果中报告。应检查返回结果，不能假定所有关联对象均已删除。

### 为每个请求绑定主体

0.106.0 引入租户本地的主体和作用域授权。从已获授权、尚未绑定主体的租户会话中创建主体和授权，随后由受信任服务为每个请求限定访问范围：

```sql
SELECT maludb_principal_upsert('agent:44', 'agent', 'Sasha', 'agent:44');
SELECT maludb_principal_grant_scope('agent:44', 'dept:3', 'read');
BEGIN;
SET LOCAL maludb_core.principal_ref = 'agent:44';
SET LOCAL maludb_core.principal_scopes = '["agent:44", "dept:3"]';
SET LOCAL maludb_core.principal_readonly = 'on';
SELECT maludb_principal_whoami();
COMMIT;
```

未设置主体时保留不受此机制限制的租户行为；未知或禁用的主体无法访问。主体拥有自身主作用域的写入权限，写授权包含读权限；请求作用域只能缩小授权范围，不能扩大。敏感度上限进一步限制检索。租户共用词汇仍可能暴露某个主体名称的存在。

这些设置是与受信任服务之间的约定，不能防范持有租户数据库登录凭据的客户端，因为客户端可以自行设置它们。此类隔离需要数据库角色和认证机制。已绑定主体的会话不能管理主体或其授权。

### 升级与备份

在主机上安装匹配的库和 SQL 文件后，更新每个数据库，并在每个已启用的租户模式中重新生成接口：

```sql
ALTER EXTENSION maludb_core UPDATE TO '0.106.0';
SELECT maludb_core.maludb_core_version();
SELECT * FROM maludb_core.enable_memory_schema('app');
SELECT schema_name, enabled_version FROM maludb_core.malu$enabled_schema;
```

仅执行扩展迁移不会替换租户所有的接口对象。0.105.2 的索引迁移在构建索引期间可能阻塞写入。0.105.0 将存储数据登记为逻辑转储内容，但有意排除数据库主密钥和认证 pepper；应单独保存这些秘密，并遵循变更日志中的恢复流程。不能假定旧版本转储包含全部记忆数据。
