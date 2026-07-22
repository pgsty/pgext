## 用法

来源：

- [官方 README](https://github.com/postgrespro/pg_backtrace/blob/de245a5f6b42ba870259ba68b4fbdd0c8b7bae69/README)
- [官方扩展 SQL](https://github.com/postgrespro/pg_backtrace/blob/de245a5f6b42ba870259ba68b4fbdd0c8b7bae69/pg_backtrace--1.0.sql)
- [官方 C 实现](https://github.com/postgrespro/pg_backtrace/blob/de245a5f6b42ba870259ba68b4fbdd0c8b7bae69/pg_backtrace.c)

`pg_backtrace` 是一个底层诊断扩展，可为指定级别的 PostgreSQL 错误附加原生调用栈，并为特定信号记录回溯。它适合在无法使用调试器或核心转储时调查后端故障或卡顿，不是常规应用可观测性工具。

### 初始化与使用

创建扩展只会安装两个函数。某个后端首次调用函数时，才会加载库并初始化钩子：

```sql
CREATE EXTENSION pg_backtrace;
SELECT pg_backtrace_init();

SET pg_backtrace.level = 'ERROR';
SELECT 1 / 0;
```

`pg_backtrace.level` 选择执行器钩子附加回溯的错误级别，默认为 `ERROR`。`pg_backtrace_init()` 用于强制加载库。`pg_backtrace_sigsegv()` 会故意触发段错误，只能作为破坏性诊断测试；绝不能在承载有效工作的后端中调用。

初始化后，扩展会为 `SIGSEGV`、`SIGBUS`、`SIGFPE` 与 `SIGINT` 安装处理器。向后端发送 `SIGINT` 可在继续执行原有信号处理器前记录当前原生调用栈。

### 注意事项

该代码会替换每个已初始化后端中的信号处理器和执行器钩子。使用前应测试它与其他基于钩子或信号的扩展是否冲突。符号解析质量取决于平台和链接选项：PostgreSQL 通常没有使用 `-rdynamic` 构建，静态符号不可见，未解析地址可能需要 `addr2line` 或匹配的调试二进制文件。原生地址和函数名可能向客户端与日志暴露敏感实现细节，因此应限制初始化权限和日志访问。该实现面向提供 `backtrace()` 的类 Unix 系统。
