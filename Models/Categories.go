package Models

import (
	"fmt"

	"alexedwards.net/CC-GO/Config"
	_ "github.com/go-sql-driver/mysql"
)

func CreateACategory(catergory *Catergory) (err error) {
	if err = Config.DB.Create(catergory).Error; err != nil {
		return err
	}
	return nil
}

func GetAllCategories(category *[]Catergory) (err error) {
	if err = Config.DB.Find(category).Error; err != nil {
		return err
	}
	return nil
}

func GetACategory(category *Catergory, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(category).Error; err != nil {
		return err
	}
	return nil
}

func UpdateACategory(category *Catergory, id string) (err error) {
	fmt.Println(category)
	Config.DB.Save(category)
	return nil
}

func DeleteACategory(category *Catergory, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(category)
	return nil
}
