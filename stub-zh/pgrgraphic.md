## 用法

来源：

- [Official README](https://github.com/asotolongo/pgrgraphic/blob/ff2a6ddcf42d4d73be4ec3c694b3e330ffcb3c0c/README.md)
- [Extension SQL](https://github.com/asotolongo/pgrgraphic/blob/ff2a6ddcf42d4d73be4ec3c694b3e330ffcb3c0c/pgrgraphic--1.0.0.sql)
- [Official examples](https://github.com/asotolongo/pgrgraphic/tree/ff2a6ddcf42d4d73be4ec3c694b3e330ffcb3c0c/ejemplo)

pgrgraphic 是一个年代较久的 PL/R 扩展，把 PostgreSQL 查询结果渲染为数据库服务器上的 PNG 或 JPEG 图表文件。它涵盖柱状图、折线图、点图、直方图、等高线图、热力图、箱线图和饼图。它依赖 PostgreSQL 9.x、R、PL/R 及 X server 时代的环境，不能安全地供不可信用户使用。

### 核心流程

服务器管理员安装 R、PL/R、lattice、gplots 和合适的图形环境后，可信用户可以把 SQL 文本查询传给图表包装函数：

```sql
CREATE EXTENSION plr;
CREATE EXTENSION pgrgraphic;

SELECT pgrgraphic.grafica_barras_simples(
  'orders_by_status',
  'Orders by status',
  'Status',
  'Count',
  'SELECT count(*) FROM orders GROUP BY status ORDER BY status',
  ARRAY['new', 'done'],
  'vertical',
  'png'
);
```

函数会在 PostgreSQL 操作系统账户的主目录或工作目录下写文件并返回文本；上游以 10 作为成功标记。

### 重要对象

- `grafica_barras_simples` 和 `grafica_barras_multiples` 渲染柱状图。
- `grafica_lineas`、`grafica_lineas_puntos` 和 `grafica_puntos` 渲染折线及点图。
- `grafica_histograma` 和 `grafica_histograma_3d` 渲染直方图。
- `grafica_caja`、`grafica_contorno`、`grafica_dispersion`、`grafica_mapa_de_calor` 和 `grafica_pie` 覆盖其他图表形式。
- 对应的 `r_*` 函数是 SQL 包装背后的 PL/R 实现。

### 安全与运维说明

PL/R 是不受信任语言，而 pgrgraphic 会接收可执行 SQL 文本，并以数据库服务端账户写入任意图表名称。不要向不可信角色授予这些函数。安装 SQL 硬编码 `pgrgraphic` 模式和 `postgres` 所有者，可能安装失败或产生意外所有权。生成文件没有生命周期管理，可能占满服务器磁盘。应优先把查询数据导出到单独隔离的可视化服务。
