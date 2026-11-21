## 用法

来源：

- [rc5 release](https://github.com/Fenoman/pg_xclaim/releases/tag/v1.0.0-rc5)

- [Official source](https://github.com/Fenoman/pg_xclaim/blob/v1.0.0-rc5/README.md)
- [Official source](https://github.com/Fenoman/pg_xclaim/blob/v1.0.0-rc5/pg_xclaim.control)
- [Official source](https://github.com/Fenoman/pg_xclaim/blob/v1.0.0-rc5/CHANGELOG.md)
- [Official source](https://github.com/Fenoman/pg_xclaim/blob/v1.0.0-rc5/sql/pg_xclaim--1.0.0-rc1--1.0.0-rc2.sql)
- [Official source](https://github.com/Fenoman/pg_xclaim/blob/v1.0.0-rc5/sql/pg_xclaim--1.0.0-rc4--1.0.0-rc5.sql)

`pg_xclaim` 1.0.0-rc5 是面向 PostgreSQL 16–18 的实验性共享内存认领机制，在独立命名空间内提供不等待、持续到事务结束的互斥。它不是持久化队列，其键也不与 PostgreSQL 咨询锁互通。

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

`xclaim.try` 接受一个 bigint 键或两个 int4 键，分别使用独立命名空间。`xclaim.try_many` 按输入顺序返回布尔数组，部分成功的申请会保留到事务结束；需要全有或全无语义时应回滚整个事务。NULL 输入会报错。

`xclaim.stats` 和 `xclaim.debug_snapshot` 是有角色限制的监控函数，详细调试需要更高权限。rc5 已从生产库和模式中移除过期申请注入器。`xclaim.session_reset` 用于异常恢复，不应作为连接池常规重置命令。应限制共享容量与每个后端的申请数量；禁用扩展会影响互斥保障。

### rc5 升级

rc5 在获取分区锁前预先建立本地申请条目，修复内存分配失败时的清理，并移除生产环境中的过期条目注入器。安装匹配文件、重启服务器后，在每个数据库执行 `ALTER EXTENSION pg_xclaim UPDATE TO '1.0.0-rc5'`。提供的升级链覆盖 rc1 至 rc5。可执行 SQL 中的版本字符串保留候选版后缀。该软件仍处于实验阶段，采用前应测试清理、取消和容量边界。
