## 用法

来源：

- [Official source](https://github.com/addiscorp/pg_harar/blob/b9f7a466d43ae65f9419013187dd1c0ff6490081/README.md)
- [Official source](https://github.com/addiscorp/pg_harar/blob/b9f7a466d43ae65f9419013187dd1c0ff6490081/pg_harar.control)
- [Official source](https://github.com/addiscorp/pg_harar/blob/b9f7a466d43ae65f9419013187dd1c0ff6490081/Cargo.toml)
- [Official source](https://github.com/addiscorp/pg_harar/blob/b9f7a466d43ae65f9419013187dd1c0ff6490081/sql/audit.sql)

`pg_harar` 0.1.0 在 PostgreSQL 17 中为 Harar 提供追加写入的哈希链审计表。令牌身份及其行级安全集成仍是计划中的功能，不属于当前审计实现。扩展不解密秘密，也不保存解密密钥。

### 基本流程

由超级用户将该非可信扩展安装到固定的 `harar` 模式。以下 SQL 流程不需要预加载或重启服务。

```sql
CREATE EXTENSION pg_harar;
INSERT INTO harar.audit (txn, org, actor, zone, via, action, resource, decision, by_role, policy_version)
VALUES ('txn-1', 'example', 'service.api', 'prod', 'broker',
        'read secret', 'DATABASE_URL', 'allow', 'reader', 1);
SELECT * FROM harar.verify_chain();
```

扩展设置 `seq`、`prev_hash`、`hash` 和数据库记录时间 `at`。可选的 `occurred_at` 是写入者提供的事件时间，也会进入哈希。`harar.verify_chain` 返回检查数量与首个断裂的序号；`first_broken` 为 NULL 表示链条与链头一致。

### 权限与并发

写入角色需要模式 USAGE 和表 INSERT 权限，通常也需要 SELECT。内部链头行仅供其所有者访问。UPDATE、DELETE 和 TRUNCATE 会被拒绝；会话使用副本行为时，审计触发器仍然启用。

插入通过同一个链头行串行执行，行锁一直持有到事务结束。应使用短事务，并配置空闲事务超时。在可重复读或可串行化隔离级别下，应重试发生序列化失败的事务。回滚的插入不会留下审计行。

### 完整性与恢复

所有者或超级用户可以禁用触发器并重写历史与链头。要发现这种篡改，需要把签名检查点保存在数据库机器之外；本地链条校验成功不能独立证明真实性。

`pg_dump` 包含审计行，但不包含链头。上游恢复流程在导入时禁用触发器，再由所有者禁用链头触发器、将链头恢复为最后一行的序号与哈希，最后验证链条。依赖备份之前，应演练完整恢复流程。
