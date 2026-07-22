## 用法

来源：

- [上游 README](https://github.com/zilder/pg_txn_status/blob/c1d49147b72a8c3230349d99a819b72a66881cde/README.md)
- [版本 0.1 SQL 定义](https://github.com/zilder/pg_txn_status/blob/c1d49147b72a8c3230349d99a819b72a66881cde/pg_txn_status--0.1.sql)
- [类型实现](https://github.com/zilder/pg_txn_status/blob/c1d49147b72a8c3230349d99a819b72a66881cde/pg_txn_status.c)

`pg_txn_status` 版本 `0.1` 提供一个单字节 `txn_status` 类型，用于存储六种应用自定义事务状态标签之一。它不会检查 PostgreSQL 事务或 WAL。

### 核心流程

```sql
CREATE EXTENSION pg_txn_status;

CREATE TABLE job_state (
  id bigint PRIMARY KEY,
  status txn_status NOT NULL
);

INSERT INTO job_state VALUES (1, 'begin'), (2, 'complete');
SELECT * FROM job_state WHERE status = 'complete'::txn_status;
```

可接受且区分大小写的值为 `begin`、`prepare`、`commit`、`rollback`、`complete` 和 `incomplete`。其他输入都会报错。

### 对象与注意事项

扩展定义输入/输出、二进制发送/接收函数，以及相等操作符 `=`。它没有定义不等、排序、哈希、B-tree 操作符类或事务管理 API。

如果模式演进、排序、索引或受维护的兼容性比单字节表示更重要，应使用 PostgreSQL 原生枚举或引用表。上游项目已废弃并使用老式 C 扩展接口；存储持久数据前，应在目标 PostgreSQL 版本上验证编译、转储恢复、复制与磁盘格式兼容性。
