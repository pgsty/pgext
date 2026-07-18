## 用法

来源：

- [上游 README](https://github.com/sweatybridge/pg_ffmpeg/blob/4a07afed5e04fb23937cf6ff8cd220c41667b2f9/README.md)
- [扩展 control 文件](https://github.com/sweatybridge/pg_ffmpeg/blob/4a07afed5e04fb23937cf6ff8cd220c41667b2f9/pg_ffmpeg.control)
- [Rust 包清单](https://github.com/sweatybridge/pg_ffmpeg/blob/4a07afed5e04fb23937cf6ff8cd220c41667b2f9/Cargo.toml)
- [过滤器安全实现](https://github.com/sweatybridge/pg_ffmpeg/blob/4a07afed5e04fb23937cf6ff8cd220c41667b2f9/src/filter_safety.rs)

`pg_ffmpeg` `0.3.4` 版在固定的 `ffmpeg` 模式中为 PostgreSQL 16–18 提供 FFmpeg 8 媒体检查与转换函数。输入输出通常是以 `bytea` 承载的完整媒体对象；扩展还提供转码、音频提取、剪裁、帧提取、HLS、GIF、波形、字幕、叠加、过滤、拼接和编码操作。

### 检查已存储的上传文件

```sql
CREATE EXTENSION pg_ffmpeg;

SELECT ffmpeg.media_info(body)
FROM media_uploads
WHERE id = 1;

SELECT ffmpeg.thumbnail(body, 1.0, 'jpeg')
FROM media_uploads
WHERE id = 1;
```

FFmpeg 会在 PostgreSQL 后端内解析攻击者可控的二进制格式，媒体任务也可能消耗大量 CPU、内存和临时存储。大型 `bytea` 值与 `concat_agg` 会放大后端内存使用。`hls` 接口可抓取 URL，并把播放列表和分片存入扩展表，因此不受限的执行权限会形成出站请求和存储滥用路径。上游过滤器白名单会阻止若干文件读取过滤器，但它并不是完整的信任边界。只向专用工作负载角色授予执行权，施加语句、资源和网络限制，尽可能在 PostgreSQL 外部验证媒体，并且绝不能向不可信 SQL 用户开放服务器端文件操作。
