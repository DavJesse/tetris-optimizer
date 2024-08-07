# tetris-solverer

## Overview

This project is a Go program that takes a list of tetrominoes from a text file and arranges them to form the smallest possible square. Each tetromino is identified by uppercase Latin letters, starting from 'A' for the first one, 'B' for the second, and so on.

## Objectives

The main objectives of this project are:

- Develop a program that receives a single argument, a path to a text file containing a list of tetrominoes.
- Assemble all tetrominoes in the smallest square possible.
- Identify each tetromino in the solution by printing them with uppercase Latin letters.
- Ensure at least one tetromino is in the text file.
- Handle bad formats or incorrect file formats gracefully by printing `ERROR`.
- Ensure the program is written in Go and follows good coding practices.
- Include unit tests for the functions.

## Features

- Compiles successfully and executes with Go.
- Assembles tetrominoes into the smallest square possible.
- Prints tetrominoes using uppercase Latin letters.
- Handles invalid input formats by printing `ERROR`.
- Includes test files for unit testing.

## Example Input and Output

### Example Input (Text File)
```text
#...
#...
#...
#...

....
....
..##
..##
```

### Example Output

```text
AABB
AABB
....
....
```

### Example with Spaces Between Tetrominoes

```text
ABB.
ABB.
A...
A...
```

## Usage

### Example Usage

1. Create a text file (e.g., `sample.txt`) with the following content:

    ```text
    ...#
    ...#
    ...#
    ...#

    ....
    ....
    ..##
    ..##

    .###
    ...#
    ....

    ....
    ..##
    .##.
    ....

    ....
    .##.
    .##.
    ....

    ....
    ....
    ##..
    .##.

    ##..
    .#..
    .#..

    ....
    ###.
    .#..
    ....
    ```

2. Run the program using Go:

    ```sh
    $ go run . sample.txt | cat -e
    ```

### Expected Output

```text
ABBBB.$
ACCCEE$
AFFCEE$
A.FFGG$
HHHDDG$
.HDD.G$
```

## Installation

1. Ensure you have Go installed on your machine.
2. Clone the repository:
- On Github

    ```sh
    git clone https://github.com/DavJesse/tetris-optimizer.git
    cd tetromino-solver
    ```
- On Gitea

    ```sh
    git clone https://learn.zone01kisumu.ke/git/davodhiambo/tetris-optimizer.git
    cd tetromino-solver
    ```

3. Run the program:

    ```sh
    go run . sample.txt
    ```


## Testing

To run the unit tests, use the following command:

```sh
go test
```

This will execute all the tests in the solver_test.go file and report the results.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Author
David Jesse Odhiambo,

Apprentice Software Developer,

Zone01 Kisumu.