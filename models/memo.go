package models

import (
	"time"

	"github.com/golang-crew/Bolierplate-CRUD-Gingonic/common"
	"github.com/golang-crew/Bolierplate-CRUD-Gingonic/requests"
)

// gen:qs
type Memos struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createAt"`
}

func GetMemo(id string) (memo Memos, err error) {
	uintID, err := common.ConvertStringToUint(id)
	if err != nil {
		return
	}

	err = MemoIDChecker(uintID)
	if err != nil {
		return
	}

	qs := NewMemosQuerySet(gGormDB)
	qs = qs.IDEq(uintID)
	err = qs.One(&memo)

	return
}

func GetMemoList() (memo []Memos, err error) {
	qs := NewMemosQuerySet(gGormDB)
	err = qs.All(&memo)
	return
}

func DeleteMemo(id string) (err error) {
	uintID, err := common.ConvertStringToUint(id)
	if err != nil {
		return
	}

	memo := &Memos{
		ID: uintID,
	}
	err = memo.Delete(gGormDB)

	return err
}

func CreateMemo(req requests.CreateMemo) (err error) {
	memo := &Memos{
		Content: req.Content,
	}

	err = memo.Create(gGormDB)

	return
}

func UpdateMemo(request requests.UpdateMemo, id string) (err error) {
	uintID, err := common.ConvertStringToUint(id)
	if err != nil {
		return
	}

	err = MemoIDChecker(uintID)
	if err != nil {
		return
	}

	_, err = NewMemosQuerySet(gGormDB).
		IDEq(uintID).
		GetUpdater().
		SetContent(request.Content).
		UpdateNum()

	return
}

func MemoIDChecker(id uint) (err error) {
	var memo Memos
	sql := "SELECT * FROM memos WHERE id = ?"
	err = gGormDB.Raw(sql, id).Scan(&memo).Error

	return
}
