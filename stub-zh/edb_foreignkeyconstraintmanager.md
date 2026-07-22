## 用法

来源：

- [官方扩展控制文件](https://github.com/vibhorkum/edb_foreignkeyconstraintmanager/blob/50a89fb337426e810eecca4f3e4f1b00d7bb1023/edb_foreignkeyconstraintmanager.control)
- [官方上游文档](https://github.com/vibhorkum/edb_foreignkeyconstraintmanager/blob/50a89fb337426e810eecca4f3e4f1b00d7bb1023/README.md)

`edb_foreignkeyconstraintmanager` 用于在 Oracle 兼容模式的 EDB Postgres Advanced Server 上管理旧式分区表的外键行为。它组合原生约束与 `refint` 触发器，并依赖 EDB 专用目录，因此不能作为社区 PostgreSQL 或现代声明式分区的可移植方案。

### 核心流程

安装所需的 `refint` 扩展和管理器，再通过 `edb_util` API 创建关系：

```sql
CREATE EXTENSION IF NOT EXISTS refint;
CREATE EXTENSION edb_foreignkeyconstraintmanager;

SELECT edb_util.create_fk_constraint(
  'parent',
  ARRAY['id'],
  'child',
  ARRAY['parent_id'],
  'restrict'
);
```

最后一个参数选择 `cascade`、`restrict` 或 `setnull` 行为。函数返回布尔值；普通关系使用原生外键，被引用关系为分区表时则建立基于触发器的约束执行机制。

### 重要对象

- `edb_util.create_fk_constraint(regclass, text[], regclass, text[], text)` 创建单列或多列关系以及所需约束或触发器。
- `edb_util.alter_table_drop_partition(name, name, name[], text)` 在维持受管关系的同时协调删除 EDB 分区。
- 受管分区触发器的名称以 `EDB_partition_` 开头。

### 兼容性与维护

实现会检查 EnterpriseDB 与 `redwood` 数据库方言，并读取 EDB 专用的分区目录。只能在该项目所针对的准确 EDB Postgres Advanced Server 世代上使用。新增分区后应重新执行 `edb_util.create_fk_constraint(...)`；删除分区时应使用配套的分区删除函数，否则约束或触发器清理可能需要人工处理。采用前应测试复合键、空值处理、所有级联模式、并发写入和分区 DDL，并只允许严格受控的管理角色执行这些动态 DDL 函数。
