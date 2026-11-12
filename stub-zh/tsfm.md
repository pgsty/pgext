## 用法

来源：

- [官方 README](https://github.com/barddoo/tsfm-pg/blob/eec8098d7f3cbda55cba1a0d207d4daa4741ad02/README.md)
- [控制文件](https://github.com/barddoo/tsfm-pg/blob/eec8098d7f3cbda55cba1a0d207d4daa4741ad02/extension/tsfm.control)
- [安装 SQL](https://github.com/barddoo/tsfm-pg/blob/eec8098d7f3cbda55cba1a0d207d4daa4741ad02/extension/tsfm--0.1.sql)
- [配置实现](https://github.com/barddoo/tsfm-pg/blob/eec8098d7f3cbda55cba1a0d207d4daa4741ad02/src/guc.zig)

`tsfm` 在 PostgreSQL 16 或 17 内运行本地时间序列预测模型，SQL/控制文件版本为 0.1。Chronos-2 支持单变量输入；Timestral 还支持多变量和协变量。管理员需要在数据库主机安装动态库并准备模型文件。SQL 调用不会下载模型，也不会将序列发送至外部服务。

### 核心用法

准备受支持的检查点文件后，由超级用户安装扩展，并在预测前检查模型：

```sql
CREATE EXTENSION tsfm;
SELECT tsfm_version();
SELECT jsonb_pretty(tsfm_model_status());
SELECT tsfm_load();

SELECT step, quantile, value
FROM tsfm_forecast(ARRAY[10,11,12,13,14]::float8[], 3)
WHERE quantile = 0.5
ORDER BY step;
```

历史数据必须按时间从旧到新排列，步数和序列编号从 1 开始。默认模型目录位于 PostgreSQL 共享数据目录下。后端加载引擎后，`tsfm.model_dir` 和 `tsfm.model` 即被固定，修改时需要重新连接。不要求共享预加载。

### 权限与配置

安装脚本撤销了 `PUBLIC` 对全部函数的权限。只向应用授予实际需要的函数权限：

```sql
GRANT EXECUTE ON FUNCTION tsfm_forecast(double precision[], integer) TO analyst;
```

模型信息与状态的包装函数还会调用底层文本辅助函数，调用者也需要对应辅助函数的权限。`tsfm_forecast_ts` 同样要求底层预测函数的执行权限。

`tsfm.model_dir`、`tsfm.model`、`tsfm.threads`、`tsfm.enabled` 和 `tsfm.stub` 是超级用户参数。`tsfm.threads` 默认为 1，在加载时确定工作线程池大小；大量连接各自启动多个线程可能造成主机资源超额使用。`tsfm.stub` 选择测试预测器，并非训练模型。`tsfm.enabled = off` 会禁止预测、预热和分位数访问。

解码参数包括 `tsfm.max_context`、`tsfm.sort_quantiles` 以及 Timestral 专用控制项。配置关键字 `tsfm."symmetric"` 时必须使用引号：

```sql
SET tsfm."symmetric" = off;
```

### 函数分类

| 函数 | 用途 |
| --- | --- |
| `tsfm_model_info()` | 返回模型及引擎实际使用的配置 |
| `tsfm_model_status()` | 检查模型文件是否可用，不加载权重 |
| `tsfm_quantiles()` | 返回模型的分位数列表 |
| `tsfm_forecast_ex()` | 使用可选协变量进行多变量预测 |
| `tsfm_forecast_batch()` | 批量处理独立序列 |
| `tsfm_forecast_matrix()` | 返回步数与分位数组成的数组 |
| `tsfm_forecast_matrix_ex()` | 返回变量、步数与分位数组成的数组 |
| `tsfm_forecast_ts()` | 附加按固定间隔递增的未来时间戳 |

扩展历史输入按变量数 × 时间步数以行优先顺序展开；批量输入按序列数 × 变量数 × 时间步数展开。过去协变量覆盖输入历史，已知协变量还需覆盖预测区间。Chronos-2 仅接受一个变量，且不接受协变量。时间戳预测假定间隔固定，不检查日历缺口。

### 资源与模型边界

输入值会转换为 float32。缺失元素和 NaN 视为缺失数据，无穷大和超出范围的值也会转为缺失值。大整数可能损失精度。文档规定预测长度上限为 1024，单次结果最多包含 1600 万个 float32 值。

推理在后端进程内运行。取消检查发生在调用之间或批次块之间，无法中断单次前向计算。应先过滤数据并限制工作量。模型检查点使用内存映射，模型状态和临时计算内存可能在事务结束后继续保留。

GPU 支持是可选项，默认构建使用 CPU。扩展的 Apache-2.0 许可证不代表所有模型都使用同一许可：Timestral 权重有独立条款，项目不分发这些权重。只有托管服务明确支持该自定义库和模型文件系统访问时，才能在那里使用此扩展。
