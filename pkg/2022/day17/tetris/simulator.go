package tetris

func Simulate(chamber *Chamber, n int) {
	for i := 0; i < n; i++ {
		chamber.SpawnRock()
		for !chamber.CurrentRockResting() {
			chamber.Jet()
			chamber.Tick()
		}
	}
}
