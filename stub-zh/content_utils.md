## 用法

来源：

- [官方扩展 SQL](https://github.com/ttfkam/pg_content_utils/blob/71efbcd005630f6d2c96e221c8a6f9aa859b6676/content_utils--1.0.0.sql)
- [官方扩展控制文件](https://github.com/ttfkam/pg_content_utils/blob/71efbcd005630f6d2c96e221c8a6f9aa859b6676/content_utils.control)
- [官方 README](https://github.com/ttfkam/pg_content_utils/blob/71efbcd005630f6d2c96e221c8a6f9aa859b6676/README.md)

`content_utils` 提供异常单词计分、按时间衰减排名以及协议相对 URL 处理等 SQL 小工具。版本 `1.0.0` 还会通过 `file_fdw` 导入服务器的系统词典，因此安装依赖特定的服务端文件布局。

### 核心流程

先安装 `file_fdw`，并确认 PostgreSQL 可以读取 `/usr/share/dict/words`，然后再创建扩展：

```sql
CREATE EXTENSION file_fdw;
CREATE EXTENSION content_utils;

SELECT count_unusual('ordinary quuxword');
SELECT match_unusual('ordinary quuxword');
SELECT reify_url(true, '//example.com/path');
```

扩展会创建 `dictionary_word_files` 外部服务器、`words_fdt` 外部表和物化视图 `words`。词典文件改变后应显式刷新快照：

```sql
REFRESH MATERIALIZED VIEW words;
```

### 函数与注意事项

`count_unusual(text)` 统计词典中不存在的词，`match_unusual(text)` 判断是否存在此类词。`rank_modifier(interval)` 及其时间戳重载计算按周衰减的分数。`reify_url(boolean, varchar)` 为协议相对 URL 添加 HTTPS 或 HTTP 前缀。

控制文件没有把 `file_fdw` 声明为自动依赖；当硬编码词典文件不存在、不可读，或不符合预期单列格式时，安装会失败。物化视图不会自动更新。依赖 `rank_modifier(interval)` 前应测试零长度与未来时间间隔，因为其计算会除以取整后的已过周数。
