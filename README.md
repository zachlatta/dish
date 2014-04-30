# Dish ![Analytics](https://ga-beacon.appspot.com/UA-34529482-6/dish/readme?pixel)

Serve files in the current directory with just four characters.

## Usage

Start an HTTP server on port 4000 in the current directory.

    $ dish

Serve `/tmp` on port 8080.

    $ dish -directory /tmp -port 8080

Run with `-help` for all usage options.

## Installation

    $ go get github.com/zachlatta/dish/dish

Head over to the [releases page](https://github.com/zachlatta/dish/releases)
if you don't have Go installed.

## License

The MIT License (MIT)

Copyright (c) 2014 Zachary Latta

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
