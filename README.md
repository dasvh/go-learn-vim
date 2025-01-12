# Go Learn Vim

## About the project

**Go Learn Vim** is an interactive terminal-based application designed to teach users the basics of Vim motions in a fun and engaging way.
Using the [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework, this project combines learning with game mechanics,
making it easier for beginners to get started with Vim's navigation system.

![Demo gif](docs/gifs/demo.gif)

### Built with

The application is written in Go and uses the following libraries:

* [Bubble Tea](https://github.com/charmbracelet/bubbletea): Core framework for building reactive terminal UIs.
* [Bubbles](https://github.com/charmbracelet/bubbles): Prebuilt components for UI interactions  
* [Lip Gloss](https://github.com/charmbracelet/lipgloss): Tools for styling and layout of terminal UIs

### Features

* **Adventure Mode**: Navigate through various levels and solve challenges while learning essential Vim motions
* **Visual Feedback**: Track your movements with visual markers that display your path on the screen
* **Dynamic Statistics**: Real-time updates on keystrokes, elapsed time, and progress
* **High Scores**: Compete for the best scores based on your efficiency in time and keystrokes
* **Modular Design**: Built with a clean and reusable component architecture for easy expansion
* **Interactive UI**: Intuitive navigation using Vim-like commands and smooth transitions between menus

## Installation

### Prerequisites

- Go 1.23.4

### Setup

*It is recommended to run the application in a dedicated terminal window (e.g. not in an IDE terminal) with a minimum size of 135x45.*

```sh
# clone the repository
git clone github.com/dasvh/go-learn-vim
cd go-learn-vim

# builds the application and saves the binary in the /tmp/bin directory
make build

# runs the application from the /tmp/bin directory
make run
```

## Development

### Project Structure

```
.
├── cmd
│   └── main.go               # application entry point
└── internal
    ├── app                   # core application logic
    │   ├── controllers       # game, level and screen business logic
    │   └── screens           # application screens
    │       ├── adventure     # adventure mode screen
    │       │   └── level     # level specific logic
    │       ├── info          # info screens
    │       ├── leaderboards  # high scores and stats
    │       ├── menus         # main menu and other menus
    │       └── selection     # player, level and game save selection
    ├── components            # reusable UI components
    ├── models                # data models for players, stats, and levels
    ├── storage               # application persistence
    ├── style                 # UI styling
    └── views                 # reusable UI views
```

### Available Make Commands

```sh
make help      # shows available commands
make build     # build the application
make run       # run the application
make test      # run tests
make audit     # run quality control checks
make tidy      # format code and tidy dependencies
```

## License

[MIT](https://github.com/dasvh/go-learn-vim/raw/main/LICENSE)

## Acknowledgments

This project was created as part of a learning exercise and serves primarily as a proof of concept rather than a production-ready application.
The primary focus was on exploring the user interface design and interaction patterns using the [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework.
As such, certain aspects of the application, such as game mechanics and backend implementation, are intentionally simplified.

Special thanks to the following resources and communities that supported the development of this project:

  * [charm_](https://charm.sh/) for the Bubble Tea framework and its related tools, which served as the backbone for building the terminal-based UI
  * [Charm-In-The-Wild](https://github.com/charm-and-friends/charm-in-the-wild) collection of community projects built with the [Charm](https://github.com/charmbracelet/) stack


## Disclaimer

This project emphasizes UI functionality and modularity. While it offers foundational mechanics for Vim-inspired learning, key features such as advanced game logic, robust error handling, and extended testing are not fully developed.
These areas can serve as potential enhancements for future iterations or real-world implementations.