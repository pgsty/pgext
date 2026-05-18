## 用法

来源：[repo README](https://github.com/EnterpriseDB/pldebugger), [v1.10 release](https://github.com/EnterpriseDB/pldebugger/releases/tag/v1.10), [extension control](https://github.com/EnterpriseDB/pldebugger/blob/master/pldbgapi.control)

`pldbgapi` 为 PL/pgSQL 函数提供 server-side API，用于交互式调试。它通常通过 **pgAdmin** 这样的 GUI client 使用。

```sql
CREATE EXTENSION pldbgapi;
```

### 使用 pgAdmin 调试

主要使用方式是通过 pgAdmin 图形界面：

- **Direct Debugging**：右键函数并选择 "Debug"，即可立即执行并逐步调试。
- **Global Breakpoints**：在函数上选择 "Set Global Breakpoint"，然后等待另一个会话（例如 Web 应用）调用该函数；调试器会拦截该调用并允许上下文内调试。

### 调试能力

通过 debug client 连接后，你可以：

- 在 PL/pgSQL 函数的指定行上 **Set breakpoints**
- 逐行 **Step through** 代码（step into、step over、step out）
- 在每一步 **Inspect variables** 及其当前值
- 查看嵌套函数调用的 **call stack**
- **Continue execution** 到下一个断点

### 架构

调试系统包含三个组件：

1. **Client GUI**（pgAdmin）：显示源码、变量和 stack
2. **Target Backend**：正在执行被调试 PL/pgSQL 代码的会话
3. **Debugging Proxy**：通过专用连接在 client 和 target 之间协调

### 支持语言

调试器适用于 PL/pgSQL functions 和 procedures。需要在每个需要调试的数据库中创建 `pldbgapi` 扩展。

### 注意事项

- package 名是 `pldebugger`，而 SQL 中创建的 extension 是 `pldbgapi`；catalog 跟踪 package version `1.10`，覆盖 PostgreSQL 14 到 18。
- v1.10 上游 release 是 PostgreSQL 兼容性更新，没有记录新的用户侧 SQL API 或调试流程。
- 上游 troubleshooting 说明必须配置 `shared_preload_libraries = '$libdir/plugin_debugger'` 并重启 PostgreSQL。缺失或错误的 preload 会阻止 global breakpoints，在某些平台上也会阻止 `pldbgapi` SQL 加载。
