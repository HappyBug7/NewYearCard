package service

import (
	"Back_end/model"
	"errors"
	"strconv"
)

type Card struct {
}

type SimplifiedCard struct {
	FromUser string
	ToUser   string
	Content  string
}

func (ca *Card) AddCard(simplecard SimplifiedCard) (err error) {
	card := model.Card{}
	card.FromUser = simplecard.FromUser
	card.ToUser = simplecard.ToUser
	card.Content = simplecard.Content
	if err := model.DB.Model(&model.Card{}).Where("from_user = ? AND to_user = ? AND content = ?", card.FromUser, card.ToUser, card.Content).FirstOrCreate(&card).Error; err != nil {
		return err
	}
	return nil
}

func (ca *Card) GetCard(EncryptedId string, d *DES) (err error, cards SimplifiedCard) {
	id, err := strconv.Atoi(d.Final_DES_Decode_process(EncryptedId))
	if err != nil {
		return errors.New("Invalid opration"), SimplifiedCard{}
	}
	card := model.Card{}
	if err := model.DB.Where("id = ?", id).First(&card).Error; err != nil {
		return err, SimplifiedCard{}
	}
	cards.Content = card.Content
	cards.FromUser = card.FromUser
	cards.ToUser = card.ToUser
	return nil, cards
}

func (ca *Card) GetEncryptedID(id string, d *DES) (EncryptedId string) {
	return d.Final_DES_Encode_process(id)
}

func (ca *Card) GetDecryptedID(id string, d *DES) (DecryptedId string) {
	return d.Final_DES_Decode_process(id)
}
