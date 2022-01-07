# goReloaded

## Intro
I was asked by [kood/Jõhvi coding school](https://kood.tech/) to share an approach to a [challenge](https://github.com/davhojt/goReloaded#challenge). I invested around 8 hours to introduce myself to [Go](https://go.dev) and complete the challenge.

The program processes the contents of a file, and writes the result to a new file.

## Usage
### Test:
```sh
$ go test
```


### Run the project:
```sh
$ go run . path_to_input path_to_output
```

## About

The string is split, and each substring becomes a token. The token stores data about the substring. For example, `Operation` tokens store a pointer to the function which applies the operation. Each token is a `struct`.

```go
type tokenKind int64

const (
	NoKind tokenKind = iota
	Word
	WhiteSpace
	Punctuation
	Quote
	Operation
)

type opData struct {
	ptr   *func(str string) string
	count int
}

type token struct {
	str  string
	kind tokenKind
	op   opData
}
```

Tokens are easier to work with when compared with an array of substrings, as observations of that substring can be done once and persisted. This reduces re-evaluation of that substring. Consider:

```go
for i := 0; i < token.op.count; i++
```

```go
for i := 0; i < getOperationCount(token); i++
```

Each word is broken down into tokens:
![](assets/img/tokenTypes.png)

Each required modification is carried out using the information in each token.

To apply operations, only `Word` and `Operation` tokens are relevant. Other types are ignored as operations are always applied to previous words irrespective of punctuation, whitespace etc.
![](assets/img/tokenOperation.png)

To use the correct article (`"A"` or `"An"`), punctuation is relevant. The `"a"` should not change to `"an"` here: `"I like the letter a, it is my favourite"`. So only `Word` and `Punctuation` tokens are relevant.
![](assets/img/tokenArticle.png)


## Process
1. `input` file is read into a string.
2. String is tokenized. Each rune is assessed. If the current rune belongs to a different type of token from the previous rune, then a new token is created. For operations (which combine characters from other token types), regex is used to assess if the current rune is the first character of a substring which represents an operation. An `Operation` token is created, and the length of that substring is skipped before normal rune processing continues.
3. `Operation` tokens are applied to relevant previous `Word` tokens.
4. Other processing takes place, correcting indefinite articles (A/An), placing `WhiteSpace` in the correct place relative to `Punctuation` and `Quote` tokens. This is done while creating a new final string.
5. String is written to `output` file.

# Challenge
This challenge was defined [here](https://git.01.kood.tech/root/public/src/branch/master/subjects/go-reloaded)

### Introduction
- Your project must be written in **Go**.
- The code must respect the [**good practices**](../good-practices/README.md).
- It is recommended to have **test files** for [unit testing](https://go.dev/doc/tutorial/add-a-test).

The tool you are about to build will receive as arguments the name of a file containing a text that needs some modifications (the input) and the name of the file the modified text should be placed in (the output). Next is a list of possible modifications that your program should execute:

- Every instance of `(hex)` should replace the word before with the decimal version of the word (in this case the word will always be a hexadecimal number). (Ex: "1E (hex) files were added" -> "30 files were added")
- Every instance of `(bin)` should replace the word before with the decimal version of the word (in this case the word will always be a binary number). (Ex: "It has been 10 (bin) years" -> "It has been 2 years")
- Every instance of `(up)` converts the word placed before in the Uppercase version of it. (Ex: "Ready, set, go (up) !" -> "Ready, set, GO !")
- Every instance of `(low)` converts the word placed before in the Lowercase version of it. (Ex: "I should stop SHOUTING (low)" -> "I should stop shouting")
- Every instance of `(cap)` transforms the previous word in the capitalized version of it. (Ex: "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge")

  - For `(low)`, `(up)`, `(cap)` if a number appears next to it, like so: `(low, <number>)` it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly. (Ex: "This is so exciting (up, 2)" -> "This is SO EXCITING")

- Every instance of the punctuations `.`, `,`, `!`, `?`, `:` and `;` should be close to the previous word and with space apart from the next one. (Ex: "I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!").
  - Except if there are groups of punctuation like: `...` or `!?`. In this case the program should format the text as in the following example: "I was thinking ... You were right" -> "I was thinking... You were right".
- The punctuation mark `'` will always be found with another instance of it and they should be placed to the right and left of the word in the middle of them, without any spaces. (Ex: "I am exactly how they describe me: ' awesome '" -> "I am exactly how they describe me: 'awesome'")
  - If there are more than one word between the two `' '` marks, the program should place the marks next to the corresponding words (Ex: "As Elton John said: ' I am the most well-known homosexual in the world '" -> "As Elton John said: 'I am the most well-known homosexual in the world'")
- Every instance of `a` should be turned into `an` if the next word begins with a vowel or an `h`. (Ex: "There it was. A amazing rock!" -> "There it was. An amazing rock!").

### Allowed packages

- Only the [standard Go](https://golang.org/pkg/) packages are allowed

### Usage

```console
$ cat sample.txt
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.

$ go run . sample.txt result.txt

$cat result.txt
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.

$ cat sample.txt
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.

$ go run . sample.txt result.txt

$cat result.txt
Simply add 66 and 2 and you will see the result is 68.

$ cat sample.txt
There is no greater agony than bearing a untold story inside you.

$ go run . sample.txt result.txt

$cat result.txt
There is no greater agony than bearing an untold story inside you.

$ cat sample.txt
Punctuation tests are ... kinda boring ,don't you think !?

$ go run . sample.txt result.txt

$cat result.txt
Punctuation tests are... kinda boring, don't you think!?
```

This project will help you learn about :

- The Go file system(**fs**) API
- String and numbers manipulation