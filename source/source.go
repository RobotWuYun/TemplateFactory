type enKaNetWorkGenshinSummary struct {
PlayerInfo struct {
Nickname             string `json:"nickname"`
Level                int64  `json:"level"`
Signature            string `json:"signature"`
WorldLevel           int64  `json:"worldLevel"`
NameCardId           int64  `json:"nameCardId"`
FinishAchievementNum int64  `json:"finishAchievementNum"`
TowerFloorIndex      int64  `json:"towerFloorIndex"`
TowerLevelIndex      int64  `json:"towerLevelIndex"`
ShowAvatarInfoList   []struct {
AvatarId int64 `json:"avatarId"`
Level    int64 `json:"level"`
} `json:"showAvatarInfoList"`
ShowNameCardIdList []int `json:"showNameCardIdList"`
ProfilePicture     struct {
AvatarId int64 `json:"avatarId"`
} `json:"profilePicture"`
} `json:"playerInfo"`
AvatarInfoList []struct {
AvatarId int64 `json:"avatarId"`
PropMap  map[string]struct {
Type int64  `json:"type"`
Ival string `json:"ival"`
Val  string `json:"val"`
} `json:"propMap"`
FightPropMap           map[string]float64 `json:"fightPropMap"`
SkillDepotId           int64              `json:"skillDepotId"`
InherentProudSkillList []int64            `json:"inherentProudSkillList"`
SkillLevelMap          map[string]int64   `json:"skillLevelMap"`
EquipList              []struct {
ItemId    int64 `json:"itemId"`
Reliquary struct {
Level            int64   `json:"level"`
MainPropId       int64   `json:"mainPropId"`
AppendPropIdList []int64 `json:"appendPropIdList"`
} `json:"reliquary,omitempty"`
Flat struct {
NameTextMapHash    string `json:"nameTextMapHash"`
SetNameTextMapHash string `json:"setNameTextMapHash,omitempty"`
RankLevel          int64  `json:"rankLevel"`
ReliquaryMainstat  struct {
MainPropId string  `json:"mainPropId"`
StatValue  float64 `json:"statValue"`
} `json:"reliquaryMainstat,omitempty"`
ReliquarySubstats []struct {
AppendPropId string  `json:"appendPropId"`
StatValue    float64 `json:"statValue"`
} `json:"reliquarySubstats,omitempty"`
ItemType    string `json:"itemType"`
Icon        string `json:"icon"`
EquipType   string `json:"equipType,omitempty"`
WeaponStats []struct {
AppendPropId string  `json:"appendPropId"`
StatValue    float64 `json:"statValue"`
} `json:"weaponStats,omitempty"`
} `json:"flat"`
Weapon struct {
Level        int64            `json:"level"`
PromoteLevel int64            `json:"promoteLevel"`
AffixMap     map[string]int64 `json:"affixMap"`
} `json:"weapon,omitempty"`
} `json:"equipList"`
TalentIdList []int64 `json:"talentIdList"`
FetterInfo   struct {
ExpLevel int64 `json:"expLevel"`
} `json:"fetterInfo"`
} `json:"avatarInfoList"`
Ttl int64  `json:"ttl"`
Uid string `json:"uid"`
}

type ShowNameCardIdList struct{
A string
V int
}
