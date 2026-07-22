## 用法

来源：

- [官方 README](https://github.com/pgstuff/sys_syn_dblink/blob/0c7a18da219a2e71b83ce374e589cfa383a72518/README.md)
- [官方用户指南](https://github.com/pgstuff/sys_syn_dblink/blob/0c7a18da219a2e71b83ce374e589cfa383a72518/doc/sys_syn_dblink.adoc)
- [官方扩展 SQL](https://github.com/pgstuff/sys_syn_dblink/blob/0c7a18da219a2e71b83ce374e589cfa383a72518/sql/sys_syn_dblink.sql)

`sys_syn_dblink` 是 `sys_syn` 队列的接收端处理器。它通过具名 `dblink` 连接在源数据库认领批次、拉取变更、进行转换并写入本地目标表，再把处理状态推回源端。

### 设置

先在源端安装 `sys_syn` 并配置输入/输出队列。在接收数据库安装声明的依赖与本扩展，打开具名连接，定义输入、输出与写入组，然后生成处理表：

```sql
CREATE EXTENSION hstore;
CREATE EXTENSION dblink;
CREATE EXTENSION sys_syn_dblink;

SELECT dblink_connect('source_syn', 'service=sys_syn_source');

INSERT INTO sys_syn_dblink.in_groups_def
VALUES ('source-cluster', 'in');
INSERT INTO sys_syn_dblink.out_groups_def
VALUES ('source-cluster', 'out');
INSERT INTO sys_syn_dblink.put_groups_def
VALUES ('put');

SELECT sys_syn_dblink.proc_table_create(
    proc_schema     => 'processor_data',
    in_table_id     => 'customer',
    out_group_id    => 'out',
    put_group_id    => 'put',
    dblink_connname => 'source_syn'
);
```

可在 `put_table_transforms` 与 `put_column_transforms` 中添加可选行以改变表/列映射。生成的工作函数实现认领、拉取、处理与状态推送阶段。官方指南要求反复执行拉取/处理/推送循环，因为每次调用只处理有限批次；认领操作还使用显式远端事务。

### 依赖与运维

接收端需要 `hstore` 和 `dblink`，远端数据库需要 `sys_syn`。时间与双时间表类型还要求指定版本或补丁的 `temporal_tables`。已复核项目文档要求 PostgreSQL 9.5 或更高版本，但 `0.0.1` 使用大量生成式 PL/pgSQL 与旧 API，因此必须验证精确的现代服务器组合。

执行生成过程的会话必须保持具名 `dblink` 连接打开。应保护连接凭据与转换定义，因为它们决定远端访问和动态生成的本地 SQL。还应监控挂起/重试状态、限制批大小、正确串行化认领，并测试事务中断，避免行丢失或重复处理。目录将该版本标为预览版。
