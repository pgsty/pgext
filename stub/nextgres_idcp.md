## Usage

Sources:

- [Alpha release documentation](https://github.com/nextgres/nextgres-idcp/blob/b90bbf19a158f937951a52edb6f888f8ffa0b30e/README.md)
- [Preload and worker registration code](https://github.com/nextgres/nextgres-idcp/blob/b90bbf19a158f937951a52edb6f888f8ffa0b30e/src/extension/entrypoint.c)
- [Example configuration](https://github.com/nextgres/nextgres-idcp/blob/b90bbf19a158f937951a52edb6f888f8ffa0b30e/nextgres_idcp.conf)

`nextgres_idcp` is an alpha in-database connection pooler. It registers background workers during shared preload and proxies client sessions to a smaller backend pool. The reviewed release targets PostgreSQL 16 on Ubuntu 22.04.

```conf
shared_preload_libraries = 'nextgres_idcp'
nextgres_idcp.thread_count = 1
nextgres_idcp.session_pool_size = 10
nextgres_idcp.port = 6543
```

Restart the cluster after enabling preload, then create the extension objects in the intended database. Upstream explicitly says this alpha is not for production: query cancellation, cleanup, and configuration remain incomplete, and it does not provide sharding or read/write load balancing. Bind the proxy only to controlled interfaces, apply PostgreSQL-grade authentication and TLS controls, and benchmark failure recovery before any non-lab use.
