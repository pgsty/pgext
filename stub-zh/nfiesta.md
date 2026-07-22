## 用法

来源：

- [已复核版本的 README](https://gitlab.com/nfiesta/nfiesta_pg/-/blob/5bf6158866a327db97a1c8513f58a06228573ecb/README.md)
- [nFIESTA 项目 Wiki](https://gitlab.com/nfiesta/nfiesta_pg/-/wikis/home)
- [SQL 源码树](https://gitlab.com/nfiesta/nfiesta_pg/-/tree/5bf6158866a327db97a1c8513f58a06228573ecb/functions)
- [扩展 control 文件](https://gitlab.com/nfiesta/nfiesta_pg/-/blob/5bf6158866a327db97a1c8513f58a06228573ecb/nfiesta.control)

`nfiesta` 是森林资源清查估算与分析系统 nFIESTA 的 PostgreSQL 估算引擎。它把国家森林清查野外样本与辅助数据（通常是遥感产品）结合，实现单阶段与两阶段估算、配置、ETL、结果和可加性工作流。

control 文件声明依赖 `postgis`、`plpython3u`、`htc` 和 `nfiesta_sdesign`。尤其是 `plpython3u` 属于不受信任过程语言；即使 `nfiesta` 本身标记为 `superuser = false`，它仍会建立数据库与服务器操作系统之间的信任边界。

```sql
CREATE EXTENSION nfiesta CASCADE;
SELECT e.extname, e.extversion
FROM pg_extension AS e
WHERE e.extname IN (
  'postgis', 'plpython3u', 'htc', 'nfiesta_sdesign', 'nfiesta'
);
```

扩展安装的是大型领域模式，而非一个独立函数。应按照精确版本的上游安装指南和最小工作示例，加载有文档的测试数据，配置估算期与模型，再把总量、比率、方差和置信结果与独立计算的参考值比较。

`nfiesta.fn_make_estimate(integer)` 会把配置好的估算结果写入 `nfiesta.t_result`。`nfiesta.fn_system_utilization()` 使用 `plpython3u` 编写，会调用 Python 的 `os.getloadavg()` 与 `os.cpu_count()`，返回服务器负载、CPU 数量和当前角色的连接上限。应限制该 API，并审计每个 Python 函数对文件、进程、环境变量和网络的访问。

统计正确性取决于抽样设计、分层、权重、辅助变量、缺失数据策略、模型配置和具体版本 SQL。应保留输入血缘与配置，隔离 ETL 和 GUI 角色，审计 security-definer 与动态 SQL 函数，并固定扩展版本以保证报告可复现。升级可能改变大量函数、视图、配置表和结果代次，必须在副本上测试迁移与回滚。
