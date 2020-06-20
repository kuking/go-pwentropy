# go-pwentropy

Creates good passwords and verifies passwords' quality (estimating its entropy).

Estimates password entropy considering common character sequences, practical enough so it does not require an external
database of sequences. Useful if you need to generate entropy for a cipher key (using some sort of password derivation
function), but want to be sure the interactive user inputs a password with enough entropy for the desired key size. 
At least, to some confidence level (it is complicated, read below.).

**Long explanation and complicated answer**

Without knowing the password generation algorithm used, and taking into consideration the adversary (and his dictionary
or knowledge on how we create passwords.) the entropy calculation will not be precise but an estimation. 

i.e. the entropy for the word '**hello**' can be calculated in multiple ways:

- 4 unique symbols (h e l o) with length 5: **10 bits** of entropy
- 26 unique symbols (lower case characters) with length 5: **23.5 bits** of entropy
- 250k unique symbols (complete English dictionary), length 1: **17.9 bits** of entropy
- 25k unique symbols (common words in the English language), length 1: **14.6 bits** of entropy
- 2 symbols (*naïve* password generator picking either: 'hello' or 'goodbye'), length 1: *1 bit* of entropy

(Using formula: `entropy = log2(uniqueSymbols^passwordLength)`)

So, when we prompt the user to enter a password, and the user enters 'hello', what entropy shall we consider fairly?

Well, the word 'hello' can be found in the [ncsc.gov.uk](https://https://www.ncsc.gov.uk/)' 
[PwnedPasswordsTop100k.txt](https://www.ncsc.gov.uk/static-assets/documents/PwnedPasswordsTop100k.txt). So, shall we
calculate the entropy of that password to be of length 1, with 100k unique symbols? (**16.6 bits** of entropy), it is
reasonable to assume the adversary could be using that weak password database.

#### 'Fair Entropy' calculation

We define three base entropy calculations:
- **EntropyByUniqueSymbols**: counts the distinct symbols in the password, i.e. for hello it would be 4. This is a lower
  bound for the entropy, but it does not takes into consideration dictionary words, or known passwords.
- **EntropyByClasses**: estimates the entropy calculating the symbols size by detecting different character classes in 
  the password i.e. if it the password has a number (or many), it adds 10 more unique symbols; if there are lower and
  upper case characters it estimates 26 unique symbols for lower case characters class and 26 symbols for upper case
  characters class; finally anything else is a 'special character' with an estimated uniqueness of 40 (as in characters
  that can be input with a standard keyboard, it is an estimation.)
- **EntropyByUniqueExclCommonSeqs**: Similar to UniqueSymbols entropy calculation, but subtracts from the length of the
  any known common char sequence. i.e. if the password is `£dictionary1`, it will assume the password length is 2 (for
  the pound symbol and the digit one.). This is can generate a very low entropy so shall be used indirectly.

We recommend using the **FairEntropy** which is a combination of the previous three:
- 50% EntropyByUniqueSymbols
- 25% EntropyByClasses
- 25% EntropyByUniqueExclCommonSeqs

It provides a fair result balancing different approaches. Examples:

| Password                                          | by Unique Symbols | by Classes | by Unique excl.Common | Fair   | 
|---------------------------------------------------|-------------------|------------|-----------------------|--------|
| `hello`                                           | 10.00             | 23.50      | 10.00                 | 13.38  |
| `this is a dictionary password`                   | 110.41            | 175.29     | 41.88                 | 109.50 |
| `dictionary`                                      | 31.70             | 47.00      | 0.00                  | 27.60  |
| `dictionarydictionary`                            | 63.40             | 94.01      | 0.00                  | 55.20  |
| `dictionary dictionary`                           | 69.76             | 126.93     | 3.32                  | 67.44  |
| `dictionarydictionarydictionary`                  | 95.10             | 141.01     | 0.00                  | 82.80  |
| `dictionary dictionary dictionary`                | 106.30            | 193.42     | 6.64                  | 103.17 |
| `4HAGK-RMYKQ`                                     | 36.54             | 68.73      | 36.54                 | 44.59  |
| `WP96N-BTY7X-DSNFQ-VAAWH`                         | 95.91             | 143.70     | 95.91                 | 107.86 |
| `WP96N-BTY7X-DSNFQ-VAAWH-QSTH5-AE7E7-VED5E-7TMWD` | 206.44            | 293.65     | 206.44                | 228.24 |

* notice, the password `hello` was not excluded by the 'common char sequences' as it is a short word with low uniqueness.
While building the database we tried to strike a balance between providing a small enough static compilable code without 
requiring to download the whole database into disk with its implied dependency. Please notice that 13 bits of entropy 
is not fair nowadays, so it can be reasonably discarded by a fair minimum entropy, i.e of 100 bits.

## Usage

```go
import pwe "github.com/kuking/go-pwentropy"

pw := pwe.PwGen(pwe.FORMAT_EASY, pwe.STRENGTH_96) 
fmt.Println("password:",pw,"has entropy of", pwe.FairEntropy(pw))
```

for a full demo look at [the demo app](demo/demo.go)
