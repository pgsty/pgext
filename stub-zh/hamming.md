## 用法

来源：

- [上游 README](https://github.com/AmineDiro/pghamming/blob/27b5119efe65f6d7f2173401ea8d20beb2d31817/README.md)
- [距离函数与操作符](https://github.com/AmineDiro/pghamming/blob/27b5119efe65f6d7f2173401ea8d20beb2d31817/src/hamming.rs)
- [索引构建实现](https://github.com/AmineDiro/pghamming/blob/27b5119efe65f6d7f2173401ea8d20beb2d31817/src/index/build.rs)

hamming 0.1.0 提供 hamming_distance(bytea, bytea) 与 <#> 操作符，用于字节级汉明距离比较。

### 顺序比较

```sql
CREATE EXTENSION hamming;
SELECT hamming_distance(decode('0f', 'hex'), decode('f0', 'hex'));
SELECT decode('0f', 'hex') <#> decode('f0', 'hex');
```

对于这两个单字节输入，两次调用都返回 8。稳定且有用的接口是通过普通扫描求值的函数与操作符。

### 注意事项

- 不要创建 ivf 索引。该访问方法是不可用的原型：构建回调并不索引元组，多项必需回调在被调用时会中止，而且缺少插入支持。
- 因此 README 中的索引示例无法提供近似最近邻搜索或持久索引存储。
- 控制文件将扩展标记为不受信任，并要求超级用户安装。安装与升级只能使用经过审查的构建版本。
- Rust 包使用 pgrx 0.10.2，且只声明支持 PostgreSQL 11 至 16。没有更新主版本的兼容证据，也没有已发布许可证。
