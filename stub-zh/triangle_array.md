## 用法

来源：

- [项目 README](https://github.com/evetion/trianglearray/blob/75b6e0ec507f6bdfd444efa008efdb328c933849/README.md)
- [扩展 control 文件](https://github.com/evetion/trianglearray/blob/75b6e0ec507f6bdfd444efa008efdb328c933849/triangle_array.control)
- [0.1 版 SQL API](https://github.com/evetion/trianglearray/blob/75b6e0ec507f6bdfd444efa008efdb328c933849/triangle_array--0.1.sql)
- [原生解码器实现](https://github.com/evetion/trianglearray/blob/75b6e0ec507f6bdfd444efa008efdb328c933849/triangle_array.c)

`triangle_array` 0.1 是一个已废弃的研究原型，用于把私有三角形数组字节布局解码为 WKB。其预期 SQL 函数为 `trianglexyz()`、`trianglez_bytea()` 和 `tinz_bytea()`，再由单独的 PostGIS 解释返回的几何字节。

### 当前源码不可部署

只能检查已有测试安装；不要假设仓库能够构建：

```sql
SELECT extversion
FROM pg_extension
WHERE extname = 'triangle_array';
```

在审查到的 head 中，C 源码把两个不兼容的 `unstitch()` 函数当作 C 支持重载来声明和定义，还存在声明结束符缺失以及 `Link` 与 `Linkq` 指针不匹配。这些是编译期缺陷。README 还说明源表名称硬编码在代码中。

### 非受信字节与内存安全边界

原生函数把调用者提供的 `bytea` 直接转换为点和三角形结构并建立索引，不验证长度、计数、偏移或顶点引用。错误甚至仅仅不匹配的数据就可能越界读取输入缓冲区，并使 PostgreSQL 后端崩溃。跨 patch 查找会针对 `triangle_array` 执行硬编码 SPI SQL，生成的 WKB 还使用主机字节序和历史几何类型选项。

不要向非受信输入开放这些函数，也不要在生产中加载未修复二进制。重新实现需要有文档的可移植 wire format、边界与溢出检查、明确字节序、fuzz 与 sanitizer、移除硬编码关系、当前 PostGIS 转换测试以及损坏安全错误。应优先使用受维护的 PostGIS TIN 或 geometry 存储，而不是保留该原型的私有二进制布局。
