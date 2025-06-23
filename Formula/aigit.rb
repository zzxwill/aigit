class Aigit < Formula
  desc "AI-powered Git commit message generator using LLM"
  homepage "https://github.com/zzxwill/aigit"
  url "https://github.com/zzxwill/aigit/archive/refs/tags/v0.0.8.tar.gz"
  sha256 "0a5163d430fb30bcdb6481b96aee8113ce070437d56863c5e9b4131c78d54e8b"
  license "Apache-2.0"
  head "https://github.com/zzxwill/aigit.git", branch: "master"

  depends_on "go" => :build

  def install
    ldflags = %W[
      -s -w
      -X main.Version=#{version}
    ]

    system "go", "build", *std_go_args(ldflags: ldflags), "./main.go"
  end

  test do
    assert_match version.to_s, shell_output("#{bin}/aigit version")
    assert_match "Generate git commit message", shell_output("#{bin}/aigit help")
  end

  def caveats
    <<~EOS
      Before using aigit, configure an AI provider:
        aigit auth add openai YOUR_API_KEY
        aigit auth add gemini YOUR_API_KEY
        aigit auth add deepseek YOUR_API_KEY
        aigit auth add doubao YOUR_API_KEY YOUR_ENDPOINT_ID

      Then use: aigit commit
    EOS
  end
end
