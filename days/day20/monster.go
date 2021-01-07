package main

type Monster struct {
	content       [][]rune
	width, height int
}

func NewMonsterFromData(contentData []string) *Monster {
	var content [][]rune
	for _, row := range contentData {
		var line []rune
		for _, letter := range row {
			line = append(line, letter)
		}
		content = append(content, line)
	}

	monster := Monster{content: content, width: len(content[0]), height: len(content)}
	return &monster
}

func (m *Monster) rotate() *Monster {
	newContent := make([][]rune, m.width)
	for n := 0; n < m.width; n++ {
		newContent[n] = make([]rune, m.height)
	}
	for rowIndex := range m.content {
		for colIndex := range m.content[rowIndex] {
			newRowIndex := colIndex
			newColIndex := m.height - 1 - rowIndex
			newContent[newRowIndex][newColIndex] = m.content[rowIndex][colIndex]
		}
	}

	return &Monster{content: newContent, width: m.height, height: m.width}
}

func (m *Monster) flip() *Monster {
	newContent := make([][]rune, m.height)
	for row := range m.content {
		newContent[row] = make([]rune, m.width)
	}
	for rowIndex := range m.content {
		for colIndex := range m.content[rowIndex] {
			newRowIndex := rowIndex
			newColIndex := m.width - 1 - colIndex
			newContent[newRowIndex][newColIndex] = m.content[rowIndex][colIndex]
		}
	}

	return &Monster{content: newContent, width: m.width, height: m.height}
}

func (m *Monster) getAllPositions() []*Monster {
	rotate1 := m.rotate()
	rotate2 := rotate1.rotate()
	rotate3 := rotate2.rotate()
	flip1 := rotate3.flip()
	flipRotate1 := flip1.rotate()
	flipRotate2 := flipRotate1.rotate()
	flipRotate3 := flipRotate2.rotate()
	return []*Monster{m, rotate1, rotate2, rotate3, flip1, flipRotate1, flipRotate2, flipRotate3}
}

func (m *Monster) countRune(r rune) int {
	count := 0
	for _, row := range m.content {
		for _, col := range row {
			if col == r {
				count += 1
			}
		}
	}
	return count
}
