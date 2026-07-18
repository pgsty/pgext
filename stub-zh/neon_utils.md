## 用法

来源：

- [Neon 官方 neon_utils 文档](https://neon.com/docs/extensions/neon-utils)

`neon_utils` 是 Neon 提供的自动伸缩观测扩展。其 `num_cpus()` 函数报告当前连接的 Neon 计算节点已分配的整数 CPU 核数，适合在对自动伸缩范围做负载测试时使用。

### 读取当前分配量

```sql
CREATE EXTENSION IF NOT EXISTS neon_utils;
SELECT num_cpus();
```

在 Neon 计算节点上施加负载时重复运行该查询，即可观察分配变化。

### 注意事项

- `neon_utils` 由服务商管理且只针对 Neon 提供文档；上游没有发布可移植源码仓库或自托管安装方式。
- `num_cpus()` 仅在启用自动伸缩的计算节点上结果正确。官方文档明确说明，它在固定规格计算节点上不会返回正确值。
- 小数 Compute Unit 分配会向上取整为整数：0.25 或 0.5 CU 都报告为 1。
- 文档没有要求预加载。可用性由 Neon 服务控制，并可能取决于计算节点的 PostgreSQL 版本和服务配置。
