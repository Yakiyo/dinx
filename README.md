# dinx
[![CI](https://github.com/Yakiyo/dinx/actions/workflows/ci.yml/badge.svg)](https://github.com/Yakiyo/dinx/actions/workflows/ci.yml)

An index for the dart sdk versions. 

This is mainly made for use in [dsm](https://github.com/Yakiyo/dsm), a version manager for the dart sdk. It scrapes google's storage api somewhat to get all available versions.

## Endpoints

Every endpoint returns json content. The result is contained within the `body` field. That is, all results are 
```json
{
    "body": // the actual result
}
```
The type of the actual result depends on the endpoint used.

If there is any error, an `error` key is provided.
```json
{
    "error": "error message"
}
```

## `/versions/all`
Returns all version

Type: `Map<string, string[]>`
```json
{
    "body": [
        "1.0.0",
        "1.1.0",
        ...
    ]
}
```

## `/versions/map`
Returns all versions, but instead of all versions merged in an array, they are each separated by their channels
Type: `Map<string, Map<string, string[]>>`
```json
{
    "body": {
        "stable": [
            "1.0.0",
            "1.1.0",
            ...
        ],
        "beta": [
            "1.0.0-1.2.beta",
            "1.1.0-2.1.beta",
            ...
        ],
        "dev": [
            "1.0.0-1.0.dev",
            "1.1.0-2.6.dev",
            ...
        ],
    }
}
```

## `/versions/{channel}`
Returns all versions of a specific channel. This will return an error if channel is invalid (not one of stable, beta, dev)

Type: `Map<string, string[]>`
```json
{
    "body": [
        "1.0.0",
        "1.0.1",
        ...
    ]
}
```

## `versions/latest/{channel}`
Returns info about the latest version of a specific channel. If channel is not provided, it defaults to stable. If provided, it must be valid, otherwise returns an error. 
Type: `Map<string, Map<string, string>`
```json
{
    "body": {
        "date": "string",
        "revision": "string",
        "version": "string",
        "channel": "string",
    }
}
```

For any queries or bugs, please open a new [issue](https://github.com/Yakiyo/dinx/issues)

## Author

**dinx** © [Yakiyo](https://github.com/Yakiyo). Authored and maintained by Yakiyo.

Released under [MIT](https://opensource.org/licenses/MIT) License
