## 用法

来源：

- [扩展 control 文件](https://github.com/baverman/pghll/blob/692b1daa98c3a6b721acf850a90f6029e42e84ca/pghll.control)
- [1.0 版 SQL API](https://github.com/baverman/pghll/blob/692b1daa98c3a6b721acf850a90f6029e42e84ca/pghll--1.0.sql)
- [HLL 解码与聚合实现](https://github.com/baverman/pghll/blob/692b1daa98c3a6b721acf850a90f6029e42e84ca/pghll.c)

`pghll` 1.0 从 `bytea` 读取一种特定的序列化 HyperLogLog 计数器格式。它可以解压计数器、估算单个计数器的基数、合并已解码计数器，或者合并压缩计数器并返回组合估算值。

### 聚合兼容的计数器

假设 `hll_samples.counter` 存放由该扩展所需格式生成的计数器：

```sql
CREATE EXTENSION pghll;

SELECT hll_count(hll_merge(hll_decode(counter)))
FROM hll_samples;

SELECT hll_sum(counter)
FROM hll_samples;
```

`hll_decode` 展开按网络字节序编码的 zlib 压缩载荷。`hll_merge` 聚合接受已解码计数器并返回已解码计数器；`hll_count` 估算其基数。`hll_sum` 聚合直接组合压缩计数器并返回估算值。

### 输入可信度与兼容性

该扩展不提供构造器、哈希函数、格式标识符或载荷校验器。其 C 代码信任解码后 `bytea` 内的头部值，并据此选择寄存器数量和内存偏移。只能传入来自已知兼容生产者的计数器；任意或截断的输入会报错，还可能进入不安全的本地代码路径。

一次聚合中的所有计数器必须使用相同精度和序列化布局。其他 PostgreSQL HLL 扩展的数据类型即使同样使用 HLL 名称，也不代表格式兼容。此仓库没有 README、release 产物、升级路径或当前 PostgreSQL 兼容矩阵，目录状态也标为 abandoned；应隔离不受信任输入，并在读取历史数据前针对实际服务器构建完成验证。
