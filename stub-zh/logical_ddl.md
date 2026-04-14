## 用法
- GitHub 仓库: [`samedyildirim/logical_ddl`](https://github.com/samedyildirim/logical_ddl)
- README: [samedyildirim/logical_ddl/blob/master/README.md](https://github.com/samedyildirim/logical_ddl/blob/master/README.md)

`logical_ddl` 会捕获表上受支持的 DDL 变更，并通过内置逻辑复制进行回放。这个项目的目标是减少手工执行 DDL 的次数，并避免发布端和订阅端之间的模式漂移。

README 指出，PostgreSQL 11 及以上版本受支持，且扩展需要在超级用户权限下运行。

### 捕获内容

支持的 DDL 操作包括：

- `ALTER TABLE ... RENAME TO ...`
- `ALTER TABLE ... RENAME COLUMN ... TO ...`
- `ALTER TABLE ... ADD COLUMN ...`
- `ALTER TABLE ... ALTER COLUMN ... TYPE ...`
- `ALTER TABLE ... DROP COLUMN ...`

README 还说明，内置类型、数组、复合类型、域和枚举都受支持，但复合类型、域和枚举的定义本身不会被复制。

### 发布端

```sql
CREATE EXTENSION logical_ddl;
INSERT INTO logical_ddl.settings VALUES (true, 'source1');
INSERT INTO logical_ddl.publish_tablelist (relid) VALUES ('table1'::regclass);
ALTER PUBLICATION log_pub_1 ADD TABLE logical_ddl.shadow_table;
```

扩展通过事件触发器捕获 DDL，把结果存入 `logical_ddl.shadow_table`，再通过逻辑复制发布这张表。

### 订阅端

```sql
CREATE EXTENSION logical_ddl;
INSERT INTO logical_ddl.settings VALUES (false, 'source1');
INSERT INTO logical_ddl.subscribe_tablelist (source, relid) VALUES ('source1', 'table1'::regclass);
ALTER SUBSCRIPTION log_sub_1 REFRESH PUBLICATION;
```

在订阅端，扩展会监听传入的 DDL 记录，并把生成的 SQL 回放到目标表上。

### 数据模型

- `logical_ddl.settings` 控制节点是发布端、订阅端，还是两者兼具。
- `logical_ddl.publish_tablelist` 控制要捕获哪些表和命令类型。
- `logical_ddl.subscribe_tablelist` 控制哪些表接收回放后的 DDL。
- `logical_ddl.shadow_table` 存储已捕获的命令。
- `logical_ddl.applied_commands` 跟踪生成的 SQL 以及执行是否失败。

### 注意事项

- DDL 支持范围仅限于上面列出的操作。
- 默认表达式、约束、索引以及 `USING` 表达式都未实现。
- README 提到该项目仍在开发中，后续可能出现不兼容变更。

### 范围

上游 README 已经足以说明扩展模型、可捕获的 DDL 类型、发布端/订阅端配置以及已知限制，因此不需要额外的文档页或主页。
