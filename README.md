# Stock Ticks Update Simulation using Raw TCP Protocol

## Building
Standing Go Build command can be used to build the program.

```bash
go build
```

## Run
Execute the generated binary. This will start the TCP Server listening on Port 9000.
Here is a sample output:
```
D:\Development\go-stocks-simulation-api>.\go-stocks-simulation-api.exe
2022/05/26 03:40:49 Server started on port 9000
```

Now go to another terminal window and execute netcat to start receiving data:

```bash
nc 172.19.112.1 9000
```

You will get each of the 10 stocks current value one by one and then after every 100 milliseconds you will receive a randomly updated stock.
Here is a sample output
```json
{
 "time": "2022-05-26T03:41:27+05:00",
 "symbol": "VRXV",
 "open": 100,
 "high": 136.7824922766078,
 "low": 62.068026488085884,
 "close": 62.068026488085884,
 "volume": 30833
}
```

Note that the `time` property in output is in RFC3339 Format.

WARNING: This program has not been stress tested.

# LICENSE - MIT
Copyright (c) 2022 Muhammad Tayyab Sheikh

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
