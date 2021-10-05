package main

import "fmt"

type Label struct {
	Name        string
	Description	string
	Color       string
}

var UrgentLabel = Label{
	Name: "Urgent",
	Description: "Urgent Item",
	Color: "darkred",
}

var ImportantLabel = Label{
	Name: "Important",
	Description: "Important Item",
	Color: "orange",
}

var WarningLabel = Label{
	Name: "Warning",
	Description: "Warning Item",
	Color: "yellow",
}

const (
	UrgentLabelKey = "Urgent"
	ImportantLabelKey = "Important"
	WarningLabelKey = "Warning"
)

var LabelMap = map[string]Label{
	UrgentLabelKey:    UrgentLabel,
	ImportantLabelKey: ImportantLabel,
	WarningLabelKey: WarningLabel,
}

func GetLabel(key string) Label {
	return LabelMap[key]
}

type Item struct {
	Name string
	Label string
}

func (s *Item) Describe() string {
	label := GetLabel(s.Label)
	return fmt.Sprintf("%s (%s) - %s", label.Name, label.Color, label.Description)
}

func main() {
	item1 := Item{
		Name:  "Item 1",
		Label: UrgentLabelKey,
	}

	item2 := Item{
		Name:  "Item 2",
		Label: ImportantLabelKey,
	}

	item3 := Item{
		Name:  "Item 3",
		Label: WarningLabelKey,
	}

	fmt.Println(item1.Name, "->", item1.Describe())
	fmt.Println(item2.Name, "->", item2.Describe())
	fmt.Println(item3.Name, "->", item3.Describe())
}