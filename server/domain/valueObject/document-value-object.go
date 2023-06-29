package valueObject

type Document struct {
	document string
}

func (d *Document) Prepare(document, personTypes string) *Document {
	// Lógica de document do documento com base no tipo de pessoa
	return &Document{
		document: document,
	}
}
