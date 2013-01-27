~Intro~ 
These are scripts that can be used to calculate unlock codes for various different mobile phones. I'll work on adding a list of supported models but for now, I'm just getting this code up here. Contributions to this repository of any kind are more than welcome, I'm new to GoLang so I'd love to see what I could have done better.


~Features~
Blackberry - Can calculate codes for 231 MEPs and ~7000 different PRDs. See findMEP.go
Huawei - Works with older USB modems.
ZTE - Three different calculators, one for old models, one calculates codes for firmware B03, the other calculates codes for firmware B04.
Alcatel - The models this will calculate codes for are commented into Alcatel.go. I'll probably move them into the README in my next commit.
JSON output - I really like the look of JSON. Some people might find it annoying though and I'll probably either remove it or make it optional later.

~Usage~
Check out main.go, use with command line arguments.

Disclaimer: This code is intended to be used for eductational purposes and comes with no guarantees of anything. In the USA it is illegal to unlock phone purchased AFTER January 26th, 2013 without the carrier's permission.