# UI Guide

This document explains the Bubble Tea TUI implementation in QnQ.

## Overview

QnQ uses [Bubble Tea](https://github.com/charmbracelet/bubbletea), a Go framework for building terminal user interfaces based on The Elm Architecture.

## The Elm Architecture (MVU Pattern)

Bubble Tea implements the Model-View-Update pattern:

```
┌─────────────────────────────────────────────────┐
│                                                 │
│  User Input → Message → Update → Model → View  │
│                            ↑          │         │
│                            └──────────┘         │
│                                                 │
└─────────────────────────────────────────────────┘
```

### Components

1. **Model**: Application state
2. **Update**: State transitions
3. **View**: Rendering logic
4. **Messages**: Events that trigger updates

## Model Structure

The model holds all application state:

```go
type model struct {
    choices  []string           // Available menu options
    cursor   int                // Current cursor position
    selected map[int]struct{}   // Selected items
}
```

### Fields Explained

- **choices**: List of menu items displayed to the user
- **cursor**: Zero-based index of the currently highlighted item
- **selected**: Set of selected item indices (using empty struct for memory efficiency)

## Initialization

The `Init()` method is called when the program starts:

```go
func (m model) Init() tea.Cmd {
    return nil  // No initial commands needed
}
```

### Initial State

```go
func initialModel() model {
    return model{
        choices:  []string{"Server", "Client", "Config", "Exit"},
        selected: make(map[int]struct{}),
    }
}
```

## Update Function

The `Update()` method handles all user input and state changes:

```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd)
```

### Message Types

#### KeyMsg

Handles keyboard input:

| Input | Action | Implementation |
|-------|--------|----------------|
| `ctrl+c`, `q` | Quit | `return m, tea.Quit` |
| `up`, `k` | Move up | `m.cursor--` (with bounds check) |
| `down`, `j` | Move down | `m.cursor++` (with bounds check) |
| `enter`, `space` | Select/Exit | Toggle selection or quit if "Exit" |

### State Transitions

```go
// Moving up
if m.cursor > 0 {
    m.cursor--
}

// Moving down
if m.cursor < len(m.choices)-1 {
    m.cursor++
}

// Toggling selection
if _, ok := m.selected[m.cursor]; ok {
    delete(m.selected, m.cursor)  // Deselect
} else {
    m.selected[m.cursor] = struct{}{}  // Select
}
```

## View Function

The `View()` method renders the UI:

```go
func (m model) View() string
```

### Rendering Logic

1. **Header**: Application title and instructions
2. **Menu Items**: List of choices with cursor and selection indicators
3. **Footer**: Help text for keyboard controls

### Visual Elements

```
QnQ - Command & Conquer DevOps Center

What would you like to do?

> [x] Server    ← Cursor and selected
  [ ] Client
  [ ] Config
  [ ] Exit

Press q to quit.
Use arrow keys or j/k to navigate, enter/space to select.
```

#### Indicators

- `>` : Cursor indicator (current position)
- `[x]`: Selected item
- `[ ]`: Unselected item

### Rendering Code

```go
for i, choice := range m.choices {
    cursor := " "
    if m.cursor == i {
        cursor = ">"  // Show cursor
    }

    checked := " "
    if _, ok := m.selected[i]; ok {
        checked = "x"  // Show selection
    }

    s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
}
```

## Program Lifecycle

```go
// Create program
p := tea.NewProgram(initialModel())

// Run program (blocks until quit)
if _, err := p.Run(); err != nil {
    return fmt.Errorf("error running TUI: %w", err)
}
```

### Lifecycle Events

1. **Start**: `initialModel()` creates initial state
2. **Init**: `Init()` called, returns initial commands
3. **Loop**: 
   - User input generates messages
   - `Update()` processes messages
   - `View()` renders new state
4. **Exit**: `tea.Quit` command terminates loop

## Keyboard Handling

### Supported Keys

```go
switch msg.String() {
case "ctrl+c", "q":      // Quit
case "up", "k":          // Navigate up
case "down", "j":        // Navigate down
case "enter", " ":       // Select/Action
}
```

### Key Design Decisions

- **Vim-style navigation**: `j`/`k` for up/down
- **Multiple quit options**: `q` and `Ctrl+C`
- **Dual selection keys**: `Enter` and `Space`

## Styling (Future Enhancement)

Currently, the UI uses plain text. Future versions will use Lip Gloss for styling:

```go
// Example with Lip Gloss (not yet implemented)
import "github.com/charmbracelet/lipgloss"

var titleStyle = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#7D56F4")).
    Background(lipgloss.Color("#1a1a1a")).
    Padding(0, 1)
```

## Best Practices

### 1. Immutable Updates

Always return a new model, don't modify in place:

```go
// Good
newModel := m
newModel.cursor++
return newModel, nil

// Also good (Go allows this)
m.cursor++
return m, nil
```

### 2. Bounds Checking

Always validate cursor position:

```go
if m.cursor > 0 {
    m.cursor--
}
```

### 3. Efficient Rendering

The `View()` function is called frequently, so:
- Avoid expensive operations
- Use string builder for complex views
- Cache computed values when possible

### 4. Error Handling

Handle errors gracefully in Update:

```go
case errMsg:
    m.err = msg
    return m, nil
```

## Advanced Features (Future)

### Multiple Views

```go
type view int

const (
    mainMenu view = iota
    serverView
    configView
)

type model struct {
    currentView view
    // ... other fields
}
```

### Animations

```go
type tickMsg time.Time

func tickCmd() tea.Cmd {
    return tea.Tick(time.Second, func(t time.Time) tea.Msg {
        return tickMsg(t)
    })
}
```

### Spinners and Progress

```go
import "github.com/charmbracelet/bubbles/spinner"

type model struct {
    spinner spinner.Model
    // ... other fields
}
```

## Testing

### Testing the Model

```go
func TestUpdate(t *testing.T) {
    m := initialModel()
    
    // Test moving down
    msg := tea.KeyMsg{Type: tea.KeyDown}
    newModel, _ := m.Update(msg)
    
    if newModel.(model).cursor != 1 {
        t.Error("cursor should move down")
    }
}
```

### Testing the View

```go
func TestView(t *testing.T) {
    m := initialModel()
    view := m.View()
    
    if !strings.Contains(view, "QnQ") {
        t.Error("view should contain title")
    }
}
```

## Performance Considerations

1. **View Rendering**: Called on every update, keep it fast
2. **State Size**: Keep model small, avoid large data structures
3. **String Building**: Use `strings.Builder` for complex views
4. **Commands**: Use commands for async operations

## Debugging Tips

### 1. Log to File

```go
f, _ := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
defer f.Close()
log.SetOutput(f)
log.Printf("Cursor: %d", m.cursor)
```

### 2. Debug View

Add debug info to the view:

```go
s += fmt.Sprintf("\nDebug: cursor=%d, selected=%v", m.cursor, m.selected)
```

### 3. Test in Isolation

Test Update and View functions independently:

```go
m := initialModel()
fmt.Println(m.View())
```

## Resources

- [Bubble Tea Documentation](https://github.com/charmbracelet/bubbletea)
- [Bubble Tea Examples](https://github.com/charmbracelet/bubbletea/tree/master/examples)
- [Lip Gloss Styling](https://github.com/charmbracelet/lipgloss)
- [Bubbles Components](https://github.com/charmbracelet/bubbles)

## Common Patterns

### Modal Dialogs

```go
type model struct {
    showModal bool
    modalText string
}
```

### List Navigation

```go
type model struct {
    items  []string
    cursor int
}

// Wrap cursor
m.cursor = (m.cursor + 1) % len(m.items)
```

### Input Forms

```go
import "github.com/charmbracelet/bubbles/textinput"

type model struct {
    textInput textinput.Model
}
```

## Conclusion

The Bubble Tea implementation in QnQ provides:
- Clean, functional architecture
- Easy to test and maintain
- Extensible for future features
- Cross-platform compatibility
- Responsive user experience
