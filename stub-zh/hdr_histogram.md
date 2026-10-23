## 用法

来源：

- [README.md](https://github.com/vovcke/pg_hdr_histogram/blob/87686cfe3a06d48addd052e50f8e41a5bdc4347b/README.md)
- [hdr_histogram--1.0.0.sql](https://github.com/vovcke/pg_hdr_histogram/blob/87686cfe3a06d48addd052e50f8e41a5bdc4347b/hdr_histogram--1.0.0.sql)
- [hdr_histogram.c](https://github.com/vovcke/pg_hdr_histogram/blob/87686cfe3a06d48addd052e50f8e41a5bdc4347b/hdr_histogram.c)
- [META.json](https://github.com/vovcke/pg_hdr_histogram/blob/87686cfe3a06d48addd052e50f8e41a5bdc4347b/META.json)
- [.github/workflows/ci.yml](https://github.com/vovcke/pg_hdr_histogram/blob/87686cfe3a06d48addd052e50f8e41a5bdc4347b/.github/workflows/ci.yml)
- [hdr_histogram.control](https://github.com/vovcke/pg_hdr_histogram/blob/87686cfe3a06d48addd052e50f8e41a5bdc4347b/hdr_histogram.control)
- [官方 1.0.0 发布归档](https://github.com/vovcke/pg_hdr_histogram/releases/download/v1.0.0/hdr_histogram-1.0.0.tar.gz)

`hdr_histogram` 存储压缩 HDR 直方图，以可配置精度计算分位数。`hdr_add` 聚合数值样本，`hdr_merge` 合并已存储的直方图。上游 CI 包含 PostgreSQL 14–18。扩展可信、可迁移，依赖 zlib；拥有数据库 CREATE 权限的角色可以安装。

Pigsty 软件包使用的官方 1.0.0 发布归档安装的 control/SQL 版本为 1.0，而 Git 标签内的 control 文件声明 1.0.0。应检查实际安装的扩展版本，不能直接从软件包版本推断。

### 核心用法

```sql
CREATE EXTENSION hdr_histogram;
SELECT hdr_percentile(hdr_add(1.0), 50);
SET hdr_histogram.merge_mode = 'lossless';
SELECT hdr_percentile(hdr_add(latency::double precision), 99)
FROM (VALUES (100), (200), (300), (500)) AS samples(latency);
CREATE TABLE latency_histograms (recorded_at timestamptz, hist hdr_histogram);
SELECT hdr_percentile(hdr_merge(hist), 99) FROM latency_histograms;
```

### 运行边界

样本会转换为 int64，因此应统一计量单位。默认最大值为 1e9、精度为三位有效数字；hdr_histogram.lowest、hdr_histogram.highest 和 hdr_histogram.sig_figs 控制新直方图参数。非有限值或超出范围的样本会报错。`hdr_percentile` 接受 0 到 100 的分位点；`hdr_min`、`hdr_max`、`hdr_mean`、`hdr_stddev`、`hdr_total_count`、`hdr_sig_figs`、`hdr_lowest` 和 `hdr_highest` 用于检查直方图。默认 merge_mode 为 lossless，不兼容的合并会报错；strict 要求参数相同，left 可能丢弃超出累加器范围的值，union 会扩展范围和精度。首次调用直方图函数会加载库；示例先调用函数，再设置会话 GUC。核心功能无需共享预加载。分位数是误差有界的估计，不是精确的顺序统计量。
