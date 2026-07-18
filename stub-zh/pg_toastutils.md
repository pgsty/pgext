## 用法

来源：

- [上游固定版本 README 与警告](https://github.com/RhodiumToad/pg-toastutils/blob/2cb6af18a866af847a672ad18b3c5e9c681b64da/README.md)
- [1.0 版本安装 SQL](https://github.com/RhodiumToad/pg-toastutils/blob/2cb6af18a866af847a672ad18b3c5e9c681b64da/pg_toastutils--1.0.sql)
- [固定版本 TOAST 实现](https://github.com/RhodiumToad/pg-toastutils/blob/2cb6af18a866af847a672ad18b3c5e9c681b64da/toastutils.c)
- [固定版本扩展控制文件](https://github.com/RhodiumToad/pg-toastutils/blob/2cb6af18a866af847a672ad18b3c5e9c681b64da/pg_toastutils.control)

pg_toastutils 1.0 是实验性的物理存储诊断模块。它分类 varlena 表示形式，暴露磁盘 TOAST 指针细节，检查某个引用的 TOAST 项，并验证表中包括死亡物理行在内的全部 TOAST 引用。

### 检查可丢弃的表

```sql
CREATE EXTENSION pg_toastutils;

CREATE TABLE toast_demo (id integer, payload text);
ALTER TABLE toast_demo
    ALTER COLUMN payload SET STORAGE EXTERNAL;

INSERT INTO toast_demo
SELECT 1, repeat(md5(random()::text), 1000);

SELECT toast_type(payload), toast_ptr_detail(payload)
FROM toast_demo;

SELECT *
FROM toast_validate_table('toast_demo', false);
```

只有关系所有者或满足条件的数据库所有者才能调用表验证器。干净结果只是诊断证据，并不是修复操作或不存在损坏的保证。

### 崩溃与阻塞边界

上游明确警告这段实验代码可能破坏、崩溃甚至毁坏数据库。它直接依赖 2018 年前后 PostgreSQL 源码中的私有 heap、TOAST、元组可见性、快照、缓冲区和索引 API。这些布局与接口会随大版本变化；绝不能加载为另一精确服务器版本构建的二进制。

验证器对基础表取得 ExclusiveLock，对其 TOAST 表取得 AccessExclusiveLock，使用 SnapshotAny 扫描物理行，并读取每个引用块、检查事务元数据。大型或损坏表可能产生很重 I/O，并阻塞需要冲突锁的读写。应在副本或恢复副本上运行，设置锁与语句超时，每次只验证一个小关系。

指针细节函数会暴露物理关系 OID 与值 OID，而且若不修改权限就默认可由 PUBLIC 执行。应撤销全部函数，只授权专门诊断角色；调查前保留取证副本，并采用受支持的 PostgreSQL 恢复流程，而不是手工修改 TOAST。该仓库被归类为 abandoned，自 2018 年后没有更新。
