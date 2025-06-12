# aigit

[English Documentation](README.md) | 中文文档

最强大的 Git 提交助手！

## 支持的大模型/AI

- [OpenAI](https://openai.com/)
- [DeepSeek](https://deepseek.com/)
- [Doubao (豆包)](https://www.volcengine.com/product/doubao) - 内置，您不需要自己携带 API Key
- [Gemini](https://gemini.google.com/)

## 快速开始

### 安装

#### 选项 1：Homebrew（推荐）

```shell
brew install https://raw.githubusercontent.com/zzxwill/aigit/master/Formula/aigit.rb
```

#### 选项 2：下载二进制文件

- 前往 [发布页面](https://github.com/zzxwill/aigit/releases) 下载适合您平台的二进制文件。

- 将二进制文件重命名为 `aigit` 并移动到 `/usr/local/bin/aigit`。

```shell
chmod +x aigit && sudo mv aigit /usr/local/bin/aigit
```

#### 选项 3：从源码构建

```shell
git clone https://github.com/zzxwill/aigit.git
cd aigit
go build -o aigit main.go
sudo mv aigit /usr/local/bin/aigit
```

### 生成提交信息

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

### 使用自己的 AI API Key 生成提交信息

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
