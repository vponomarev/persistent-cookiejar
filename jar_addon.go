package persistent_cookiejar

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func (j *Jar) ExportEntries() (e map[string]map[string]entry) {
	return j.entries
}

func (j *Jar) ImportEntries(e map[string]map[string]entry) {
	j.entries = e
}

func (j *Jar) SaveEntriesYAML(filename string) (ok bool, err error) {
	f, err := os.Create(filename)
	if err != nil {
		err = fmt.Errorf("cannot write file [%s]", filename)
		return false, err
	}
	defer f.Close()

	output, err := yaml.Marshal(j.entries)
	if err != nil {
		err = fmt.Errorf("cannot generate YAML output")
		return false, err
	}
	f.Write(output)

	return true, nil
}

func (j *Jar) LoadEntriesYAML(filename string) (ok bool, err error) {
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		err = fmt.Errorf("cannot read file [%s]", filename)
		return false, err
	}

	if err = yaml.Unmarshal(source, &j.entries); err != nil {
		err = fmt.Errorf("error parsing file [%s]: %v", filename, err)
		return false, err
	}

	return true, nil
}
