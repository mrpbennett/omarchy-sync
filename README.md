# 🔄 omarchy-sync

A seamless configuration synchronization tool for Omarchy (Arch + Hyprland) systems. Keep your dotfiles in sync across multiple machines with an elegant CLI interface.

## 🎯 Project Vision

**omarchy-sync** solves the common problem of maintaining consistent configurations across multiple machines. Whether you're working on your desktop, laptop, or a fresh Omarchy installation, your carefully crafted configs stay in sync automatically.

### The Problem

Managing dotfiles across machines is tedious:

- Manual copying leads to configuration drift
- Traditional git workflows require constant manual commits
- Symlink managers need careful setup on each machine
- Conflicts arise when both machines are modified
- Sensitive data accidentally gets committed

### The Solution

omarchy-sync provides:

- **One-time setup** - Interactive wizard to configure everything
- **Automatic syncing** - Background daemon watches for changes
- **Smart conflict resolution** - Handles simultaneous edits gracefully
- **Secure by default** - Never commits secrets or sensitive data
- **Beautiful UX** - Charm-based TUI that's a joy to use

## 🏗️ Architecture

### Design Principles

1. **Simplicity First** - The tool should "just work" without deep git knowledge
2. **Non-intrusive** - Runs in the background without user intervention
3. **Safe** - Always backs up before overwriting, filters sensitive data
4. **Flexible** - Works with any directory structure, not just `.config`

### How It Works

```
┌─────────────┐                    ┌──────────────┐
│   Main PC   │                    │  Remote PC   │
│             │                    │              │
│  1. Watch   │──┐              ┌──│  1. Poll     │
│     changes │  │              │  │     repo     │
│             │  │              │  │              │
│  2. Commit  │  │   GitHub    │  │  2. Pull     │
│     & Push  │──┼─►  Repo   ──┼──│     changes  │
│             │  │              │  │              │
│  3. Auto-   │  │              │  │  3. Apply    │
│     sync    │──┘              └──│     with     │
│             │                    │     stow     │
└─────────────┘                    └──────────────┘
```

#### Main Machine Workflow

1. **File Watcher** monitors selected directories using `fsnotify`
2. **Change Detection** identifies modified, added, or deleted files
3. **Smart Filtering** excludes cache, binaries, and secrets
4. **Batch Commits** groups changes over a time window (e.g., 5 minutes)
5. **Auto Push** commits and pushes to GitHub repository

#### Remote Machine Workflow

1. **Periodic Polling** checks for updates (configurable interval)
2. **Pull Changes** fetches latest commits from GitHub
3. **Conflict Detection** identifies local modifications
4. **Stow Integration** creates symlinks to synced configs
5. **Rollback Support** keeps backups of overwritten files

## 🚀 Features

### Core Functionality

- ✅ Interactive setup wizard
- ✅ Directory selection with multi-select
- ✅ Automatic GitHub repository creation
- ✅ SSH and token-based authentication
- ✅ Background daemon with systemd integration
- ✅ Real-time file watching
- ✅ Intelligent change batching
- ✅ GNU Stow integration for symlink management

### Smart Features

- 🔒 **Secret Detection** - Automatically excludes API keys, tokens, passwords
- 🎯 **Pattern Matching** - Include/exclude specific file types
- 🔄 **Conflict Resolution** - Choose main-wins, remote-wins, or manual merge
- 📦 **Selective Sync** - Sync only what you need per machine
- 📊 **Sync Status** - Visual dashboard of sync state
- 🔙 **Rollback** - Undo syncs that broke something
- 🌐 **Offline Queue** - Commits are queued when offline, pushed when connected

### Developer Experience

- 💻 Beautiful TUI with Charm libraries (Bubble Tea, Lipgloss)
- 📝 Comprehensive logging
- 🧪 Dry-run mode for testing
- 🛠️ Easy debugging with verbose flags
- 📚 Well-documented configuration file

## 📋 Planned Commands

```bash
# Initial setup
omarchy-sync init                    # Interactive setup wizard

# Directory management
omarchy-sync add ~/.config/hypr      # Add directory to sync
omarchy-sync add ~/.config/waybar --pattern "*.ini,*.json"
omarchy-sync remove ~/.config/hypr   # Remove from sync
omarchy-sync list                    # Show all synced directories

# Manual sync operations
omarchy-sync push                    # Force push changes now
omarchy-sync pull                    # Force pull changes now
omarchy-sync status                  # Show sync status and conflicts

# Daemon management
omarchy-sync daemon start            # Start background sync
omarchy-sync daemon stop             # Stop daemon
omarchy-sync daemon status           # Check daemon status
omarchy-sync daemon logs             # View daemon logs

# Configuration
omarchy-sync config edit             # Edit config in $EDITOR
omarchy-sync config show             # Display current config
omarchy-sync config validate         # Check config validity

# Advanced
omarchy-sync rollback <commit>       # Restore previous state
omarchy-sync conflicts list          # Show current conflicts
omarchy-sync conflicts resolve       # Interactive conflict resolver
omarchy-sync doctor                  # Check system health
```

