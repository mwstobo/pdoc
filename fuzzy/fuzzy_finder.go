package fuzzy

type FuzzyFinder interface {
	Find(query string, targets []string) []string
}
