## 用法

来源：

- [postvec/README.md](https://github.com/univec-ai/postvec/blob/6ca448ca9b90cb8c45bef9cce8019a7053bf0bd9/postvec/README.md)
- [postvec/postvec.control](https://github.com/univec-ai/postvec/blob/6ca448ca9b90cb8c45bef9cce8019a7053bf0bd9/postvec/postvec.control)
- [postvec/Cargo.toml](https://github.com/univec-ai/postvec/blob/6ca448ca9b90cb8c45bef9cce8019a7053bf0bd9/postvec/Cargo.toml)
- [postvec/sql/managed/functions.sql](https://github.com/univec-ai/postvec/blob/6ca448ca9b90cb8c45bef9cce8019a7053bf0bd9/postvec/sql/managed/functions.sql)
- [postvec/sql/managed/lifecycle.sql](https://github.com/univec-ai/postvec/blob/6ca448ca9b90cb8c45bef9cce8019a7053bf0bd9/postvec/sql/managed/lifecycle.sql)
- [postvec/sql/postvec--0.2.0--0.3.0.sql](https://github.com/univec-ai/postvec/blob/6ca448ca9b90cb8c45bef9cce8019a7053bf0bd9/postvec/sql/postvec--0.2.0--0.3.0.sql)

`postvec` 0.3.0 异步维护嵌入向量列，并融合词法与语义搜索。所核验发布为 postvec-v0.3.0-1；软件包后缀与扩展版本分别计数。

### 启用与核心工作流

需要 PostgreSQL 16–18 和 `vector` 0.8 或更新版本。将 `postvec` 加入 `shared_preload_libraries`，配置目标数据库与推理后端后重启。发布的软件包包含内嵌与远程模式；源码构建若未启用内嵌特性，就必须使用远程推理。注册表前应确保所选模型可用。

```sql
CREATE EXTENSION postvec CASCADE;
SELECT postvec.refresh_models();
CREATE TABLE docs (id bigserial PRIMARY KEY, body text);
SELECT postvec.enable('public.docs', 'body',
  model => 'sentence-transformers-all-minilm-l6-v2', create_fts_index => true);
INSERT INTO docs(body) VALUES ('PostgreSQL extension development');
SELECT * FROM postvec.status();
SELECT d.body, s.rrf_score
FROM postvec.search('public.docs', 'body', 'database extensions') s
JOIN docs d ON d.id = s.pk_value::bigint ORDER BY s.rrf_score DESC;
```

### 对象与边界

`postvec.enable` 安装触发器、受管向量列与回填任务；处理完成后才会得到完整搜索结果。`postvec.adopt` 接管已有向量，`postvec.disable` 停止管理。`postvec.create_vector_index` 创建向量索引，自动建索引可能阻塞写入。`postvec.retry_dead` 重试符合条件的死信任务。

管理操作需要表所有者或超级用户权限。工作进程使用提升后的权限，应审查行安全限制。远程推理会将源文本发送到数据库主机之外。扩展使用 PostgreSQL 许可证，独立服务端和下载模型各有其许可条款。

### 0.3.0 升级

更新共享库后，执行 `ALTER EXTENSION postvec UPDATE TO '0.3.0'` 应用 SQL 迁移。迁移会加固触发器的 `search_path`，并规范化受日期、时间格式影响的主键。它可能重置回填检查点并重新排队任务。存在歧义的旧键作为证据保留在 `jobs_dead` 中，在身份明确前自动重试会拒绝处理。应为额外回填预留资源，并检查迁移结果，再确认全部向量已更新。
