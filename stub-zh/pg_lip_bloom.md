## 用法

来源：

- [README.md](https://github.com/yunjiazhang/adaptiveness_vs_learning/blob/6df122faaf37e58cdd5f37954540fba9ef4f68dc/README.md)
- [pg_lip_bloom/pg_lip_bloom--1.0.sql](https://github.com/yunjiazhang/adaptiveness_vs_learning/blob/6df122faaf37e58cdd5f37954540fba9ef4f68dc/pg_lip_bloom/pg_lip_bloom--1.0.sql)
- [pg_lip_bloom/pg_lip_validation.sql](https://github.com/yunjiazhang/adaptiveness_vs_learning/blob/6df122faaf37e58cdd5f37954540fba9ef4f68dc/pg_lip_bloom/pg_lip_validation.sql)
- [LICENSE](https://github.com/yunjiazhang/adaptiveness_vs_learning/blob/6df122faaf37e58cdd5f37954540fba9ef4f68dc/LICENSE)
- [pg_lip_bloom/pg_lip_bloom.control](https://github.com/yunjiazhang/adaptiveness_vs_learning/blob/6df122faaf37e58cdd5f37954540fba9ef4f68dc/pg_lip_bloom/pg_lip_bloom.control)

`pg_lip_bloom` 1.0 为上游自适应查询处理实验中重写的连接查询提供布隆过滤函数。它不会自动优化任意 SQL，实际安装 SQL 也没有注册 control 描述中提到的 Bloom 索引访问方法。

### 核心用法

```sql
CREATE EXTENSION pg_lip_bloom;
SELECT pg_lip_bloom_init(1);
SELECT pg_lip_bloom_add(0, 42);
SELECT pg_lip_bloom_probe(0, 42);
SELECT pg_lip_bloom_info();
```

### 运行边界

`pg_lip_bloom_init(integer)` 分配过滤器槽位，`pg_lip_bloom_add(integer, integer)` 向指定过滤器加入值，`pg_lip_bloom_probe(integer, integer)` 检查成员关系。`pg_lip_bloom_info()` 报告过滤器状态，`pg_lip_bloom_set_dynamic(integer)` 选择实验性的动态模式。概率预过滤后仍需保留原始连接和谓词。上游环境是 PostgreSQL 12.5，实验依赖查询重写和会话状态。默认仅超级用户可安装；它是研究代码，不是透明替换的生产优化器。
