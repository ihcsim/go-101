package magnifier

func Magnify(i uint8) string {
	index := i - 'A'
	return magnifyCharacters[index]
}

var magnifyCharacters = []string{
	"A MAGNIFIED A",
	"A MAGNIFIED B",
	"A MAGNIFIED C",
	"A MAGNIFIED D",
	"A MAGNIFIED E",
	"A MAGNIFIED F",
	"A MAGNIFIED G",
	"A MAGNIFIED H",
	"A MAGNIFIED I",
	"A MAGNIFIED J",
	"A MAGNIFIED K",
	"A MAGNIFIED L",
	"A MAGNIFIED M",
	"A MAGNIFIED N",
	"A MAGNIFIED O",
	"A MAGNIFIED P",
	"A MAGNIFIED Q",
	"A MAGNIFIED R",
	"A MAGNIFIED S",
	"A MAGNIFIED T",
	"A MAGNIFIED U",
	"A MAGNIFIED V",
	"A MAGNIFIED W",
	"A MAGNIFIED X",
	"A MAGNIFIED Y",
	"A MAGNIFIED Z",
}
