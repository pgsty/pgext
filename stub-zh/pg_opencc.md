## 用法

来源：

- [官方 README](https://github.com/VitoVan/pg_opencc/blob/main/README.md)
- [扩展 SQL](https://github.com/VitoVan/pg_opencc/blob/main/pg_opencc--1.0.sql)
- [C 实现](https://github.com/VitoVan/pg_opencc/blob/main/pg_opencc.c)
- [OpenCC 项目](https://github.com/BYVoid/OpenCC)

`pg_opencc` 1.0 把 OpenCC 预设配置封装为 PostgreSQL 文本函数，用于简体中文、各地区繁体中文与日文新字体。数据库主机上的 OpenCC 词典与预设受控时，可用它执行词汇级文字规范化。

### 核心流程

在数据库主机安装匹配的 OpenCC 库和预设数据，创建扩展，然后调用目标转换：

```sql
CREATE EXTENSION pg_opencc;

SELECT opencc_s2t('汉字');
SELECT opencc_s2twp('鼠标和软件');
SELECT opencc_tw2sp('滑鼠和軟體');
```

转换可能识别词组，并不保证往返后恢复原文；应按产品术语与地区编辑规范检查输出。

### 函数索引

- 基础转换对：`opencc_s2t`、`opencc_t2s`。
- 台湾与香港变体：`opencc_s2tw`、`opencc_tw2s`、`opencc_s2hk`、`opencc_hk2s`、`opencc_t2tw`、`opencc_tw2t`、`opencc_t2hk`、`opencc_hk2t`。
- 词组转换对：`opencc_s2twp`、`opencc_tw2sp`。
- 日文变体：`opencc_t2jp`、`opencc_jp2t`。

每个函数都接受并返回文本，并在调用时打开相应的 OpenCC JSON 预设。

### 运维说明

SQL 把这些函数标记为不可变，但结果依赖主机安装的 OpenCC 库、配置文件与词典。因此，OpenCC 升级可能改变表达式索引或存储生成列的语义；受控更新词典后应重建派生值。C 实现每次调用都会打开转换器，却没有释放 `opencc_convert_utf8` 返回的缓冲区，从而每次调用都泄漏进程堆内存。在修复并验证该缺陷前，应避免高频使用。上游只报告测试过 OpenCC 1.1.1/1.1.2 与 PostgreSQL 13。
