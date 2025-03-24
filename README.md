# sub2sing-box

A tool for converting subscriptions/node connections into **sing-box** configuration.

## templates by me | [click me](https://github.com/legiz-ru/sb-rule-sets/blob/main/.github/sub2sing-box/README.md)

## Proxy panel with localhost sub2sing-box

  - [x-ui-pro by legiz](https://github.com/legiz-ru/x-ui-pro/blob/master/README.md#install-x-ui-pro)
  - [xui-reverse-proxy by cortez24rus](https://github.com/cortez24rus/xui-reverse-proxy)

## Console Commands

Use `sub2sing-box <command> -h` to view help information for each command.

## Configuration

Example:

```json
{
  "subscriptions": ["Subscription URL 1", "Subscription URL 2"],
  "proxies": ["Proxy 1", "Proxy 2"],
  "template": "Template path or URL",
  "delete": "Remaining traffic",
  "rename": { "Original text": "New text" },
  "group-type": "selector",
  "sort": "name",
  "sort-type": "asc",
  "output": "./config.json"
}
```

Save the above JSON content as `sub2sing-box.json`, then execute:

```
sub2sing-box convert -c ./sub2sing-box.json
```

This will generate a **sing-box** configuration without the need to repeatedly set parameters.

## Templates

### Default Template

The default template is located in the `templates` directory and uses the `tun` configuration. This configuration is for reference only—please modify it according to your actual needs.

### Placeholders

The following placeholders can be used in the template:

- `<all-proxy-tags>`: Inserts all node tags
- `<all-country-tags>`: Inserts all country tags
- `<Country/Region Code>`: Inserts all node tags for a specific country/region, e.g., `<tw>`

#### Example Usage of Placeholders:

Assuming we have nodes:

- US
- SG
- TW

```json
{
  "type": "selector",
  "tag": "Node Selection",
  "outbounds": ["<all-proxy-tags>", "direct"],
  "interrupt_exist_connections": true
}

// After conversion

{
  "type": "selector",
  "tag": "Node Selection",
  "outbounds": ["US", "SG", "TW", "direct"],
  "interrupt_exist_connections": true
}
```

```json
{
  "type": "selector",
  "tag": "Node Selection",
  "outbounds": ["<all-country-tags>", "direct"],
  "interrupt_exist_connections": true
}

// After conversion

{
  "type": "selector",
  "tag": "Node Selection",
  "outbounds": ["United States (US)", "Singapore (SG)", "Taiwan (TW)", "direct"],
  "interrupt_exist_connections": true
}

// Where "United States (US)", "Singapore (SG)", and "Taiwan (TW)" are policy groups, each containing US, SG, and TW nodes.
```

```json
{
  "type": "selector",
  "tag": "Bahamut",
  "outbounds": ["<tw>", "direct"],
  "interrupt_exist_connections": true
}

// After conversion

{
  "type": "selector",
  "tag": "Bahamut",
  "outbounds": ["Taiwan (TW)", "direct"],
  "interrupt_exist_connections": true
}
```

## Docker Usage

```
docker run -p 127.0.0.1:8080:8080 ghcr.io/legiz-ru/sub2sing-box:latest
```

You can mount directories to add custom templates.

## Server Mode API

### GET /convert

| Query | Description |
| ----- | ----------- |
| data  | Same as the configuration above, but must be encoded using [base64 URL safe encoding](<https://gchq.github.io/CyberChef/#recipe=To_Base64('A-Za-z0-9%2B/%3D')>) |
