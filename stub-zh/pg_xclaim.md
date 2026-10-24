## 用法

来源：

- [Official source](https://github.com/Fenoman/pg_xclaim/blob/49e5bc3982de87bf94b3108fa7ce29ae813a07b2/README.md)
- [Official source](https://github.com/Fenoman/pg_xclaim/blob/49e5bc3982de87bf94b3108fa7ce29ae813a07b2/pg_xclaim.control)
- [Official source](https://github.com/Fenoman/pg_xclaim/blob/49e5bc3982de87bf94b3108fa7ce29ae813a07b2/CHANGELOG.md)
- [Official source](https://github.com/Fenoman/pg_xclaim/blob/49e5bc3982de87bf94b3108fa7ce29ae813a07b2/sql/pg_xclaim--1.0.0-rc1--1.0.0-rc2.sql)
- [Official source](https://github.com/Fenoman/pg_xclaim/blob/49e5bc3982de87bf94b3108fa7ce29ae813a07b2/sql/pg_xclaim--1.0.0-rc2--1.0.0-rc3.sql)

`pg_xclaim` 1.0.0-rc3 是面向 PostgreSQL 16–18 的实验性共享内存认领机制，在独立命名空间内提供不等待、持续到事务结束的互斥。它不是持久化队列，其键也不与 PostgreSQL 咨询锁互通。

### 基本流程

将动态库加入已有预加载列表并保留其他项目，重启 PostgreSQL 后由超级用户创建扩展。副本也应保留预加载：恢复期间调用会失败，但提升后即可使用。

```ini
shared_preload_libraries = 'pg_xclaim'
```

```sql
CREATE EXTENSION pg_xclaim;
BEGIN;
SELECT xclaim.try(1, 100);
SELECT xclaim.count();
COMMIT;
```

`xclaim.try` 返回真，表示当前顶层事务持有认领；同一所有者重复获取会成功，顶层事务提交或回滚时释放。认领不会随保存点回滚释放，这与咨询事务锁不同。应统一迁移所有参与调用的位置，避免混用彼此独立的锁命名空间。

### 对象与限制

`xclaim.try` 接受一个 bigint 键或两个 int4 键，两种形式使用不同命名空间。`xclaim.try_many` 按输入顺序返回布尔数组，部分成功的认领会保留到事务结束。如果需要全部成功或全部失败，应在任一失败时回滚整个事务。NULL 输入会报错。

`xclaim.stats` 和 `xclaim.debug_snapshot` 是有角色限制的监控函数；详细调试与过期认领注入是超级用户工具，注入仅用于测试。`xclaim.session_reset` 用于异常恢复，不应当作常规连接池重置命令。按文档限制共享容量及每个后端的认领数量；禁用扩展可能取消互斥保护。

### rc3 升级

rc2 修复后端退出清理和分配失败处理；rc3 增加 Tantor SE 18 构建兼容性，没有新增 SQL 对象。安装匹配文件、重启服务，再在各数据库执行 `ALTER EXTENSION pg_xclaim UPDATE TO '1.0.0-rc3'`。该版本仍是候选版，采用前应测试故障与容量边界。
