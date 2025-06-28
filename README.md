## slog_sample

No authorization required for health endpoint.
```bash
{"time":"2025-06-28T19:40:24.401936542+09:00","level":"INFO","msg":"request completed","http_info":{"method":"GET","path":"/api/v1/health","status":200,"latency":"65.755µs","user_agent":"curl/7.81.0","referer":"","remote_addr":"127.0.0.1:48788","request_id":"5cfa5a42-d282-4683-9ea0-31fe06b6c220"},"system_info":{"environment":"dev","service":"dev-slog-server","hostname":"ulala"},"auth_info":{"role":"anonymous","tenant_id":"default","member_id":"unknown"}}
```

When both authorization and logic succeed.
```bash
{"time":"2025-06-28T19:40:24.409490993+09:00","level":"INFO","msg":"request completed","http_info":{"method":"GET","path":"/api/v1/users/me","status":200,"latency":"40.765µs","user_agent":"curl/7.81.0","referer":"","remote_addr":"127.0.0.1:48798","request_id":"1692da63-f965-4f6e-9e9c-a94768b03326"},"system_info":{"environment":"dev","service":"dev-slog-server","hostname":"ulala"},"auth_info":{"role":"user","tenant_id":"tenant123","member_id":"member456"}}
```

When authorization fails.
```bash
{"time":"2025-06-28T19:40:24.415056794+09:00","level":"WARN","msg":"request failed","http_info":{"method":"GET","path":"/api/v1/users/profile/me","status":400,"latency":"37.811µs","user_agent":"curl/7.81.0","referer":"","remote_addr":"127.0.0.1:48808","request_id":"c3099314-dcea-4c40-a88c-7dd7eb9037d0"},"system_info":{"environment":"dev","service":"dev-slog-server","hostname":"ulala"},"auth_info":{"role":"user","tenant_id":"tenant123","member_id":"member456"}}
```

When authorization succeeds but unexpected error occurs in logic.
```bash
{"time":"2025-06-28T19:40:24.420764071+09:00","level":"ERROR","msg":"request failed","http_info":{"method":"GET","path":"/api/v1/products/10010000","status":500,"latency":"14.573µs","user_agent":"curl/7.81.0","referer":"","remote_addr":"127.0.0.1:48816","request_id":"edb29ba2-0b56-45d0-aa52-b5c91fe9ef2a"},"system_info":{"environment":"dev","service":"dev-slog-server","hostname":"ulala"},"auth_info":{"role":"anonymous","tenant_id":"default","member_id":"unknown"}}
```