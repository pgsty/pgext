## 用法

来源：

- [官方双语 README](https://github.com/maxoodf/pg_mystem/blob/38554744c42e15c6f10132dcb314118a39bc85f1/README.md)
- [1.0.1 版 SQL API](https://github.com/maxoodf/pg_mystem/blob/38554744c42e15c6f10132dcb314118a39bc85f1/pg_mystem--1.0.1.sql)
- [扩展控制文件](https://github.com/maxoodf/pg_mystem/blob/38554744c42e15c6f10132dcb314118a39bc85f1/pg_mystem.control)

`pg_mystem` 将文本发送给长期运行的 Yandex Mystem 工作进程，并从 SQL 返回俄语形态规范化结果。它是外部 `mystem` 可执行程序的桥梁，而不是 PostgreSQL 内建的全文检索词典。

### 核心流程

把适合当前体系结构的 `mystem` 二进制文件安装到 PostgreSQL 共享数据目录。将 `pg_mystem` 加入 `shared_preload_libraries`，为 `max_worker_processes` 预留足够名额，重启服务端，然后创建扩展。

```sql
CREATE EXTENSION pg_mystem;

SELECT mystem_convert('Ехал грека через реку, сунул грека руку в реку');
```

唯一的 SQL 入口是 `mystem_convert(text) RETURNS text`。扩展将它声明为 `STRICT`，所以 NULL 文档会直接返回 NULL，不调用工作进程。

### 配置与安全

`DOC_LEN_MAX` 与 `MYSTEM_PROCS` 是 Makefile 编译期设置，并非运行时 GUC；修改任一项都需要重新构建并安装扩展。上游建议至少保留 `MYSTEM_PROCS + 1` 个工作进程槽位，但其吞吐估算高度依赖硬件，因此应使用代表性文本做基准测试，而不能只按该估算确定容量。

PostgreSQL 服务端账号会执行另行安装的第三方二进制文件和后台工作进程。应固定并校验该二进制文件、限制其文件系统权限，并考虑工作进程故障与队列饱和，同时测试非俄语文本、编码、超长输入、超时和重启行为。虽然 SQL 声明把 `mystem_convert()` 标记为 `IMMUTABLE`，结果仍取决于外部二进制与词典版本；依赖发生变化时，应重建索引或刷新已存储的派生值。
