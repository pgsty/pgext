## 用法

来源：

- [官方 README](https://github.com/vidardb/pgrocks-fdw/blob/8f89fcad7e9c471e8dacfd167457716b7cf3d6e5/README.md)
- [0.0.1 版 SQL API](https://github.com/vidardb/pgrocks-fdw/blob/8f89fcad7e9c471e8dacfd167457716b7cf3d6e5/kv_fdw--0.0.1.sql)
- [扩展控制文件](https://github.com/vidardb/pgrocks-fdw/blob/8f89fcad7e9c471e8dacfd167457716b7cf3d6e5/kv_fdw.control)

`kv_fdw` 将 RocksDB 作为 PostgreSQL 外部表的内嵌存储，因而不需要单独的远程数据库服务。它支持基本读写，但上游明确说明 `0.0.1` 版只面向 PostgreSQL 13，且尚未达到产品可用状态。

### 核心流程

安装兼容的 RocksDB 库，将 `kv_fdw` 加入 `shared_preload_libraries`，重启 PostgreSQL 后再创建扩展与服务器。

```sql
CREATE EXTENSION kv_fdw;
CREATE SERVER kv_server FOREIGN DATA WRAPPER kv_fdw;
CREATE FOREIGN TABLE student (
    id integer,
    name text
) SERVER kv_server;

INSERT INTO student VALUES (20757123, 'Rafferty');
SELECT * FROM student WHERE id = 20757123;
```

外部表的第一个属性就是键。要表达复合键，上游建议将 PostgreSQL 复合类型作为第一个属性。文档中的流程没有为该包装器提供服务器选项。

### 限制与恢复

`kv_fdw` 不支持回滚或中止。外部表创建后，文档所述实现不能增加或删除列，也没有二级索引。由此不能套用普通 PostgreSQL 事务预期：应演练故障与崩溃恢复，避免将关键关系变更与 RocksDB 写入混在同一业务步骤中，并保留独立备份策略。

删除外部表不代表其磁盘 RocksDB 数据已经移除。应遵循上游卸载流程，并谨慎检查 `data_directory`。应固定 RocksDB `6.11.4`、PostgreSQL 13 以及审阅构建所用的扩展提交；ABI 或存储格式漂移可能导致启动或数据访问失败。
