# Vex

## Overview

Vex is a domain-specific language (DSL) for text processing that is designed to be very simple. It aims for optimal performance, scalability, and maintainability.

My motivation for this project stems from my strong admiration for [awk](https://www.gnu.org/software/gawk/manual/gawk.html) and [sed](https://www.gnu.org/software/sed/manual/sed.html), which are highly effective text processing tools for UNIX. I am working on this project solely for my thesis, and my tool is not intended to replace awk or sed; rather, it aims to provide an alternative solution for text processing on different operating systems.

This project is still under development, and I will soon provide clearer documentation and additional features. For now, you can try out Vex's capabilities.

## Usage

Vex follows a simple principle:

```shell
vex [command] [pattern]
```

### 1. Find and Replace Files

To perform a find and replace operation on a file, use:

```shell
vex replace "foo:bar=input.txt"
```

Where:

- `foo` is the text you want to find.
- `bar` is the text you want to replace it with.
- `input.txt` is the file where the operation will take place.
  
This command is equivalent to the format:

```shell
[find:replace=<textinput>]
```

By default, the `replace` command is **case sensitive**; however, you can make it case insensitive by passing an argument:

```shell
vex replace "foo:bar=input.txt" -i
```

- `-i` stands for case insensitive.

### 2. File Input/Output

To load a text file into memory, you can use the following command:

```shell
vex load ":=input.txt"
```

To save the current content to a file, use:

```shell
vex save ":=output.txt"
```

### 3. Filtering

To filter a word in your file, you can run the following command:

```shell
vex filter "Hello=world.js"
```

This command is equivalent to the format:

```shell
[word=<textinput>]
```

Example Output:

```shell
Pattern: "Hello"
File: world.js
Total Matches: 3

Match 1: Line 5: console.log("Hello")
Match 2: Line 6: console.log("Hello")
Match 3: Line 7: console.log("Hello")
```

> [!TIP]
> This feature works best for source code!

**More features coming soon! ☀️**

## Installation

You can install Vex by running a following command:

```shell
go install github.com/ezrantn/vex@latest
```

Check if Vex already installed in your system:

```shell
vex version
```

## License

This tool is open-source and available under the [MIT](https://github.com/ezrantn/vex/blob/main/LICENSE) License.

## Contributions

I kindly request that you do not contribute to this project at this time, as it is part of my thesis, and I want to avoid any potential issues of academic dishonesty. 

However, if you have any suggestions or improvements, please feel free to reach out to me at [ezrantn@proton.me](mailto:ezrantn@proton.me)