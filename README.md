# 🧠 Seiri (整理)

### **"The Second Brain for the Distracted Mind."**

**Seiri** is a modular, high-performance knowledge engine built for people whose brains have too many tabs open. It’s a PARA-method organizer that lives in your terminal today and will haunt your desktop tomorrow.

If your "system" for taking notes is currently a graveyard of `untitled_final_v2.md` files scattered across your `Downloads` folder, **Seiri** is here to slap some order into your life.

> **Why Go?** Because it's fast.
> **Why the Terminal?** Because clicking buttons is a distraction.
> **Why for ADHD?** Because if the interface isn't pretty and the feedback isn't instant, we're going to go watch 4 hours of cat documentaries instead of studying/working.

--- 

## Key Philosophy

* **Keyboard-First:** Deep integration with Vim-style keybindings ($h, j, k, l$).
* **Modern Brutalist:** High-contrast, bold aesthetics using the **Catppuccin Mocha** palette.
* **Modular Architecture:** A "Headless" engine that separates file logic from the interface, allowing for TUI, Desktop, and Web versions to share the same brain.
* **Local-First:** Your notes are plain Markdown files on your disk. No proprietary databases, no cloud lock-in.

---

## 1. The Core Engine (Platform Agnostic)

* **Logic-First Design:** The `/core` package contains zero UI code. It handles file system operations, metadata parsing, and search, making it compatible with any future interface.
* **Concurrency:** Leverages Go Goroutines for background file indexing and Git synchronization, ensuring the UI never flickers or freezes.
* **Cross-Platform:** Native support for **macOS** (development) and **Linux** (deployment on CachyOS/Arch) via Go's powerful cross-compilation.

## 2. The Interfaces

* **TUI (Current):** A terminal interface built with `bubbletea` and `lipgloss`. It follows the Model-View-Update (MVU) pattern for a snappy, state-driven experience.
* **Desktop (Future):** Built with `Wails`, allowing the same Go backend to power a modern web-tech frontend (Astro/React) for a full GUI experience.

---

## 📂 Project Structure (For the Curious)

```text
seiri/
├── main.go               # Entry point (TUI Launcher)
├── config/
│   └── config.go         # YAML loader (handles ~/.config/seiri/config.yml)
├── core/                 # THE BRAIN: Pure Go logic
│   ├── para.go           # PARA folder management
│   ├── search.go         # Fuzzy matching logic
│   ├── git.go            # Native Git wrapper (os/exec)
│   └── types.go          # Shared data structures
├── internal/             # TUI-SPECIFIC LOGIC (Bubble Tea)
│   ├── app.go            # State initialization
│   ├── update.go         # Keybindings & Events
│   └── view.go           # Layout & Styling
├── ui/                   # SHARED AESTHETICS
│   ├── theme.go          # Catppuccin palette & Lipgloss styles
│   └── kanban.go         # TUI-based column rendering
└── desktop/              # FUTURE: Wails GUI integration

```

---

## The Roadmap

### Phase 1: Foundation (Modular Core)

* [x] Initialize Go environment and Git repository.
* [ ] **Core Logic:** Implement file scanning and PARA categorization logic.
* [ ] **Config System:** Build a robust YAML loader for cross-platform note paths.
* [ ] **TUI Skeleton:** Display basic file lists in a multi-column layout.

### Phase 2: Interaction (MVU Pattern)

* [ ] **Navigation:** Smooth Vim-style movement between Kanban columns.
* [ ] **Markdown Preview:** Integrate `glamour` for rich terminal rendering of notes.
* [ ] **Live Editing:** Support for toggling Markdown checkboxes directly from the TUI.

### Phase 3: Utility (Intelligence)

* [ ] **Global Search:** Blazing fast fuzzy search across the entire notes directory.
* [ ] **Triage Mode:** Quick-move workflow to process the `00_Inbox` into relevant PARA folders.
* [ ] **Git Sync:** Background auto-commits and remote syncing.

### Phase 4: Expansion (The Second Face)

* [ ] **Polish:** Finalize the "Modern Brutalist" aesthetic with thick borders and icons.
* [ ] **GUI Prototype:** Initialize `Wails` to bridge the Go core to a web-based frontend.

---

## 🛠️ Installation (For the Brave)

**On your Mac:**

```bash
go build -o seiri main.go
./seiri

```

---

*Made with ❤️, caffeine, and a total lack of focus by a CS student.*
