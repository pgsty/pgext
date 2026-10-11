## 用法

来源：

- [README.md](https://github.com/MisterRaindrop/cloudberry-datafusion/blob/cc65001b59e2dcc7ba1d76ad825b3c92729b8511/README.md)
- [datafusion_executor--1.0.sql](https://github.com/MisterRaindrop/cloudberry-datafusion/blob/cc65001b59e2dcc7ba1d76ad825b3c92729b8511/datafusion_executor--1.0.sql)
- [datafusion_executor.control](https://github.com/MisterRaindrop/cloudberry-datafusion/blob/cc65001b59e2dcc7ba1d76ad825b3c92729b8511/datafusion_executor.control)

`datafusion_executor` 将符合条件的 Cloudberry 执行计划片段交给 Rust DataFusion 引擎处理，通过原生 C 接口与数据库连接。它面向 Cloudberry 内核，尚未确认可直接用于标准 PostgreSQL。

### 核心用法

```ini
shared_preload_libraries = 'datafusion_executor'
```

```sql
CREATE EXTENSION datafusion_executor;
SELECT datafusion_version();
SET datafusion.mode = 'explain';
```

### 运行边界

将 `datafusion_executor` 加入 `shared_preload_libraries` 后重启协调节点和分段进程，确保各后端都安装执行器钩子。上游也记录了对新会话生效的数据库级会话预加载方式。SQL 对象需要管理员安装。`datafusion.mode` 默认为 `off`，`explain` 只诊断可接管的片段，`on` 才启用替代执行器；不支持的片段继续使用原执行器。`datafusion.worker_threads` 控制每个后端的 Rust 工作线程数，按 CPU 数自动配置可能导致多后端争用。内存纳入 Cloudberry 的限额，但溢写文件尚不受其普通工作文件限额约束。浮点聚合的末尾数位可能不同。用于制造错误和测试运行时的诊断函数已撤销 PUBLIC 执行权限。
