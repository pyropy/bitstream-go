# Bitstream Go

Implementation of [bitstream protocol](https://robinlinus.com/bitstream.pdf) in Go.

# Example usage

You can find an example of how to use this library in the [example](example) directory.

# Deviations from the original protocol

- Original protocol uses CompactInt to encode integers, and this implementation uses LittleEndian.
- Original protocol signs encrypted root + payment hash using Schnorr signature, and this implementation signs encrypted root only.