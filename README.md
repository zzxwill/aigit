# aigit

[中文文档 (Chinese Documentation)](./README_CN.md) | English Documentation

The most powerful git commit assistant ever!

It's a command-line tool that streamlines the git commit process by automatically generating meaningful and standardized commit messages, including title and body.

`aigit commit` is as simple as `git commit`.

## Supported 🤖 AI Providers

- [OpenAI](https://openai.com/)
- [DeepSeek](https://deepseek.com/)
- [Doubao (豆包)](https://www.volcengine.com/product/doubao) - Built-in, you don't need to bring your own key
- [Gemini](https://gemini.google.com/)

## Getting Started

### Installation

#### Option 1: Homebrew (Recommended)

```shell
# Add the repository as a tap (use full URL)
brew tap zzxwill/aigit https://github.com/zzxwill/aigit.git

# Install stable version (from releases)
brew install aigit

# Install development version (from dev branch)
brew install --HEAD aigit

# Alternative: Install from local formula file
# curl -O https://raw.githubusercontent.com/zzxwill/aigit/master/Formula/aigit.rb
# brew install --formula aigit.rb
```

#### Option 2: Download Binary

- Go to the [releases page](https://github.com/zzxwill/aigit/releases) and download the binary for your platform.

- Rename the binary to `aigit` and move it to `/usr/local/bin/aigit`.

```shell
chmod +x aigit && sudo mv aigit /usr/local/bin/aigit
```

#### Option 3: Build from Source

```shell
git clone https://github.com/zzxwill/aigit.git
cd aigit
go build -o aigit main.go
sudo mv aigit /usr/local/bin/aigit
```

### Generate commit message

```shell
$ aigit commit

🤖 Generating commit message...

📝 Generated commit message:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
feat(llm): add support for volcengine-go-sdk

This commit adds support for the volcengine-go-sdk for integrating with Doubao LLM service.

The following changes were made:

- Provider type and APIKey field were added to the llm.Config struct.
- generateDoubaoCommitMessage function was updated to use the volcengine-go-sdk.
- The client is initialized with the apiKey and endpointId.
- A prompt is constructed and sent to the CreateChatCompletion API.
- The first choice's message is returned as the commit message.
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

🤔 What would you like to do?
1. Commit this message (default)
2. Regenerate message

Enter your choice (press Enter for default):

✅ Successfully committed changes!
```

### Generate commit message with your own AI API Key

```shell
$ aigit auth add gemini AIzaSyCb56bjWn02e2v4s_TxHMDnHbSJQSx_tu8
Successfully added API key for gemini

$ aigit auth add doubao 6e3e438c-a380-4ed5-b597-e01cb82bc4df ep-20250110202503-fdkgq
Successfully added API key for doubao

$ aigit auth ls
Configured providers:
  gemini *default
  doubao

$ aigit commit

🤖 Generating commit message...

📝 Generated commit message:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
feat(llm): add support for volcengine-go-sdk

This commit adds support for the volcengine-go-sdk for integrating with Doubao LLM service.

The following changes were made:

- Provider type and APIKey field were added to the llm.Config struct.
- generateDoubaoCommitMessage function was updated to use the volcengine-go-sdk.
- The client is initialized with the apiKey and endpointId.
- A prompt is constructed and sent to the CreateChatCompletion API.
- The first choice's message is returned as the commit message.
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

🤔 What would you like to do?
1. Commit this message (default)
2. Regenerate message

Enter your choice (press Enter for default): 2

🤖 Regenerating commit message...

📝 Generated commit message:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
feat(llm): add support for volcengine-go-sdk

This commit adds support for the volcengine-go-sdk for integrating with Doubao LLM service.

The following changes were made:

- Provider type and APIKey field were added to the llm.Config struct.
- generateDoubaoCommitMessage function was updated to use the volcengine-go-sdk.
- The client is initialized with the apiKey and endpointId.
- A prompt is constructed and sent to the CreateChatCompletion API.
- The first choice's message is returned as the commit message.
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

🤔 What would you like to do?
1. Commit this message (default)
2. Regenerate message

Enter your choice (press Enter for default): 1

✅ Successfully committed changes!

```
