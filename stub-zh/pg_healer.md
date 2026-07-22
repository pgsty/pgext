## 用法

来源：

- [官方 README](https://github.com/turnstep/pg_healer/blob/d313cccc2944a3c4a54757c5fa2693e9ae21d637/README.md)
- [官方控制文件](https://github.com/turnstep/pg_healer/blob/d313cccc2944a3c4a54757c5fa2693e9ae21d637/pg_healer.control)
- [官方安装 SQL](https://github.com/turnstep/pg_healer/blob/d313cccc2944a3c4a54757c5fa2693e9ae21d637/pg_healer--1.0.sql)
- [官方实现](https://github.com/turnstep/pg_healer/blob/d313cccc2944a3c4a54757c5fa2693e9ae21d637/pg_healer.c)
- [作者的设计文章](https://www.endpointdev.com/blog/2016/09/pghealer-repairing-postgres-problems/)

`pg_healer` 是一个已废弃的概念验证项目：它挂接 PostgreSQL 错误，并尝试用文件系统备份中的数据块替换损坏的堆页面。它绕过 PostgreSQL 的缓冲区与 WAL 机制直接修改关系文件，因此只能在一次性副本上用于研究损坏问题，不能当作恢复工具。

### 核心流程

必须预加载该动态库。控制文件没有声明对不受信任语言 `plperlu` 的依赖，因此要显式安装该语言；重启 PostgreSQL 后再安装扩展。

```conf
shared_preload_libraries = 'pg_healer'
```

```sql
CREATE EXTENSION plperlu;
CREATE EXTENSION pg_healer;
REVOKE ALL ON FUNCTION pg_healer_cauldron() FROM PUBLIC;

SELECT * FROM pg_healer_config;
SELECT * FROM pg_healer_log ORDER BY happened DESC;
```

错误钩子会监听内部数据损坏错误，查询 `pg_healer_config`，并尝试用 `PGDATA/pg_healer` 下副本中的数据块替换报告的数据块。预期的备份入口是 `pg_healer_cauldron()`，但它只执行原始文件复制，无法生成具备在线一致性的 PostgreSQL 备份。

### 对象

- `pg_healer_config` 按关系启用或禁用修复尝试。
- `pg_healer_log` 记录修复尝试。
- `pg_healer_cauldron()` 调用 Perl 备份辅助程序。
- `pg_healer_remove_from_buffer(regclass)` 尝试从共享缓冲区中逐出关系；C 实现要求超级用户权限。
- `pg_healer_corrupt(regclass, text)` 是故意制造损坏的辅助函数，绝不能用于有价值的数据。

### 运维边界

已复核的 `1.0` 版本即使作为自动修复原型也不安全。SQL 将 `pg_healer_corrupt` 声明为两个参数，但 C 函数读取了第三个参数；数据块编号解析又从单词 `block` 开始，而不是从后续数字开始。备份辅助程序的校验和验证尚未完成，直接文件写入也没有与 PostgreSQL 缓冲区、检查点、WAL 或并发访问协调。

不要在生产系统中保留预加载设置或任何扩展对象。遇到真实损坏时，应停止写入、保存原始文件、制作经过验证的物理副本，并使用受支持的恢复、备份或专业取证流程。
