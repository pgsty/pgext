## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/gerdansantos/pg_elephant/blob/ba00d779ccc306f36db14ec86a89208d279c2328/README.md)
- [1.0 版本 SQL 对象](https://github.com/gerdansantos/pg_elephant/blob/ba00d779ccc306f36db14ec86a89208d279c2328/pg_elephant--1.0.sql)
- [C 实现](https://github.com/gerdansantos/pg_elephant/blob/ba00d779ccc306f36db14ec86a89208d279c2328/pg_elephant.c)

`pg_elephant` 明确是一个 PostgreSQL C 扩展玩具和构建示例，只安装一个演示函数。

```sql
CREATE EXTENSION pg_elephant;
SELECT pg_elephant();
```

它只应用于可丢弃开发集群，验证 PGXS 编译、扩展打包和函数加载。尽管口号是“大象永不遗忘”，它并不提供持久化、缓存、审计历史或任何生产功能。

README 针对老旧 PostgreSQL 9.5 开发包，没有当前兼容、维护或安全契约。不要为示例向生产攻击面增加原生库。若把代码改作模板，发布前应替换名称与元数据，加入严格输入验证、正确 volatility/parallel/security 标记、回归测试、受支持版本 CI、升级路径和真正的许可证审查。
