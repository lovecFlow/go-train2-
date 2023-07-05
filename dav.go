package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(IntsCopy([]int{1, 2, 3, 4}, 5))
	fmt.Println(UniqueUserIDs([]int64{21, 3, 5, 3, 2, 61, 2, 16}))
	fmt.Println(MostPopularWord([]string{"hello", "world", "hello"}))
	fmt.Println(shiftASCII("anc", 2))
	fmt.Println(isASCII("hello \U0001F970"))
	fmt.Println(latinLetters("abcd"))
}

//Реализуйте функцию IntsCopy(src []int, maxLen int) []int, которая создает копию слайса src с длиной maxLen. Если maxLen равен нулю или отрицательный, то функция возвращает пустой слайс []int{}. Если maxLen больше длины src, то возвращается полная копия src.

func IntsCopy(src []int, maxLen int) []int {
	emptySlice := make([]int, 0)
	if maxLen <= 0 {
		return emptySlice
	} else if maxLen > len(src) {
		return src
	} else {
		scrCp := make([]int, maxLen)
		copy(scrCp, src)
		return scrCp
	}
}

//Реализуйте функцию UniqueUserIDs(userIDs []int64) []int64, которая возвращает слайс, состоящий из уникальных идентификаторов userIDs. Порядок слайса должен сохраниться.

func UniqueUserIDs(userIDs []int64) []int64 {
	m := make(map[int64]struct{}) //Создали мапу (struct -- это тип данных, который занимает 0 байт используется, когда нужно
	//проверять в мапе только наличие ключа)

	uniqUserIDs := make([]int64, 0) //Создаём слайс для уникальных ID
	for _, uid := range userIDs {   //Создаём цикл с ренджом userIDs
		if _, ok := m[uid]; ok { //Проверяем наличие элемента по ключу
			continue
		}
		uniqUserIDs = append(uniqUserIDs, uid) //Доваляем элементы uid в uniqUsersUIs
		m[uid] = struct{}{}                    //Записываем мапу в структуру
	}
	return uniqUserIDs
}

//Реализуйте функцию MostPopularWord(words []string) string, которая возвращает самое часто встречаемое слово в слайсе.
//Если таких слов несколько, то возвращается первое из них.

func MostPopularWord(words []string) string {
	m := make(map[string]int, 0)
	MosPopWord := ""
	max := 0

	for _, word := range words {
		m[word]++
		if m[word] > max {
			max = m[word]
			MosPopWord = word
		}
	}
	return MosPopWord
}

func shiftASCII(s string, step int) string {
	if step == 0 {
		return s
	}

	NewL := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		NewL[i] = byte(int(s[i]) + step)
	}

	return string(NewL)
}

func isASCII(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func latinLetters(s string) string {
	sb := &strings.Builder{}
	for _, r := range s {
		if unicode.Is(unicode.Latin, r) {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func generateSelfStory(name string, age int, money float64) string {
	return fmt.Sprintf("Hello! My name is %s. I'm %d y.o. And I also have %f in my wallet right now.", name, age, money)
}
