package handlers

import (
	auth "api_getway/genproto/authentication_service"
	item "api_getway/genproto/item_service"
)

type Handlers interface {
	Auth() AuthHendler 
	Item() Itemhendler
}

type hendlers struct {
	authHendler AuthHendler
	itemHendler Itemhendler
}

func NewHandler(AuthSer auth.EcoServiceClient, ItemSer item.EcoExchangeServiceClient) Handlers {
	return &hendlers{
		authHendler: NewAuthHendler(AuthSer),
		itemHendler: NewItemHendler(ItemSer),
	}
}

func (h *hendlers) Auth() AuthHendler {
	return h.authHendler
}

func (h *hendlers) Item() Itemhendler {
	return h.itemHendler
}
