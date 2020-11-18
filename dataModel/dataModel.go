package dataModel

type Oplog struct {
	DocumentKey       OplogDocumentKeyObject       `bson:"documentKey"`
	Namespace         OplogNamespaceObject         `bson:"ns"`
	OperationType     string                       `bson:"operationType"`
	UpdateDescription OplogUpdateDescriptionObject `bson:"updateDescription"`
}

type OplogDocumentKeyObject struct {
	DocumentId string `bson:"_id"`
}

type OplogNamespaceObject struct {
	Database   string `bson:"db"`
	Collection string `bson:"coll"`
}
type OplogUpdateDescriptionObject struct {
	UpdatedFields map[string]interface{} `bson:"updatedFields"`
}

type ChangeStreamDocument struct {
	Id                  string `bson:"_id"`
	Database            string `bson:"database"`
	Collection          string `bson:"collection"`
	OperationType       string `bson:"operationType"`
	TriggerUpdatedField string `bson:"triggerUpdatedField"`
}
