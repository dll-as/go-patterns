package main

import "fmt"

// Memento
type EditorMemento struct {
	content string
}

// Originator
type TextEditor struct {
	content string
}

func (e *TextEditor) Write(text string) {
	e.content += text
}

func (e *TextEditor) Save() *EditorMemento {
	return &EditorMemento{content: e.content}
}

func (e *TextEditor) Restore(m *EditorMemento) {
	e.content = m.content
}

func (e *TextEditor) Read() string {
	return e.content
}

// Caretaker
type History struct {
	mementos []*EditorMemento
}

func (h *History) Push(m *EditorMemento) {
	h.mementos = append(h.mementos, m)
}

func (h *History) Pop() *EditorMemento {
	if len(h.mementos) == 0 {
		return nil
	}

	last := h.mementos[len(h.mementos)-1]
	h.mementos = h.mementos[:len(h.mementos)-1]
	return last
}

func main() {
	editor := &TextEditor{}
	history := &History{}

	editor.Write("Hello ")
	history.Push(editor.Save())

	editor.Write("World!")
	history.Push(editor.Save())

	fmt.Println("Current:", editor.Read()) // Hello World!

	editor.Restore(history.Pop())
	fmt.Println("Undo 1:", editor.Read()) // Hello

	editor.Restore(history.Pop())
	fmt.Println("Undo 2:", editor.Read()) // (empty)
}
