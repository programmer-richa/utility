package forms

// FieldBuilder is a helper to construct the form Field
type FieldBuilder struct {
	formatters []FormatterFunc
	loader     LoaderFunc
	validators []ValidatorFunc
	required   bool
	empty      interface{}
	label      string
	min        int
	max        int
	fieldType  string
}

// WithFormatters specifies the name of the formatting function of the field.
func (fb *FieldBuilder) WithFormatters(formatters ...FormatterFunc) *FieldBuilder {
	fb.formatters = append(fb.formatters, formatters...)
	return fb
}

// WithValidators specifies the name of the validator function to attach
// the validator of the field.
func (fb *FieldBuilder) WithValidators(validators ...ValidatorFunc) *FieldBuilder {
	fb.validators = append(fb.validators, validators...)
	return fb
}

// Loader specifies the name of the loader function to attach
// the datatype of the field.
func (fb *FieldBuilder) Loader(loader LoaderFunc) *FieldBuilder {
	fb.loader = loader
	return fb
}

// Required specifies that the field is mandatory
func (fb *FieldBuilder) Required() *FieldBuilder {
	fb.required = true
	return fb
}

// Min sets min value acceptable by the FieldBuilder.
func (fb *FieldBuilder) Min(value int) *FieldBuilder {
	fb.min = value
	return fb
}

// Max sets max value acceptable by the FieldBuilder.
func (fb *FieldBuilder) Max(value int) *FieldBuilder {
	fb.max = value
	return fb
}

// FieldType sets type (text, textarea, email etc) of the FieldBuilder.
func (fb *FieldBuilder) FieldType(value string) *FieldBuilder {
	fb.fieldType = value
	return fb
}

// Label sets label of the FieldBuilder.
func (fb *FieldBuilder) Label(value string) *FieldBuilder {
	fb.label = value
	return fb
}

// Empty sets value of empty attribute of the FieldBuilder.
func (fb *FieldBuilder) Empty(value interface{}) *FieldBuilder {
	fb.empty = value
	return fb
}

// Build returns a pointer to Field
func (fb *FieldBuilder) Build() *Field {
	return &Field{
		Formatters: fb.formatters,
		Loader:     fb.loader,
		Validators: fb.validators,
		Required:   fb.required,
		Empty:      fb.empty,
		Max:        fb.max,
		Min:        fb.min,
		Label:      fb.label,
		FieldType:  fb.fieldType,
	}
}
