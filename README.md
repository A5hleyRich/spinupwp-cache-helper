# SpinupWP Cache Helper

## Installation

For x86 (Intel/AMD):

```shell
curl -O https://raw.githubusercontent.com/A5hleyRich/spinupwp-cache-helper/main/builds/cache-amd64
chmod +x cache-amd64
sudo mv cache-amd64 /usr/local/bin/cache
```

For Arm64 (Ampere):

```shell
curl -O https://raw.githubusercontent.com/A5hleyRich/spinupwp-cache-helper/main/builds/cache-arm64
chmod +x cache-arm64
sudo mv cache-arm64 /usr/local/bin/cache
```

## Usage

### Cache Purging

```shell
cache purge
```

### Cache Warming

```shell
cache warm
```

To purge the cache before warming:

```shell
cache warm --purge
```
