# go-sro

A custom server/backend implementation of the game Silkroad Online
written in Go.

## Quickstart

You can use tmuxp to start a tmux session that launches the emulator and opens a neovim session. You can run the game client by running the `start-game-client.sh` file. Don't forget to set you WINEPREFIX.

```bash
tmuxp load .
```

### Game Client Binaries

This server needs the Silkroad Online binaries to run. You probably want them anyway. They have to be stored in the `./game_client_binaries/` directory.
Getting these binaries is a big pain, especially since you need the specific version this emulator is built for.

## Architecture

![architecture diagram](http://plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/ferdoran/go-sro/main/_docs/architecture.puml)

The backend architecture consists of several components that may interact with each other.
There are the 3 front-/client-facing servers (_Gateway Server_, _Agent Server_ and _Download / Patch Server_)
and 3 backend servers (multiple _Game Servers_, _Shard Server_, _Chat Server_).

All of them handle different kind of aspects to the game:

- **Download/Patch Server**: Provide updates and patches to the clients.
- **Gateway Server**: Perform authentication and transfer to the specific realm or shard.
- **Agent Server**: Proxy server for the client through which all network traffic is sent.
  Takes care routing network traffic to the correct servers.
- **Game Server**: Handles core game logic (navigation, AI, combat, ...)
  and game objects (players, pets, NPCs, ...).
  Usually there are multiple game servers, each handling a different region of the overall map
- **Shard Server**: Handles all region-independent logic (guild, party, market, events, ...)
- **Chat Server**: Handles overall chat messages (except local/region chat).
  Could also be handled by **Shard Server**
