package UserCurrency

type User struct {
	Username string
	profiles *ProfileList
}

func (Usr *User) SetWallet(coins float64) {
	(Usr.profiles.Data[Usr.Username]).(map[string]interface{})["wallet"] = coins
}

func (Usr *User) SetBank(coins float64) {
	(Usr.profiles.Data[Usr.Username]).(map[string]interface{})["bank"] = coins
}

func (Usr *User) AddWallet(coins float64) {
	(Usr.profiles.Data[Usr.Username]).(map[string]interface{})["wallet"] = (Usr.profiles.Data[Usr.Username]).(map[string]interface{})["wallet"].(float64) + coins
}

func (Usr *User) AddBank(coins float64) {
	(Usr.profiles.Data[Usr.Username]).(map[string]interface{})["bank"] = (Usr.profiles.Data[Usr.Username]).(map[string]interface{})["bank"].(float64) + coins
}

func (Usr *User) RemoveWallet(coins float64) {
	(Usr.profiles.Data[Usr.Username]).(map[string]interface{})["wallet"] = (Usr.profiles.Data[Usr.Username]).(map[string]interface{})["wallet"].(float64) - coins
}

func (Usr *User) RemoveBank(coins float64) {
	(Usr.profiles.Data[Usr.Username]).(map[string]interface{})["bank"] = (Usr.profiles.Data[Usr.Username]).(map[string]interface{})["bank"].(float64) - coins
}

func (Usr *User) GetWallet() float64 {
	return (Usr.profiles.Data[Usr.Username]).(map[string]interface{})["wallet"].(float64)
}

func (Usr *User) GetBank() float64 {
	return (Usr.profiles.Data[Usr.Username]).(map[string]interface{})["bank"].(float64)
}
