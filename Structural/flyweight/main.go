package main

import "fmt"

// Flyweight (shared state)
type CharacterStyle struct {
	Font  string
	Color string
}

// Context (unique state per object)
type Character struct {
	char     rune
	position int
	style    *CharacterStyle
}

func (c *Character) Render() {
	fmt.Printf("Character '%c' at %d with style %v\n", c.char, c.position, c.style)
}

// Flyweight Factory
type CharacterFactory struct {
	styles map[string]*CharacterStyle
}

func NewCharacterFactory() *CharacterFactory {
	return &CharacterFactory{styles: make(map[string]*CharacterStyle)}
}

func (f *CharacterFactory) GetStyle(font, color string) *CharacterStyle {
	key := font + ":" + color
	if style, ok := f.styles[key]; ok {
		return style
	}

	style := &CharacterStyle{Font: font, Color: color}
	f.styles[key] = style
	return style
}

func main() {
	factory := NewCharacterFactory()

	text := "hello"
	chars := []*Character{}

	for i, ch := range text {
		style := factory.GetStyle("Arial", "Red") // shared style
		chars = append(chars, &Character{
			char:     ch,
			position: i,
			style:    style,
		})
	}

	// render all characters
	for _, c := range chars {
		c.Render()
	}
}
