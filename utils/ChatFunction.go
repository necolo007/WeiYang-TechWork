package utils

import (
	"WeiYangWork/Model"
	"WeiYangWork/global"
)

func CreateId(uid, toUid string) string {
	return uid + "->" + toUid
}

// IsUserInTeam checks if a user is a member of a team
func IsUserInTeam(user *Model.UserClaims, teamID uint) bool {
	var team Model.Team
	err := global.Db.Model(&Model.Team{}).Where("id = ?", teamID).Preload("Member").Find(&team).Error
	if err != nil {
		return false
	}

	for _, member := range team.Member {
		if member.ID == user.UserId || user.Username == team.Leader {
			return true
		}
	}
	return false
}
