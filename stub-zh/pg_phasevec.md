## 用法

来源：

- [README](https://github.com/Manuelreyesbravo/pg_phasevec/blob/2beae566ff111c205072463a7974099d6ca562d6/README.md)
- [Control](https://github.com/Manuelreyesbravo/pg_phasevec/blob/2beae566ff111c205072463a7974099d6ca562d6/pg_phasevec.control)
- [Base SQL](https://github.com/Manuelreyesbravo/pg_phasevec/blob/2beae566ff111c205072463a7974099d6ca562d6/sql/pg_phasevec--0.1.0.sql)
- [0.16.0 SQL upgrade](https://github.com/Manuelreyesbravo/pg_phasevec/blob/2beae566ff111c205072463a7974099d6ca562d6/sql/pg_phasevec--0.15.0--0.16.0.sql)
- [PGXN release](https://pgxn.org/dist/pg_phasevec/0.16.0/)

`pg_phasevec` 在 PostgreSQL 行旁保存带来源的向量读数。每个读数记录已封存的读取器身份、精确文本哈希和时间；过期或来自其他读取器的读数不会参与检索。本版本面向 PostgreSQL 18 及更高版本，不发起网络调用。

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
