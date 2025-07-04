<div align="center">
    <img src=".github/assets/logo.png" alt="logo" width="300" />
</div>

Starling is a fast, lightweight distributed state-sharing system, written in Go. It uses the RAFT consensus algorithm to reliably synchronize values and state across applications and services. Designed for simplicity, speed, and cloud-native environments, Starling makes it easy to coordinate distributed workloads and share configuration or ephemeral state with minimal overhead.

## Starling Endpoints

### RAFT Internal Related API (v2: private)

```sh
curl -X GET -H "X-Client: client" -H "X-Session-Key: session_key" "/api/v2/raft/stats"
curl -X GET -H "X-Client: client" -H "X-Session-Key: session_key" "/api/v2/raft/votes"
curl -X GET -H "X-Client: client" -H "X-Session-Key: session_key" "/api/v2/raft/leaders"
```

### Application Interface (v1: semi-private)

```sh
curl -X GET -H "X-Session-Key: session_key" "/api/v1/whoisleader"
curl -X GET -H "X-Session-Key: session_key" "/api/v1/whoami?id=pid"
```

### Helpers (v0: public)

```sh
curl -X GET "/api/v0/readme"
curl -X GET "/api/v0/healthz"
curl -X GET "/api/v0/readyz"
```
