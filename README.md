# Hello BranchKit (Go)

A minimal [BranchKit](https://branchkit.dev) plugin that types a greeting at the cursor.

This is a **starter template** — clone it, customize it, and use it as the foundation for your own plugin.

## What it does

- **Voice**: Say "hello branchkit" → types "Hello, BranchKit!" at your cursor
- **Voice**: Say "hello Alice" → types "Hello, Alice!"
- **Keybind**: Alt+Shift+H → types "Hello, BranchKit!"

## Setup

### 1. Clone

```bash
git clone https://github.com/branchkit/branchkit-plugin-helloworld-go.git my-plugin
cd my-plugin
```

### 2. Generate typed params (optional)

If you add or change fields in `action_types`, re-run the codegen tool:

```bash
go install github.com/branchkit/branchkit-gen@latest
branchkit-gen --plugin .
```

This updates `src/actions_gen.go` with typed param structs derived from your `plugin.json`.

### 3. Build

```bash
cd src && go build -o ../hello-plugin .
```

### 4. Install

Copy the plugin into BranchKit's plugin directory:

```bash
cp -r . ~/Library/Application\ Support/BranchKitDev/plugins/helloworld/
```

Restart BranchKit. Your plugin loads automatically.

## Files

| File | Purpose |
|---|---|
| `plugin.json` | Manifest — declares actions, keybinds, voice commands |
| `commands.json` | Voice command patterns that trigger actions |
| `src/main.go` | Handler logic — your plugin's behavior |
| `src/actions_gen.go` | Generated typed params (from `branchkit-gen`) |

## Customizing

1. Change `"id"` and `"action_prefix"` in `plugin.json` to your plugin's name
2. Add actions to `action_types` with the fields they accept
3. Run `branchkit-gen --plugin .` to regenerate `src/actions_gen.go`
4. Write handler logic in `src/main.go`
5. Add voice commands to `commands.json` and keybinds to `plugin.json`

## Learn more

- [Your First Plugin](https://branchkit.dev/guide/getting-started/your-first-plugin) — step-by-step tutorial
- [Actions](https://branchkit.dev/guide/concepts/actions) — action routing, typed params, validation
- [Plugin SDK (Go)](https://github.com/branchkit/plugin-sdk-go) — full SDK reference
- [branchkit-gen](https://github.com/branchkit/branchkit-gen) — codegen tool

## License

MIT
