package samples

import (
	"fmt"
	"strings"
	"testing"

	dc0d "github.com/dc0d/caseconv"
	ettle "github.com/ettle/strcase"
	"github.com/gobeam/stringy"
	iancoleman "github.com/iancoleman/strcase"
	mantidtech "github.com/mantidtech/wordcase"
	nikitaksv "github.com/nikitaksv/strcase"
	parithiban "github.com/parithiban/stringcases"
	pascaldekloe "github.com/pascaldekloe/name"
	stoewer "github.com/stoewer/go-strcase"
)

// FIXME pattern names
const (
	linePattern  = "| %-67s | %-19s | %-19s | %-19s | %-19s |\n"
	linePattern2 = "| %-14s | %-19s | %-19s | %-19s | %-19s |\n"
)

func TestSummary(t *testing.T) {
	samples := []string{"GooID", "HTTPStatusCode", "FooBAR", "URL", "ID", "hostIP", "JSON", "JSONName", "NameJSON", "UneTête"}

	extractors := []func(string) Info{
		infoEttle,
		infoStringy,
		infoStoewer,
		infoIancoleman,
		infoPascaldekloe,
		infoNikitaksv,
		infoMantidtech,
		infoDc0d,
		infoParithiban,
	}

	fmt.Printf(linePattern, "Lib", "Snake", "Kebab", "Pascal Case", "Camel Case")
	fmt.Printf("|%s|%[2]s|%[2]s|%[2]s|%[2]s|\n", strings.Repeat("-", 69), strings.Repeat("-", 21))

	for _, value := range samples {
		fmt.Printf(linePattern, "**"+value+"**", "-", "-", "-", "-")

		for _, extractor := range extractors {
			extractor(value).PrintWithLibName()
		}
	}

	fmt.Println()

	for _, extractor := range extractors {
		fmt.Println("### ", extractor("noop").Name)
		fmt.Println()

		fmt.Printf(linePattern2, "Source", "Snake", "Kebab", "Pascal Case", "Camel Case")
		fmt.Printf("|%s|%[2]s|%[2]s|%[2]s|%[2]s|\n", strings.Repeat("-", 16), strings.Repeat("-", 21))

		for _, value := range samples {
			extractor(value).PrintWithValue()
		}

		fmt.Println()
	}
}

type Info struct {
	Name       string
	Value      string
	Snake      string
	UpperSnake string
	Kebab      string
	UpperKebab string
	Pascal     string
	Camel      string
}

func (i Info) PrintWithLibName() {
	fmt.Printf(linePattern, fmt.Sprintf("[%s](https://%s)", strings.TrimPrefix(i.Name, "github.com/"), i.Name),
		i.Snake,
		i.Kebab,
		i.Pascal,
		i.Camel,
	)
}

func (i Info) PrintWithValue() {
	fmt.Printf(linePattern2, i.Value,
		i.Snake,
		i.Kebab,
		i.Pascal,
		i.Camel,
	)
}

func infoDc0d(value string) Info {
	return Info{
		Name:       "github.com/dc0d/caseconv",
		Value:      value,
		Snake:      dc0d.ToSnake(value),
		UpperSnake: "",
		Kebab:      dc0d.ToKebab(value),
		UpperKebab: "",
		Pascal:     dc0d.ToPascal(value),
		Camel:      dc0d.ToCamel(value),
	}
}

func infoMantidtech(value string) Info {
	return Info{
		Name:       "github.com/mantidtech/wordcase",
		Value:      value,
		Snake:      mantidtech.SnakeCase(value),
		UpperSnake: mantidtech.ScreamingSnakeCase(value),
		Kebab:      mantidtech.KebabCase(value),
		UpperKebab: "",
		Pascal:     mantidtech.PascalCase(value),
		Camel:      mantidtech.CamelCase(value),
	}
}

func infoParithiban(value string) Info {
	return Info{
		Name:       "github.com/parithiban/stringcases",
		Value:      value,
		Snake:      parithiban.ConvertToSnake(value),
		UpperSnake: "",
		Kebab:      parithiban.ConvertToKebab(value),
		UpperKebab: "",
		Pascal:     parithiban.ConvertToPascal(value),
		Camel:      parithiban.ConvertToCamel(value),
	}
}

func infoNikitaksv(value string) Info {
	return Info{
		Name:       "github.com/nikitaksv/strcase",
		Value:      value,
		Snake:      nikitaksv.ToSnakeCase(value),
		UpperSnake: "",
		Kebab:      nikitaksv.ToKebabCase(value),
		UpperKebab: "",
		Pascal:     nikitaksv.ToPascalCase(value),
		Camel:      nikitaksv.ToCamelCase(value),
	}
}

func infoEttle(value string) Info {
	// TODO ToGOSnake
	return Info{
		Name:       "github.com/ettle/strcase",
		Value:      value,
		Snake:      ettle.ToSnake(value),
		UpperSnake: ettle.ToSNAKE(value),
		Kebab:      ettle.ToKebab(value),
		UpperKebab: ettle.ToKEBAB(value),
		Pascal:     ettle.ToPascal(value),
		Camel:      ettle.ToCamel(value),
	}
}

func infoPascaldekloe(value string) Info {
	return Info{
		Name:       "github.com/pascaldekloe/name",
		Value:      value,
		Snake:      pascaldekloe.SnakeCase(value),
		UpperSnake: "",
		Kebab:      "",
		UpperKebab: "",
		Pascal:     pascaldekloe.CamelCase(value, true),
		Camel:      pascaldekloe.CamelCase(value, false),
	}
}

func infoStoewer(value string) Info {
	return Info{
		Name:       "github.com/stoewer/go-strcase",
		Value:      value,
		Snake:      stoewer.SnakeCase(value),
		UpperSnake: stoewer.UpperSnakeCase(value),
		Kebab:      stoewer.KebabCase(value),
		UpperKebab: stoewer.UpperKebabCase(value),
		Pascal:     stoewer.UpperCamelCase(value),
		Camel:      stoewer.LowerCamelCase(value),
	}
}

func infoStringy(value string) Info {
	str := stringy.New(value)

	return Info{
		Name:       "github.com/gobeam/stringy",
		Value:      value,
		Snake:      str.SnakeCase().Get(),
		UpperSnake: str.SnakeCase().ToUpper(),
		Kebab:      str.KebabCase().Get(),
		UpperKebab: str.KebabCase().ToUpper(),
		Pascal:     str.CamelCase(),
		Camel:      "",
	}
}

func infoIancoleman(value string) Info {
	return Info{
		Name:       "github.com/iancoleman/strcase",
		Value:      value,
		Snake:      iancoleman.ToSnake(value),
		UpperSnake: iancoleman.ToScreamingSnake(value),
		Kebab:      iancoleman.ToKebab(value),
		UpperKebab: iancoleman.ToScreamingKebab(value),
		Pascal:     iancoleman.ToCamel(value),
		Camel:      iancoleman.ToLowerCamel(value),
	}
}
