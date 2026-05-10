func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        fmt.Printf("Error loading config: %v\n", err)
        os.Exit(1)
    }

    // Pass the config into your model
    m := app.NewModel(cfg.NotesDir)
    
    p := tea.NewProgram(m)
    // ...
}