## 🔧 Configuration Structure

```json
{
  "machine": {
    "type": "main",
    "name": "desktop-omarchy",
    "hostname": "arch-desktop"
  },
  "repository": {
    "url": "git@github.com:username/omarchy-dots.git",
    "branch": "main",
    "auth_method": "ssh"
  },
  "sync": {
    "interval": 300,
    "batch_window": 300,
    "auto_push": true,
    "auto_pull": true
  },
  "directories": [
    {
      "path": "~/.config/hypr",
      "patterns": ["*"],
      "exclude": ["*.log", "cache/*"],
      "stow": true
    },
    {
      "path": "~/.config/waybar",
      "patterns": ["*.json", "*.css"],
      "stow": true
    }
  ],
  "filters": {
    "exclude_patterns": [
      "*.cache",
      "*.log",
      "**/cache/**",
      "**/Cache/**",
      "**/.git/**"
    ],
    "secret_patterns": [
      "**/token*",
      "**/*secret*",
      "**/*password*",
      "**/*.key",
      "**/*.pem"
    ]
  },
  "stow": {
    "target_dir": "~",
    "stow_dir": "~/.omarchy-sync/stow",
    "backup": true,
    "backup_dir": "~/.omarchy-sync/backups"
  }
}
```

## 🛠️ Technology Stack

- **Language**: Go 1.21+
- **TUI Framework**: [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- **Styling**: [Lip Gloss](https://github.com/charmbracelet/lipgloss)
- **Forms**: [Huh](https://github.com/charmbracelet/huh)
- **Git Operations**: [go-git](https://github.com/go-git/go-git)
- **File Watching**: [fsnotify](https://github.com/fsnotify/fsnotify)
- **CLI Framework**: [Cobra](https://github.com/spf13/cobra)

## 📦 Installation (Planned)

```bash
# From AUR (future)
yay -S omarchy-sync

# From source
git clone https://github.com/yourusername/omarchy-sync
cd omarchy-sync
go build -o omarchy-sync
sudo mv omarchy-sync /usr/local/bin/

# Via go install
go install github.com/yourusername/omarchy-sync@latest
```

## 🚦 Getting Started

```bash
# 1. Initialize on your main machine
omarchy-sync init

# Follow the interactive wizard:
# - Select directories to sync
# - Choose machine type (main)
# - Set up GitHub authentication
# - Configure sync preferences

# 2. Start the daemon
omarchy-sync daemon start

# 3. On your remote machine
omarchy-sync init

# Follow the wizard:
# - Enter the same repository URL
# - Choose machine type (remote)
# - Authenticate to GitHub
# - Pull existing configs

# 4. Start daemon on remote
omarchy-sync daemon start

# Done! Your configs now sync automatically
```

## 🎨 Design Philosophy

### User Experience

- **Progressive Disclosure**: Simple by default, powerful when needed
- **Sensible Defaults**: Works out of the box for 90% of use cases
- **Clear Feedback**: Always show what's happening and why
- **Fail Safely**: Never lose data, always backup first

### Code Philosophy

- **Modular Design**: Each component is independently testable
- **Error Handling**: Graceful degradation, never panic
- **Documentation**: Code should be self-documenting with clear comments
- **Performance**: Efficient file watching, minimal resource usage

## 🗺️ Roadmap

### Phase 1: Foundation (Current)

- [x] Basic TUI framework
- [ ] Configuration file structure
- [ ] Git repository initialization
- [ ] Directory selection wizard
- [ ] Machine type configuration

### Phase 2: Core Sync

- [ ] File watching implementation
- [ ] Git commit/push operations
- [ ] Pull and apply changes
- [ ] Stow integration
- [ ] Secret detection and filtering

### Phase 3: Daemon

- [ ] Background service
- [ ] Systemd integration
- [ ] Conflict detection
- [ ] Queue system for offline commits
- [ ] Health monitoring

### Phase 4: Advanced Features

- [ ] Conflict resolution UI
- [ ] Rollback functionality
- [ ] Machine-specific templating
- [ ] Multi-repo support
- [ ] Sync statistics and analytics

### Phase 5: Polish

- [ ] Comprehensive testing
- [ ] Documentation
- [ ] AUR package
- [ ] Example configurations
- [ ] Video tutorials

## 🤝 Contributing

This project is in early development. Contributions, ideas, and feedback are welcome!

## 📄 License

MIT License - See LICENSE file for details

## 🙏 Acknowledgments

- [Charm](https://charm.sh/) for amazing TUI libraries
- [Hyprland](https://hyprland.org/) community
- [Arch Linux](https://archlinux.org/) and the dotfiles management community

---

**Status**: 🚧 In Development

Built with ❤️ for the Omarchy community
