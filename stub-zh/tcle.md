## 用法

来源：

- [官方 README](https://github.com/julmon/tcle/blob/b709024ea638df07eff74205999f645471a8d3b7/README.md)
- [官方 1.0 版本扩展 SQL](https://github.com/julmon/tcle/blob/b709024ea638df07eff74205999f645471a8d3b7/tcle--1.0.sql)
- [官方表访问方法实现](https://github.com/julmon/tcle/blob/b709024ea638df07eff74205999f645471a8d3b7/tcleheap.c)

`tcle` 是一个实验性的透明单元级加密表访问方法。它扩展 PostgreSQL 堆访问方法，并在元组跨越存储边界时，使用 AES-256-CBC 加密由特殊类型声明的列。上游明确要求不要用于生产，目录也将其标为已放弃。

### 设置与核心流程

PostgreSQL 必须启用 OpenSSL 构建。将 `tcle` 加入 `shared_preload_libraries`，重启集群，创建扩展，并在访问加密表前设置数据库主口令：

```sql
CREATE EXTENSION tcle;
SELECT tcle_set_passphrase('replace-with-a-strong-passphrase');

CREATE TABLE private_records (
    id encrypt_numeric,
    label encrypt_text,
    observed_at encrypt_timestamptz
) USING tcleam;

INSERT INTO private_records
VALUES (1, 'confidential', now());
SELECT * FROM private_records;
```

`encrypt_text` 使用 plain 存储而非 TOAST，因此限制约 2 KB。`encrypt_numeric` 与 `encrypt_timestamptz` 分别表示数值和带时区时间戳。每张表都有随机表密钥，该密钥由数据库主密钥加密存储。主密钥从所给口令派生并保存在共享内存，而不是磁盘上。

主口令轮换必须在事务块之外执行：

```sql
SELECT tcle_change_passphrase(
    'replace-with-a-strong-passphrase',
    'replace-with-the-new-passphrase'
);
```

### 关键限制

已复核上游版本只支持 PostgreSQL 12 和 13，并自称实验性项目。它不支持对加密数据建立索引。明文仍可能出现在临时文件、SQL 语句与日志、客户端流量、转储和备份中；项目还警告主口令可能泄漏到日志。轮换表密钥需要重建/复制表，并重新创建约束、索引与触发器。

丢失主口令或密钥会导致数据无法恢复。应使用可丢弃数据测试重启、备份恢复、崩溃恢复、口令轮换、`VACUUM FULL`、复制与密钥丢失流程。真实工作负载应优先采用仍在维护的加密方案和外部密钥管理系统。
