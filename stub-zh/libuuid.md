## 用法

来源：

- [项目 README](https://github.com/petere/pglibuuid/blob/e05be1a67b06163c71571b5a9f926c20d70d11d7/README.md)
- [扩展 control 文件](https://github.com/petere/pglibuuid/blob/e05be1a67b06163c71571b5a9f926c20d70d11d7/libuuid.control)
- [1.0 版 SQL API](https://github.com/petere/pglibuuid/blob/e05be1a67b06163c71571b5a9f926c20d70d11d7/libuuid--1.0.sql)
- [libuuid 包装实现](https://github.com/petere/pglibuuid/blob/e05be1a67b06163c71571b5a9f926c20d70d11d7/libuuid.c)

`libuuid` 1.0 是已归档的操作系统 libuuid 库 C 包装器。它增加 `uuid_generate()`、`uuid_generate_random()` 和 `uuid_generate_time()`，均返回 PostgreSQL 内置 `uuid` 类型。

### 生成数值

```sql
CREATE EXTENSION libuuid;

SELECT uuid_generate_random();
SELECT uuid_generate_time();
SELECT uuid_generate();
```

`uuid_generate_random()` 请求随机 UUID，`uuid_generate_time()` 请求基于时间的 UUID。通用 `uuid_generate()` 把策略选择交给 libuuid，因此要求特定 UUID 版本的应用应调用明确函数并验证结果。

### 优先使用当前 PostgreSQL 功能

基于时间的 UUID 会暴露时间及节点相关信息，并且比随机 UUID 更可预测；不要把它当作秘密。随机 UUID 的碰撞抗性也取决于主机库和熵源。生成函数不能替代唯一约束。

该仓库已经归档，在 1.0 后没有升级路径，并依赖 PostgreSQL 服务器内部接口与平台库。现代 PostgreSQL 在核心中提供 `gen_random_uuid()`，更新的大版本还有更多内置 UUID 功能。新模式应优先使用内置功能。退役此扩展前，应比较 UUID 版本要求和默认值，修改应用与列默认值，然后确认不存在剩余函数依赖。
