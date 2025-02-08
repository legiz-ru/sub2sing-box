# sub2sing-box

A tool to convert subscription/node connections into sing-box configurations.

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

Save the above JSON content as `sub2sing-box.json` and execute:

```
sub2sing-box convert -c ./sub2sing-box.json
```

to generate the sing-box configuration without needing to set parameters every time.

## Templates

### Default Template

The default template is located in the `templates` directory and uses the `tun` configuration. This configuration is for reference only—please modify it according to your needs.

### Placeholders

The following placeholders can be used in the template:

- `<all-proxy-tags>`: Inserts all proxy tags
- `<all-country-tags>`: Inserts all country tags
- `<Country (Region) Two-Letter Code>`: Inserts all node tags for a specified country (region), e.g., `<tw>`

#### Placeholder Usage Example:

Assume there are nodes:

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

// Where "United States (US)", "Singapore (SG)", and "Taiwan (TW)" are strategy groups containing US, SG, and TW nodes respectively
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
docker run -p 8080:8080 nite07/sub2sing-box:latest
```

You can mount a directory to add custom templates.

## Server Mode API

### GET /convert

| query | Description                                                                                                                    |
| ----- | ------------------------------------------------------------------------------------------------------------------------------ |
| data  | Same as the configuration above, but needs to be encoded using [base64 URL safe encoding](<https://gchq.github.io/CyberChef/#recipe=To_Base64('A-Za-z0-9%2B/%3D')>) |
