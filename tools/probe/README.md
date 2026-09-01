# probe

Checks whether this platform's regional endpoints are actually up, and measures how long recovery takes when they go down.

## Build

	cd tools/probe
	go build -o probe ./cmd/probe

## check

Single-shot by default — exits non-zero if anything's unhealthy, so it works as a CI gate. Pass `--interval` to poll continuously instead (runs until Ctrl-C, or `--duration` if set).

	probe check --url https://demo.tendo.pro --url https://demo-us-east-2.tendo.pro
	probe check --url https://app.tendo.pro --interval 10s

## recover

Watches one URL, starting from a required healthy baseline, and reports how long an outage lasted once one starts — the real number behind a resilience claim, not a guess. Start this in one terminal, then trigger `chaos run node-failure ...` in another.

	probe recover --url https://app.tendo.pro --max-wait 5m

Requires `--settle` (default 3) consecutive successful checks before declaring recovery, so one flaky response mid-outage doesn't get reported as the moment things came back.