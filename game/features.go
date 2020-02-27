package game

type Feature uint64

const (
	NoFeature Feature = iota
	LadybugPieceFeature
	PillBugPieceFeature
	MosquitoPieceFeature

	TournamentQueensRuleFeature
)

var featureMap = map[Feature]bool{
	LadybugPieceFeature:         false,
	PillBugPieceFeature:         false,
	MosquitoPieceFeature:        false,
	TournamentQueensRuleFeature: false,
}

func copyFeatureMap() (features map[Feature]bool) {
	features = make(map[Feature]bool)
	for k, v := range featureMap {
		features[k] = v
	}
	return features
}
