

## Usage

> [bgw_replstatus: lightweight replication status checker for load balancers](https://github.com/mhagander/bgw_replstatus)

bgw_replstatus is a background worker that listens on a TCP port and immediately reports whether the node is a `MASTER` or `STANDBY`. Designed for load balancer health checks without requiring a full PostgreSQL connection.

### How It Works

Any TCP connection to the configured port receives an immediate text response (`MASTER` or `STANDBY`) and the connection is closed. No request is needed.

### Quick Test

```bash
nc localhost 5400
# Output: MASTER (or STANDBY)
```

### HAProxy Configuration Example

```
frontend test
    bind 127.0.0.1:5999
    default_backend pgcluster

backend pgcluster
    mode tcp
    option tcp-check
    tcp-check expect string MASTER
    server s1 127.0.0.1:5500 check port 5400
    server s2 127.0.0.1:5501 check port 5401 backup
    server s3 127.0.0.1:5502 check port 5402 backup
```

### Configuration

| Parameter | Default | Description |
|-----------|---------|-------------|
| `bgw_replstatus.port` | 5400 | TCP port to listen on |
| `bgw_replstatus.bind` | (wildcard) | IP address to bind to |

**Security note**: There is no source verification. Protect the port with a host firewall.
