# wc

`wc` is a Golang program that mimics the functionality of the UNIX [wc](https://linux.die.net/man/1/wc) command-line utility. It reads one or more input files and prints the number of lines, words, and bytes contained in each file, along with a total line for all files combined. If no files are specified, it reads from the standard input.

This project was originally created as part of a coding challenge from [codingchallenges.fyi](https://codingchallenges.fyi/challenges/challenge-wc).

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
$ ./wc examples/test1.txt
    7145   58164  342190 examples/test1.txt
```

Count only lines and words in a file:

```sh
$ ./wc -l -w examples/test1.txt
    7145   58164 examples/test1.txt
```

Count lines, words, and bytes in multiple files:

```sh
$ ./wc test1.txt test2.txt
    7145   58164  342190 examples/test1.txt
   22315  215838 1253970 examples/test2.txt
   29460  274002 1596160 total
```

Read from standard input:

```sh
$ echo "Hello World" | ./wc
       1       2      12
```

## Performance

[hyperfine](https://github.com/sharkdp/hyperfine) is used to perform benchmarks.

To run the pre-defined benchmark:

```sh
make benchmark
```

## Contributing

Contributions are welcome! If you find a bug or want to add a new feature, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## TODOs

* [ ] Add unit tests
* [ ] Add support for `-libxo` flag (generate output via libxo)
* [ ] Improve performance
