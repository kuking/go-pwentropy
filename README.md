# go-pwentropy
Estimates password entropy considering common character sequences, practical enough so it does not require an external
database of sequences. Useful if you need to generate entropy for a cipher key (using some sort of password derivation
function), but want to be sure the interactive user inputs a password with enough entropy for the desired key size. 
At least, to some hopefully confidence level (it is complicated, read below.).

#### Long explanation and complicated answer
Without knowing the password generation algorithm used, and taking into consideration the adversary (and his dictionary
or knowledge on how we create passwords.) the entropy calculation will not be precise but an estimation. 

i.e. the entropy for the word '**hello**' can be calculated in multiple ways:

- 4 unique symbols (h e l o), length 5: **10 bits** of entropy
- 26 unique symbols (lower case characters), length 5: **23.5 bits** of entropy
- 250k unique symbols (complete English dictionary), length 1: **17.9 bits** of entropy
- 25k unique symbols (common words in the English language), length 1: **14.6 bits** of entropy
- 2 symbols (*naïve* password generator picking either: 'hello' or 'goodbye'), length 1: *1 bit* of entropy

So, when we prompt the user to enter a password, and the user enters 'hello', what entropy shall we estimate?

Well, the word 'hello' can be found in the [ncsc.gov.uk](https://https://www.ncsc.gov.uk/)' 
[PwnedPasswordsTop100k.txt](https://www.ncsc.gov.uk/static-assets/documents/PwnedPasswordsTop100k.txt). So, shall we
calculate the entropy of that password to be of length 1, with 100k unique symbols? (16.6 bits of entropy), given that
the password database can be what the adversary is *actually* using to guess it. 

#### Approach

We define three different entropy calculations:
- EntropyByUniqueSymbols: counts the distinct symbols in the password, i.e. for hello it would be 4. This is a lower bound for
  the entropy, but it does not takes into consideration dictionary words, or known passwords.
- EntropyByClasses: calculates the entropy calculating the symbol size by detecting character classes in the password 
  (i.e. upper case, lower case, numbers, etc.)
- EntropyByUniqueExclCommonSeqs: Similar to UniqueSymbols entropy calculation, but substract from the length of the
  password the known common char sequences.

We recommend using the FairEntropy which is a combination of the previous three:
- 50% EntropyByUniqueSymbols
- 25% EntropyByClasses
- 25% EntropyByUniqueExclCommonSeqs

It provides a conservative result. Examples:

| Password                                          | by Unique Symbols | by Classes | by Unique excl.Common | Fair   | 
|---------------------------------------------------|-------------------|------------|-----------------------|--------|
| `this is a dictionary password`                   | 110.41            | 175.29     | 41.88                 | 109.50 |
| `dictionary`                                      | 31.70             | 47.00      | 0.00                  | 27.60  |
| `dictionarydictionary`                            | 63.40             | 94.01      | 0.00                  | 55.20  |
| `dictionary dictionary`                           | 69.76             | 126.93     | 3.32                  | 67.44  |
| `dictionarydictionarydictionary`                  | 95.10             | 141.01     | 0.00                  | 82.80  |
| `dictionary dictionary dictionary`                | 106.30            | 193.42     | 6.64                  | 103.17 |
| `4HAGK-RMYKQ`                                     | 36.54             | 68.73      | 36.54                 | 44.59  |
| `WP96N-BTY7X-DSNFQ-VAAWH`                         | 95.91             | 143.70     | 95.91                 | 107.86 |
| `WP96N-BTY7X-DSNFQ-VAAWH-QSTH5-AE7E7-VED5E-7TMWD` | 206.44            | 293.65     | 206.44                | 228.24 |
