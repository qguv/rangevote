package rangevote

type election struct {
	name      string
	minweight int
	maxweight int
	cstatus   map[candidate]status
}

type status struct {
	weight int
	count  int
}

type candidate struct {
	name string
	id   int
}
