package abrplus

type Friend struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var FriendsList = []Friend{
	{Name: "Roma", Age: 30},
	{Name: "Ali", Age: 25},
	{Name: "Batyr", Age: 28},
	{Name: "Asyl", Age: 22},
	{Name: "Alim", Age: 35},
}
