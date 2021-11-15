# UserCurrency
Simple user currency module that uses JSON.

# Install
```
go get github.com/NeutronX-dev/UserCurrency
```
# Functions
* `func UserCurrency.Read(path_to_json string) (*ProfileList, error)`
# Types & Methods
* `ProfileList`
* * `func (Profiles *ProfileList) UserExists(user string) bool`
* * `func (Profiles *ProfileList) CreateUser(user string)`
* * `func (Profiles *ProfileList) GetUser(user string) User`
* * `func (Profiles *ProfileList) SaveData() error`
* `User`
* * `func (Usr *User) SetWallet(coins float64)`
* * `func (Usr *User) SetBank(coins float64)`
* * `func (Usr *User) AddWallet(coins float64)`
* * `func (Usr *User) AddBank(coins float64)`
* * `func (Usr *User) RemoveWallet(coins float64)`
* * `func (Usr *User) RemoveBank(coins float64)`

# Example
## JSON (`users.json`) before executing:
```json
{}
```
## Code:
```go
package main

import (
	"github.com/NeutronX-dev/UserCurrency"
)

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	Users, err := UserCurrency.Read("./users.json")
	HandleError(err)

	if !Users.UserExists("Someone") {
		Users.CreateUser("Someone")
	}

	User := Users.GetUser("Someone")
	User.SetWallet(1000)

	err = Users.SaveData()
	// Note: Data will always be avaliable but never
	//       saved into the JSON until you call this
	//       method. So make sure to call it.
	HandleError(err)
}

```
## JSON (`users.json`) after executing:
```json
{
	"Someone": {
		"bank": 0,
		"wallet": 1000
	}
}
```

# LICENSE
![gnu-logo](logos/gplv3-88x31.png)

This program is free software: you can redistribute it and/or modify
it under the terms of the [GNU General Public License](https://github.com/NeutronX-dev/ws.js/blob/main/LICENSE) as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.