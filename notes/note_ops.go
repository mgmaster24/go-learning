package gol_notes

type NoteOps interface {
	Save(fileName string) error
	Print()
}
