package maps

const (
	ErrKeyNotFound       = DictionaryErr("could not find the word you were looking for")
	ErrKeyAlreadyExist   = DictionaryErr("cannot add word because it already exists")
	ErrUpdateKeyNotFound = DictionaryErr("cannot update not existing key")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	v, ok := d[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return v, nil
}

func (d Dictionary) Add(key, value string) error {
	_, ok := d[key]
	if !ok {
		d[key] = value
		return nil
	}
	return ErrKeyAlreadyExist
}

func (d Dictionary) Update(key, value string) error {
	_, ok := d[key]
	if !ok {
		return ErrUpdateKeyNotFound
	}
	d[key] = value
	return nil
}
