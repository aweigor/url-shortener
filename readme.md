# Url Shortener

## Basic features

1. Configure short link.
2. Process short link.
3. Monitor short links.

## Configuration

```
forward_ref: destination url
chain_links: optional array of additional links to visit
```

## Classes and types

### Short Link

```
url: short url
forward_ref: redirection url
enter_ref: (optional) url to visit after redirection
generator_rule: url generator rule
```

### Generator rules

1. SLUG8 - (default) Human-readable URL 8 characters long
