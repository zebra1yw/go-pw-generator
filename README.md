# go-pw-generator
go-pw-generator is a customizable command-line tool written in Go. It’s designed to generate secure, random passwords directly from your terminal.

# Features
Customizable Length: You can specify the length of the generated password.
Character Sets: Choose from a variety of character sets for your password, including alphanumeric, symbols, and more.
Easy to Use: Simply run the command and get a secure password instantly.

# Installation
To install go-pw-generator, you can use the go get command:
```
go get github.com/zebra1yw/go-pw-generator
```

# Usage
To generate a password, run the following command:

To generate a 6-character password: 
```
pwGen -l 6
```

You can also specify the length of the password,  to generate a 12-character password with symbols, numbers, and mixed-case letters, simply type:

```
 pwGen -l 12 -u -d -s
```

# Contributing
Contributions are welcome! Please feel free to submit a Pull Request.

# License
go-pw-generator is licensed under the MIT License.


