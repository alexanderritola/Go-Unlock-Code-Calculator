# Unlock Code Calculator in GoLang #

## Intro 
These are scripts that can be used to calculate unlock codes for various different mobile phones. I'll work on adding a list of supported models but for now, I'm just getting this code up here. Contributions to this repository of any kind are more than welcome, I'm new to GoLang so I'd love to see what I could have done better.


## Features 
**Blackberry** - Can calculate codes for 231 MEPs and ~7000 different PRDs. See findMEP.go

**Huawei** - Works with older USB modems.

**ZTE** - Three different calculators, one for old models, one calculates codes for firmware B03, the other calculates codes for firmware B04.

**Alcatel** - The models this will calculate codes for are commented into Alcatel.go. I'll probably move them into the README in my next commit.

**JSON output** - I really like the look of JSON. Some people might find it annoying though and I'll probably either remove it or make it optional later.

## Usage 
Check out main.go, use with command line arguments.

##Disclaimer
This code is intended to be used for eductational purposes and comes with no guarantees of anything. In the USA it is illegal to unlock phone purchased AFTER January 26th, 2013 without the carrier's permission.

##Legal
The MIT License (MIT)
Copyright (c) <year> <copyright holders>
 
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
