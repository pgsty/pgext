## 用法

来源：

- [官方 README](https://github.com/lanterndata/lantern/blob/7df590ce19080df7dd8a8bbb0f9875fc324bb2ea/lantern_extras/README.md)
- [官方扩展控制文件](https://github.com/lanterndata/lantern/blob/7df590ce19080df7dd8a8bbb0f9875fc324bb2ea/lantern_extras/lantern_extras.control)
- [官方 Rust 软件包清单](https://github.com/lanterndata/lantern/blob/7df590ce19080df7dd8a8bbb0f9875fc324bb2ea/lantern_extras/Cargo.toml)

`lantern_extras` 提供服务器端辅助函数，用于加载向量基准数据、运行本地 ONNX 嵌入模型、调用远程嵌入或补全 API，以及调度整表 AI 作业。这些函数不依赖主 Lantern 扩展，但会在 PostgreSQL 服务边界内执行文件、网络、模型推理与后台任务。

### 核心流程

以超级用户创建扩展，然后查看已安装版本内置的运行时和模型：

```sql
CREATE EXTENSION lantern_extras;

SELECT get_available_runtimes();
SELECT get_available_models();
```

对受支持的本地 ONNX 模型使用 `ort`，或者明确选择远程运行时：

```sql
SELECT llm_embedding(
  model   => 'BAAI/bge-base-en',
  input   => 'PostgreSQL vector search',
  runtime => 'ort'
);

SELECT llm_embedding(
  model     => 'text-embedding-3-small',
  input     => 'PostgreSQL vector search',
  runtime   => 'openai',
  api_token => '<api-token>'
);
```

扩展还能解析服务器本地的 `.fvecs` 文件并流式加载基准数据集：

```sql
SELECT *
FROM parse_fvecs('/srv/vector/siftsmall_base.fvecs', 10000);
```

### 作业与守护进程

`add_embedding_job` 与 `add_completion_job` 根据源表填充目标列。使用 `get_embedding_job_status`、`get_embedding_jobs` 和 `get_completion_jobs` 监控作业；`cancel_embedding_job` 与 `resume_embedding_job` 控制嵌入任务。

由 SQL 管理的 Lantern 守护进程被上游明确标为实验功能。文档所述配置需要预加载项、`lantern_extras.enable_daemon`、为预加载变更重启服务器，并可选设置 `lantern_extras.daemon_databases`。应按已安装软件包核对模块文件名，因为当前 README 把预加载库写作 `lantern_extra.so`，而扩展名是 `lantern_extras`。

### 安全与容量边界

- 调用可能下载大型归档或模型，并在数据库后端执行 CPU 密集型推理。应限制并发、语句时间、磁盘与内存使用，并把实验与延迟敏感工作负载隔离。
- 本地路径指数据库服务器上的路径，由 PostgreSQL 操作系统账户读取。远程 URL 和 API 调用会从服务器发起网络流量；应相应限制 SQL 权限和出站连接。
- 远程运行时会把输入数据发送给所选提供商。使用前应进行数据分级，并审查服务方的数据保留、地域、费用、速率限制和故障行为。
- `lantern_extras.llm_token` 可提供 API 凭据。应把该设置视为机密，明确谁能检查会话或配置值，并使用窄权限凭据，不要把密钥写入持久 SQL 或日志。
- 模型名、运行时、预处理、输出维数和模型修订共同构成应用契约。应记录这些元数据，并在混合或索引结果前验证向量维数。
- 控制文件把扩展标记为非可信且只能由超级用户安装。只向确实需要这些能力的角色授予文件、网络和作业函数执行权限。
