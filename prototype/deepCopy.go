package prototype

type Filev2 struct {
	name string
	data *Data // Con trỏ đến đối tượng khác
}

type Data struct {
}

func (f *Filev2) clone() *Filev2 {
	clonedData := *f.data                                      // Sao chép đối tượng Data
	return &Filev2{name: f.name + "_clone", data: &clonedData} // Deep copy
}

func (f *Filev2) Clone() *Filev2 {
	return f.clone()
}
