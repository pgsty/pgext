## 用法

来源：

- [官方上游 README](https://github.com/masahikosawada/pg_logical_ddl/blob/0c459994931bb61a8f9637d7b2e54ff899ef82fa/README.md)
- [官方扩展控制文件 (pg_logical_ddl.control)](https://github.com/masahikosawada/pg_logical_ddl/blob/0c459994931bb61a8f9637d7b2e54ff899ef82fa/pg_logical_ddl.control)
- [官方实现源代码](https://github.com/masahikosawada/pg_logical_ddl/blob/0c459994931bb61a8f9637d7b2e54ff899ef82fa/pg_logical_ddl.c)

`pg_logical_ddl` — 实验性 DDL 重播实现，用于 PostgreSQL。在迁移、转换或集成相应数据时使用。上游明确表示该扩展尚未准备好生产环境。

### 核心工作流

构建带有上游链接的、正在进行中的 apply-worker 消息补丁的 PostgreSQL。安装库，在发布者和订阅者上预加载，并重启两个服务器：

```ini
shared_preload_libraries = 'pg_logical_ddl'
```

在发布者上选择要捕获的 DDL 命令标签：

```ini
pg_logical_ddl.log_command_tags = 'CREATE TABLE, ALTER TABLE, DROP TABLE'
```

创建常规逻辑复制对象，启用订阅上的消息传递：

```sql
-- publisher
CREATE PUBLICATION ddl_pub FOR ALL TABLES;

-- subscriber
CREATE SUBSCRIPTION ddl_sub
  CONNECTION 'host=publisher dbname=app user=replicator'
  PUBLICATION ddl_pub
  WITH (message = true);
```

在发布者上执行的 DDL 作为逻辑消息写入，并由订阅者的 apply 工作者使用捕获的角色和 `search_path` 重新执行。

### 重要设置

- `pg_logical_ddl.log_command_tags` 是一个不区分大小写的逗号分隔列表。其空默认值禁用 DDL 捕获。

### 要求与注意事项

- 经审核的控制文件、注册表或目录证据标识版本 `1.0`。
- 控制文件将扩展标记为可重定位。
- 上游明确表示该项目尚未准备好生产环境。
- 上游将项目的一部分或全部标记为实验性。
- 上游将项目描述为概念验证。
- 所需的核心补丁和 `message = true` 订阅选项在发布的 PostgreSQL 版本中不可用。
- 没有循环预防。在订阅上禁用捕获并使用单向复制。
- DDL 文本在语句级别重放。环境相关的名称、函数、权限和 `search_path` 内容可能会有所不同。
- 仅普通 `CREATE TABLE` 在审核的实现中自动注册到订阅中。
