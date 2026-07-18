## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/CyberDem0n/bg_mon/blob/d215e94f88a9088bdde0ce54f74d7d8bb2deb8e0/README.md)
- [Background-worker and HTTP implementation](https://github.com/CyberDem0n/bg_mon/blob/d215e94f88a9088bdde0ce54f74d7d8bb2deb8e0/bg_mon.c)
- [MIT license](https://github.com/CyberDem0n/bg_mon/blob/d215e94f88a9088bdde0ce54f74d7d8bb2deb8e0/LICENSE)

`bg_mon` is a headless background-worker module: it has no SQL extension to create. Add its library to `shared_preload_libraries`, configure it in `postgresql.conf`, and restart PostgreSQL.

```conf
shared_preload_libraries = 'bg_mon'
bg_mon.listen_address = '127.0.0.1'
bg_mon.port = 8080
bg_mon.naptime = 1
bg_mon.history_buckets = 20
```

The worker combines backend activity, process, memory, system, network, disk, and partition statistics, then serves current JSON at `/`, historical buckets at `/<bucket>`, and a simple browser UI at `/ui`. It requires libevent; Brotli history compression is optional at build time.

The embedded HTTP server has no documented authentication or TLS. Keep the default loopback binding and place an authenticated proxy in front if remote access is required. Binding `bg_mon.listen_address` to `0.0.0.0` can expose database activity and host telemetry to the network. Validate collection overhead, endpoint permissions, and memory retained by history on a representative load.
