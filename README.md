# AI-Code-Critic 🚀

![Golang](https://img.shields.io/badge/Go-00add8.svg?labelColor=171e21&style=for-the-badge&logo=go)

![Build](https://github.com/kmesiab/ai-code-critic/actions/workflows/go.yml/badge.svg)

![Logo](./assets/logo.png)

## Overview 🌟

AI-Code-Critic is a desktop application designed to automate code reviews
across various programming languages, utilizing OpenAI's ChatGPT API.
It offers intelligent insights and suggestions to improve code quality
and developer efficiency.

## Features 🛠️

- **Language-Agnostic Analysis**: Compatible with multiple programming
languages.
- **AI-Powered Insights**: Employs ChatGPT for in-depth code analysis.
- **User-Friendly Interface**: Simple and intuitive GUI for effortless
usage, built with Fyne.

## Installation 🔧

To install AI-Code-Critic, you need to have Go installed on your machine.
Follow these steps:

```bash
go install github.com/kmesiab/ai-code-critic
```

To run the program:

```bash
./ai-code-critic -f input.diff
```

## Usage 💡

Usage guidelines for effectively utilizing the application will be available
here.

## Development and Testing 🧪

### Building the Project 🏗️

```bash
make build
```

### Running Tests ✔️

```bash
make test
make test-verbose
make test-race
```

### Installing Tools 🛠️

```bash
make install-tools
```

### Linting 🧹

```bash
make lint
make lint-markdown
```

## Contributing 🤝

### Forking and Sending a Pull Request

1. **Fork the Repository**: Click the 'Fork' button at the top right of this
page.
2. **Clone Your Fork**:

   ```bash
   git clone https://github.com/kmesiab/ai-code-critic
   cd ai-code-critic
   ```

3. **Create a New Branch**:

   ```bash
   git checkout -b your-branch-name
   ```

4. **Make Your Changes**: Implement your changes or fix issues.
5. **Commit and Push**:

   ```bash
   git commit -m "Add your commit message"
   git push origin your-branch-name
   ```

6. **Create a Pull Request**: Go to your fork on GitHub and click the
'Compare & pull request' button.

Please ensure your code adheres to the project's standards and guidelines.

## License 📝

Information regarding the licensing of AI-Code-Critic will be included here.

---

*Note: This project is under active development. Additional features
and documentation will be updated in due course.* 🌈
