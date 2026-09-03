## 用法

来源：

- [官方文档](https://github.com/USC-InfoLab/fft_approximate/blob/f330f53d42e0e4fecb4dd1ccd29e480fc5407d72/README.md)
- [扩展控制文件](https://github.com/USC-InfoLab/fft_approximate/blob/f330f53d42e0e4fecb4dd1ccd29e480fc5407d72/fft_approximate.control)
- [官方仓库](https://github.com/USC-InfoLab/fft_approximate)

`fft_approximate` 对时序数据执行基于 FFT 的近似区间聚合查询。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `fft_approximate`：

```sql
CREATE EXTENSION fft_approximate;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT approximate_avg(c, 0, 10, 10 ORDER BY k) FROM coeffs;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `re` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `approximate_avg` | AGGREGATE | 扩展提供的聚合函数。 |
| `complex` | TYPE | 扩展创建的用户数据类型。 |
| `im` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `complex_add` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `complex_avg` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `complex_avg_accum` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `complex_in` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 应将扩展升级视为数据库变更：先审查上游升级路径、权限、锁以及备份恢复行为。
