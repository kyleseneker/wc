# wc

`wc` is a Golang program that mimics the functionality of the UNIX `wc` (word count) command-line tool. It reads one or more input files and prints the number of lines, words, and bytes contained in each file, along with a total line for all files combined. If no files are specified, it reads from the standard input.

This project was created as a coding challenge from [codingchallenges.fyi](https://codingchallenges.fyi/challenges/challenge-wc).

## Features

* Count the number of lines, words, and bytes in a file.
* Handle multiple input files and provide a total count.
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
    make build
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

To count lines, words, and bytes in a multiple files:

```sh
./wc file1 file2 file3
```

To read from standard input:

```sh
./wc
```

### Options

* `-L`: Display length of longest line in bytes
  * Returns length in bytes by default
  * Returns length in characters if the `-m` option is provided

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

Count lines, words, and bytes in multiple files:

```sh
./wc file1.txt file2.txt
```

Read from standard input:

```sh
echo "Hello World" | ./wc
```

## Contributing

Contributions are welcome! If you find a bug or want to add a new feature, please open an issue or submit a pull request.

## TODOs

* [ ] Add unit tests
* [ ] Add support for `-libxo` flag (generate output via libxo)
