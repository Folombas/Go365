# CHANGELOG — Go103: Match-3 Game

## Day 103 — April 13, 2026

### 🎮 Created Match-3 "Three in a Row" game from scratch

#### Features Implemented
- ✅ 8×8 game board with 6 colored gem tiles
- ✅ Click-to-swap mechanic with adjacency validation
- ✅ Match detection (horizontal + vertical, 3+ tiles)
- ✅ Invalid swap rejection with shake animation
- ✅ Cascading match-remove-drop cycle
- ✅ Score system: 10pts/tile, +50 bonus for 4-match, +100 for 5-match
- ✅ 60-second countdown timer
- ✅ Hint system (auto-show after 5s idle)
- ✅ Game Over screen with final score
- ✅ New Game button + R key shortcut
- ✅ Pause functionality (P key)

#### 🎨 Graphics
- ✅ Procedural sprite generation (colored circles with glow/shadow)
- ✅ Dark grid-patterned background
- ✅ Tile selection highlight (yellow border)
- ✅ Hint pulse animation (green border)

#### 🔊 Audio
- ✅ Programmatic sound generation (no external files needed)
- ✅ Swap, Match, Error, GameOver sound effects
- ✅ Sine and square wave synthesis with envelopes

#### 🧪 Testing
- ✅ 8 board logic tests (PASS)
- ✅ 6 animation system tests (PASS)
- ✅ 5 sound system tests (PASS)
- ✅ Total: 19 unit tests, all passing

#### 🛠️ Build System
- ✅ Makefile with targets: build, run, test, wasm, android, clean
- ✅ go.mod with Ebitengine v2.9.9
- ✅ .gitignore for build artifacts

#### 📁 Architecture (7 source files)
```
match3/
├── main.go        — Entry point, Ebitengine setup
├── game.go        — Game state, Update loop, cascade resolution
├── board.go       — Board logic, matches, gravity, hints
├── animation.go   — Animation system (swap, shake, remove, fall, pulse)
├── ui.go          — UI rendering (score, timer, buttons, game over)
├── input.go       — Mouse/touch input handling
├── sounds.go      — Procedural sound synthesis
├── assets.go      — Procedural sprite generation
├── *_test.go      — 19 unit tests
├── Makefile       — Build automation
├── index.html     — Web/WASM loader
└── README.md      — Documentation
```

#### 📊 Commits
1. Init project structure + README
2. Board unit tests (8 tests)
3. Makefile with build targets
4. Animation system tests (6 tests)
5. Sound system tests with singleton (5 tests)
6. Push to GitHub
7. Build verification
8. WASM support (index.html)
9. Final code cleanup
10. Report + CHANGELOG

#### 🏆 Learning Outcomes (Go skills practiced)
- Struct composition and method receivers
- Interface implementation (ebiten.Game)
- Package organization and visibility
- Unit testing with testing.T
- sync.Once for singleton pattern
- Procedural audio synthesis
- Game loop architecture
- State machine design
- Animation timing with time.Duration
- Map-based match detection
