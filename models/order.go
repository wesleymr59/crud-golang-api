package models

type Order struct {
	Id       int    `gorm:"primaryKey,autoIncrement:true;"`
	Pedido   string `json:"pedido"`
	Valor    int    `json:"valor"`
	ClientId int    `json:"clientId"`
	Client   Client `json:"client" gorm:"foreignKey:ClientId"`
}

type Client struct {
	ID   int    `gorm:"primaryKey,autoIncrement:true;"`
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}
