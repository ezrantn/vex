# vex

Vex is a domain-specific language (DSL) for text processing that combines DSL features with a command-line interface (CLI). It is designed for simplicity, optimal performance, scalability, and maintainability.

Inspired by tools like [awk](https://www.gnu.org/software/gawk/manual/gawk.html) and [sed](https://www.gnu.org/software/sed/manual/sed.html), Vex offers an alternative solution for text processing across different operating systems.

The project is under development, and clearer documentation and additional features will be provided soon. For now, you can explore Vex's capabilities.

## Principle

The heart of Vex lies in its straightforward syntax:

```shell
vex [command] [pattern]
```

## Installation

You can install Vex by running a following command:

```shell
go install github.com/ezrantn/vex@latest
```

Check if Vex already installed in your system:

```shell
vex version
```

## Usage

### 1. Find and Replace

Vex makes finding and replacing text in files simple and efficient. Here’s how you can use it:

#### Single Find and Replace

To perform a single find-and-replace operation, use the following syntax:

```shell
vex replace "foo:bar=input.txt"
```

**Explanation:**

- `foo`: The text you want to find.
- `bar`: The text you want to replace it with.
- `input.txt`: The file where the operation will be performed.

**Command Format**

```shell
[find:replace=<textinput>]
```

#### Batch Find and Replace

Vex also supports multiple find-and-replace operations in a single command. Here’s the syntax:

```shell
vex replace "foo1,foo2,foo3:bar1,bar2,bar3=input.txt"
```

**Explanation:**

- `foo1, foo2, foo3`: The words you want to find.
- `bar1, bar2, bar3`: The corresponding words to replace them with.
- `input.txt`: The file where the batch operation will be performed.

**Operation:**

- foo1 → bar1
- foo2 → bar2
- foo3 → bar3

#### Case Sensitivity

By default, the replace command is **case-sensitive**. To make it **case-insensitive**, use the `-i` flag:

```shell
vex replace "foo:bar=input.txt" -i
```

**Key Points:**

- `-i` allows matching text regardless of case (e.g., foo, Foo, and FOO are treated as the same).
  
### 2. Input / Output

To load a text file into memory, you can use the following command:

```shell
vex load ":=input.txt"
```

To save the current content to a file, use:

```shell
vex save ":=output.txt"
```

### 3. Filtering

Vex makes it easy to filter specific words in your files. Use the following command to search for a word:

```shell
vex filter "Hello=world.js"
```

#### Command Format

The general format for the filter command is:

```shell
[word=<textinput>]
```

**Example Output** 

Here’s what you might see when running the command:


```shell
Pattern: "Hello"
File: world.js
Total Matches: 3

Match 1: Line 5: console.log("Hello")
Match 2: Line 6: console.log("Hello")
Match 3: Line 7: console.log("Hello")
```

This output shows the word you’re searching for, the file name, and a detailed list of all matches, including their line numbers and content.

### 4. Pattern Matching

Perform regex searches within text:

```shell
vex match "[0-9]=data.txt"
```

### 5. Basic Statistics

Count occurrences of a term in a text:

```shell
vex count "success=input.txt"
```

**More features yet to come!**

## License

This tool is open-source and available under the [MIT License](https://github.com/ezrantn/vex/blob/main/LICENSE).

## Contributions

Contributions are welcome! Please feel free to submit a pull request.