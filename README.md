# Lestrade: Colorful Text Printing for Go

![Lestrade Icon](palette.png)


Lestrade is a small Go library that allows you to print text in different colors. It's a fun and handy tool for adding some visual flair to your command-line applications. With Lestrade, you can make your text stand out and enhance the user experience.

## Table of Contents


- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgments](#acknowledgments)


## Usage:
```go

package main

import (
	"fmt"
	"github.com/BlyDoesCoding/lestrade"
)

func main() {
	fmt.Println(lestrade.Red("This is red text"))
	fmt.Println(lestrade.Green("This is green text"))
	fmt.Println(lestrade.Blue("This is blue text"))
}

```
## Contributing

Contributions to Lestrade are welcome! If you'd like to contribute to this project, please follow these guidelines:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and test them thoroughly.
4. Submit a pull request, describing your changes in detail.


## License

This project is licensed under the [BSD 2-Clause License](LICENSE). See the [LICENSE](LICENSE) file for details.


## Acknowledgments

- Special thanks to the Go community for their support and contributions.

Enjoy adding colorful text to your Go applications with Lestrade!

