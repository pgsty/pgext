## 用法

来源：

- [README.md](https://api.pgxn.org/src/pg_phasevec/pg_phasevec-0.18.0/README.md)
- [pg_phasevec.control](https://api.pgxn.org/src/pg_phasevec/pg_phasevec-0.18.0/pg_phasevec.control)
- [sql/pg_phasevec--0.1.0.sql](https://api.pgxn.org/src/pg_phasevec/pg_phasevec-0.18.0/sql/pg_phasevec--0.1.0.sql)
- [sql/pg_phasevec--0.16.0--0.17.0.sql](https://api.pgxn.org/src/pg_phasevec/pg_phasevec-0.18.0/sql/pg_phasevec--0.16.0--0.17.0.sql)
- [sql/pg_phasevec--0.17.0--0.18.0.sql](https://api.pgxn.org/src/pg_phasevec/pg_phasevec-0.18.0/sql/pg_phasevec--0.17.0--0.18.0.sql)
- [src/phasevec_ivf.c](https://api.pgxn.org/src/pg_phasevec/pg_phasevec-0.18.0/src/phasevec_ivf.c)

`pg_phasevec` 0.18.0 在 PostgreSQL 行旁存储带来源的向量读数。每个读数标识已封存的读取器、精确文本哈希与时间，过期或来自其他读取器的读数不会参与检索。读取器在本地运行，由运维方提供。

### 启用与兼容性

```sql
CREATE EXTENSION pg_phasevec;
SELECT extversion FROM pg_extension WHERE extname = 'pg_phasevec';
```

安装需要超级用户，模式固定为 `phasevec`，不依赖其他扩展。上游记录了 PostgreSQL 18–19 测试。0.17.0 修复了 0.10.0–0.16.0 的 PostgreSQL 18 编译问题：PostgreSQL 18 使用后端独立 IVF 缓存，`phasevec.shared_ivf_cache` 关闭且不可启用；共享 IVF 要求 PostgreSQL 19。本次源码核验不代表已重新构建本地二进制。

### 读取器与检索

用模型摘要和维度注册读取器。watch 指定源表、键与文本表达式。`phasevec.next` 将工作交给读取器，`phasevec.record` 将结果代码与文本哈希一同保存。`phasevec.search`、`find` 和 `alike` 检索兼容读数；`coverage` 报告缺口，`reader_switch` 将旧读取器结果标为 foreign。扩展不提供嵌入模型或持久读取器服务。

`ivf_build`、`ivf_maintain` 和 `ivf_absorb` 显式维护近似索引。源行修改时，其读数在同一事务内变为 stale。coverage 阈值可能拒绝查询未完全填充的 watch；是否允许部分结果由调用方决定。

### 召回率与摘要

`phasevec.recall_approve` 保存所有者认可的基线，`phasevec.recall_check` 比较当前召回率。检查采用已保存的检索参数与抽样键，大样本可能代价较高；超过 100000 个有效读数时，参考检索本身也经过预筛选。缺少基线或样本不可用时会报告无法测量。这些函数及基线表不向 `PUBLIC` 开放，应显式授予监控权限。

0.18.0 新增 `phasevec.brief(watch, budget)`，默认字符预算为 8000。它汇总仍然有效的读数，并以调用者权限获取源文本，不会额外赋予源表访问权。

### 升级与信任边界

安装匹配的库与 SQL 文件，遵循所提供的升级链。在使用共享 IVF 的服务器上跨越 0.17.0 时，缓存格式变更后需要重启以释放旧共享内存段。新缓存也区分扩展安装实例，避免删除重建扩展后复用陈旧状态。

维护函数不向 `PUBLIC` 开放。读取器封存与模型摘要只能证明来源，不能保证模型质量；读取器清单、维护权限和源表权限应纳入数据库运维控制。
