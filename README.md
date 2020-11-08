# ByBit API Golang SDK
Fork of [frankrap's implementation](https://github.com/frankrap/bybit-api) which essential bugfixes.

## How to use
Refer to the original project README

## How to run tests
The original project hardcoded api credentials in sources.
For this to work, there is a file you need to manually create as it's in gitignore: `rest/.env`.

Put the following contents in it:
```.env
BASE_URL=https://api.bybit.com/
API_KEY=...
API_KEY=...
```
