package env

var config *ENV

func Init(filename string) (*ENV, error) {
	val, err := loadENV(filename)
	if err != nil {
		return nil, err
	}
	config = val
	return config, nil
}

func Get() *ENV {
	return config
}
