# Datadog Buildpack

## Mandatory Env Vars

These env vars must be provided. See [https://docs.datadoghq.com/getting_started/tagging/unified_service_tagging/#overview](https://docs.datadoghq.com/getting_started/tagging/unified_service_tagging/#overview)

Example:

```sh
DD_SERVICE: "example-service"
DD_ENV: "example-env"
DD_VERSION: "0.0.1"
```

## Extra Tags

You can set extra tags for logs, metrics, etc. by setting the `DD_EXTRA_TAGS` env var. The format is a space seperated list with `key:value` pairs.

They get merged with `DD_TAGS` that is prefilled with Cloud Foundry specific tags.

Example:

```sh
DD_EXTRA_TAGS: "team:example scope:example product_billing_number:9999"
```

## Logs

Log forwarding is disabled by default. To enable it set the env var `DD_LOGS_ENABLED` to `"true"`.

This configures the Datadog Agent to listen on a TCP socket on port `10514`. See [https://docs.datadoghq.com/agent/logs/?tab=tcpudp#custom-log-collection](https://docs.datadoghq.com/agent/logs/?tab=tcpudp#custom-log-collection)

STDOUT and STDERR of the main process is redirected to a named pipe that is read by the `log-forwarder` (a small Go application) that forwards the logs to the Datadog Agent TCP socket.

You should set the `DD_LOGS_CONFIG_SOURCE` env var. An empty value is permitted but not recommended.

To set a different `service` tag only for logs use the `DD_LOGS_CONFIG_SERVICE` env var. By default it uses the value from `DD_SERVICE`.

To set additional `tags` only for logs use the `DD_LOGS_CONFIG_TAGS` env var. The format is a space seperated list with `key:value` pairs.

## APM

APM is enabled by default. To disable APM set the env var `DD_APM_ENABLED` to `"false"`.

The `trace-agent` listens on `localhost:8126` and on a UNIX Domain Socket. The env var `DD_APM_RECEIVER_SOCKET` contains the path to the socket.

## Custom Metrics

`dogstatsd` listens on `localhost:8125` and a UNIX Domain Socket. The env var `DD_DOGSTATSD_SOCKET` contains the path to the socket.

To disable listening on the UDP socket set the env var `DD_DOGSTATSD_PORT` to `0`.

## Differences to the official Buildpack

`RUN_AGENT` unnecessary, does nothing.

`DD_ENABLE_CHECKS` unnecessary, does nothing.

`DD_SPARSE_APP_LOGS` unnecessary, does nothing. Sparse logs work out-of-the-box.

Datadog has deprecated the old TCP intake. Use the new HTTP intake.
Set these env vars:
```yaml
DD_LOGS_CONFIG_LOGS_DD_URL: "agent-http-intake.logs.datadoghq.eu:443"
DD_LOGS_CONFIG_USE_HTTP: "true"
DD_LOGS_CONFIG_USE_COMPRESSION: "true"
```

