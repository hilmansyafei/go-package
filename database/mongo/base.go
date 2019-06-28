package mongo

// BaseStruct struct
type BaseStruct struct {
	CreatedAt string `json:"createdAt" bson:"createdAt"`
	UpdatedAt string `json:"updatedAt" bson:"updatedAt"`
}
