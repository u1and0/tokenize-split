# tokenize-split

`tokenize-split` is a command-line tool that tokenizes input text and splits it into multiple files, with each file containing a maximum number of tokens specified by the user.

## Installation

1. Install Go: https://golang.org/doc/install
2. Run the following command:

```
$ go install github.com/u1and0/tokenize-split@latest
```

3. The `tokenize-split` command should now be available in your `$GOPATH/bin` directory.

## Usage

```
$ cat input.txt | tokenize-split -n 4096 -m gpt-3.5-turbo -v
```

This command reads input text from `stdin`, tokenizes it using OpenAI's GPT-3 tokenizer, and splits it into multiple files, each containing at most 4096 tokens. The resulting files are saved in the current directory with names like `tokenized_0.txt`, `tokenized_1.txt`, etc.

### Options

The following options are available:

- `-n <num_tokens>`: Maximum number of tokens per file. Defaults to 4096.
- `-m <model_name>`: ChatGPT Model name. Defaults to "gpt-3.5-turbo".
- `-v`: Show tokens size and file name.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
