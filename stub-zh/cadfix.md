## 用法

来源：

- [官方上游 README](https://github.com/zqqiang/node-cad/blob/739ff348b4d2c77b275c3a0fe87682c14ffd8181/README.md)
- [官方扩展控制文件 (cadfix.control)](https://github.com/zqqiang/node-cad/blob/739ff348b4d2c77b275c3a0fe87682c14ffd8181/db/cadfix/cadfix.control)
- [官方实现源代码](https://github.com/zqqiang/node-cad/blob/739ff348b4d2c77b275c3a0fe87682c14ffd8181/db/cadfix/cadfix.c)

`cadfix` 是历史 `node-cad` 原型中的 PostgreSQL 侧 C 库。它通过 Open CASCADE 加载 CAD 数据，并暴露由随附的 Node.js 应用程序使用的辅助函数；它不是一个自包含的现代 PostgreSQL 扩展。

### 核心工作流

上游 README 安装了 `cadfix` 共享库并手动注册其入口点：

```sql
CREATE FUNCTION cadinit(cstring)
RETURNS integer AS 'cadfix' LANGUAGE C;

CREATE FUNCTION full_edge(cstring, cstring)
RETURNS integer AS 'cadfix' LANGUAGE C;

SELECT cadinit('path/to/cad.step');
SELECT full_edge('evaluate', 'path/to/evaluate.csv');
SELECT full_edge('import', 'path/to/import.csv');
```

`cadinit` 打开一个 CAD 文件。`full_edge` 支持上游的 `evaluate` 和 `import` 模式，用于边数据。

### 要求与注意事项

- 审查过的控制文件标识版本 `1.0`，名称为 `$libdir/cadfix`，并且是非可重定位的，但仓库中没有提供版本化的扩展 SQL。
- 创建 C 语言函数并加载任意服务器库需要超级用户的权限。
- 文档中的构建目标是 PostgreSQL 9.5.1、Visual Studio 2012、32 位 Node.js 4.3.2 和捆绑的 Open CASCADE 工作流。将其视为历史源代码，而不是与当前 PostgreSQL 兼容性的证据。
- 文件路径由服务器进程解释。在使用不受信任的 CAD 输入进行测试之前，请验证所有权、权限、文件格式和错误处理。
