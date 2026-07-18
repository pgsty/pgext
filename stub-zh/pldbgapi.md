


## 用法

来源：[repo README](https://github.com/EnterpriseDB/pldebugger), [v1.10 release](https://github.com/EnterpriseDB/pldebugger/releases/tag/v1.10), [extension control](https://github.com/EnterpriseDB/pldebugger/blob/master/pldbgapi.control)

`pldbgapi` 为 PL/pgSQL 函数提供服务端 API，用于交互式调试。它通常通过 **pgAdmin** 这样的图形客户端使用。

```sql
CREATE EXTENSION pldbgapi;
```

### 使用 pgAdmin 调试

主要使用方式是通过 pgAdmin 图形界面：

- **直接调试**：右键函数并选择“Debug”，即可立即执行并逐步调试。
- **全局断点**：在函数上选择“Set Global Breakpoint”，然后等待另一个会话（例如 Web 应用）调用该函数；调试器会拦截该调用并允许在上下文中调试。

### 调试能力

通过调试客户端连接后，可以：

- 在 PL/pgSQL 函数的指定行上设置断点
- 逐行调试代码（步入、步过、步出）
- 在每一步检查变量及其当前值
- 查看嵌套函数调用的调用栈
- 继续执行到下一个断点

### 架构

调试系统包含三个组件：

1. **图形客户端**（pgAdmin）：显示源码、变量和调用栈
2. **目标后端**：正在执行被调试 PL/pgSQL 代码的会话
3. **调试代理**：通过专用连接协调客户端和目标后端

### 支持语言

调试器适用于 PL/pgSQL 函数和过程。需要在每个需要调试的数据库中创建 `pldbgapi` 扩展。

### 注意事项

- 软件包名是 `pldebugger`，而 SQL 中创建的扩展是 `pldbgapi`；目录跟踪的软件包版本为 `1.10`，覆盖 PostgreSQL 14 到 18。
- v1.10 上游版本是 PostgreSQL 兼容性更新，没有记录新的用户侧 SQL API 或调试流程。
- 上游故障排查说明必须配置 `shared_preload_libraries = '$libdir/plugin_debugger'` 并重启 PostgreSQL。缺失或错误的预加载会阻止全局断点，在某些平台上也会阻止 `pldbgapi` SQL 加载。
