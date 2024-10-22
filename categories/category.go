package categories

import(
	"webapp/types"
)

func GetCategories() []types.Category{
	return []types.Category{
		{Link: "/engineering", Label: "Engineering"},
	}

}