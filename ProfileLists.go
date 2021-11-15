package UserCurrency

import (
	"encoding/json"
	"io/ioutil"
)

type ProfileList struct {
	Data map[string]interface{}
	Path string
}

// Check if user Exists
func (Profiles *ProfileList) UserExists(user string) bool {
	return Profiles.Data[user] != nil
}

// Makes a Profile
func (Profiles *ProfileList) CreateUser(user string) {
	usr := make(map[string]interface{})
	usr["wallet"] = 0
	usr["bank"] = 0
	Profiles.Data[user] = usr
}

func (Profiles *ProfileList) GetUser(user string) User {
	return User{Username: user, profiles: Profiles}
}

// Save all data into the JSON file.
func (Profiles *ProfileList) SaveData() error {
	data, err := json.MarshalIndent(Profiles.Data, "", "\t")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(Profiles.Path, data, 0777)
	if err != nil {
		return err
	}
	return nil
}

func Read(path string /* path to json */) (*ProfileList, error) {
	var res ProfileList = ProfileList{Data: make(map[string]interface{}), Path: path}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return &res, err
	}
	err = json.Unmarshal(data, &(res.Data))
	if err != nil {
		return &res, err
	}
	return &res, nil
}
