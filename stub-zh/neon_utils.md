## 用法

来源：

- [Neon 官方文档](https://neon.com/docs/extensions/neon-utils)

`neon_utils` 是 Neon 托管扩展，用于报告启用自动伸缩的 Neon compute 当前分配的 CPU 容量。它只在 Neon 内有意义，不是可移植的社区软件包，也不是通用操作系统监控 API。

### 安装并读取容量

在 Neon 数据库中创建扩展，然后调用其唯一有文档的函数：

```sql
CREATE EXTENSION IF NOT EXISTS neon_utils;

SELECT num_cpus();
```

`num_cpus()` 返回当前分配的 CPU 核数。可将它与工作负载指标一起轮询，观察 compute 如何在配置的自动伸缩最小值与最大值之间变化。

### 解释限制

Neon 使用可以为小数的 Compute Units 衡量计算容量。`num_cpus()` 不报告小数大小：例如，Neon 文档说明 0.25 或 0.5 CU 的 compute 会返回 `1`。因此该值是运维信号，而不是准确的计费、内存或配额度量。

该函数只在启用 Autoscaling 的 compute 上正确工作；在固定大小 compute 上不会返回正确值。使用结果前应在 Neon 控制台确认 compute 配置。

### 监控说明

应以合理间隔采样，并记录时间戳、工作负载延迟和配置的自动伸缩边界；单个值不能解释伸缩事件的原因。不要把舍入后的结果作为硬性调度或准入控制输入。可用性、版本、权限和行为遵循 Neon 服务，因此 compute 或平台升级后应重新核对官方页面。公开官方页面记录了函数与限制，但没有公布扩展版本；需要版本证据时，应在目标 Neon 数据库中核对实际安装的 `extversion`。
