
# Cards: A GoLang Library for Playing Card Decks

[![Go Reference](https://pkg.go.dev/badge/github.com/ehsun7b/cards.svg)](https://pkg.go.dev/github.com/ehsun7b/cards)
[![Go Report Card](https://goreportcard.com/badge/github.com/ehsun7b/cards)](https://goreportcard.com/report/github.com/ehsun7b/cards)

## Overview

**Cards** is a lightweight, easy-to-use GoLang library for creating and managing playing card decks. Whether you're building a classic game of **Poker**, a thrilling round of **Blackjack**, or a strategic game of **Bridge**, **Cards** provides all the tools you need to shuffle, deal, and manipulate playing cards with ease.

This library is designed with flexibility in mind, making it perfect for developing a wide range of card games, including but not limited to:

- **Poker** (Texas Hold'em, Omaha, Seven-Card Stud)
- **Blackjack**
- **Bridge**
- **Solitaire**
- **Hearts**
- **Spades**
- **Rummy**
- **Gin Rummy**
- **Crazy Eights**
- **Go Fish**
- **War**
- **Baccarat**
- **Euchre**

## Features

- **Deck Initialization**: Create standard 52-card decks or custom decks with jokers, wildcards, or other special cards.
- **Shuffling**: Shuffle your deck using Go's built-in random number generator for reproducible shuffling.
- **Dealing**: Deal cards to players or hands with customizable options.
- **Card Operations**: Perform operations such as drawing, discarding, or peeking at cards.
- **Support for Multiple Decks**: Easily manage multiple decks for games requiring more than one deck.
- **Extensible**: Easily extend the library to include custom cards or special rules.

## Installation

To install the library, use `go get`:

\```bash
go get github.com/ehsun7b/cards
\```

## Usage

### Basic Example

Here's a quick example of how to use **Cards** to create, shuffle, and deal a deck of cards:

\```go
package main

import (
    "fmt"
    "math/rand"
    "time"
    "github.com/ehsun7b/cards"
)

func main() {
    // Initialize a new deck with jokers
    deck := cards.NewDeck(true)

    // Create a random source with a seed for reproducible shuffling
    r := rand.New(rand.NewSource(time.Now().UnixNano()))

    // Shuffle the deck
    shuffledDeck := cards.Shuffle(deck, r)

    // Deal 5 cards
    hand, remainingDeck, err := cards.Deal(shuffledDeck, 5)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Display the hand
    fmt.Println("Hand:")
    for _, card := range hand {
        fmt.Println(card)
    }

    // Display the remaining deck
    fmt.Println("\nRemaining Deck:")
    for _, card := range remainingDeck {
        fmt.Println(card)
    }
}
\```

### Creating Custom Decks

You can create custom decks with jokers or special cards:

\```go
deck := cards.NewDeck(true) // Creates a deck with jokers
\```

### Advanced Shuffling Techniques

**Cards** allows you to implement and use shuffling techniques with a random source:

\```go
r := rand.New(rand.NewSource(time.Now().UnixNano()))
shuffledDeck := cards.Shuffle(deck, r)
\```

## Error Handling in Dealing

The `Deal` function returns an error if the hand size is invalid (too large or non-positive):

\```go
handSize := 60
hand, remainingDeck, err := cards.Deal(deck, handSize)
if err != nil {
    fmt.Println("Error:", err)
    // Handle error (e.g., retry with a valid hand size)
}
\```

## Documentation

For detailed documentation and examples, please visit the [GoDoc](https://pkg.go.dev/github.com/ehsun7b/cards) page.

## Contributing

We welcome contributions from the community! If you'd like to contribute, please fork the repository and submit a pull request. Make sure to follow the coding standards and include tests with your submissions.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Playing Card Wikipedia](https://en.wikipedia.org/wiki/Playing_card)
- [Shuffling Algorithms](https://en.wikipedia.org/wiki/Shuffling)

## Contact

Feel free to reach out with any questions or suggestions. You can contact the project maintainer at [your.email@example.com](mailto:your.email@example.com).
