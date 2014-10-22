package rangevote

import "math"

func (e election) vote(c candidate, w int) bool {
	if w > e.maxweight || w < e.minweight {
		return false
	}

	cstatus := e.cstatus[c]
	cstatus.weight += w
	cstatus.count++
	e.cstatus[c] = cstatus

	return true
}

func (e election) list() (candidates []candidate) {
	for c, _ := range e.cstatus {
		candidates = append(candidates, c)
	}
	return
}

/* election.winner determines the winning candidate of an election. A winner must also satisfy the so-called 'rule d' where a candidate is disqualified if she doesn't have half of the score sum of the sum-maximizing candidate. The election is run again without her. */
func (e election) winner() (winner candidate, ok bool) {

	// will be used to mark not-obvious winning situations
	ok = true

	// calculate the maximum sum for rule-d
	var maxweight int
	for _, s := range e.cstatus {
		if s.weight > maxweight {
			maxweight = s.weight
		}
	}

	// calculate the winner
	var winscore, thisscore float64
	for c, s := range e.cstatus {

		// ignore candidates who violate rule-d
		if s.weight < (maxweight / 2) {
			continue
		}

		// calculate the candidate's average score
		thisscore = float64(s.weight) / float64(s.count)

		// if any scores are quite close (within rounding error), mark a potential tie to be broken
		if math.Abs(thisscore-winscore) > 1e-4 {
			ok = false
		}

		// if we're better than the past winner, overwrite
		if thisscore > winscore {
			winner = c
			winscore = thisscore
		}
	}
	return
}

func NewElection(name string, minweight, maxweight int, candidates []candidate) *election {
	cstatus := make(map[candidate]status, len(candidates))
	for _, c := range candidates {
		cstatus[c] = status{}
	}
	return &election{
		name:    name,
		minweight: minweight,
		maxweight: maxweight,
		cstatus: cstatus,
	}
}
