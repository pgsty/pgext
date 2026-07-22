## 用法

来源：

- [官方 README](https://github.com/disqus/pg_audit/blob/9c4c21c1a4bd94554fdcbc474043eaba68fb44cd/README.md)
- [官方 control 文件](https://github.com/disqus/pg_audit/blob/9c4c21c1a4bd94554fdcbc474043eaba68fb44cd/pg_audit.control)
- [官方 0.1.3 版 SQL](https://github.com/disqus/pg_audit/blob/9c4c21c1a4bd94554fdcbc474043eaba68fb44cd/sql/pg_audit--0.1.3.sql)

Disqus 的 `pg_audit` 是一个为每张表生成审计表和行触发器的 SQL 生成器。它不是名称相近的 pgaudit 服务器审计扩展：这里把行镜像写入普通 PostgreSQL 表，而不是生成语句审计日志。

### 核心流程

安装后，`create_audit()` 以文本形式返回 DDL。先保存并审查这些文本，再作为独立事务执行：

```sh
psql -AtX -o actually_create_audit.sql \
  -c 'SELECT create_audit();' -U app_owner appdb

psql -X -1f actually_create_audit.sql -U app_owner appdb
```

调用 `create_audit(true)` 还会生成在 `replica` 模式下启用的触发器；无参数形式使用普通触发器设置。

### 生成的对象

对于符合条件的应用模式，生成器会创建对应的 `<schema>_audit` 模式和审计表。每条审计记录包含源字段的新旧值对，以及操作、时间戳、当前用户、会话用户等元数据；同时生成行触发器、辅助索引和现有数据的初始 `INSERT` 快照。

### API

- `create_audit() RETURNS SETOF text`：生成默认审计 DDL。
- `create_audit(enable_replica boolean) RETURNS SETOF text`：控制生成的触发器是否也在 replica 模式触发。

### 注意事项

必须遵循官方流程，在执行前审查生成的 SQL。生成器会排除 PostgreSQL 系统模式和审计模式，但会广泛处理安装角色可见的其他表。扩展依赖 PL/pgSQL，执行角色还需有权创建模式、表、索引和触发器。

版本 `0.1.3` 没有完整的模式演进流程；其 SQL 明确没有处理审计配置后源表新增字段的情况。审计行仍是普通可变数据，因此访问控制、保留策略与防篡改能力需要另行设计。
