# Discord Bot Template

A streamlined, feature-rich template for building Discord bots using the [**DiscordGo**](https://github.com/bwmarrin/discordgo) library in Go. This template provides a modular structure, making it easy to add commands, manage events, and maintain scalability as your bot grows.

## Features

- **Command System**: Easily register and manage commands.
- **Slash Command Support**: Utilize Discord's modern interaction system.
- **Utility Functions**: Predefined utilities for sending embeds, handling errors, and more.
- **Scalable Structure**: Designed for modularity and ease of expansion.

---

## Installation

1. **Clone the Repository**:
    ```bash
    git clone https://github.com/kqcl/discordgo-template.git
    cd discord-bot-template
    ```

2. **Install Dependencies**:
    Ensure you have [**Go**](https://go.dev) installed and run:
    ```bash
    go mod tidy
    ```
    Alternatively you can use [**make**](https://www.gnu.org/software/make/):
    ```bash
    make setup
    ```

3. **If you haven't, create a bot [here](https://discord.com/developer)** (optional)

4. **Configure Environment Variables**:
    Rename the `.env.example` file to `.env` and set the following environment variables:
    ```env
    DISCORD_TOKEN=your-bot-token
    APPLICATION_ID=your-application-id
    GUILD_ID=your-guild-id
    ```

5. **Run the Bot**:
    ```bash
    go run main.go
    ```
    Or with **make**:
    ```bash
    make
    ```

---

## File Structure

```plaintext
├── README.md
├── cmd
│   ├── command.go
│   ├── echo.go
│   └── ping.go
├── go.mod
├── go.sum
├── handlers
│   └── interactionCreate.go
├── helpers
│   ├── helpers.go
│   └── register.go
├── main.go
├── makefile
└── util
    ├── commands.go
    └── env.go
```


## Adding Commands

1. **Create a Command File:** Create a new .go file in the cmd/ directory.
2. **Define the Command:** Use the `RegisterCommand` function in `init()` to define your command:

```go
func init() {
    RegisterCommand(Command{
        Name:        "example",
        Description: "An example command",
        Options:     []discordgo.ApplicationCommandOption{},
        Callback: func(s *discordgo.Session, i *discordgo.InteractionCreate) error {
            return util.SimpleResponse("Hello, World!", s, i)
        },
    })
}
```

3. **Compile and Run:** The bot will automatically detect and register your new command.

## Contributing

1. Fork the repository
2. Create a new branch for your feature:
```bash
git checkout -b feature/new-feature
```
3. Commit changes and push:
```bash
git add .
git commit -m "Add new feature"
git push origin feature/new-feature
```
4. Create a pull request.

## License
This project is licensed under the MIT license. See the `LICENSE` file for details.

## Support
For issues or questions, please open an issue on the [Github repository](https://github.com/kqcl/discordgo-template).
