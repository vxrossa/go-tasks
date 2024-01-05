### Go Cipher

To use a cipher:
1. In the `input.txt` file specify two arguments:
    - The first argument is a string that will be encoded or decoded
    - The second argument is a combination for the Cipher
2. The cipher uses these rules:
    - C is for a Caesar cipher, R is for ROT13 cipher and A is for Atbash
    - 1 is for encoding, 0 is for decoding
    - If the wrong letter or number is specified, the program will throw an error

#### Example:

*"Hello world"* and **C1-C0-R1-R0** will result in *"Hello world"* because it has been encoded and decoded twice.