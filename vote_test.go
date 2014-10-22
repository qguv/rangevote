package rangevote

import "testing"

func TestObviousElection(t *testing.T) {
	bush := candidate{"G. H. W. Bush", 1}
	wbush := candidate{"G. W. Bush", 2}
	marx := candidate{"Karl Marx", 3}
	trotsky := candidate{"Leon Trotsky", 4}
	ponies := candidate{"Vermin Supreme", 5}
	satan := candidate{"Satan", 666}

	e := NewElection(
		"U.S. President",
		0, 9,
		[]candidate{
			bush,
			wbush,
			marx,
			trotsky,
			ponies,
			satan,
		},
	)
	e.vote(marx, 9)
	e.vote(bush, 9)
	e.vote(marx, 9)
	e.vote(marx, 9)
	e.vote(marx, 9)
	winner, _ := e.winner()
	if winner != marx {
		t.Fail()
	}
}
