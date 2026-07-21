## 用法

来源：

- [pgMemento v0.7.4 README](https://github.com/pgMemento/pgMemento/blob/v0.7.4/README.md)
- [pgMemento v0.7.4 control file](https://github.com/pgMemento/pgMemento/blob/v0.7.4/extension/pgmemento.control)
- [pgMemento wiki](https://github.com/pgMemento/pgMemento/wiki)
- [Changes from v0.7.3 to v0.7.4](https://github.com/pgMemento/pgMemento/compare/v0.7.3...v0.7.4)

pgmemento 是一个基于触发器的 PostgreSQL 审计跟踪扩展。它记录 DML 变更作为 JSONB 差异，按事务和表事件分组，并追踪审计后的模式变更，提供恢复和回滚辅助功能。当需要在 PostgreSQL 内查询行历史和事务上下文时，请使用此扩展。

### 创建与初始化

    CREATE EXTENSION pgmemento;
    SELECT pgmemento.init('public');

init 会为选定模式中的可跟踪表添加审计标识，并增加用于追踪行历史的 pgmemento_audit_id 身份列。首次在测试副本上运行此命令：审计表定义变更并安装事件和行触发器。

使用 start 和 stop 控制模式下的审计活动，当有意移除 pgMemento 的功能时，请参阅文档中的 drop 函数。不要手动删除扩展的触发器或审计标识符。

### 检查审计跟踪

中心数据模型包括：

- transaction_log：事务元数据和可选的应用程序上下文。
- table_event_log：事务内的表级事件。
- row_log：与表事件关联的 JSONB 行差异。
- audited_schema 和 audited_table 元数据：追踪的模式、表、列及其版本。

典型的调查会将事务与其表事件及行差异进行连接：

    SELECT t.id,
           t.txid_time,
           e.table_operation,
           r.audit_id,
           r.old_data,
           r.new_data
    FROM pgmemento.transaction_log AS t
    JOIN pgmemento.table_event_log AS e
      ON e.transaction_id = t.id
    JOIN pgmemento.row_log AS r
      ON r.event_key = e.event_key
    WHERE t.id = 12345;

在嵌入此查询之前，请先检查安装的视图和列名，因为审计模式细节可能因 pgmemento 版本而异。

### 恢复与回滚

pgmemento 提供了恢复函数，可以从审计跟踪中重构表状态，并提供了 revert_transaction 或相关辅助函数来应用补偿变更。将这些操作视为恢复操作：

1. 备份并验证备份；
2. 识别确切的事务及其依赖变更；
3. 尽可能预览重构的数据；
4. 在受控事务中运行此操作；
5. 验证约束、序列和应用程序不变量。

### 版本 0.7.4

版本 0.7.4 改变了行序列化方式，以避免 PostgreSQL 的 jsonb_build_object 参数限制对非常宽的负载，并增加了对 PostgreSQL 15 的支持。升级时，请在测试特定版本的升级脚本后使用 ALTER EXTENSION pgmemento UPDATE。

### 运营边界

- 审计触发器会增加每个跟踪变更的延迟和写入量。监控 row_log 增长及索引维护。
- 审计跟踪位于同一数据库中，不是备份、WAL 归档或防篡改外部审计存储的替代品。
- 模式初始化和 DDL 跟踪会改变应用程序表。与 pgmemento 协调迁移，而不是绕过其事件触发器。
- 限制直接写入 pgmemento 模式，并保护可能包含用户或应用信息的事务元数据。
