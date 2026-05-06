# Meibel CLI

The official command-line interface for the [Meibel API](https://docs.meibel.ai).

## Installation

Install from Git (v2):

```bash
go install github.com/meibel-ai/meibel-cli@v2.0.0
```

Or clone and build:

```bash
git clone --branch v2.0.0 https://github.com/meibel-ai/meibel-cli.git
cd meibel-cli
go build -o meibel .
```

## Configuration

Set your API key:

```bash
meibel config set api-key your-api-key
```

Or use an environment variable:

```bash
export MEIBEL_API_KEY=your-api-key
```

## Usage

```bash
# List datasources
meibel datasources list

# Parse a document
meibel documents parse --file document.pdf

# Upload content to a datasource
meibel datasources content upload --file data.csv

# List data elements
meibel datasources data-elements list --datasource-id ds-123

# Chat with an agent
meibel agents sessions send-chat-message --session-id sess-123 --message "Hello"

# Stream a chat response
meibel agents sessions send-chat-message-stream --session-id sess-123 --message "Hello"
```

## Command Structure

Commands follow the nested resource hierarchy:

```
meibel
├── datasources
│   ├── list / create / get / update / delete
│   ├── content
│   │   ├── upload / list
│   │   └── trigger-ingest / get-ingest-status
│   ├── data-elements
│   │   ├── list / get / update / search
│   ├── downloads
│   │   ├── create / process / download
│   └── table-descriptions
│       ├── list-tables / list-columns
│       └── update-table-descriptions / update-column-descriptions
├── agents
│   ├── list / create / get / update / delete / publish
│   └── sessions
│       ├── create / list
│       └── send-chat-message / send-chat-message-stream
├── documents
│   ├── parse / process / get-result
├── sessions
│   ├── get / get-messages
└── ...
```

## Documentation

- [API Reference](https://docs.meibel.ai/api-reference/overview)
- [CLI Guide](https://docs.meibel.ai/sdk/cli)

## License

MIT
