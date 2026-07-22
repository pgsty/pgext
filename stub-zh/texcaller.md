## 用法

来源：

- [官方 PostgreSQL 绑定文档](https://vog.github.io/texcaller/group__postgresql.html)
- [扩展 SQL](https://github.com/vog/texcaller/blob/fc5fbc9862a6df3e907e2ee409e5d995aa175ef6/postgresql/texcaller.sql)
- [转换实现](https://github.com/vog/texcaller/blob/fc5fbc9862a6df3e907e2ee409e5d995aa175ef6/c/texcaller.c)

`texcaller` 在 PostgreSQL 后端内编译 TeX 家族源代码，并以 `bytea` 返回生成的 DVI 或 PDF 文档。它只适合严格受控的文档生成负载：每次调用都会创建临时文件，并以数据库服务器的操作系统身份启动外部 TeX 可执行文件。

### 核心流程

安装扩展，转义所有用户提供的纯文本片段，组装完整文档，再以有限的运行次数编译：

```sql
CREATE EXTENSION texcaller;

SELECT texcaller_convert(
  '\documentclass{article}' ||
  '\begin{document}' ||
  texcaller_escape_latex('A value with $ and % characters') ||
  '\end{document}',
  'LaTeX',
  'PDF',
  5
);
```

`texcaller_convert` 是 strict 函数并返回二进制输出。`max_runs` 必须至少为 2；编译会重复执行，直到辅助文件稳定或达到上限。

### 格式与函数

- `texcaller_convert(source, source_format, result_format, max_runs)` 负责编译文档。生成 PDF 时，源格式可以是 `TeX`、`LaTeX`、`XeTeX`、`XeLaTeX`、`LuaTeX` 与 `LuaLaTeX`；生成 DVI 时只支持 `TeX` 或 `LaTeX`。
- `texcaller_escape_latex` 为插入 LaTeX 源码的纯文本引用特殊字符。它不会验证整个模板，也不会建立安全沙箱。

所需可执行文件必须存在于 PostgreSQL 服务器进程的路径中；根据请求组合，实现会调用 `tex`、`latex`、`pdftex`、`pdflatex`、`xetex`、`xelatex`、`luatex` 或 `lualatex`。

### 安全与运维

子进程以批处理、遇错停止、文件行诊断和 `-no-shell-escape` 选项启动。禁用 shell escape 很重要，但不会让任意 TeX 变得安全：TeX 程序可能消耗大量计算资源，也可能访问已安装引擎与操作系统策略允许的文件。应撤销 public 执行权，只授予专用可信角色：

```sql
REVOKE EXECUTE ON FUNCTION
  texcaller_convert(text, text, text, integer)
FROM PUBLIC;

GRANT EXECUTE ON FUNCTION
  texcaller_convert(text, text, text, integer)
TO document_renderer;
```

每次调用都会派生子进程，并在 `TMPDIR` 或 `/tmp` 下使用私有目录。应设置语句超时、输入大小限制、并发控制、文件系统隔离，并监控 CPU、内存、磁盘与进程消耗。不要在已转义片段周围拼接不受信的 TeX 命令。

扩展控制文件报告版本 0、允许重定位，并且没有预加载设置。上游代码较旧；应在目标主机验证编译器二进制、字体、TeX 包、PostgreSQL 兼容性、清理行为以及逐字节输出。即使 SQL 输入不变，文档输出也可能随 TeX 安装环境变化。
