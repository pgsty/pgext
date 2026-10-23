## 用法

来源：

- [Official source](https://github.com/saaros/pgbloomfun/blob/89ff830fed3f37fadd1e30ae9b64fc50f0abf60e/README)
- [Official source](https://github.com/saaros/pgbloomfun/blob/89ff830fed3f37fadd1e30ae9b64fc50f0abf60e/Makefile)
- [Official source](https://github.com/saaros/pgbloomfun/blob/89ff830fed3f37fadd1e30ae9b64fc50f0abf60e/ext/pgbloomfun.control)
- [Official source](https://github.com/saaros/pgbloomfun/blob/89ff830fed3f37fadd1e30ae9b64fc50f0abf60e/ext/pgbloomfun.sql)

`pgbloomfun` 0.1 将可增长的布隆过滤器存储在字节数组中，提供函数而非索引访问方法。这是 2013 年的旧源码，尚未确认兼容当前 PostgreSQL 版本。

### 核心用法

```sql
CREATE EXTENSION pgbloomfun;
SELECT bloom_check(bloom_add(bloom_init(500, 3, 0.01), 'hello'), 'hello');
```

`bloom_init(initial_capacity, growth_factor, error_rate)` 创建过滤器。`bloom_add` 返回更新后的字节数组，维护持久化过滤器时必须保存返回值。`bloom_check` 测试可能存在的成员，`bloom_stats` 返回过滤器统计信息。

### 限制

肯定结果可能是假阳性，需要精确核对。过滤器不会自动跟踪表的变化。安装涉及 C 函数，需要超级用户；扩展可迁移模式，没有指定预加载。构建元数据将控制版本替换为 0.1，并在支持时选择 JSON 输出。将旧过滤器移到新服务器前，应验证恢复流程和扩展兼容性。
