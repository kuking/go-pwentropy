# go-pwentropy
Estimates password entropy and excludes common passwords. Useful if you need to generate entropy for a key but you
want to let the user input a password but you want to be sure the password provided has enough entropy for the
desired key size. At least, to some extend (it is compliated.).

**Long explanation and complicated answer**

Without knowing the password generation algorithm used, and taking into consideration the adversary (and his dictionary
or knowledge on how we create passwords.) the entropy calculation will not be precise but an
estimation. i.e. the entropy for the word '**hello**' can be calculated in multiple ways:

- 4 unique symbols (h e l o), length 5: 10 bits of entropy
- 26 unique symbols (lower case characters), length 5: 23.5 bits of entropy
- 250k unique symbols (english dictionary), length 1: 17.9 bits of entropy
- 25k unique symbols (common words in the english dictionary), length 1: 14.6 bits of entropy
- 2 symbols (naïve password generator picking one of the following words: 'hello', 'goodbye'), length 1: 1 bit of entropy

So, when we prompt the user to enter a password, and the user enters 'hello', what entropy will we estimate?

(work in progress, come back soon)
