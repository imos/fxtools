package cookiejar

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func (j *Jar) Load(filename string) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, j)
}

func (j *Jar) Save(filename string) error {
	buf, err := json.Marshal(j)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, buf, os.ModePerm)
}
