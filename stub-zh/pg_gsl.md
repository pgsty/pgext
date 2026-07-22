## 用法

来源：

- [官方 README](https://github.com/ferndale-hall/pg_gsl/blob/18630157063cd7a71e5ea6b3a3a3e84bfe9f19f3/README.md)
- [扩展 SQL 定义](https://github.com/ferndale-hall/pg_gsl/blob/18630157063cd7a71e5ea6b3a3a3e84bfe9f19f3/sql/pg_gsl.sql)
- [PGXN 元数据](https://github.com/ferndale-hall/pg_gsl/blob/18630157063cd7a71e5ea6b3a3a3e84bfe9f19f3/META.json)

`pg_gsl` 0.0.2 将 GNU Scientific Library（GSL）的一小部分 FFT 能力暴露给 PostgreSQL。它适合在 SQL 中变换一维 `double precision[]`、查看或截断频谱，并重建信号。它只封装了两个 GSL 例程，并不是通用的 GSL 绑定。

### 核心流程

先把 GSL 共享库安装到 PostgreSQL 可加载的位置，再安装扩展文件并在目标数据库中创建扩展：

```sql
CREATE EXTENSION pg_gsl;
SELECT pg_gsl_x_version();
```

对信号做变换、生成归一化频谱、将变换结果末尾置零，然后执行逆变换：

```sql
WITH signal(x) AS (
  VALUES (ARRAY[1.0, 0.0, -1.0, 0.0]::double precision[])
), transformed AS (
  SELECT pg_gsl_fft_real_transform(x) AS fft FROM signal
)
SELECT pg_gsl_x_fftToSpectrum(fft) AS spectrum,
       pg_gsl_fft_halfcomplex_inverse(
         pg_gsl_x_fftTruncate(fft, 2)
       ) AS filtered_signal
FROM transformed;
```

正向变换返回 GSL 的 half-complex 数组表示。应把该表示直接交给逆变换函数；`pg_gsl_x_fftToSpectrum` 只用于查看频谱，不能作为逆变换输入。

### 函数索引

- `pg_gsl_fft_real_transform(double precision[])` 执行实数 FFT。
- `pg_gsl_fft_halfcomplex_inverse(double precision[])` 从 half-complex 数组重建信号。
- `pg_gsl_x_fftToSpectrum(double precision[])` 将变换结果转换成长度为输入一半的归一化频谱。
- `pg_gsl_x_fftTruncate(double precision[], integer)` 将数组最后 `n` 个元素置零，从而在本项目使用的表示中抑制较高频率分量。
- `pg_gsl_x_version()` 返回扩展版本。

### 依赖与注意事项

`pg_gsl` 动态链接 GSL。上游构建方式要求 PostgreSQL 的 `pkglibdir` 中可找到 `libgsl.so`；无需设置 `shared_preload_libraries`。数组形状与长度会直接传给底层 FFT 代码，因此应用应先校验输入，并用代表性信号验证结果后再用于生产。0.0.2 是 2016 年发布的早期、窄范围版本。
