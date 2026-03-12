

## 用法

> [GitHub: amutu/zhparser](https://github.com/amutu/zhparser)

`zhparser` 是基于 SCWS（Simple Chinese Word Segmentation）分词库的 PostgreSQL 中文全文搜索扩展。

## 功能特性

- 为 PostgreSQL 全文搜索提供中文分词
- 基于 SCWS（简易中文分词）库
- 支持自定义词典（TXT 和 XDB 格式）
- 数据库级自定义词表（v2.1 起）
- 多个可调参数控制分词行为

## 快速开始

```sql
-- 创建扩展
CREATE EXTENSION zhparser;

-- 创建使用 zhparser 的文本搜索配置
CREATE TEXT SEARCH CONFIGURATION chinese (PARSER = zhparser);

-- 添加词类映射
ALTER TEXT SEARCH CONFIGURATION chinese ADD MAPPING FOR n,v,a,i,e,l WITH simple;

-- 测试中文分词
SELECT to_tsvector('chinese', '小明硕士毕业于中国科学院计算所，后在日本京都大学深造');

-- 创建表和中文全文搜索索引
CREATE TABLE articles (id serial PRIMARY KEY, title text, body text);

CREATE INDEX articles_body_idx ON articles
  USING gin (to_tsvector('chinese', body));

-- 中文全文搜索查询
SELECT * FROM articles
  WHERE to_tsvector('chinese', body) @@ to_tsquery('chinese', '中国');
```

## 配置参数

zhparser 提供多个 GUC 参数控制分词行为：

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `zhparser.punctuation_ignore` | `off` | 忽略所有标点符号 |
| `zhparser.seg_with_duality` | `off` | 对长词进行二元分词 |
| `zhparser.dict_in_memory` | `off` | 将整个词典加载到内存 |
| `zhparser.multi_short` | `off` | 短词复合分词 |
| `zhparser.multi_duality` | `off` | 二元复合分词 |
| `zhparser.multi_zmain` | `off` | 首次复合分词中的关键词 |
| `zhparser.multi_zall` | `off` | 使用所有复合分词 |

## 词类

zhparser 支持 SCWS 的以下词类：

| 代码 | 说明 |
|------|------|
| `a` | 形容词 |
| `b` | 区别词 |
| `c` | 连词 |
| `d` | 副词 |
| `e` | 叹词 |
| `f` | 方位词 |
| `g` | 词根 |
| `h` | 前缀 |
| `i` | 成语 |
| `j` | 简称 |
| `k` | 后缀 |
| `l` | 临时习语 |
| `m` | 数词 |
| `n` | 名词 |
| `o` | 拟声词 |
| `p` | 介词 |
| `q` | 量词 |
| `r` | 代词 |
| `s` | 处所词 |
| `t` | 时间词 |
| `u` | 助词 |
| `v` | 动词 |
| `w` | 标点符号 |
| `x` | 未知 |
| `y` | 语气词 |
| `z` | 状态词 |

## 自定义词典

### 基于文件的词典

将自定义词典文件放在共享目录（通常为 `$SHAREDIR/tsearch_data/`）：

- TXT 格式：每行一个词
- XDB 格式：编译后的 SCWS 词典格式

自定义词典优先于内置词典。

### 数据库级自定义词（v2.1+）

```sql
-- 通过 zhparser 内置表添加自定义词
INSERT INTO zhparser.zhprs_custom_word VALUES ('中国科学院计算所');

-- 重新加载自定义词典（同步后需重新连接才能生效）
SELECT sync_zhprs_custom_word();

-- 验证带自定义词的分词结果
SELECT to_tsvector('chinese', '小明硕士毕业于中国科学院计算所');
```

## Docker 快速启动

```bash
docker run --name pgzhparser -d \
  -e POSTGRES_PASSWORD=somepassword \
  zhparser/zhparser:bookworm-16
```
