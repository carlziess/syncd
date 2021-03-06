
// Copyright 2018 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

import (
    "time"

    "github.com/dreamans/syncd/model"
)

type ProjectUser struct {
    ID          int      `gorm:"primary_key"`
    SpaceId     int      `gorm:"type:int(11);not null;default:0"`
    UserId      int      `gorm:"type:int(11);not null;default:0"`
    Ctime       int      `gorm:"type:int(11);not null;default:0"`
}

const (
    TableName = "project_user"
)

func Create(data *ProjectUser) bool {
    data.Ctime = int(time.Now().Unix())
    return model.Create(TableName, data)
}

func List(query model.QueryParam) ([]ProjectUser, bool) {
    var data []ProjectUser
    ok := model.GetMulti(TableName, &data, query)
    return data, ok
}

func GetSpaceIdsByUserId(userId int) ([]int, bool) {
    userList, ok := List(model.QueryParam{
        Fields: "space_id",
        Where: []model.WhereParam{
            model.WhereParam{
                Field: "user_id",
                Prepare: userId,
            },
        },
    })
    if !ok {
        return nil, false
    }
    var spaceIds []int
    for _, l := range userList {
        spaceIds = append(spaceIds, l.SpaceId)
    }
    return spaceIds, true
}

func Total(query model.QueryParam) (int, bool) {
    var count int
    ok := model.Count(TableName, &count, query)
    return count, ok
}

func Get(id int) (ProjectUser, bool){
    var data ProjectUser
    ok := model.GetByPk(TableName, &data, id)
    return data, ok
}

func GetOne(query model.QueryParam) (ProjectUser, bool) {
    var data ProjectUser
    ok := model.GetOne(TableName, &data, query)
    return data, ok
}

func Delete(id int) bool {
    ok := model.DeleteByPk(TableName, ProjectUser{ID: id})
    return ok
}
