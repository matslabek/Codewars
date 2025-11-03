package kata

// https://www.codewars.com/kata/525c65e51bf619685c000059/train/go
func Cakes(recipe map[string]int, available map[string]int) int {
	if len(recipe) > len(available) {
		return 0
	}

	var requiredIngredients []string
	for r := range recipe {
		requiredIngredients = append(requiredIngredients, r)
	}
	var howManyCakes int
	for i := 0; i < len(requiredIngredients); i++ {
		howManyFromThisIngredient := available[requiredIngredients[i]] / recipe[requiredIngredients[i]]
		if howManyFromThisIngredient == 0 {
			return 0
		}
		if i == 0 {
			howManyCakes = howManyFromThisIngredient
		}
		if howManyCakes > howManyFromThisIngredient {
			howManyCakes = howManyFromThisIngredient
		}
	}
	return howManyCakes
}
