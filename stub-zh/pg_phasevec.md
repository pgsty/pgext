## 用法

来源：

- [README](https://github.com/Manuelreyesbravo/pg_phasevec/blob/v0.16.0/README.md)
- [Control](https://github.com/Manuelreyesbravo/pg_phasevec/blob/v0.16.0/pg_phasevec.control)
- [Base SQL](https://github.com/Manuelreyesbravo/pg_phasevec/blob/v0.16.0/sql/pg_phasevec--0.1.0.sql)
- [0.16.0 SQL upgrade](https://github.com/Manuelreyesbravo/pg_phasevec/blob/v0.16.0/sql/pg_phasevec--0.15.0--0.16.0.sql)
- [PGXN release](https://pgxn.org/dist/pg_phasevec/0.16.0/)
- [IVF implementation](https://github.com/Manuelreyesbravo/pg_phasevec/blob/v0.16.0/src/phasevec_ivf.c)
- [PostgreSQL 18 DSM interface](https://github.com/postgres/postgres/blob/REL_18_STABLE/src/include/storage/dsm_registry.h)
- [PostgreSQL 18 LWLock interface](https://github.com/postgres/postgres/blob/REL_18_STABLE/src/include/storage/lwlock.h)

`pg_phasevec` 在 PostgreSQL 行旁保存带来源的向量读数。每个读数记录已封存的读取器身份、精确文本哈希和时间；过期或来自其他读取器的读数不会参与检索。上游声明要求 PostgreSQL 18 及更高版本，不发起网络调用；但 0.16.0 源码存在下述 PostgreSQL 18 兼容性问题。

### PostgreSQL 18 兼容性

0.16.0 无法在 PostgreSQL 18 上编译：IVF 实现以不兼容的签名调用 `LWLockNewTrancheId` 和 `GetNamedDSMSegment`，还调用了 PostgreSQL 18 未提供的 `GetNamedDSA`。原生编译与 LLVM 编译均失败。因此，声明的最低版本不能证明 PostgreSQL 18 实际可用；应等待上游兼容性修复，再使用下述安装流程。此处尚未验证其他 PostgreSQL 版本。

### 封存读取器并检索

```sql
CREATE EXTENSION pg_phasevec;
SELECT extversion FROM pg_extension WHERE extname = 'pg_phasevec';

```

读取器通过模型摘要和维度注册。watch 指定源表、键和文本表达式；`phasevec.next` 将待处理行交给本地读取器，`phasevec.record` 将结果和文本哈希一起保存。`phasevec.search`、`find` 和 `alike` 只查询兼容读数。`coverage` 报告缺口，`reader_switch` 会把旧读取器的读数标为 foreign。

### 索引与边界

`ivf_build`、`ivf_maintain` 和 `ivf_absorb` 提供显式 IVF 维护。索引是近似的，coverage 阈值可以拒绝从未填满的 watch 返回结果。被监视的行修改后，读数会在同一事务中变为 stale。扩展不提供嵌入模型、读取器服务、网络传输或跨系统持久排序；读取器封存、权限和模型来源由调用方负责。模式固定为 `phasevec`，扩展不可迁移。

### 升级与安全

需要超级用户安装。源码包含初始安装及通向 0.16.0 的 SQL 升级链。应安装匹配的库和 SQL 文件并遵循升级链；0.16.0 加固内存分配边界、游标 NULL 处理和维护权限。模型摘要、读取器清单和源表访问应与数据库处于同一信任边界。
