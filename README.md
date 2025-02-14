
# Log Tailscale network flow logs local

This is a logger that listens for incoming network flow log messages, and sends them decompressed to stdout.
Relates to [Headscale issue #1687](https://github.com/juanfont/headscale/issues/1687)

Headscale and Tailscale:
* Headscale: Currently headscale does not support Network flow logging. In order to enable this in headscale, an example branch is available: [headscale-network-flow-logs](https://github.com/jdewinne/headscale/tree/headscale-network-flow-logs)
* Tailscale: The tailscale daemon does not support a custom endpoint for Network flow logging. In order to enable this in tailscaled, an example branch is available; [headscale-network-flow-logs](https://github.com/jdewinne/tailscale/tree/headscale-network-flow-logs)

Example output:

```
Received Decompressed Body: [{"logtail":{"client_time":"2025-02-14T00:50:06.4918846Z","proc_id":2604043402,"proc_seq":82},"nodeId":"1","start":"2025-02-14T00:50:03.665123895Z","end":"2025-02-14T00:50:06.491818802Z","exitTraffic":[{"dst":"100.64.0.1:0","txPkts":1,"txBytes":27}],"physicalTraffic":[{"src":"100.64.0.1:0","dst":"192.168.1.2:41641","txPkts":1,"txBytes":64,"rxPkts":2,"rxBytes":208}]}
]
Received Decompressed Body: [{"logtail":{"client_time":"2025-02-14T00:50:16.491717132Z","proc_id":2604043402,"proc_seq":83},"nodeId":"1","start":"2025-02-14T00:50:13.66691621Z","end":"2025-02-14T00:50:16.491625064Z","physicalTraffic":[{"src":"100.64.0.1:0","dst":"192.168.1.2:41641","rxPkts":1,"rxBytes":32}]}
]
Received Decompressed Body: [{"logtail":{"client_time":"2025-02-14T00:50:31.491756877Z","proc_id":2604043402,"proc_seq":84},"nodeId":"1","start":"2025-02-14T00:50:29.43977283Z","end":"2025-02-14T00:50:31.491661763Z","virtualTraffic":[{"proto":6,"src":"100.64.0.100:61843","dst":"100.64.0.1:32856","txPkts":2,"txBytes":104},{"proto":6,"src":"100.64.0.100:61843","dst":"100.64.0.1:32846","txPkts":2,"txBytes":104}],"physicalTraffic":[{"src":"100.64.0.1:0","dst":"192.168.1.2:41641","txPkts":4,"txBytes":384,"rxPkts":4,"rxBytes":384}]}
```