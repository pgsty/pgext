

## 用法

有时需要在将 `bytea` 对象返回给客户端之前对其进行压缩。

有时从客户端接收到压缩过的 `bytea`，需要先解压才能进行处理。

本扩展正是为此而设计的。

本扩展**不适用于**存储压缩场景。PostgreSQL 本身已经具备[元组压缩](https://www.postgresql.org/docs/current/storage-toast.html)机制，当元组足够大时会自动进行压缩，手动使用本函数预压缩数据并不会进一步减小存储空间。


* `gzip(uncompressed BYTEA, [compression_level INTEGER])` 返回 `BYTEA`
* `gzip(uncompressed TEXT, [compression_level INTEGER])` 返回 `BYTEA`
* `gunzip(compressed BYTEA)` 返回 `BYTEA`


### 示例

    > SELECT gzip('this is my this is my this is my this is my text');

                                       gzip
    --------------------------------------------------------------------------
     \x1f8b08000000000000132bc9c82c5600a2dc4a851282ccd48a12002e7a22ff30000000

等等，压缩后的输出怎么反而更长了？！其实并非如此，只是**看起来**更长罢了，因为十六进制表示中每个字节需要两个十六进制字符。原始字符串的十六进制表示如下：

    > SELECT 'this is my this is my this is my this is my text'::bytea;

                                                   bytea
    ----------------------------------------------------------------------------------------------------
     \x74686973206973206d792074686973206973206d792074686973206973206d792074686973206973206d792074657874

对于真正较长且重复度高的内容，压缩效果自然非常显著：

    > SELECT gzip(repeat('this is my ', 100));

                                                   bytea
    ----------------------------------------------------------------------------------------------------
     \x1f8b08000000000000132bc9c82c5600a2dc4a859251e628739439ca24970900d1341c5c4c040000

要将 `bytea` 转换回等价的 `text`，必须使用 `encode()` 函数并指定 `escape` 编码方式。

    > SELECT encode('test text'::bytea, 'escape');
       encode
    -----------
     test text

    > SELECT encode(gunzip(gzip('this text has been compressed and then decompressed')), 'escape')

                          encode
    -----------------------------------------------------
     this text has been compressed and then decompressed
