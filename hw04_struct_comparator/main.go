package main

import (
	"fmt"
)

type Book struct {
	title, author  string
	id, year, size int
	rate           float32
}

func (b *Book) GetID() int              { return b.id }
func (b *Book) SetID(id int)            { b.id = id }
func (b *Book) GetYear() int            { return b.year }
func (b *Book) SetYear(year int)        { b.year = year }
func (b *Book) GetSize() int            { return b.size }
func (b *Book) SetSize(size int)        { b.size = size }
func (b *Book) GetRate() float32        { return b.rate }
func (b *Book) SetRate(rate float32)    { b.rate = rate }
func (b *Book) GetTitle() string        { return b.title }
func (b *Book) SetTitle(title string)   { b.title = title }
func (b *Book) GetAuthor() string       { return b.author }
func (b *Book) SetAuthor(author string) { b.author = author }

type ByWhat int

const (
	ByYear ByWhat = iota
	BySize
	ByRate
)

type BookComparator struct {
	comparison ByWhat
}

func NewBookComparator(comparison ByWhat) *BookComparator {
	return &BookComparator{comparison}
}

func (bc *BookComparator) Compare(book1, book2 *Book) bool {
	switch bc.comparison {
	case ByYear:
		return book1.GetYear() > book2.GetYear()
	case BySize:
		return book1.GetSize() > book2.GetSize()
	case ByRate:
		return book1.GetRate() > book2.GetRate()
	default:
		return false
	}
}

func main() {
	book1 := Book{
		"Бойцовский клуб", "Чак Паланик", 1, 2002, 256, 7.8,
	}
	book2 := Book{
		"Капиталъ", "Карл Маркс", 2, 1867, 480, 8.57,
	}
	fmt.Printf("Первая книга: %s. Автор: %s\n", book1.GetTitle(), book1.GetAuthor())
	fmt.Printf("Вторая книга: %s. Автор: %s\n", book2.GetTitle(), book2.GetAuthor())
	comparator := NewBookComparator(ByYear)
	fmt.Println("Сравнение по году:", comparator.Compare(&book1, &book2))

	comparator = NewBookComparator(BySize)
	fmt.Println("Сравнение по страницам:", comparator.Compare(&book1, &book2))

	comparator = NewBookComparator(ByRate)
	fmt.Println("Сравнение по рейтингу:", comparator.Compare(&book1, &book2))
}
