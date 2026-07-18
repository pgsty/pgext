## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/petere/autopex/blob/886690f4cdf4977eb1917cde6696b7dac3fcc43c/README.md)
- [0 版本 SQL 对象](https://github.com/petere/autopex/blob/886690f4cdf4977eb1917cde6696b7dac3fcc43c/autopex--0.sql)
- [事件触发器实现](https://github.com/petere/autopex/blob/886690f4cdf4977eb1917cde6696b7dac3fcc43c/autopex.c)

`autopex` 是一个已归档的 PostgreSQL 9.3 时代实验，通过事件触发器拦截 `CREATE EXTENSION`。配合外部 Pex 工具，当 PostgreSQL 本地找不到扩展时，它会尝试自动下载、构建和安装。

```sql
CREATE EXTENSION autopex;
SELECT evtname, evtevent, evtenabled
FROM pg_event_trigger
WHERE evtname = 'autopex';
```

操作系统 PostgreSQL 账户必须安装 Pex 和构建环境。该设计会把数据库 DDL 请求转化为网络访问、任意上游源码获取、编译以及写入服务器安装目录，因此把数据库超级用户权限扩展成操作系统供应链执行能力。

不要在生产或当前 PostgreSQL 上使用 `autopex`。应改用经审查、可复现并由配置管理安装的软件包。若仅为保留历史测试环境，应禁止不可信 `CREATE` 权限、隔离出站网络、固定每个工件和校验和，并只在可丢弃主机运行。删除扩展后，还应确认其事件触发器已移除再继续执行 DDL。
