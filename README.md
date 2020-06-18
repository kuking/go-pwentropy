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


