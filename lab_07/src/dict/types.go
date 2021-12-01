package dict

// Dict used to represent dictionary with custom types.
type Dict map[string]interface{}

// DictArray used to represent array of Dict instances.
type DictArray []Dict

// Freq used to represent frequency analyser type.
type Freq struct {
	l    string
	cnt  int
	darr DictArray
}

// FreqArray used to represent array of Freq instances.
type FreqArray []Freq
