# Typing Test Application
This project is a command-line typing test application written in Go. It presents users with random writing prompts and measures typing speed, accuracy, and other metrics.

## Features
- Random Writing Prompts: Utilizes the select_random_writing_prompt() function to present a random prompt from prompts.txt.
- Interactive Typing Interface: Incorporates a text area using the textarea model from the Bubbletea framework.
- Real-time Metrics Tracking: Records keystrokes and computes metrics such as accuracy and words per minute.
- Performance Analysis: Displays detailed statistics after completion, including time taken and mistakes corrected

## Getting Started
### Prerequisites
Go 1.23.3 or later installed on your machine.
### Installation
Clone the repository:
```zsh
git clone https://github.com/matwate/typing-test.git
```
Navigate to the project directory:
```zsh
cd typing-test
```
Install dependencies:
```zsh
go mod tidy
```

Execute the application with:
```zsh
go run main.go
```
The program will display a random writing prompt for you to type. As you type, it tracks your input and calculates performance metrics.




