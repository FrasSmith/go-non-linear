# go-non-linear

A long time ago, I used to play with a BASIC program on my Acorn Archimedes that plotted non-linear graphs on the screen. It was very slow. I was reminded of this recently when this [thread](https://twitter.com/PaulBecherer/status/1336772062149533698) about it was posted on Twitter.

I decided it was time to revisit the code in a modern language.

This is my golang version of the code. I'm happy to say that it's significantly faster. On my system it can generate a 4K image with 8 million points in under 3 seconds. An 8k image with 32 million points takes under 15 seconds.