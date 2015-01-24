# ltsv2json

Command to convert from LTSV to JSON

## Installation

    $ go get github.com/yuya-takeyama

## Usage

This example shows way to filter your LTSV file with [jq](http://stedolan.github.io/jq/)

```
$ tail -f access_log.ltsv | ltsv2json | jq 'select(.status == "200") | .'
```

## Contributing

1. Fork it ( https://github.com/yuya-takeyama/ltsv2json/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
