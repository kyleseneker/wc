# wc

`wc` is a Golang program that mimics the functionality of the UNIX `wc` (word count) command-line tool. It reads an input file and prints the number of lines, words, and bytes contained in the file. If no file is specified, it reads from the standard input.

This project was created as a coding challenge from [codingchallenges.fyi](https://codingchallenges.fyi/challenges/challenge-wc).

## Features

* Count the number of lines, words, and bytes in a file.
* Read from standard input if no file is specified

## Installation

To install the `wc` tool, you need to have Golang installed on your system. Follow these steps:

1. Clone the repository:

    ```sh
    git clone https://github.com/kyleseneker/wc.git
    ```

1. Navigate to the project directory:

    ```sh
    cd wc
    ```

1. Build the program:

    ```sh
    go build -o wc
    ```

1. (Optional) Move the executable to a directory in your PATH:

    ```sh
    mv wc /usr/local/bin
    ```

## Usage

The `wc` command can be used with various options to specify which counts to display.

### Basic Usage

To count lines, words, and bytes in a single file:

```sh
./wc filename
```

To read from standard input:

```sh
./wc
```

### Options

* `-l`: Display the number of lines

* `-w`: Display the number of words

* `-c`: Display the number of bytes
  * This will cancel out any prior usage of the `-m` option

* `-m`: Display the number of characters
  * If the current locale does not support multibyte characters, this is equivalent to the `-c` option
  * This will cancel out any prior usage of the `-c` option

You can combine these options to display specific counts:

```sh
./wc -l -w filename
```

If no options are provided, `wc` displays lines, words, and bytes by default.

### Examples

Count lines, words, and bytes in a file:

```sh
./wc examples/test.txt
```

Count only lines and words in a file:

```sh
./wc -l -w examples/test.txt
```

Read from standard input:

```sh
echo "Hello World" | ./wc
```

## Contributing

Contributions are welcome! If you find a bug or want to add a new feature, please open an issue or submit a pull request.

## TODOs

* [ ] Add unit tests
* [ ] Add support for `-L` flag (length of the line containing the most bytes or characters)
* [ ] Add support for `-libxo` flag (generate output via libxo)