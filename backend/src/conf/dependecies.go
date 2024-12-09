package conf

import (
	model "backend/src/models"
	"backend/src/services"
	"backend/src/utils"
)

type Conf struct {
	servicePlayer *services.ServicePlayer
	playerModel   model.PlayerModel
}

var conf *Conf

func Init() *Conf {
	return conf
}

func (c *Conf) Player() *services.ServicePlayer {
	if c.servicePlayer != nil {
		return c.servicePlayer
	}
	return services.NewPlayer(c.getPlayerModel(), utils.JsonWebToken{})
}

func (c *Conf) getPlayerModel() model.PlayerModel {
	if c.playerModel != nil {
		return c.playerModel
	}
	return model.NewPlayerModel()
}
